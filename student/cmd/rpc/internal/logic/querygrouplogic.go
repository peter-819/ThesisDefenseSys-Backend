package logic

import (
	"context"

	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/internal/utils"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGroupLogic {
	return &QueryGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryGroupLogic) QueryGroup(in *student.QueryGroupRequest) (*student.Group, error) {
	// todo: add your logic here and delete this line
	group, err := l.svcCtx.GroupModel.QueryGroup(in.GroupId)
	if err != nil {
		return nil, err
	}
	return utils.GroupBtoJ(group), nil
}
