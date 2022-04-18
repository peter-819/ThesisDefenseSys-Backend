package logic

import (
	"context"
	"TDS-backend/common/errorx"

	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/schedule/cmd/api/internal/utils"
	"TDS-backend/schedule/model"

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
	role := l.ctx.Value("role").(string)
	bSecretary := l.ctx.Value("is_secretary").(bool)
	if role != "Teacher" || !bSecretary {
		return nil, errorx.NewDefaultError("没有权限")
	}
	var schedules []model.Defense

	schedules, err = l.svcCtx.DefenseModel.GetAllDefense()
	for _, s := range schedules {
		resp.Defenses = append(resp.Defenses, *utils.DefenseBtoJ(&s))
	}
	
	return resp, err
}
