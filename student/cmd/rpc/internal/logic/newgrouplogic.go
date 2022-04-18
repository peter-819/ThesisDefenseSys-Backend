package logic

import (
	"context"

	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNewGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewGroupLogic {
	return &NewGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NewGroupLogic) NewGroup(in *student.NewGroupRequest) (*student.NewGroupResponse, error) {
	// todo: add your logic here and delete this line
	groupId, err := l.svcCtx.GroupModel.NewGroup(in.Name)
	if err != nil {
		return nil, err
	}
	return &student.NewGroupResponse{
		Id: groupId,
	}, nil
}
