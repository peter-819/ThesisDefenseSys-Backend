package logic

import (
	"context"

	"TDS-backend/user/cmd/rpc/internal/svc"
	"TDS-backend/user/cmd/rpc/types/user"
	"TDS-backend/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTeacherInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetTeacherInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTeacherInfoLogic {
	return &SetTeacherInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetTeacherInfoLogic) SetTeacherInfo(in *user.SetTeacherInfoRequest) (*user.SetTeacherInfoResponse, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.TeacherModel.SetTeacherInfo(in.Id, model.Teacher{
		ID: in.Info.Id,
		Name: in.Info.Name,
		IsSecretary: in.Info.IsSecretary,
	})
	return &user.SetTeacherInfoResponse{}, err
}
