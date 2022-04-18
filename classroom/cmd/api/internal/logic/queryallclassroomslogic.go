package logic

import (
	"context"

	"TDS-backend/classroom/cmd/api/internal/svc"
	"TDS-backend/classroom/cmd/api/internal/types"
	"TDS-backend/classroom/cmd/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllClassroomsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAllClassroomsLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryAllClassroomsLogic {
	return QueryAllClassroomsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAllClassroomsLogic) QueryAllClassrooms() (resp *types.QueryClassroomsReply, err error) {
	// todo: add your logic here and delete this line
	resp = &types.QueryClassroomsReply{}
	classrooms, err := l.svcCtx.ClassroomModel.QueryAllClassrooms()
	if err != nil {
		return nil, err
	}
	for _, c := range classrooms {
		resp.Classrooms = append(resp.Classrooms, *utils.ClassroomBtoJ(&c))
	}
	return resp, nil
}
