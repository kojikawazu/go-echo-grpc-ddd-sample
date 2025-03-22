package interfaces_todo

import (
	domain_todo "backend/internal/domain/todo"
	pkg_logger "backend/internal/pkg/logger"
	pkg_timer "backend/internal/pkg/timer"
	usecase_todo "backend/internal/usecase/todo"
	pb "backend/proto/github.com/grpc/backend/proto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Todoハンドラー層
type TodoHandler struct {
	logger *pkg_logger.AppLogger
	timer  *pkg_timer.TimerPkg
	pb.UnimplementedTodoServiceServer
	todoUsecase usecase_todo.ITodoUsecase
}

// Todoハンドラー層のインスタンス化
func NewTodoHandler(l *pkg_logger.AppLogger, todoUsecase usecase_todo.ITodoUsecase) *TodoHandler {
	return &TodoHandler{logger: l, todoUsecase: todoUsecase, timer: pkg_timer.NewTimerPkg()}
}

// Todo情報を取得する
func (h *TodoHandler) GetAllTodos(ctx context.Context, req *emptypb.Empty) (*pb.TodoList, error) {
	h.logger.InfoLog.Println("GetTodo called")
	h.timer.Start()

	// Todo情報を取得する(usecase層)
	todos, err := h.todoUsecase.GetAllTodos()
	if err != nil {
		h.logger.ErrorLog.Printf("Failed to get todos: %v", err)
		h.logger.PrintDuration("GetAllTodos", h.timer.GetDuration())
		return nil, err
	}

	pbTodos := make([]*pb.Todo, len(todos))
	for i, todo := range todos {
		pbTodos[i] = &pb.Todo{
			Id:          todo.ID,
			Description: todo.Description,
			Completed:   todo.Completed,
			UserId:      todo.UserId,
			CreatedAt:   timestamppb.New(todo.CreatedAt),
			UpdatedAt:   timestamppb.New(todo.UpdatedAt),
		}
	}

	h.logger.InfoLog.Printf("GetTodo success: %v todos", len(pbTodos))
	h.logger.PrintDuration("GetAllTodos", h.timer.GetDuration())
	return &pb.TodoList{Todos: pbTodos}, nil
}

// Todoを取得する
func (h *TodoHandler) GetTodoById(ctx context.Context, req *pb.GetTodoByIdRequest) (*pb.Todo, error) {
	h.logger.InfoLog.Println("GetTodoById called")
	h.timer.Start()

	// Todoを取得する(usecase層)
	todo, err := h.todoUsecase.GetTodoById(req.Id)
	if err != nil {
		switch err.Error() {
		case "id is empty":
			h.logger.ErrorLog.Printf("Failed to get todo: %v", err)
			h.logger.PrintDuration("GetTodoById", h.timer.GetDuration())
			return nil, status.Errorf(codes.InvalidArgument, "id is empty")
		default:
			h.logger.ErrorLog.Printf("Failed to get todo: %v", err)
			h.logger.PrintDuration("GetTodoById", h.timer.GetDuration())
			return nil, err
		}
	}

	pbTodo := &pb.Todo{
		Id:          todo.ID,
		Description: todo.Description,
		Completed:   todo.Completed,
		UserId:      todo.UserId,
		CreatedAt:   timestamppb.New(todo.CreatedAt),
		UpdatedAt:   timestamppb.New(todo.UpdatedAt),
	}

	h.logger.InfoLog.Printf("GetTodoById success: %v", pbTodo)
	h.logger.PrintDuration("GetTodoById", h.timer.GetDuration())
	return pbTodo, nil
}

// 特定のユーザーのTodoを取得する
func (h *TodoHandler) GetTodoByUserId(ctx context.Context, req *pb.GetTodoByUserIdRequest) (*pb.TodoList, error) {
	h.logger.InfoLog.Println("GetTodoByUserId called")
	h.timer.Start()

	// 特定のユーザーのTodoを取得する(usecase層)
	todos, err := h.todoUsecase.GetTodoByUserId(req.UserId)
	if err != nil {
		switch err.Error() {
		case "user_id is empty":
			h.logger.ErrorLog.Printf("Failed to get todos: %v", err)
			h.logger.PrintDuration("GetTodoByUserId", h.timer.GetDuration())
			return nil, status.Errorf(codes.InvalidArgument, "user_id is empty")
		default:
			h.logger.ErrorLog.Printf("Failed to get todos: %v", err)
			h.logger.PrintDuration("GetTodoByUserId", h.timer.GetDuration())
			return nil, err
		}
	}

	pbTodos := make([]*pb.Todo, len(todos))
	for i, todo := range todos {
		pbTodos[i] = &pb.Todo{
			Id:          todo.ID,
			Description: todo.Description,
			Completed:   todo.Completed,
			UserId:      todo.UserId,
			CreatedAt:   timestamppb.New(todo.CreatedAt),
			UpdatedAt:   timestamppb.New(todo.UpdatedAt),
		}
	}

	h.logger.InfoLog.Printf("GetTodoByUserId success: %v todos", len(pbTodos))
	h.logger.PrintDuration("GetTodoByUserId", h.timer.GetDuration())
	return &pb.TodoList{Todos: pbTodos}, nil
}

// Todoを作成する
func (h *TodoHandler) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.Todo, error) {
	h.logger.InfoLog.Println("CreateTodo called")
	h.timer.Start()

	// Todoを作成する(usecase層)
	todo := domain_todo.Todo{
		Description: req.Description,
		UserId:      req.UserId,
	}
	createdTodo, err := h.todoUsecase.CreateTodo(todo)
	if err != nil {
		switch err.Error() {
		case "description is empty":
			h.logger.ErrorLog.Printf("Failed to create todo: %v", err)
			h.logger.PrintDuration("CreateTodo", h.timer.GetDuration())
			return nil, status.Errorf(codes.InvalidArgument, "description is empty")
		case "user_id is empty":
			h.logger.ErrorLog.Printf("Failed to create todo: %v", err)
			h.logger.PrintDuration("CreateTodo", h.timer.GetDuration())
			return nil, status.Errorf(codes.InvalidArgument, "user_id is empty")
		default:
			h.logger.ErrorLog.Printf("Failed to create todo: %v", err)
			h.logger.PrintDuration("CreateTodo", h.timer.GetDuration())
			return nil, err
		}
	}

	pbTodo := &pb.Todo{
		Id:          createdTodo.ID,
		Description: createdTodo.Description,
		Completed:   createdTodo.Completed,
		UserId:      createdTodo.UserId,
		CreatedAt:   timestamppb.New(createdTodo.CreatedAt),
		UpdatedAt:   timestamppb.New(createdTodo.UpdatedAt),
	}

	h.logger.InfoLog.Printf("CreateTodo success: %v", pbTodo)
	h.logger.PrintDuration("CreateTodo", h.timer.GetDuration())
	return pbTodo, nil
}

