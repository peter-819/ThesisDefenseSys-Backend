package logic

import (
	"context"

	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/internal/utils"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyGroupLogic {
	return &ModifyGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyGroupLogic) ModifyGroup(in *student.ModifyGroupRequest) (*student.EmptyResponse, error) {
	// todo: add your logic here and delete this line
	new_group, err := utils.GroupJtoB(in.NewGroup)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.GroupModel.ModifyGroup(in.Id, new_group)
	if err != nil {
		return nil, err
	}
	return &student.EmptyResponse{}, nil
}
