package logic

import (
	"context"

	"TDS-backend/classroom/cmd/api/internal/svc"
	"TDS-backend/classroom/cmd/api/internal/types"
	"TDS-backend/classroom/cmd/rpc/classroomservice"
	"TDS-backend/common/typex"

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
	resp = &types.QueryClassroomsReply{}
	rpcRes, err := l.svcCtx.ClassroomRpc.QueryAllClassrooms(l.ctx, &classroomservice.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	err = typex.Convert(&rpcRes.Classrooms, &resp.Classrooms)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
