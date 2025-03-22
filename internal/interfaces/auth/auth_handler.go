package interfaces_auth

import (
	"backend/config"
	pkg_logger "backend/internal/pkg/logger"
	pkg_timer "backend/internal/pkg/timer"
	usecase_auth "backend/internal/usecase/auth"
	pb "backend/proto/github.com/grpc/backend/proto"
	"context"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
)

// 認証ハンドラー層
type AuthHandler struct {
	logger    *pkg_logger.AppLogger
	timer     *pkg_timer.TimerPkg
	AppConfig *config.AppConfig
	pb.UnimplementedAuthServiceServer
	authUsecase usecase_auth.IAuthUsecase
}

// 認証ハンドラー層のインスタンス化
func NewAuthHandler(l *pkg_logger.AppLogger, ac *config.AppConfig, authUsecase usecase_auth.IAuthUsecase) *AuthHandler {
	return &AuthHandler{logger: l, AppConfig: ac, authUsecase: authUsecase, timer: pkg_timer.NewTimerPkg()}
}

// ログイン
func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	h.logger.InfoLog.Println("Login called")
	h.timer.Start()

	// ログイン(usecase層)
	token, err := h.authUsecase.Login(req.Email, req.Password)
	if err != nil {
		switch err.Error() {
		case "invalid email or password":
			h.logger.ErrorLog.Printf("Login failed: %v", err)
			h.logger.PrintDuration("Login", h.timer.GetDuration())
			return nil, status.Errorf(codes.Unauthenticated, "invalid email or password")
		case "invalid email format":
			h.logger.ErrorLog.Printf("Login failed: %v", err)
			h.logger.PrintDuration("Login", h.timer.GetDuration())
			return nil, status.Errorf(codes.InvalidArgument, "invalid email format")
		default:
			h.logger.ErrorLog.Printf("Login failed: %v", err)
			h.logger.PrintDuration("Login", h.timer.GetDuration())
			return nil, status.Errorf(codes.Internal, "failed to login")
		}
	}

	// トークンを生成
	tokenString, err := h.GenerateToken(token)
	if err != nil {
		h.logger.ErrorLog.Printf("Failed to generate token: %v", err)
		h.logger.PrintDuration("Login", h.timer.GetDuration())
		return nil, err
	}

	h.logger.InfoLog.Println("Login successful")
	h.logger.PrintDuration("Login", h.timer.GetDuration())
	return &pb.LoginResponse{Token: tokenString}, nil
}

// JWTトークンを生成
func (h *AuthHandler) GenerateToken(id string) (string, error) {
	h.logger.InfoLog.Println("Generating token...")
	h.timer.Start()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": h.AppConfig.UserRole,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	// JWTトークンをシグネーション
	tokenString, err := token.SignedString([]byte(h.AppConfig.JWTSecret))
	if err != nil {
		h.logger.ErrorLog.Printf("Failed to sign token: %v", err)
		h.logger.PrintDuration("GenerateToken", h.timer.GetDuration())
		return "", err
	}

	h.logger.InfoLog.Println("Token generated successfully")
	h.logger.PrintDuration("GenerateToken", h.timer.GetDuration())
	return tokenString, nil
}

// 認証インターセプター
func (h *AuthHandler) AuthInterceptor(jwtSecret string, requiredRole string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		h.timer.Start()

		// 認可スキップ対象のメソッド
		if info.FullMethod == "/pb.AuthService/Login" {
			h.logger.InfoLog.Println("Login method called")
			h.logger.PrintDuration("AuthInterceptor", h.timer.GetDuration())
			return handler(ctx, req)
		}

		// メタデータから Authorization ヘッダーを取得
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			h.logger.ErrorLog.Println("Missing metadata")
			h.logger.PrintDuration("AuthInterceptor", h.timer.GetDuration())
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		authHeaders := md["authorization"]
		if len(authHeaders) == 0 {
			h.logger.ErrorLog.Println("Missing authorization header")
			h.logger.PrintDuration("AuthInterceptor", h.timer.GetDuration())
			return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
		}

		// Bearer トークンからJWT抽出
		tokenString := strings.TrimPrefix(authHeaders[0], "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				h.logger.ErrorLog.Println("Invalid token")
				h.logger.PrintDuration("AuthInterceptor", h.timer.GetDuration())
				return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			h.logger.ErrorLog.Println("Invalid token")
			h.logger.PrintDuration("AuthInterceptor", h.timer.GetDuration())
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			h.logger.ErrorLog.Println("Invalid claims")
			h.logger.PrintDuration("AuthInterceptor", h.timer.GetDuration())
			return nil, status.Errorf(codes.Unauthenticated, "invalid claims")
		}

		role := claims["role"].(string)
		if role != requiredRole {
			h.logger.ErrorLog.Println("Permission denied")
			h.logger.PrintDuration("AuthInterceptor", h.timer.GetDuration())
			return nil, status.Errorf(codes.PermissionDenied, "permission denied")
		}

		// 必要なら ID を context に追加してハンドラーに渡す
		userID := claims["id"].(string)
		ctx = context.WithValue(ctx, h.AppConfig.UserID, userID)

		h.logger.InfoLog.Println("AuthInterceptor successful")
		h.logger.PrintDuration("AuthInterceptor", h.timer.GetDuration())
		return handler(ctx, req)
	}
}
