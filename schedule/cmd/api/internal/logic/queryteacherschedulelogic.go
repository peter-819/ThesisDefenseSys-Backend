package logic

import (
	"context"

	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryteacherscheduleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryteacherscheduleLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryteacherscheduleLogic {
	return QueryteacherscheduleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryteacherscheduleLogic) Queryteacherschedule(req types.GetTScheduleReq) (resp *types.GetTScheduleReply, err error) {
	// todo: add your logic here and delete this line
	resp = &types.GetTScheduleReply{}
	role := l.ctx.Value("role").(string)
	if role != "Teacher" && role != "Secretary" {
		return nil, errorx.NewDefaultError("没有权限")
	}
	id := l.ctx.Value("id").(string)
	
	schedule, err := l.svcCtx.ScheduleModel.GetTeacherSchedule(id)
	for _, s := range schedule {
		resp.Schedules = append(resp.Schedules, types.TSchedule{
			ID: s.ID,
			Name: s.Name,
			StartTime: s.StartTime.String(),
			EndTime: s.EndTime.String(), 
		})
	}
	return resp, err
}
