package logic

import (
	"context"

	"TDS-backend/classroom/cmd/rpc/internal/svc"
	"TDS-backend/classroom/cmd/rpc/internal/utils"
	"TDS-backend/classroom/cmd/rpc/types/classroom"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryClassroomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryClassroomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryClassroomLogic {
	return &QueryClassroomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryClassroomLogic) QueryClassroom(in *classroom.QueryClassroomRequest) (*classroom.Classroom, error) {
	// todo: add your logic here and delete this line
	classroom, err := l.svcCtx.ClassroomModel.FindClassroomByIdString(in.Id)
	if err != nil {
		return nil, err
	}
	out := utils.ClassroomBtoJ(classroom)
	return out, nil
}
