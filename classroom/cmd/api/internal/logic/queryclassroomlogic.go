package logic

import (
	"context"

	"TDS-backend/classroom/cmd/api/internal/svc"
	"TDS-backend/classroom/cmd/api/internal/types"
	"TDS-backend/classroom/cmd/rpc/classroomservice"
	"TDS-backend/common/typex"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryClassroomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryClassroomLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryClassroomLogic {
	return QueryClassroomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryClassroomLogic) QueryClassroom(req types.QueryClassroomReq) (resp *types.Classroom, err error) {
	rpcRes, err := l.svcCtx.ClassroomRpc.QueryClassroom(l.ctx, &classroomservice.QueryClassroomRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.Classroom{}
	err = typex.Convert(rpcRes, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
