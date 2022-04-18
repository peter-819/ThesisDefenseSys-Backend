package logic

import (
	"context"

	"TDS-backend/teacher/cmd/rpc/internal/svc"
	"TDS-backend/teacher/cmd/rpc/types/teacher"
	"TDS-backend/teacher/cmd/rpc/internal/utils"
	// "TDS-backend/teacher/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTeacherLogic {
	return &AddTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTeacherLogic) AddTeacher(in *teacher.AddTeacherRequest) (*teacher.EmptyResponse, error) {
	// todo: add your logic here and delete this line
	t, err := utils.TeacherJtoB(in.Teacher)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.TeacherModel.AddTeacher(t)
	return &teacher.EmptyResponse{}, err
}
