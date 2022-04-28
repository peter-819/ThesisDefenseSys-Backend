package logic

import (
	"context"

	"TDS-backend/classroom/cmd/rpc/internal/svc"
	"TDS-backend/classroom/cmd/rpc/types/classroom"
	"TDS-backend/common/errorx"
	"TDS-backend/common/timex"
	"TDS-backend/common/typex"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAvailableByTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAvailableByTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAvailableByTimeLogic {
	return &QueryAvailableByTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAvailableByTimeLogic) QueryAvailableByTime(in *classroom.QueryByTimeRequest) (*classroom.QueryClassroomsResponse, error) {
	start_time, err := timex.GStringToTime(in.StartTime)
	if err != nil {
		return nil, errorx.NewDefaultError("教室开始时间格式错误: " + err.Error())
	}
	end_time, err := timex.GStringToTime(in.EndTime)
	if err != nil {
		return nil, errorx.NewDefaultError("教室结束时间格式错误: " + err.Error())
	}
	out := &classroom.QueryClassroomsResponse{}
	classrooms, err := l.svcCtx.ClassroomModel.QueryAvailableByTime(start_time, end_time)
	if err != nil {
		return nil, err
	}
	typex.Convert(&classrooms, &out.Classrooms)
	return out, nil
}
