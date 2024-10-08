// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: budgets_service.proto

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

// BudgetsServiceClient is the client API for BudgetsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BudgetsServiceClient interface {
	Create(ctx context.Context, in *CreateBudget, opts ...grpc.CallOption) (*Budget, error)
	GetById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Budget, error)
	GetAll(ctx context.Context, in *BudgetFilter, opts ...grpc.CallOption) (*Budgets, error)
	Update(ctx context.Context, in *Budget, opts ...grpc.CallOption) (*Budget, error)
	Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error)
}

type budgetsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBudgetsServiceClient(cc grpc.ClientConnInterface) BudgetsServiceClient {
	return &budgetsServiceClient{cc}
}

func (c *budgetsServiceClient) Create(ctx context.Context, in *CreateBudget, opts ...grpc.CallOption) (*Budget, error) {
	out := new(Budget)
	err := c.cc.Invoke(ctx, "/budgeting.BudgetsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetsServiceClient) GetById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Budget, error) {
	out := new(Budget)
	err := c.cc.Invoke(ctx, "/budgeting.BudgetsService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetsServiceClient) GetAll(ctx context.Context, in *BudgetFilter, opts ...grpc.CallOption) (*Budgets, error) {
	out := new(Budgets)
	err := c.cc.Invoke(ctx, "/budgeting.BudgetsService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetsServiceClient) Update(ctx context.Context, in *Budget, opts ...grpc.CallOption) (*Budget, error) {
	out := new(Budget)
	err := c.cc.Invoke(ctx, "/budgeting.BudgetsService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetsServiceClient) Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/budgeting.BudgetsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BudgetsServiceServer is the server API for BudgetsService service.
// All implementations must embed UnimplementedBudgetsServiceServer
// for forward compatibility
type BudgetsServiceServer interface {
	Create(context.Context, *CreateBudget) (*Budget, error)
	GetById(context.Context, *PrimaryKey) (*Budget, error)
	GetAll(context.Context, *BudgetFilter) (*Budgets, error)
	Update(context.Context, *Budget) (*Budget, error)
	Delete(context.Context, *PrimaryKey) (*Void, error)
	mustEmbedUnimplementedBudgetsServiceServer()
}

// UnimplementedBudgetsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBudgetsServiceServer struct {
}

func (UnimplementedBudgetsServiceServer) Create(context.Context, *CreateBudget) (*Budget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBudgetsServiceServer) GetById(context.Context, *PrimaryKey) (*Budget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedBudgetsServiceServer) GetAll(context.Context, *BudgetFilter) (*Budgets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedBudgetsServiceServer) Update(context.Context, *Budget) (*Budget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBudgetsServiceServer) Delete(context.Context, *PrimaryKey) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBudgetsServiceServer) mustEmbedUnimplementedBudgetsServiceServer() {}

// UnsafeBudgetsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BudgetsServiceServer will
// result in compilation errors.
type UnsafeBudgetsServiceServer interface {
	mustEmbedUnimplementedBudgetsServiceServer()
}

func RegisterBudgetsServiceServer(s grpc.ServiceRegistrar, srv BudgetsServiceServer) {
	s.RegisterService(&BudgetsService_ServiceDesc, srv)
}

func _BudgetsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBudget)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting.BudgetsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetsServiceServer).Create(ctx, req.(*CreateBudget))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetsService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetsServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting.BudgetsService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetsServiceServer).GetById(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetsService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BudgetFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetsServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting.BudgetsService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetsServiceServer).GetAll(ctx, req.(*BudgetFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Budget)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting.BudgetsService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetsServiceServer).Update(ctx, req.(*Budget))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting.BudgetsService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetsServiceServer).Delete(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// BudgetsService_ServiceDesc is the grpc.ServiceDesc for BudgetsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BudgetsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "budgeting.BudgetsService",
	HandlerType: (*BudgetsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BudgetsService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _BudgetsService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BudgetsService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BudgetsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BudgetsService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "budgets_service.proto",
}
