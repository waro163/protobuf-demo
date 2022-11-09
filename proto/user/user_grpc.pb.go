// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: proto/user/user.proto

package user

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

// GetUserClient is the client API for GetUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetUserClient interface {
	GetOneUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error)
	GetBulkUsers(ctx context.Context, opts ...grpc.CallOption) (GetUser_GetBulkUsersClient, error)
}

type getUserClient struct {
	cc grpc.ClientConnInterface
}

func NewGetUserClient(cc grpc.ClientConnInterface) GetUserClient {
	return &getUserClient{cc}
}

func (c *getUserClient) GetOneUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.user.GetUser/GetOneUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getUserClient) GetBulkUsers(ctx context.Context, opts ...grpc.CallOption) (GetUser_GetBulkUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &GetUser_ServiceDesc.Streams[0], "/proto.user.GetUser/GetBulkUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &getUserGetBulkUsersClient{stream}
	return x, nil
}

type GetUser_GetBulkUsersClient interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ClientStream
}

type getUserGetBulkUsersClient struct {
	grpc.ClientStream
}

func (x *getUserGetBulkUsersClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *getUserGetBulkUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GetUserServer is the server API for GetUser service.
// All implementations must embed UnimplementedGetUserServer
// for forward compatibility
type GetUserServer interface {
	GetOneUser(context.Context, *UserID) (*User, error)
	GetBulkUsers(GetUser_GetBulkUsersServer) error
	mustEmbedUnimplementedGetUserServer()
}

// UnimplementedGetUserServer must be embedded to have forward compatible implementations.
type UnimplementedGetUserServer struct {
}

func (UnimplementedGetUserServer) GetOneUser(context.Context, *UserID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneUser not implemented")
}
func (UnimplementedGetUserServer) GetBulkUsers(GetUser_GetBulkUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBulkUsers not implemented")
}
func (UnimplementedGetUserServer) mustEmbedUnimplementedGetUserServer() {}

// UnsafeGetUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetUserServer will
// result in compilation errors.
type UnsafeGetUserServer interface {
	mustEmbedUnimplementedGetUserServer()
}

func RegisterGetUserServer(s grpc.ServiceRegistrar, srv GetUserServer) {
	s.RegisterService(&GetUser_ServiceDesc, srv)
}

func _GetUser_GetOneUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetUserServer).GetOneUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.user.GetUser/GetOneUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetUserServer).GetOneUser(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetUser_GetBulkUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GetUserServer).GetBulkUsers(&getUserGetBulkUsersServer{stream})
}

type GetUser_GetBulkUsersServer interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ServerStream
}

type getUserGetBulkUsersServer struct {
	grpc.ServerStream
}

func (x *getUserGetBulkUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *getUserGetBulkUsersServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GetUser_ServiceDesc is the grpc.ServiceDesc for GetUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.user.GetUser",
	HandlerType: (*GetUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOneUser",
			Handler:    _GetUser_GetOneUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBulkUsers",
			Handler:       _GetUser_GetBulkUsers_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/user/user.proto",
}