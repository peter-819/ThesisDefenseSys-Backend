package logic

import (
	"context"

	"TDS-backend/common/typex"
	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllStudentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllStudentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllStudentsLogic {
	return &QueryAllStudentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllStudentsLogic) QueryAllStudents(in *student.EmptyRequest) (*student.QueryStudentsResponse, error) {
	students, err := l.svcCtx.StudentModel.QueryAllStudents()
	if err != nil {
		return nil, err
	}
	resp := &student.QueryStudentsResponse{}
	for _, s := range students {
		converted := &student.Student{}
		typex.Convert(s, converted)
		resp.List = append(resp.List, converted)
	}
	return resp, nil
}
