package logic

import (
	"context"

	"TDS-backend/classroom/cmd/rpc/internal/svc"
	"TDS-backend/classroom/cmd/rpc/types/classroom"
	"TDS-backend/common/typex"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllClassroomsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllClassroomsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllClassroomsLogic {
	return &QueryAllClassroomsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllClassroomsLogic) QueryAllClassrooms(in *classroom.EmptyRequest) (*classroom.QueryClassroomsResponse, error) {
	// todo: add your logic here and delete this line
	out := &classroom.QueryClassroomsResponse{}
	classrooms, err := l.svcCtx.ClassroomModel.QueryAllClassrooms()
	if err != nil {
		return nil, err
	}
	err = typex.Convert(&classrooms, &out.Classrooms)
	if err != nil {
		return nil, err
	}
	return out, nil
}
