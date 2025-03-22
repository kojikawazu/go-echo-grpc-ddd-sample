package interfaces_user

import (
	pkg_logger "backend/internal/pkg/logger"
	pkg_timer "backend/internal/pkg/timer"
	usecase_user "backend/internal/usecase/user"
	pb "backend/proto/github.com/grpc/backend/proto"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

// ユーザーハンドラー層
type UserHandler struct {
	logger *pkg_logger.AppLogger
	timer  *pkg_timer.TimerPkg
	pb.UnimplementedUserServiceServer
	userUsecase usecase_user.IUserUsecase
}

// ユーザーハンドラー層のインスタンス化
func NewUserHandler(l *pkg_logger.AppLogger, userUsecase usecase_user.IUserUsecase) *UserHandler {
	return &UserHandler{logger: l, userUsecase: userUsecase, timer: pkg_timer.NewTimerPkg()}
}

// ユーザー情報を取得する
func (h *UserHandler) GetAllUsers(ctx context.Context, req *emptypb.Empty) (*pb.UserList, error) {
	h.logger.InfoLog.Println("GetUser called")
	h.timer.Start()

	// ユーザー情報を取得する(usecase層)
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		h.logger.ErrorLog.Printf("Failed to get users: %v", err)
		h.logger.InfoLog.Printf("GetUser duration: %v", h.timer.GetDuration())
		return nil, err
	}

	pbUsers := make([]*pb.User, len(users))
	for i, user := range users {
		pbUsers[i] = &pb.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}
	}

	h.logger.InfoLog.Printf("GetUser success: %v users", len(pbUsers))
	h.logger.InfoLog.Printf("GetUser duration: %v", h.timer.GetDuration())
	return &pb.UserList{Users: pbUsers}, nil
}
