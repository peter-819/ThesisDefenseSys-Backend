package logic

import (
	"context"

	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/schedule/model"
	"TDS-backend/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerydefensescheduleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerydefensescheduleLogic(ctx context.Context, svcCtx *svc.ServiceContext) QuerydefensescheduleLogic {
	return QuerydefensescheduleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerydefensescheduleLogic) Querydefenseschedule(req types.GetDScheduleReq) (resp *types.GetDScheduleReply, err error) {
	// todo: add your logic here and delete this line
	resp = &types.GetDScheduleReply{}
	role := l.ctx.Value("role").(string)
	if role != "Teacher" && role != "Student" {
		return nil, errorx.NewDefaultError("没有权限")
	}
	id := l.ctx.Value("id").(string)
	
	var result []model.DSchedule

	if role == "Teacher" {
		result, err = l.svcCtx.ScheduleModel.GetTeacherDefense(id)
	} else if role == "Student"{
		result, err = l.svcCtx.ScheduleModel.GetStudentDefense(id)
	}
	
	for _, s := range result {
		resp.Schedules = append(resp.Schedules, types.DSchedule {
			Name: s.Name,
			StartTime: s.StartTime.String(),
			EndTime: s.EndTime.String(),
			Classroom: s.Classroom,
		})
	}
	return resp, err
}
