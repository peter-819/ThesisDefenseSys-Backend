package logic

import (
	"context"
	"fmt"

	"TDS-backend/common/authx"
	"TDS-backend/route/cmd/api/internal/svc"
	"TDS-backend/route/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RouteheaderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRouteheaderLogic(ctx context.Context, svcCtx *svc.ServiceContext) RouteheaderLogic {
	return RouteheaderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RouteheaderLogic) Routeheader(req types.HeaderReq) (resp *types.HeaderReply, err error) {
	// todo: add your logic here and delete this line
	resp = &types.HeaderReply{}
	fmt.Println("x-user-id: ", l.ctx.Value(authx.IdCtxKey).(string))
	fmt.Println("x-user-role: ", l.ctx.Value(authx.RoleCtxKey).(string))
	err = l.svcCtx.RouteModel.GetHeader(l.ctx.Value(authx.RoleCtxKey).(string), resp)
	if err != nil {
		return resp, err
	}
	return resp, err
}
