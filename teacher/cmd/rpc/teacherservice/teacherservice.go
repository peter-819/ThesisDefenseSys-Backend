// Code generated by goctl. DO NOT EDIT!
// Source: teacher.proto

package teacherservice

import (
	"context"

	"TDS-backend/teacher/cmd/rpc/types/teacher"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddTeacherRequest             = teacher.AddTeacherRequest
	EmptyResponse                 = teacher.EmptyResponse
	ModifyTeacherRequest          = teacher.ModifyTeacherRequest
	QueryAllTeacherRequest        = teacher.QueryAllTeacherRequest
	QueryAvailableTeachersRequest = teacher.QueryAvailableTeachersRequest
	QueryTeacherRequest           = teacher.QueryTeacherRequest
	QueryTeacherResponse          = teacher.QueryTeacherResponse
	QueryTeachersResponse         = teacher.QueryTeachersResponse
	RemoveTeacherRequest          = teacher.RemoveTeacherRequest
	Schedule                      = teacher.Schedule
	Teacher                       = teacher.Teacher

	TeacherService interface {
		AddTeacher(ctx context.Context, in *AddTeacherRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
		RemoveTeacher(ctx context.Context, in *RemoveTeacherRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
		ModifyTeacher(ctx context.Context, in *ModifyTeacherRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
		QueryAllTeacher(ctx context.Context, in *QueryAllTeacherRequest, opts ...grpc.CallOption) (*QueryTeachersResponse, error)
		QueryTeacher(ctx context.Context, in *QueryTeacherRequest, opts ...grpc.CallOption) (*QueryTeacherResponse, error)
		QueryAvailableTeachers(ctx context.Context, in *QueryAvailableTeachersRequest, opts ...grpc.CallOption) (*QueryTeachersResponse, error)
	}

	defaultTeacherService struct {
		cli zrpc.Client
	}
)

func NewTeacherService(cli zrpc.Client) TeacherService {
	return &defaultTeacherService{
		cli: cli,
	}
}

func (m *defaultTeacherService) AddTeacher(ctx context.Context, in *AddTeacherRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	client := teacher.NewTeacherServiceClient(m.cli.Conn())
	return client.AddTeacher(ctx, in, opts...)
}

func (m *defaultTeacherService) RemoveTeacher(ctx context.Context, in *RemoveTeacherRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	client := teacher.NewTeacherServiceClient(m.cli.Conn())
	return client.RemoveTeacher(ctx, in, opts...)
}

func (m *defaultTeacherService) ModifyTeacher(ctx context.Context, in *ModifyTeacherRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	client := teacher.NewTeacherServiceClient(m.cli.Conn())
	return client.ModifyTeacher(ctx, in, opts...)
}

func (m *defaultTeacherService) QueryAllTeacher(ctx context.Context, in *QueryAllTeacherRequest, opts ...grpc.CallOption) (*QueryTeachersResponse, error) {
	client := teacher.NewTeacherServiceClient(m.cli.Conn())
	return client.QueryAllTeacher(ctx, in, opts...)
}

func (m *defaultTeacherService) QueryTeacher(ctx context.Context, in *QueryTeacherRequest, opts ...grpc.CallOption) (*QueryTeacherResponse, error) {
	client := teacher.NewTeacherServiceClient(m.cli.Conn())
	return client.QueryTeacher(ctx, in, opts...)
}

func (m *defaultTeacherService) QueryAvailableTeachers(ctx context.Context, in *QueryAvailableTeachersRequest, opts ...grpc.CallOption) (*QueryTeachersResponse, error) {
	client := teacher.NewTeacherServiceClient(m.cli.Conn())
	return client.QueryAvailableTeachers(ctx, in, opts...)
}
