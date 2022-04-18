package logic

import (
	"context"

	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveGroupLogic {
	return &RemoveGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveGroupLogic) RemoveGroup(in *student.QueryGroupRequest) (*student.EmptyResponse, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.GroupModel.RemoveGroup(in.GroupId)
	return &student.EmptyResponse{}, err
}
