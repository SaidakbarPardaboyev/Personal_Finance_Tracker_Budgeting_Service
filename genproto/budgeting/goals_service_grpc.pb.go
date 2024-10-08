// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: goals_service.proto

package budgeting

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GoalsServiceClient is the client API for GoalsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoalsServiceClient interface {
	Create(ctx context.Context, in *CreateGoal, opts ...grpc.CallOption) (*Goal, error)
	GetById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Goal, error)
	GetAll(ctx context.Context, in *GoalFilter, opts ...grpc.CallOption) (*Goals, error)
	Update(ctx context.Context, in *Goal, opts ...grpc.CallOption) (*Goal, error)
	Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error)
}

type goalsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoalsServiceClient(cc grpc.ClientConnInterface) GoalsServiceClient {
	return &goalsServiceClient{cc}
}

func (c *goalsServiceClient) Create(ctx context.Context, in *CreateGoal, opts ...grpc.CallOption) (*Goal, error) {
	out := new(Goal)
	err := c.cc.Invoke(ctx, "/budgeting.GoalsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goalsServiceClient) GetById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Goal, error) {
	out := new(Goal)
	err := c.cc.Invoke(ctx, "/budgeting.GoalsService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goalsServiceClient) GetAll(ctx context.Context, in *GoalFilter, opts ...grpc.CallOption) (*Goals, error) {
	out := new(Goals)
	err := c.cc.Invoke(ctx, "/budgeting.GoalsService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goalsServiceClient) Update(ctx context.Context, in *Goal, opts ...grpc.CallOption) (*Goal, error) {
	out := new(Goal)
	err := c.cc.Invoke(ctx, "/budgeting.GoalsService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goalsServiceClient) Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/budgeting.GoalsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoalsServiceServer is the server API for GoalsService service.
// All implementations must embed UnimplementedGoalsServiceServer
// for forward compatibility
type GoalsServiceServer interface {
	Create(context.Context, *CreateGoal) (*Goal, error)
	GetById(context.Context, *PrimaryKey) (*Goal, error)
	GetAll(context.Context, *GoalFilter) (*Goals, error)
	Update(context.Context, *Goal) (*Goal, error)
	Delete(context.Context, *PrimaryKey) (*Void, error)
	mustEmbedUnimplementedGoalsServiceServer()
}

// UnimplementedGoalsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoalsServiceServer struct {
}

func (UnimplementedGoalsServiceServer) Create(context.Context, *CreateGoal) (*Goal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGoalsServiceServer) GetById(context.Context, *PrimaryKey) (*Goal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedGoalsServiceServer) GetAll(context.Context, *GoalFilter) (*Goals, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedGoalsServiceServer) Update(context.Context, *Goal) (*Goal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGoalsServiceServer) Delete(context.Context, *PrimaryKey) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGoalsServiceServer) mustEmbedUnimplementedGoalsServiceServer() {}

// UnsafeGoalsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoalsServiceServer will
// result in compilation errors.
type UnsafeGoalsServiceServer interface {
	mustEmbedUnimplementedGoalsServiceServer()
}

func RegisterGoalsServiceServer(s grpc.ServiceRegistrar, srv GoalsServiceServer) {
	s.RegisterService(&GoalsService_ServiceDesc, srv)
}

func _GoalsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting.GoalsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalsServiceServer).Create(ctx, req.(*CreateGoal))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoalsService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalsServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting.GoalsService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalsServiceServer).GetById(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoalsService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoalFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalsServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting.GoalsService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalsServiceServer).GetAll(ctx, req.(*GoalFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoalsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Goal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting.GoalsService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalsServiceServer).Update(ctx, req.(*Goal))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoalsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting.GoalsService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalsServiceServer).Delete(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// GoalsService_ServiceDesc is the grpc.ServiceDesc for GoalsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoalsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "budgeting.GoalsService",
	HandlerType: (*GoalsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GoalsService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _GoalsService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _GoalsService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GoalsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GoalsService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goals_service.proto",
}
