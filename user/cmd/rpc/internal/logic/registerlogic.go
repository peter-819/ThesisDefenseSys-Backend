package logic

import (
	"context"

	"TDS-backend/user/cmd/rpc/internal/svc"
	"TDS-backend/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
	"TDS-backend/common/errorx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	err := l.svcCtx.UserModel.AddUser(in.Id, in.Name, in.Role, in.Password)
	if err != nil {
		return nil, err
	}
	u, err := l.svcCtx.UserModel.GetUser(in.Id)
	if err != nil {
		return nil, errorx.NewDefaultError("创建失败")
	}
	return &user.RegisterResponse{
		Id: u.ID,
	}, nil
}
