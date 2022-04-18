package logic

import (
	"context"

	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/types/student"

	"TDS-backend/student/cmd/rpc/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryNongroupedStudentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryNongroupedStudentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryNongroupedStudentsLogic {
	return &QueryNongroupedStudentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryNongroupedStudentsLogic) QueryNongroupedStudents(in *student.EmptyRequest) (*student.QueryStudentsResponse, error) {
	// todo: add your logic here and delete this line
	students, err := l.svcCtx.StudentModel.QueryNongroupedStudents()
	if err != nil {
		return nil, err
	}
	resp := &student.QueryStudentsResponse{}
	for _, s := range students {
		resp.List = append(resp.List, utils.StudentBtoJ(&s))
	}
	return resp, nil
}
