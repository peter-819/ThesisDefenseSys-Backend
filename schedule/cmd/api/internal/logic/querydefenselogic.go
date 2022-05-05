package logic

import (
	"context"

	"TDS-backend/common/typex"
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

func (l *QueryDefenseLogic) QueryDefense(req types.QueryDefenseReq) (resp *types.Defense, err error) {
	// todo: add your logic here and delete this line
	mdefense, err := l.svcCtx.DefenseModel.QueryDefense(req.Id)
	if err != nil {
		return nil, err
	}
	defense := &types.Defense{}
	err = typex.Convert(mdefense, defense)
	if err != nil {
		return nil, err
	}
	return defense, nil
}
