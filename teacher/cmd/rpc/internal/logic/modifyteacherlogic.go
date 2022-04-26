package logic

import (
	"context"

	"TDS-backend/teacher/cmd/rpc/internal/svc"
	"TDS-backend/teacher/cmd/rpc/types/teacher"
	"TDS-backend/teacher/cmd/rpc/internal/utils"
	// "TDS-backend/teacher/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyTeacherLogic {
	return &ModifyTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyTeacherLogic) ModifyTeacher(in *teacher.ModifyTeacherRequest) (*teacher.EmptyResponse, error) {
	// todo: add your logic here and delete this line
	t, err := utils.TeacherJtoB(in.NewTeacher)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.TeacherModel.ModifyTeacher(in.Id, t)
	return &teacher.EmptyResponse{}, err
}