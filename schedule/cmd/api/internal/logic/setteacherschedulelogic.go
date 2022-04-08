package logic

import (
	"context"
	"time"
	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/schedule/model"
	"TDS-backend/common/errorx"
	

	"github.com/zeromicro/go-zero/core/logx"
)

type SetteacherscheduleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetteacherscheduleLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetteacherscheduleLogic {
	return SetteacherscheduleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetteacherscheduleLogic) Setteacherschedule(req types.SetTScheduleReq) (resp *types.SetTScheduleReply, err error) {
	resp = &types.SetTScheduleReply{}
	role := l.ctx.Value("role").(string)
	if role != "Teacher" {
		return nil, errorx.NewDefaultError("没有权限")
	}
	id := l.ctx.Value("id").(string)
	
	var newSchedules []model.TSchedule
	for _, s := range req.Schedules {
		start_time, _ := time.Parse("2006-01-02T15:04:05", s.StartTime)
		end_time, _ := time.Parse("2006-01-02T15:04:05", s.EndTime)
		
		newSchedules = append(newSchedules, model.TSchedule{
			Name: s.Name,
			ID: s.ID,
			StartTime: start_time,
			EndTime: end_time,
		})
	}
	
	err = l.svcCtx.ScheduleModel.ResetTeacherSchedule(id, newSchedules)

	return resp, err
}
