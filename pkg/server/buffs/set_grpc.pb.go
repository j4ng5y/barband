// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: set.proto

package buffs

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

// SetServiceClient is the client API for SetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SetServiceClient interface {
	CreateSet(ctx context.Context, in *CreateSetRequest, opts ...grpc.CallOption) (*Set, error)
	ReadSet(ctx context.Context, in *ReadSetRequest, opts ...grpc.CallOption) (*Set, error)
	UpdateSet(ctx context.Context, in *UpdateSetRequest, opts ...grpc.CallOption) (*Set, error)
	DeleteSet(ctx context.Context, in *DeleteSetRequest, opts ...grpc.CallOption) (*Empty, error)
}

type setServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSetServiceClient(cc grpc.ClientConnInterface) SetServiceClient {
	return &setServiceClient{cc}
}

func (c *setServiceClient) CreateSet(ctx context.Context, in *CreateSetRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/barband.set.SetService/CreateSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setServiceClient) ReadSet(ctx context.Context, in *ReadSetRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/barband.set.SetService/ReadSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setServiceClient) UpdateSet(ctx context.Context, in *UpdateSetRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/barband.set.SetService/UpdateSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setServiceClient) DeleteSet(ctx context.Context, in *DeleteSetRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/barband.set.SetService/DeleteSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SetServiceServer is the server API for SetService service.
// All implementations must embed UnimplementedSetServiceServer
// for forward compatibility
type SetServiceServer interface {
	CreateSet(context.Context, *CreateSetRequest) (*Set, error)
	ReadSet(context.Context, *ReadSetRequest) (*Set, error)
	UpdateSet(context.Context, *UpdateSetRequest) (*Set, error)
	DeleteSet(context.Context, *DeleteSetRequest) (*Empty, error)
	mustEmbedUnimplementedSetServiceServer()
}

// UnimplementedSetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSetServiceServer struct {
}

func (UnimplementedSetServiceServer) CreateSet(context.Context, *CreateSetRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSet not implemented")
}
func (UnimplementedSetServiceServer) ReadSet(context.Context, *ReadSetRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadSet not implemented")
}
func (UnimplementedSetServiceServer) UpdateSet(context.Context, *UpdateSetRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSet not implemented")
}
func (UnimplementedSetServiceServer) DeleteSet(context.Context, *DeleteSetRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSet not implemented")
}
func (UnimplementedSetServiceServer) mustEmbedUnimplementedSetServiceServer() {}

// UnsafeSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SetServiceServer will
// result in compilation errors.
type UnsafeSetServiceServer interface {
	mustEmbedUnimplementedSetServiceServer()
}

func RegisterSetServiceServer(s grpc.ServiceRegistrar, srv SetServiceServer) {
	s.RegisterService(&SetService_ServiceDesc, srv)
}

func _SetService_CreateSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetServiceServer).CreateSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/barband.set.SetService/CreateSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetServiceServer).CreateSet(ctx, req.(*CreateSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SetService_ReadSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetServiceServer).ReadSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/barband.set.SetService/ReadSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetServiceServer).ReadSet(ctx, req.(*ReadSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SetService_UpdateSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetServiceServer).UpdateSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/barband.set.SetService/UpdateSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetServiceServer).UpdateSet(ctx, req.(*UpdateSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SetService_DeleteSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetServiceServer).DeleteSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/barband.set.SetService/DeleteSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetServiceServer).DeleteSet(ctx, req.(*DeleteSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SetService_ServiceDesc is the grpc.ServiceDesc for SetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "barband.set.SetService",
	HandlerType: (*SetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSet",
			Handler:    _SetService_CreateSet_Handler,
		},
		{
			MethodName: "ReadSet",
			Handler:    _SetService_ReadSet_Handler,
		},
		{
			MethodName: "UpdateSet",
			Handler:    _SetService_UpdateSet_Handler,
		},
		{
			MethodName: "DeleteSet",
			Handler:    _SetService_DeleteSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "set.proto",
}
