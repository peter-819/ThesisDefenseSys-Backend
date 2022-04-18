package logic

import (
	"context"

	"TDS-backend/user/cmd/rpc/internal/svc"
	"TDS-backend/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherLogic {
	return &GetTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTeacherLogic) GetTeacher(in *user.GetTeacherByIdRequest) (*user.TeacherInfo, error) {
	// todo: add your logic here and delete this line
	teacher, err := l.svcCtx.TeacherModel.GetTeacher(in.Id)
	return &user.TeacherInfo{
		Id: teacher.ID,
		Name: teacher.Name,
		IsSecretary: teacher.IsSecretary,
	}, err
}
