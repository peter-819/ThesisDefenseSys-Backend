package logic

import (
	"context"

	"TDS-backend/common/errorx"
	"TDS-backend/user/cmd/rpc/internal/svc"
	"TDS-backend/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetSecretaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetSecretaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSecretaryLogic {
	return &SetSecretaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetSecretaryLogic) SetSecretary(in *user.SetSecretaryRequest) (*user.EmptyResponse, error) {
	// todo: add your logic here and delete this line
	u, err := l.svcCtx.UserModel.GetUser(in.Id)
	if err != nil {
		return nil, err
	}
	if u.Role != "Teacher" && u.Role != "Secretary" {
		return nil, errorx.NewDefaultError("该用户不是教师")
	}
	if in.IsSecretary {
		u.Role = "Secretary"
	} else {
		u.Role = "Teacher"
	}
	err = l.svcCtx.UserModel.ModifyUser(in.Id, u)
	if err != nil {
		return nil, err
	}
	return &user.EmptyResponse{}, nil
}
