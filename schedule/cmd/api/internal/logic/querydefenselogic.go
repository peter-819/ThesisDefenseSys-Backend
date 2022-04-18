package logic

import (
	"context"

	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDefenseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDefenseLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryDefenseLogic {
	return QueryDefenseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDefenseLogic) QueryDefense(req types.QueryDefenseReq) (resp *types.QueryDefenseReply, err error) {
	// todo: add your logic here and delete this line

	return
}
