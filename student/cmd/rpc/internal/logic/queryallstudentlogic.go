package logic

import (
	"context"

	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllStudentLogic {
	return &QueryAllStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllStudentLogic) QueryAllStudent(in *student.EmptyRequest) (*student.QueryStudentsResponse, error) {
	// todo: add your logic here and delete this line

	return &student.QueryStudentsResponse{}, nil
}
