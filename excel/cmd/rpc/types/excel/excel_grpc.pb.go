// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package excel

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	ParseExcel(ctx context.Context, opts ...grpc.CallOption) (User_ParseExcelClient, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) ParseExcel(ctx context.Context, opts ...grpc.CallOption) (User_ParseExcelClient, error) {
	stream, err := c.cc.NewStream(ctx, &User_ServiceDesc.Streams[0], "/excel.User/ParseExcel", opts...)
	if err != nil {
		return nil, err
	}
	x := &userParseExcelClient{stream}
	return x, nil
}

type User_ParseExcelClient interface {
	Send(*ExcelRequest) error
	CloseAndRecv() (*ExcelResponse, error)
	grpc.ClientStream
}

type userParseExcelClient struct {
	grpc.ClientStream
}

func (x *userParseExcelClient) Send(m *ExcelRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userParseExcelClient) CloseAndRecv() (*ExcelResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ExcelResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	ParseExcel(User_ParseExcelServer) error
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) ParseExcel(User_ParseExcelServer) error {
	return status.Errorf(codes.Unimplemented, "method ParseExcel not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_ParseExcel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServer).ParseExcel(&userParseExcelServer{stream})
}

type User_ParseExcelServer interface {
	SendAndClose(*ExcelResponse) error
	Recv() (*ExcelRequest, error)
	grpc.ServerStream
}

type userParseExcelServer struct {
	grpc.ServerStream
}

func (x *userParseExcelServer) SendAndClose(m *ExcelResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userParseExcelServer) Recv() (*ExcelRequest, error) {
	m := new(ExcelRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "excel.User",
	HandlerType: (*UserServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ParseExcel",
			Handler:       _User_ParseExcel_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "excel.proto",
}