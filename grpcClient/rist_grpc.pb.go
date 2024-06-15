// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: rist.proto

package grpcClient

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

// RistAppClient is the client API for RistApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RistAppClient interface {
	StartRistApp(ctx context.Context, in *RistAppconfig, opts ...grpc.CallOption) (*AppStatusResponse, error)
	StopRistApp(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AppStatusResponse, error)
	StatusRistApp(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AppStatusResponse, error)
	UpdateRistApp(ctx context.Context, in *RistAppconfig, opts ...grpc.CallOption) (*AppStatusResponse, error)
}

type ristAppClient struct {
	cc grpc.ClientConnInterface
}

func NewRistAppClient(cc grpc.ClientConnInterface) RistAppClient {
	return &ristAppClient{cc}
}

func (c *ristAppClient) StartRistApp(ctx context.Context, in *RistAppconfig, opts ...grpc.CallOption) (*AppStatusResponse, error) {
	out := new(AppStatusResponse)
	err := c.cc.Invoke(ctx, "/RistApp.RistApp/StartRistApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ristAppClient) StopRistApp(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AppStatusResponse, error) {
	out := new(AppStatusResponse)
	err := c.cc.Invoke(ctx, "/RistApp.RistApp/StopRistApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ristAppClient) StatusRistApp(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AppStatusResponse, error) {
	out := new(AppStatusResponse)
	err := c.cc.Invoke(ctx, "/RistApp.RistApp/StatusRistApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ristAppClient) UpdateRistApp(ctx context.Context, in *RistAppconfig, opts ...grpc.CallOption) (*AppStatusResponse, error) {
	out := new(AppStatusResponse)
	err := c.cc.Invoke(ctx, "/RistApp.RistApp/UpdateRistApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RistAppServer is the server API for RistApp service.
// All implementations must embed UnimplementedRistAppServer
// for forward compatibility
type RistAppServer interface {
	StartRistApp(context.Context, *RistAppconfig) (*AppStatusResponse, error)
	StopRistApp(context.Context, *Empty) (*AppStatusResponse, error)
	StatusRistApp(context.Context, *Empty) (*AppStatusResponse, error)
	UpdateRistApp(context.Context, *RistAppconfig) (*AppStatusResponse, error)
	mustEmbedUnimplementedRistAppServer()
}

// UnimplementedRistAppServer must be embedded to have forward compatible implementations.
type UnimplementedRistAppServer struct {
}

func (UnimplementedRistAppServer) StartRistApp(context.Context, *RistAppconfig) (*AppStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRistApp not implemented")
}
func (UnimplementedRistAppServer) StopRistApp(context.Context, *Empty) (*AppStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopRistApp not implemented")
}
func (UnimplementedRistAppServer) StatusRistApp(context.Context, *Empty) (*AppStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusRistApp not implemented")
}
func (UnimplementedRistAppServer) UpdateRistApp(context.Context, *RistAppconfig) (*AppStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRistApp not implemented")
}
func (UnimplementedRistAppServer) mustEmbedUnimplementedRistAppServer() {}

// UnsafeRistAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RistAppServer will
// result in compilation errors.
type UnsafeRistAppServer interface {
	mustEmbedUnimplementedRistAppServer()
}

func RegisterRistAppServer(s grpc.ServiceRegistrar, srv RistAppServer) {
	s.RegisterService(&RistApp_ServiceDesc, srv)
}

func _RistApp_StartRistApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RistAppconfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RistAppServer).StartRistApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RistApp.RistApp/StartRistApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RistAppServer).StartRistApp(ctx, req.(*RistAppconfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _RistApp_StopRistApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RistAppServer).StopRistApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RistApp.RistApp/StopRistApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RistAppServer).StopRistApp(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RistApp_StatusRistApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RistAppServer).StatusRistApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RistApp.RistApp/StatusRistApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RistAppServer).StatusRistApp(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RistApp_UpdateRistApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RistAppconfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RistAppServer).UpdateRistApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RistApp.RistApp/UpdateRistApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RistAppServer).UpdateRistApp(ctx, req.(*RistAppconfig))
	}
	return interceptor(ctx, in, info, handler)
}

// RistApp_ServiceDesc is the grpc.ServiceDesc for RistApp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RistApp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RistApp.RistApp",
	HandlerType: (*RistAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartRistApp",
			Handler:    _RistApp_StartRistApp_Handler,
		},
		{
			MethodName: "StopRistApp",
			Handler:    _RistApp_StopRistApp_Handler,
		},
		{
			MethodName: "StatusRistApp",
			Handler:    _RistApp_StatusRistApp_Handler,
		},
		{
			MethodName: "UpdateRistApp",
			Handler:    _RistApp_UpdateRistApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rist.proto",
}
