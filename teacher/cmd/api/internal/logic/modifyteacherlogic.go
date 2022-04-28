package logic

import (
	"context"

	"TDS-backend/common/typex"
	"TDS-backend/teacher/cmd/api/internal/svc"
	"TDS-backend/teacher/cmd/api/internal/types"
	"TDS-backend/teacher/cmd/rpc/teacherservice"

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
	newTeacher := &teacherservice.Teacher{}
	typex.Convert(req.NewTeacher, newTeacher)
	_, err := l.svcCtx.TeacherRpc.ModifyTeacher(l.ctx, &teacherservice.ModifyTeacherRequest{
		Id:         req.Id,
		NewTeacher: newTeacher,
	})
	return err
}
