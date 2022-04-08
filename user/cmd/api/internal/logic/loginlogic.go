package logic

import (
	"context"
	"fmt"

	"TDS-backend/user/cmd/api/internal/svc"
	"TDS-backend/user/cmd/api/internal/types"
	"TDS-backend/user/cmd/rpc/user"
	"TDS-backend/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
        Id: req.UserID,
    })
	if err != nil {
		fmt.Println("API失败",err)
		return nil, err
	}
	if userInfo.Password != req.Password {
		return nil, errorx.NewDefaultError("密码错误")
	}
	tokenRes, err := l.svcCtx.UserRpc.GetToken(l.ctx, &user.TokenRequest{
		Id: userInfo.Id,
		Role: userInfo.Role,
		ExpireTime: l.svcCtx.Config.Auth.AccessExpire,
		SecretKey: l.svcCtx.Config.Auth.AccessSecret,
	})
	return &types.LoginReply{
		Token: tokenRes.Token,
		Role: userInfo.Role,
	}, err
}
