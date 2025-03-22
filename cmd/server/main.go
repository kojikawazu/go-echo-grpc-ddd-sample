package main

import (
	"backend/config"
	infrastructure_auth "backend/internal/infrastructure/auth"
	infrastructure_todo "backend/internal/infrastructure/todo"
	infrastructure_user "backend/internal/infrastructure/user"
	interfaces_auth "backend/internal/interfaces/auth"
	interfaces_todo "backend/internal/interfaces/todo"
	interfaces_user "backend/internal/interfaces/user"
	pkg_logger "backend/internal/pkg/logger"
	pkg_supabase "backend/internal/pkg/supabase"
	usecase_auth "backend/internal/usecase/auth"
	usecase_todo "backend/internal/usecase/todo"
	usecase_user "backend/internal/usecase/user"
	pb "backend/proto/github.com/grpc/backend/proto"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

// main関数のセットアップ
func setUp(l *pkg_logger.AppLogger, appConfig *config.AppConfig, sc *pkg_supabase.SupabaseClient) (*grpc.Server, error) {
	// Supabaseの接続
	err := sc.InitSupabase(l)
	if err != nil {
		l.ErrorLog.Fatalf("Failed to initialize Supabase: %v", err)
	}
	// テストクエリ
	err = sc.TestQuery(l)
	if err != nil {
		l.ErrorLog.Fatalf("Failed to test query: %v", err)
	}

	// DI
	// repository層
	userRepository := infrastructure_user.NewUserRepository(l, sc)
	todoRepository := infrastructure_todo.NewTodoRepository(l, sc)
	authRepository := infrastructure_auth.NewAuthRepository(l, sc)
	// usecase層
	userUsecase := usecase_user.NewUserUsecase(l, userRepository)
	todoUsecase := usecase_todo.NewTodoUsecase(l, todoRepository)
	authUsecase := usecase_auth.NewAuthUsecase(l, authRepository)
	// handler層
	userHandler := interfaces_user.NewUserHandler(l, userUsecase)
	todoHandler := interfaces_todo.NewTodoHandler(l, todoUsecase)
	authHandler := interfaces_auth.NewAuthHandler(l, appConfig, authUsecase)

	// gRPCサーバーのインスタンス化
	server := grpc.NewServer(
		grpc.UnaryInterceptor(authHandler.AuthInterceptor(appConfig.JWTSecret, appConfig.UserRole)),
	)

	// gRPCサーバーにハンドラーを登録
	pb.RegisterUserServiceServer(server, userHandler)
	pb.RegisterTodoServiceServer(server, todoHandler)
	pb.RegisterAuthServiceServer(server, authHandler)

	return server, nil
}

// アプリケーションのメイン関数
func main() {
	// 環境変数の読み込み
	appConfig := config.NewAppConfig()
	appConfig.SetUpEnv()

	// ログ設定
	logger := pkg_logger.NewAppLogger()
	logger.SetUpLogger()

	// Supabaseの初期化
	supabaseClient := pkg_supabase.NewSupabaseClient()

	// Echoのインスタンス化
	e := echo.New()

	// gRPCサーバーの設定
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.ErrorLog.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}

	// セットアップ
	server, err := setUp(logger, appConfig, supabaseClient)
	if err != nil {
		logger.ErrorLog.Fatalf("failed to set up: %v", err)
		os.Exit(1)
	}

	// シグナルハンドラーの設定
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 終了ゴルーチン
	go func() {
		<-quit
		logger.InfoLog.Println("Shutting down server...")

		// Echoサーバーのシャットダウン
		if err := e.Close(); err != nil {
			logger.ErrorLog.Printf("Echo shutdown failed: %v", err)
		}

		// gRPCサーバーのシャットダウン
		server.GracefulStop()

		// Supabaseコネクションプールのクローズ
		supabaseClient.ClosePool(logger)
	}()

	// gRPCサーバーの起動
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}
	logger.InfoLog.Printf("Starting gRPC server on port %s...", grpcPort)
	go func() {
		if err := server.Serve(lis); err != nil {
			logger.ErrorLog.Fatalf("gRPC server failed: %v", err)
		}
	}()

	// サーバーの起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
		logger.ErrorLog.Fatalf("Echo server failed: %v", err)
	}
}
