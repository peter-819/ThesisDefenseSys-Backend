package logic

import (
	"context"

	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/internal/utils"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryStudentLogic {
	return &QueryStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryStudentLogic) QueryStudent(in *student.QueryStudentRequest) (*student.QueryStudentResponse, error) {
	// todo: add your logic here and delete this line
	s, err := l.svcCtx.StudentModel.QueryStudent(in.Id)
	if err != nil {
		return nil, err
	}
	out := &student.QueryStudentResponse{}
	out.Student = utils.StudentBtoJ(s)
	return out, nil
}
