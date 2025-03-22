// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: internal/interfaces/todo/todo.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TodoService_GetAllTodos_FullMethodName     = "/pb.TodoService/GetAllTodos"
	TodoService_GetTodoById_FullMethodName     = "/pb.TodoService/GetTodoById"
	TodoService_GetTodoByUserId_FullMethodName = "/pb.TodoService/GetTodoByUserId"
	TodoService_CreateTodo_FullMethodName      = "/pb.TodoService/CreateTodo"
	TodoService_UpdateTodo_FullMethodName      = "/pb.TodoService/UpdateTodo"
	TodoService_DeleteTodo_FullMethodName      = "/pb.TodoService/DeleteTodo"
)

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	GetAllTodos(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TodoList, error)
	GetTodoById(ctx context.Context, in *GetTodoByIdRequest, opts ...grpc.CallOption) (*Todo, error)
	GetTodoByUserId(ctx context.Context, in *GetTodoByUserIdRequest, opts ...grpc.CallOption) (*TodoList, error)
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*Todo, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*Todo, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) GetAllTodos(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TodoList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodoList)
	err := c.cc.Invoke(ctx, TodoService_GetAllTodos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodoById(ctx context.Context, in *GetTodoByIdRequest, opts ...grpc.CallOption) (*Todo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Todo)
	err := c.cc.Invoke(ctx, TodoService_GetTodoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodoByUserId(ctx context.Context, in *GetTodoByUserIdRequest, opts ...grpc.CallOption) (*TodoList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodoList)
	err := c.cc.Invoke(ctx, TodoService_GetTodoByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*Todo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Todo)
	err := c.cc.Invoke(ctx, TodoService_CreateTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*Todo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Todo)
	err := c.cc.Invoke(ctx, TodoService_UpdateTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TodoService_DeleteTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility.
type TodoServiceServer interface {
	GetAllTodos(context.Context, *emptypb.Empty) (*TodoList, error)
	GetTodoById(context.Context, *GetTodoByIdRequest) (*Todo, error)
	GetTodoByUserId(context.Context, *GetTodoByUserIdRequest) (*TodoList, error)
	CreateTodo(context.Context, *CreateTodoRequest) (*Todo, error)
	UpdateTodo(context.Context, *UpdateTodoRequest) (*Todo, error)
	DeleteTodo(context.Context, *DeleteTodoRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTodoServiceServer struct{}

func (UnimplementedTodoServiceServer) GetAllTodos(context.Context, *emptypb.Empty) (*TodoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTodos not implemented")
}
func (UnimplementedTodoServiceServer) GetTodoById(context.Context, *GetTodoByIdRequest) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoById not implemented")
}
func (UnimplementedTodoServiceServer) GetTodoByUserId(context.Context, *GetTodoByUserIdRequest) (*TodoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoByUserId not implemented")
}
func (UnimplementedTodoServiceServer) CreateTodo(context.Context, *CreateTodoRequest) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodoServiceServer) UpdateTodo(context.Context, *UpdateTodoRequest) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (UnimplementedTodoServiceServer) DeleteTodo(context.Context, *DeleteTodoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}
func (UnimplementedTodoServiceServer) testEmbeddedByValue()                     {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	// If the following call pancis, it indicates UnimplementedTodoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_GetAllTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetAllTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_GetAllTodos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetAllTodos(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_GetTodoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodoById(ctx, req.(*GetTodoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodoByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodoByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_GetTodoByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodoByUserId(ctx, req.(*GetTodoByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_CreateTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_UpdateTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodo(ctx, req.(*UpdateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_DeleteTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllTodos",
			Handler:    _TodoService_GetAllTodos_Handler,
		},
		{
			MethodName: "GetTodoById",
			Handler:    _TodoService_GetTodoById_Handler,
		},
		{
			MethodName: "GetTodoByUserId",
			Handler:    _TodoService_GetTodoByUserId_Handler,
		},
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoService_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/interfaces/todo/todo.proto",
}
