package logic

import (
	"context"

	"TDS-backend/TDS/internal/svc"
	"TDS-backend/TDS/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"TDS-backend/TDS/internal/models"
	"fmt"
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
	// data := serverdata.GetServerData()
	fmt.Println(l.ctx.Value("username"))
	routes, err := models.QueryRoutesMenu(l.svcCtx.Database)
	return routes,nil
}
