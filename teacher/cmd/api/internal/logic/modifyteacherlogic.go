package logic

import (
	"context"

	"TDS-backend/common/errorx"
	"TDS-backend/common/timex"
	"TDS-backend/common/typex"
	"TDS-backend/teacher/cmd/api/internal/svc"
	"TDS-backend/teacher/cmd/api/internal/types"
	"TDS-backend/teacher/cmd/rpc/teacherservice"
	"TDS-backend/teacher/cmd/rpc/types/teacher"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) ModifyTeacherLogic {
	return ModifyTeacherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyTeacherLogic) ModifyTeacher(req types.ModifyTeacherReq) error {
	// 已安排答辩的时间无法删除修改
	teaRes, err := l.svcCtx.TeacherRpc.QueryTeacher(l.ctx, &teacher.QueryTeacherRequest{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	for _, defense := range teaRes.Teacher.Defenses {
		bInc := false
		for _, schedule := range req.NewTeacher.Schedules {
			inc, err := timex.IsIncludedString(schedule.StartTime, schedule.EndTime, defense.StartTime, defense.EndTime)
			if err != nil {
				return err
			}
			bInc = bInc || inc
		}
		if !bInc {
			return errorx.NewDefaultError("已安排答辩的日期无法修改")
		}
	}

	newTeacher := new(teacherservice.Teacher)
	typex.Convert(req.NewTeacher, newTeacher)
	_, err = l.svcCtx.TeacherRpc.ModifyTeacher(l.ctx, &teacherservice.ModifyTeacherRequest{
		Id:         req.Id,
		NewTeacher: newTeacher,
	})
	return err
}
