package logic

import (
	"context"

	"TDS-backend/user/cmd/api/internal/svc"
	"TDS-backend/user/cmd/api/internal/types"
	"TDS-backend/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.QueryUserReq) (resp *types.UserInfo, err error) {
	// todo: add your logic here and delete this line
	u, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfo{
		Id:   u.Id,
		Name: u.Name,
		Role: u.Role,
	}, nil
}
