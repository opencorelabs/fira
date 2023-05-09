// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: protos/fira/v1/api.proto

package v1

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

const (
	FiraService_GetApiInfo_FullMethodName        = "/protos.fira.v1.FiraService/GetApiInfo"
	FiraService_CreateLinkSession_FullMethodName = "/protos.fira.v1.FiraService/CreateLinkSession"
	FiraService_GetLinkSession_FullMethodName    = "/protos.fira.v1.FiraService/GetLinkSession"
)

// FiraServiceClient is the client API for FiraService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FiraServiceClient interface {
	GetApiInfo(ctx context.Context, in *GetApiInfoRequest, opts ...grpc.CallOption) (*GetApiInfoResponse, error)
	// Create a new link session. This will return a URL to redirect the user to where they will be able to select a
	// financial institution to connect to and log in.
	CreateLinkSession(ctx context.Context, in *CreateLinkSessionRequest, opts ...grpc.CallOption) (*CreateLinkSessionResponse, error)
	// Retrieve the status of a link session.
	GetLinkSession(ctx context.Context, in *GetLinkSessionRequest, opts ...grpc.CallOption) (*GetLinkSessionResponse, error)
}

type firaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFiraServiceClient(cc grpc.ClientConnInterface) FiraServiceClient {
	return &firaServiceClient{cc}
}

func (c *firaServiceClient) GetApiInfo(ctx context.Context, in *GetApiInfoRequest, opts ...grpc.CallOption) (*GetApiInfoResponse, error) {
	out := new(GetApiInfoResponse)
	err := c.cc.Invoke(ctx, FiraService_GetApiInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firaServiceClient) CreateLinkSession(ctx context.Context, in *CreateLinkSessionRequest, opts ...grpc.CallOption) (*CreateLinkSessionResponse, error) {
	out := new(CreateLinkSessionResponse)
	err := c.cc.Invoke(ctx, FiraService_CreateLinkSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firaServiceClient) GetLinkSession(ctx context.Context, in *GetLinkSessionRequest, opts ...grpc.CallOption) (*GetLinkSessionResponse, error) {
	out := new(GetLinkSessionResponse)
	err := c.cc.Invoke(ctx, FiraService_GetLinkSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FiraServiceServer is the server API for FiraService service.
// All implementations must embed UnimplementedFiraServiceServer
// for forward compatibility
type FiraServiceServer interface {
	GetApiInfo(context.Context, *GetApiInfoRequest) (*GetApiInfoResponse, error)
	// Create a new link session. This will return a URL to redirect the user to where they will be able to select a
	// financial institution to connect to and log in.
	CreateLinkSession(context.Context, *CreateLinkSessionRequest) (*CreateLinkSessionResponse, error)
	// Retrieve the status of a link session.
	GetLinkSession(context.Context, *GetLinkSessionRequest) (*GetLinkSessionResponse, error)
	mustEmbedUnimplementedFiraServiceServer()
}

// UnimplementedFiraServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFiraServiceServer struct {
}

func (UnimplementedFiraServiceServer) GetApiInfo(context.Context, *GetApiInfoRequest) (*GetApiInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiInfo not implemented")
}
func (UnimplementedFiraServiceServer) CreateLinkSession(context.Context, *CreateLinkSessionRequest) (*CreateLinkSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLinkSession not implemented")
}
func (UnimplementedFiraServiceServer) GetLinkSession(context.Context, *GetLinkSessionRequest) (*GetLinkSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkSession not implemented")
}
func (UnimplementedFiraServiceServer) mustEmbedUnimplementedFiraServiceServer() {}

// UnsafeFiraServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FiraServiceServer will
// result in compilation errors.
type UnsafeFiraServiceServer interface {
	mustEmbedUnimplementedFiraServiceServer()
}

func RegisterFiraServiceServer(s grpc.ServiceRegistrar, srv FiraServiceServer) {
	s.RegisterService(&FiraService_ServiceDesc, srv)
}

func _FiraService_GetApiInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FiraServiceServer).GetApiInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FiraService_GetApiInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FiraServiceServer).GetApiInfo(ctx, req.(*GetApiInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FiraService_CreateLinkSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLinkSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FiraServiceServer).CreateLinkSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FiraService_CreateLinkSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FiraServiceServer).CreateLinkSession(ctx, req.(*CreateLinkSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FiraService_GetLinkSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FiraServiceServer).GetLinkSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FiraService_GetLinkSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FiraServiceServer).GetLinkSession(ctx, req.(*GetLinkSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FiraService_ServiceDesc is the grpc.ServiceDesc for FiraService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FiraService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.fira.v1.FiraService",
	HandlerType: (*FiraServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApiInfo",
			Handler:    _FiraService_GetApiInfo_Handler,
		},
		{
			MethodName: "CreateLinkSession",
			Handler:    _FiraService_CreateLinkSession_Handler,
		},
		{
			MethodName: "GetLinkSession",
			Handler:    _FiraService_GetLinkSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/fira/v1/api.proto",
}
