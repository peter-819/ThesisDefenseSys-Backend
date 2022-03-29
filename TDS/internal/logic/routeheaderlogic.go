package logic

import (
	"context"
	"fmt"
	
	"TDS-backend/TDS/internal/svc"
	"TDS-backend/TDS/internal/types"
	"TDS-backend/TDS/internal/models"

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
	// userInfo, err := models.QueryUserInfo(l.svcCtx.Database, req.Token)
	// if err != nil {
	// 	fmt.Println("unknown user")
	// 	return nil,err
	// }
	header, err := models.QueryRouteHeader(l.svcCtx.Database, "teacher")
	if err != nil {
		fmt.Println("query header failed")
	}
	return header,err
}
