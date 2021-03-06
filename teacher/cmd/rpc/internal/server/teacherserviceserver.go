// Code generated by goctl. DO NOT EDIT!
// Source: teacher.proto

package server

import (
	"context"

	"TDS-backend/teacher/cmd/rpc/internal/logic"
	"TDS-backend/teacher/cmd/rpc/internal/svc"
	"TDS-backend/teacher/cmd/rpc/types/teacher"
)

type TeacherServiceServer struct {
	svcCtx *svc.ServiceContext
	teacher.UnimplementedTeacherServiceServer
}

func NewTeacherServiceServer(svcCtx *svc.ServiceContext) *TeacherServiceServer {
	return &TeacherServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *TeacherServiceServer) AddTeacher(ctx context.Context, in *teacher.AddTeacherRequest) (*teacher.EmptyResponse, error) {
	l := logic.NewAddTeacherLogic(ctx, s.svcCtx)
	return l.AddTeacher(in)
}

func (s *TeacherServiceServer) RemoveTeacher(ctx context.Context, in *teacher.RemoveTeacherRequest) (*teacher.EmptyResponse, error) {
	l := logic.NewRemoveTeacherLogic(ctx, s.svcCtx)
	return l.RemoveTeacher(in)
}

func (s *TeacherServiceServer) ModifyTeacher(ctx context.Context, in *teacher.ModifyTeacherRequest) (*teacher.EmptyResponse, error) {
	l := logic.NewModifyTeacherLogic(ctx, s.svcCtx)
	return l.ModifyTeacher(in)
}

func (s *TeacherServiceServer) QueryAllTeacher(ctx context.Context, in *teacher.QueryAllTeacherRequest) (*teacher.QueryTeachersResponse, error) {
	l := logic.NewQueryAllTeacherLogic(ctx, s.svcCtx)
	return l.QueryAllTeacher(in)
}

func (s *TeacherServiceServer) QueryTeacher(ctx context.Context, in *teacher.QueryTeacherRequest) (*teacher.QueryTeacherResponse, error) {
	l := logic.NewQueryTeacherLogic(ctx, s.svcCtx)
	return l.QueryTeacher(in)
}

func (s *TeacherServiceServer) QueryAvailableTeachers(ctx context.Context, in *teacher.QueryAvailableTeachersRequest) (*teacher.QueryTeachersResponse, error) {
	l := logic.NewQueryAvailableTeachersLogic(ctx, s.svcCtx)
	return l.QueryAvailableTeachers(in)
}
