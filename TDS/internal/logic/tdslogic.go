package logic

import (
	"context"

	"TDS-backend/TDS/internal/svc"
	"TDS-backend/TDS/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TDSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTDSLogic(ctx context.Context, svcCtx *svc.ServiceContext) TDSLogic {
	return TDSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TDSLogic) TDS(req types.Request) (resp *types.Response, err error) {
    return &types.Response{
        Message: "Hello go-zero",
    }, nil
}
