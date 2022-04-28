package logic

import (
	"context"

	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryNondefensedStudentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryNondefensedStudentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryNondefensedStudentsLogic {
	return &QueryNondefensedStudentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryNondefensedStudentsLogic) QueryNondefensedStudents(in *student.EmptyRequest) (*student.QueryStudentsResponse, error) {
	// todo: add your logic here and delete this line

	return &student.QueryStudentsResponse{}, nil
}
