package logic

import (
	"context"

	"TDS-backend/classroom/cmd/rpc/internal/svc"
	"TDS-backend/classroom/cmd/rpc/types/classroom"
	"TDS-backend/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveClassroomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveClassroomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveClassroomLogic {
	return &RemoveClassroomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveClassroomLogic) RemoveClassroom(in *classroom.RemoveClassroomRequest) (*classroom.EmptyResponse, error) {
	// todo: add your logic here and delete this line
	cla, err := l.svcCtx.ClassroomModel.FindClassroomByIdString(in.Id)
	if err != nil {
		return nil, errorx.NewDefaultError("找不到教室")
	}
	if len(cla.Defenses) != 0 {
		return nil, errorx.NewDefaultError("请先移除使用该教室的答辩")
	}
	err = l.svcCtx.ClassroomModel.RemoveClassroom(in.Id)
	return &classroom.EmptyResponse{}, err
}
