package logic

import (
	"context"

	"TDS-backend/classroom/cmd/api/internal/svc"
	"TDS-backend/classroom/cmd/api/internal/types"
	"TDS-backend/classroom/cmd/rpc/types/classroom"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveClassroomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveClassroomLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveClassroomLogic {
	return RemoveClassroomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveClassroomLogic) RemoveClassroom(req types.RemoveClassroomReq) error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.ClassroomRpc.RemoveClassroom(l.ctx, &classroom.RemoveClassroomRequest{
		Id: req.Id,
	})
	return err
}
