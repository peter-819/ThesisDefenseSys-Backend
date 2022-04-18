package utils

import (
	"TDS-backend/common/errorx"
	"time"
	"TDS-backend/teacher/cmd/api/internal/types"
	"TDS-backend/teacher/model"
)

func ScheduleJtoB(in *types.Schedule) (out *model.Schedule, err error){
	start_time, err := time.Parse("2006-01-02T15:04:05", in.StartTime)
	if err != nil {
		return nil, errorx.NewDefaultError("教师日程开始时间格式错误: " + err.Error())
	}
	end_time, err := time.Parse("2006-01-02T15:04:05", in.EndTime)
	if err != nil {
		return nil, errorx.NewDefaultError("教师日程结束时间格式错误: " + err.Error())
	}	

	out = &model.Schedule {
		ScheduleId: in.ScheduleId,
		Name: in.Name,
		StartTime: start_time,
		EndTime: end_time,
	}
	return out, nil
}

func ScheduleBtoJ(in *model.Schedule) (out *types.Schedule){
	out = &types.Schedule {
		ScheduleId: in.ScheduleId,
		Name: in.Name,
		StartTime: in.StartTime.String(),
		EndTime: in.EndTime.String(),
	}
	return out
}

func TeacherJtoB(in *types.Teacher) (out *model.Teacher, err error) {
	schedules := []model.Schedule{}
	for _, s := range in.Schedules {
		schedule, err := ScheduleJtoB(&s)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, *schedule)
	}
	out = &model.Teacher{
		Id: in.Id,
		Name: in.Name,
		IsSecretary: in.IsSecretary,
		MaxDefensePerWeek: in.MaxDefensePerWeek,
		PreferKeywords: in.PreferKeywords,
		Schedules:schedules,
	}
	return out, nil
}

func TeacherBtoJ(in *model.Teacher) (out *types.Teacher) {
	schedules := []types.Schedule{}
	for _, s := range in.Schedules {
		schedule := ScheduleBtoJ(&s)
		schedules = append(schedules, *schedule)
	}
	out = &types.Teacher{
		Id: in.Id,
		Name: in.Name,
		IsSecretary: in.IsSecretary,
		MaxDefensePerWeek: in.MaxDefensePerWeek,
		PreferKeywords: in.PreferKeywords,
		Schedules:schedules,
	}
	return out
}