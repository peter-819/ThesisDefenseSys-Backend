package logic

import (
	"context"

	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryStudentsContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryStudentsContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryStudentsContentLogic {
	return QueryStudentsContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryStudentsContentLogic) QueryStudentsContent(req types.QueryStudentsContentReq) (resp *types.QueryGroupContentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
