// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: position.proto

package proto

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

// PositionServiceClient is the client API for PositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionServiceClient interface {
	Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Position, error)
	GetAll(ctx context.Context, in *GetAllPositionsRequest, opts ...grpc.CallOption) (*PositionsList, error)
	Create(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Position, error)
	Update(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Position, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error)
}

type positionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionServiceClient(cc grpc.ClientConnInterface) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Position, error) {
	out := new(Position)
	err := c.cc.Invoke(ctx, "/employees_api.proto.PositionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) GetAll(ctx context.Context, in *GetAllPositionsRequest, opts ...grpc.CallOption) (*PositionsList, error) {
	out := new(PositionsList)
	err := c.cc.Invoke(ctx, "/employees_api.proto.PositionService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Create(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Position, error) {
	out := new(Position)
	err := c.cc.Invoke(ctx, "/employees_api.proto.PositionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Update(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Position, error) {
	out := new(Position)
	err := c.cc.Invoke(ctx, "/employees_api.proto.PositionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/employees_api.proto.PositionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionServiceServer is the server API for PositionService service.
// All implementations must embed UnimplementedPositionServiceServer
// for forward compatibility
type PositionServiceServer interface {
	Get(context.Context, *Id) (*Position, error)
	GetAll(context.Context, *GetAllPositionsRequest) (*PositionsList, error)
	Create(context.Context, *Position) (*Position, error)
	Update(context.Context, *Position) (*Position, error)
	Delete(context.Context, *Id) (*Status, error)
	mustEmbedUnimplementedPositionServiceServer()
}

// UnimplementedPositionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPositionServiceServer struct {
}

func (UnimplementedPositionServiceServer) Get(context.Context, *Id) (*Position, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPositionServiceServer) GetAll(context.Context, *GetAllPositionsRequest) (*PositionsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPositionServiceServer) Create(context.Context, *Position) (*Position, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPositionServiceServer) Update(context.Context, *Position) (*Position, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPositionServiceServer) Delete(context.Context, *Id) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPositionServiceServer) mustEmbedUnimplementedPositionServiceServer() {}

// UnsafePositionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionServiceServer will
// result in compilation errors.
type UnsafePositionServiceServer interface {
	mustEmbedUnimplementedPositionServiceServer()
}

func RegisterPositionServiceServer(s grpc.ServiceRegistrar, srv PositionServiceServer) {
	s.RegisterService(&PositionService_ServiceDesc, srv)
}

func _PositionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employees_api.proto.PositionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Get(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employees_api.proto.PositionService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).GetAll(ctx, req.(*GetAllPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employees_api.proto.PositionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Create(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employees_api.proto.PositionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Update(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employees_api.proto.PositionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// PositionService_ServiceDesc is the grpc.ServiceDesc for PositionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PositionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "employees_api.proto.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PositionService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PositionService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PositionService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PositionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PositionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "position.proto",
}
