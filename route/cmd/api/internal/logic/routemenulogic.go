package logic

import (
	"context"

	"TDS-backend/route/cmd/api/internal/svc"
	"TDS-backend/route/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoutemenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoutemenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoutemenuLogic {
	return RoutemenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoutemenuLogic) Routemenu() (resp *types.RouteMenuReply, err error) {
	// todo: add your logic here and delete this line
	resp = &types.RouteMenuReply{}
	err = l.svcCtx.RouteModel.GetMenu(resp)
	
	return resp, err
}
