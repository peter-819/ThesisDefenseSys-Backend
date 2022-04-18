package logic

import (
	"context"

	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/internal/utils"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStudentLogic {
	return &AddStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddStudentLogic) AddStudent(in *student.AddStudentRequest) (*student.EmptyResponse, error) {
	// todo: add your logic here and delete this line
	t, err := utils.StudentJtoB(in.Student)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.StudentModel.AddStudent(t)
	return &student.EmptyResponse{}, err
}
