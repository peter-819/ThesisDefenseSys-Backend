package logic

import (
	"context"

	"TDS-backend/common/typex"
	"TDS-backend/schedule/cmd/rpc/internal/svc"
	"TDS-backend/schedule/cmd/rpc/types/defense"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllDefensesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllDefensesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllDefensesLogic {
	return &QueryAllDefensesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllDefensesLogic) QueryAllDefenses(in *defense.EmptyRequest) (*defense.QueryAllDefensesReply, error) {
	// todo: add your logic here and delete this line
	defenses, err := l.svcCtx.DefenseModel.GetAllDefense()
	if err != nil {
		return nil, err
	}
	outDefenses := []*defense.Defense{}
	err = typex.Convert(&defenses, &outDefenses)
	return &defense.QueryAllDefensesReply{
		Defenses: outDefenses,
	}, err
}
