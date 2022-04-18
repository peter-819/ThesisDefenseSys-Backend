package logic

import (
	"context"

	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) NewGroupLogic {
	return NewGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewGroupLogic) NewGroup(req types.NewGroupReq) (resp *types.NewGroupResp, err error) {
	// todo: add your logic here and delete this line
	rpcRes, err := l.svcCtx.StudentRpc.NewGroup(l.ctx, &student.NewGroupRequest{
		Name: req.Name,
	})

	return &types.NewGroupResp{
		GroupId: rpcRes.Id,
	}, err
}
