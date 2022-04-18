package logic

import (
	"context"

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
	err = l.svcCtx.RouteModel.GetHeader(l.ctx.Value("role").(string), resp)
	if err != nil {
		return resp, err
	}
	if l.ctx.Value("is_secretary").(bool) && l.ctx.Value("role").(string) == "Teacher" {
		secretaryResp := &types.HeaderReply{}
		err = l.svcCtx.RouteModel.GetHeader("Secretary", secretaryResp)
		if err != nil {
			return resp, err
		}
		resp.Headers = append(resp.Headers, secretaryResp.Headers...)
	}
	return resp,err
}
