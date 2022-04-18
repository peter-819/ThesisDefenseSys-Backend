package logic

import (
	"context"

	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGroupContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryGroupContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryGroupContentLogic {
	return QueryGroupContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryGroupContentLogic) QueryGroupContent(req types.QueryGroupReq) (resp *types.QueryGroupContentResp, err error) {
	// todo: add your logic here and delete this line
	rpcRes, err := l.svcCtx.StudentRpc.QueryGroupContent(l.ctx, &student.QueryGroupRequest{
		GroupId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.QueryGroupContentResp{
		Keywords: rpcRes.Keywords,
		Mentors:  rpcRes.Mentors,
	}, nil
}
