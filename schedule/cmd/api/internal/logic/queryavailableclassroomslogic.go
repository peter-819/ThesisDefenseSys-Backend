package logic

import (
	"context"

	"TDS-backend/common/errorx"
	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/classroom/cmd/rpc/types/classroom"
	"TDS-backend/common/timex"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAvailableClassroomsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAvailableClassroomsLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryAvailableClassroomsLogic {
	return QueryAvailableClassroomsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAvailableClassroomsLogic) QueryAvailableClassrooms(req types.QueryAvailableClassroomsReq) (resp *types.QueryAvailableClassroomsResp, err error) {
	// todo: add your logic here and delete this line
	defense, err := l.svcCtx.DefenseModel.QueryDefense(req.DefenseId)
	if err != nil {
		return nil, err
	}
	classroomsRes, err := l.svcCtx.ClassroomRpc.QueryAvailableByTime(l.ctx, &classroom.QueryByTimeRequest{
		StartTime: timex.TimeToString(defense.StartTime),
		EndTime: timex.TimeToString(defense.EndTime),
	})
	if err != nil {
		return nil, err
	}
	if len(classroomsRes.Classrooms) == 0 {
		return nil, errorx.NewDefaultError("无可用教室")
	}
	resp = &types.QueryAvailableClassroomsResp{}
	for _, c := range classroomsRes.Classrooms {
		resp.Classrooms = append(resp.Classrooms, types.ClassroomInfo{
			Name: c.Location + " " + c.Name,
			Id: c.Id,
		})
	}
	return resp, nil
}
