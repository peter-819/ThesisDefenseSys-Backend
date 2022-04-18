package logic

import (
	"context"

	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryGroupLogic {
	return QueryGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryGroupLogic) QueryGroup(req types.QueryGroupReq) (resp *types.Group, err error) {
	// todo: add your logic here and delete this line
	group, err := l.svcCtx.StudentRpc.QueryGroup(l.ctx, &student.QueryGroupRequest{
		GroupId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.Group{
		Id:        group.Id,
		Name:      group.Name,
		Members:   group.Members,
		DefenseId: group.DefenseId,
	}, nil
}