// Todoを更新する
func (h *TodoHandler) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.Todo, error) {
	h.logger.InfoLog.Println("UpdateTodo called")
	h.timer.Start()

	// Todoを更新する(usecase層)
	todo := domain_todo.Todo{
		ID:          req.Id,
		Description: req.Description,
		Completed:   req.Completed,
		UserId:      req.UserId,
	}
	updatedTodo, err := h.todoUsecase.UpdateTodo(todo)
	if err != nil {
		switch err.Error() {
		case "id is empty":
			h.logger.ErrorLog.Printf("Failed to update todo: %v", err)
			h.logger.PrintDuration("UpdateTodo", h.timer.GetDuration())
			return nil, status.Errorf(codes.InvalidArgument, "id is empty")
		case "description is empty":
			h.logger.ErrorLog.Printf("Failed to update todo: %v", err)
			h.logger.PrintDuration("UpdateTodo", h.timer.GetDuration())
			return nil, status.Errorf(codes.InvalidArgument, "description is empty")
		case "user_id is empty":
			h.logger.ErrorLog.Printf("Failed to update todo: %v", err)
			h.logger.PrintDuration("UpdateTodo", h.timer.GetDuration())
			return nil, status.Errorf(codes.InvalidArgument, "user_id is empty")
		default:
			h.logger.ErrorLog.Printf("Failed to update todo: %v", err)
			h.logger.PrintDuration("UpdateTodo", h.timer.GetDuration())
			return nil, err
		}
	}

	pbTodo := &pb.Todo{
		Id:          updatedTodo.ID,
		Description: updatedTodo.Description,
		Completed:   updatedTodo.Completed,
		UserId:      updatedTodo.UserId,
		CreatedAt:   timestamppb.New(updatedTodo.CreatedAt),
		UpdatedAt:   timestamppb.New(updatedTodo.UpdatedAt),
	}

	h.logger.InfoLog.Printf("UpdateTodo success: %v", pbTodo)
	h.logger.PrintDuration("UpdateTodo", h.timer.GetDuration())
	return pbTodo, nil
}

// Todoを削除する
func (h *TodoHandler) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*emptypb.Empty, error) {
	h.logger.InfoLog.Println("DeleteTodo called")
	h.timer.Start()

	// Todoを削除する(usecase層)
	err := h.todoUsecase.DeleteTodo(req.Id)
	if err != nil {
		switch err.Error() {
		case "id is empty":
			h.logger.ErrorLog.Printf("Failed to delete todo: %v", err)
			h.logger.PrintDuration("DeleteTodo", h.timer.GetDuration())
			return nil, status.Errorf(codes.InvalidArgument, "id is empty")
		default:
			h.logger.ErrorLog.Printf("Failed to delete todo: %v", err)
			h.logger.PrintDuration("DeleteTodo", h.timer.GetDuration())
			return nil, err
		}
	}

	h.logger.InfoLog.Println("DeleteTodo success")
	h.logger.PrintDuration("DeleteTodo", h.timer.GetDuration())
	return &emptypb.Empty{}, nil
}
