// Code generated by goctl. DO NOT EDIT!
// Source: student.proto

package server

import (
	"context"

	"TDS-backend/student/cmd/rpc/internal/logic"
	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/types/student"
)

type StudentServiceServer struct {
	svcCtx *svc.ServiceContext
	student.UnimplementedStudentServiceServer
}

func NewStudentServiceServer(svcCtx *svc.ServiceContext) *StudentServiceServer {
	return &StudentServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *StudentServiceServer) AddStudent(ctx context.Context, in *student.AddStudentRequest) (*student.EmptyResponse, error) {
	l := logic.NewAddStudentLogic(ctx, s.svcCtx)
	return l.AddStudent(in)
}

func (s *StudentServiceServer) RemoveStudent(ctx context.Context, in *student.RemoveStudentRequest) (*student.EmptyResponse, error) {
	l := logic.NewRemoveStudentLogic(ctx, s.svcCtx)
	return l.RemoveStudent(in)
}

func (s *StudentServiceServer) ModifyStudent(ctx context.Context, in *student.ModifyStudentRequest) (*student.EmptyResponse, error) {
	l := logic.NewModifyStudentLogic(ctx, s.svcCtx)
	return l.ModifyStudent(in)
}

func (s *StudentServiceServer) QueryAllStudents(ctx context.Context, in *student.EmptyRequest) (*student.QueryStudentsResponse, error) {
	l := logic.NewQueryAllStudentsLogic(ctx, s.svcCtx)
	return l.QueryAllStudents(in)
}

func (s *StudentServiceServer) QueryStudent(ctx context.Context, in *student.QueryStudentRequest) (*student.QueryStudentResponse, error) {
	l := logic.NewQueryStudentLogic(ctx, s.svcCtx)
	return l.QueryStudent(in)
}

func (s *StudentServiceServer) QueryNongroupedStudents(ctx context.Context, in *student.EmptyRequest) (*student.QueryStudentsResponse, error) {
	l := logic.NewQueryNongroupedStudentsLogic(ctx, s.svcCtx)
	return l.QueryNongroupedStudents(in)
}

func (s *StudentServiceServer) QueryStudentsBatch(ctx context.Context, in *student.QueryStudentsBatchRequest) (*student.QueryStudentsResponse, error) {
	l := logic.NewQueryStudentsBatchLogic(ctx, s.svcCtx)
	return l.QueryStudentsBatch(in)
}

func (s *StudentServiceServer) NewGroup(ctx context.Context, in *student.NewGroupRequest) (*student.NewGroupResponse, error) {
	l := logic.NewNewGroupLogic(ctx, s.svcCtx)
	return l.NewGroup(in)
}

func (s *StudentServiceServer) QueryGroup(ctx context.Context, in *student.QueryGroupRequest) (*student.Group, error) {
	l := logic.NewQueryGroupLogic(ctx, s.svcCtx)
	return l.QueryGroup(in)
}

func (s *StudentServiceServer) QueryAllGroups(ctx context.Context, in *student.EmptyRequest) (*student.QueryGroupsResponse, error) {
	l := logic.NewQueryAllGroupsLogic(ctx, s.svcCtx)
	return l.QueryAllGroups(in)
}

func (s *StudentServiceServer) ModifyGroup(ctx context.Context, in *student.ModifyGroupRequest) (*student.EmptyResponse, error) {
	l := logic.NewModifyGroupLogic(ctx, s.svcCtx)
	return l.ModifyGroup(in)
}

func (s *StudentServiceServer) RemoveGroup(ctx context.Context, in *student.QueryGroupRequest) (*student.EmptyResponse, error) {
	l := logic.NewRemoveGroupLogic(ctx, s.svcCtx)
	return l.RemoveGroup(in)
}

func (s *StudentServiceServer) QueryGroupContent(ctx context.Context, in *student.QueryGroupRequest) (*student.QueryGroupContentResponse, error) {
	l := logic.NewQueryGroupContentLogic(ctx, s.svcCtx)
	return l.QueryGroupContent(in)
}
