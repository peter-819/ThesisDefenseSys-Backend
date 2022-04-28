package logic

import (
	"context"

	"TDS-backend/common/typex"
	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryStudentsBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryStudentsBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryStudentsBatchLogic {
	return &QueryStudentsBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryStudentsBatchLogic) QueryStudentsBatch(in *student.QueryStudentsBatchRequest) (*student.QueryStudentsResponse, error) {
	students, err := l.svcCtx.StudentModel.QueryStudentsBatch(in.Ids)
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
