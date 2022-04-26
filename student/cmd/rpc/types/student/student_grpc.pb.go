// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package student

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

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	AddStudent(ctx context.Context, in *AddStudentRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	RemoveStudent(ctx context.Context, in *RemoveStudentRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	ModifyStudent(ctx context.Context, in *ModifyStudentRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	QueryAllStudents(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*QueryStudentsResponse, error)
	QueryStudent(ctx context.Context, in *QueryStudentRequest, opts ...grpc.CallOption) (*QueryStudentResponse, error)
	QueryNongroupedStudents(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*QueryStudentsResponse, error)
	QueryStudentsBatch(ctx context.Context, in *QueryStudentsBatchRequest, opts ...grpc.CallOption) (*QueryStudentsResponse, error)
	NewGroup(ctx context.Context, in *NewGroupRequest, opts ...grpc.CallOption) (*NewGroupResponse, error)
	QueryGroup(ctx context.Context, in *QueryGroupRequest, opts ...grpc.CallOption) (*Group, error)
	QueryAllGroups(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*QueryGroupsResponse, error)
	ModifyGroup(ctx context.Context, in *ModifyGroupRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	RemoveGroup(ctx context.Context, in *QueryGroupRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	QueryGroupContent(ctx context.Context, in *QueryGroupRequest, opts ...grpc.CallOption) (*QueryGroupContentResponse, error)
	QueryStudentsContent(ctx context.Context, in *QueryStudentsContentRequest, opts ...grpc.CallOption) (*QueryGroupContentResponse, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) AddStudent(ctx context.Context, in *AddStudentRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/AddStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) RemoveStudent(ctx context.Context, in *RemoveStudentRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/RemoveStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) ModifyStudent(ctx context.Context, in *ModifyStudentRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/ModifyStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) QueryAllStudents(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*QueryStudentsResponse, error) {
	out := new(QueryStudentsResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/QueryAllStudents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) QueryStudent(ctx context.Context, in *QueryStudentRequest, opts ...grpc.CallOption) (*QueryStudentResponse, error) {
	out := new(QueryStudentResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/QueryStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) QueryNongroupedStudents(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*QueryStudentsResponse, error) {
	out := new(QueryStudentsResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/QueryNongroupedStudents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) QueryStudentsBatch(ctx context.Context, in *QueryStudentsBatchRequest, opts ...grpc.CallOption) (*QueryStudentsResponse, error) {
	out := new(QueryStudentsResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/QueryStudentsBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) NewGroup(ctx context.Context, in *NewGroupRequest, opts ...grpc.CallOption) (*NewGroupResponse, error) {
	out := new(NewGroupResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/NewGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) QueryGroup(ctx context.Context, in *QueryGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/student.StudentService/QueryGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) QueryAllGroups(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*QueryGroupsResponse, error) {
	out := new(QueryGroupsResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/QueryAllGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) ModifyGroup(ctx context.Context, in *ModifyGroupRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/ModifyGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) RemoveGroup(ctx context.Context, in *QueryGroupRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/RemoveGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) QueryGroupContent(ctx context.Context, in *QueryGroupRequest, opts ...grpc.CallOption) (*QueryGroupContentResponse, error) {
	out := new(QueryGroupContentResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/QueryGroupContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) QueryStudentsContent(ctx context.Context, in *QueryStudentsContentRequest, opts ...grpc.CallOption) (*QueryGroupContentResponse, error) {
	out := new(QueryGroupContentResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/QueryStudentsContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	AddStudent(context.Context, *AddStudentRequest) (*EmptyResponse, error)
	RemoveStudent(context.Context, *RemoveStudentRequest) (*EmptyResponse, error)
	ModifyStudent(context.Context, *ModifyStudentRequest) (*EmptyResponse, error)
	QueryAllStudents(context.Context, *EmptyRequest) (*QueryStudentsResponse, error)
	QueryStudent(context.Context, *QueryStudentRequest) (*QueryStudentResponse, error)
	QueryNongroupedStudents(context.Context, *EmptyRequest) (*QueryStudentsResponse, error)
	QueryStudentsBatch(context.Context, *QueryStudentsBatchRequest) (*QueryStudentsResponse, error)
	NewGroup(context.Context, *NewGroupRequest) (*NewGroupResponse, error)
	QueryGroup(context.Context, *QueryGroupRequest) (*Group, error)
	QueryAllGroups(context.Context, *EmptyRequest) (*QueryGroupsResponse, error)
	ModifyGroup(context.Context, *ModifyGroupRequest) (*EmptyResponse, error)
	RemoveGroup(context.Context, *QueryGroupRequest) (*EmptyResponse, error)
	QueryGroupContent(context.Context, *QueryGroupRequest) (*QueryGroupContentResponse, error)
	QueryStudentsContent(context.Context, *QueryStudentsContentRequest) (*QueryGroupContentResponse, error)
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) AddStudent(context.Context, *AddStudentRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStudent not implemented")
}
func (UnimplementedStudentServiceServer) RemoveStudent(context.Context, *RemoveStudentRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveStudent not implemented")
}
func (UnimplementedStudentServiceServer) ModifyStudent(context.Context, *ModifyStudentRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyStudent not implemented")
}
func (UnimplementedStudentServiceServer) QueryAllStudents(context.Context, *EmptyRequest) (*QueryStudentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAllStudents not implemented")
}
func (UnimplementedStudentServiceServer) QueryStudent(context.Context, *QueryStudentRequest) (*QueryStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStudent not implemented")
}
func (UnimplementedStudentServiceServer) QueryNongroupedStudents(context.Context, *EmptyRequest) (*QueryStudentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryNongroupedStudents not implemented")
}
func (UnimplementedStudentServiceServer) QueryStudentsBatch(context.Context, *QueryStudentsBatchRequest) (*QueryStudentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStudentsBatch not implemented")
}
func (UnimplementedStudentServiceServer) NewGroup(context.Context, *NewGroupRequest) (*NewGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewGroup not implemented")
}
func (UnimplementedStudentServiceServer) QueryGroup(context.Context, *QueryGroupRequest) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryGroup not implemented")
}
func (UnimplementedStudentServiceServer) QueryAllGroups(context.Context, *EmptyRequest) (*QueryGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAllGroups not implemented")
}
func (UnimplementedStudentServiceServer) ModifyGroup(context.Context, *ModifyGroupRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyGroup not implemented")
}
func (UnimplementedStudentServiceServer) RemoveGroup(context.Context, *QueryGroupRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroup not implemented")
}
func (UnimplementedStudentServiceServer) QueryGroupContent(context.Context, *QueryGroupRequest) (*QueryGroupContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryGroupContent not implemented")
}
func (UnimplementedStudentServiceServer) QueryStudentsContent(context.Context, *QueryStudentsContentRequest) (*QueryGroupContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStudentsContent not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_AddStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).AddStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/AddStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).AddStudent(ctx, req.(*AddStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_RemoveStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).RemoveStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/RemoveStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).RemoveStudent(ctx, req.(*RemoveStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_ModifyStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).ModifyStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/ModifyStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).ModifyStudent(ctx, req.(*ModifyStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_QueryAllStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).QueryAllStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/QueryAllStudents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).QueryAllStudents(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_QueryStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).QueryStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/QueryStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).QueryStudent(ctx, req.(*QueryStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_QueryNongroupedStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).QueryNongroupedStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/QueryNongroupedStudents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).QueryNongroupedStudents(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_QueryStudentsBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStudentsBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).QueryStudentsBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/QueryStudentsBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).QueryStudentsBatch(ctx, req.(*QueryStudentsBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_NewGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).NewGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/NewGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).NewGroup(ctx, req.(*NewGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_QueryGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).QueryGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/QueryGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).QueryGroup(ctx, req.(*QueryGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_QueryAllGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).QueryAllGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/QueryAllGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).QueryAllGroups(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_ModifyGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).ModifyGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/ModifyGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).ModifyGroup(ctx, req.(*ModifyGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_RemoveGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).RemoveGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/RemoveGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).RemoveGroup(ctx, req.(*QueryGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_QueryGroupContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).QueryGroupContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/QueryGroupContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).QueryGroupContent(ctx, req.(*QueryGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_QueryStudentsContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStudentsContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).QueryStudentsContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/QueryStudentsContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).QueryStudentsContent(ctx, req.(*QueryStudentsContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddStudent",
			Handler:    _StudentService_AddStudent_Handler,
		},
		{
			MethodName: "RemoveStudent",
			Handler:    _StudentService_RemoveStudent_Handler,
		},
		{
			MethodName: "ModifyStudent",
			Handler:    _StudentService_ModifyStudent_Handler,
		},
		{
			MethodName: "QueryAllStudents",
			Handler:    _StudentService_QueryAllStudents_Handler,
		},
		{
			MethodName: "QueryStudent",
			Handler:    _StudentService_QueryStudent_Handler,
		},
		{
			MethodName: "QueryNongroupedStudents",
			Handler:    _StudentService_QueryNongroupedStudents_Handler,
		},
		{
			MethodName: "QueryStudentsBatch",
			Handler:    _StudentService_QueryStudentsBatch_Handler,
		},
		{
			MethodName: "NewGroup",
			Handler:    _StudentService_NewGroup_Handler,
		},
		{
			MethodName: "QueryGroup",
			Handler:    _StudentService_QueryGroup_Handler,
		},
		{
			MethodName: "QueryAllGroups",
			Handler:    _StudentService_QueryAllGroups_Handler,
		},
		{
			MethodName: "ModifyGroup",
			Handler:    _StudentService_ModifyGroup_Handler,
		},
		{
			MethodName: "RemoveGroup",
			Handler:    _StudentService_RemoveGroup_Handler,
		},
		{
			MethodName: "QueryGroupContent",
			Handler:    _StudentService_QueryGroupContent_Handler,
		},
		{
			MethodName: "QueryStudentsContent",
			Handler:    _StudentService_QueryStudentsContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student.proto",
}