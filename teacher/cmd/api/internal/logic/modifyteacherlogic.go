package logic

import (
	"context"

	"TDS-backend/teacher/cmd/api/internal/svc"
	"TDS-backend/teacher/cmd/api/internal/types"
	"TDS-backend/teacher/cmd/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) ModifyTeacherLogic {
	return ModifyTeacherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyTeacherLogic) ModifyTeacher(req types.ModifyTeacherReq) error {
	// todo: add your logic here and delete this line
	newTeacher, err := utils.TeacherJtoB(&req.NewTeacher)
	if err != nil {
		return err
	}
	err = l.svcCtx.TeacherModel.ModifyTeacher(req.Id, newTeacher)
	return err
}
