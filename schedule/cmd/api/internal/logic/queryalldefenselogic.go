package logic

import (
	"TDS-backend/common/typex"
	"context"

	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllDefenseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAllDefenseLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryAllDefenseLogic {
	return QueryAllDefenseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAllDefenseLogic) QueryAllDefense() (resp *types.QueryDefenseReply, err error) {
	resp = &types.QueryDefenseReply{}

	schedules, err := l.svcCtx.DefenseModel.GetAllDefense()
	typex.Convert(&schedules, &resp.Defenses)
	return resp, err
}
