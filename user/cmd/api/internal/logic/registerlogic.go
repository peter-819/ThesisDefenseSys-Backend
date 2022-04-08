package logic

import (
	"context"

	"TDS-backend/user/cmd/api/internal/svc"
	"TDS-backend/user/cmd/api/internal/types"

	"TDS-backend/user/cmd/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (resp *types.RegisterReply, err error) {
	_, err = l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Id:req.UserID,
		Name: req.Name,
		Role: req.Role,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	tokenRes, err := l.svcCtx.UserRpc.GetToken(l.ctx, &user.TokenRequest{
		Id: req.UserID,
		Role: req.Role,
		ExpireTime: l.svcCtx.Config.Auth.AccessExpire,
		SecretKey: l.svcCtx.Config.Auth.AccessSecret,
	})
	return &types.RegisterReply{
		Token: tokenRes.Token,
	}, err
}
