package logic

import (
	"context"

	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveDefenseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveDefenseLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveDefenseLogic {
	return RemoveDefenseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveDefenseLogic) RemoveDefense(req types.RemoveDefenseReq) error {
	// todo: add your logic here and delete this line
	defense, err := l.svcCtx.DefenseModel.QueryDefense(req.ID)
	if err != nil {
		return err
	}
	groupRes, err := l.svcCtx.StudentRpc.QueryGroup(l.ctx, &student.QueryGroupRequest{
		GroupId: defense.GroupId,
	})
	if err != nil {
		return err
	}
	groupRes.DefenseId = ""
	_, err = l.svcCtx.StudentRpc.ModifyGroup(l.ctx, &student.ModifyGroupRequest{
		Id:       defense.GroupId,
		NewGroup: groupRes,
	})
	if err != nil {
		return err
	}
	err = l.svcCtx.DefenseModel.RemoveDefense(req.ID)
	return err
}
