package logic

import (
	"context"

	"TDS-backend/teacher/cmd/rpc/internal/svc"
	"TDS-backend/teacher/cmd/rpc/types/teacher"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveTeacherLogic {
	return &RemoveTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveTeacherLogic) RemoveTeacher(in *teacher.RemoveTeacherRequest) (*teacher.EmptyResponse, error) {
	// todo: add your logic here and delete this line

	return &teacher.EmptyResponse{}, nil
}
