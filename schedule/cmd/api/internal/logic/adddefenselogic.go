package logic

import (
	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/schedule/model"
	"TDS-backend/student/cmd/rpc/studentservice"
	"TDS-backend/student/cmd/rpc/types/student"
	"context"
	"fmt"

	"TDS-backend/classroom/cmd/rpc/classroomservice"
	"TDS-backend/classroom/cmd/rpc/types/classroom"
	"TDS-backend/teacher/cmd/rpc/teacherservice"
	"TDS-backend/teacher/cmd/rpc/types/teacher"

	"TDS-backend/common/errorx"
	"TDS-backend/common/timex"
	"TDS-backend/common/typex"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDefenseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDefenseLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddDefenseLogic {
	return AddDefenseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDefenseLogic) AddDefense(req types.Defense) (resp *types.AddDefenseReply, err error) {
	if len(req.Students) == 0 {
		return nil, errorx.NewDefaultError("学生不可为空")
	}
	if req.StartTime == "" || req.EndTime == "" {
		return nil, errorx.NewDefaultError("时间段不可为空")
	}
	_, err = timex.GStringToTime(req.StartTime)
	if err != nil {
		return nil, errorx.NewDefaultError("开始时间错误")
	}
	_, err = timex.GStringToTime(req.EndTime)
	if err != nil {
		return nil, errorx.NewDefaultError("结束时间错误")
	}

	var cla *classroom.Classroom
	var teas []*teacher.Teacher
	var stus []*student.Student
	//classroom check
	classroomId := req.Classroom
	if classroomId != "" {
		claRes, err := l.svcCtx.ClassroomRpc.QueryClassroom(l.ctx, &classroom.QueryClassroomRequest{
			Id: classroomId,
		})
		if err != nil {
			return nil, errorx.NewDefaultError("非法教室ID: " + err.Error())
		}
		fmt.Println("claRes: ", claRes)
		inc, err := timex.IsIncludedString(claRes.StartTime, claRes.EndTime, req.StartTime, req.EndTime)
		if err != nil {
			return nil, err
		}
		if !inc {
			return nil, errorx.NewDefaultError("教室不可用")
		}
		for _, d := range claRes.Defenses {
			inter, err := timex.IsIntersectedString(req.StartTime, req.EndTime, d.StartTime, d.EndTime)
			if err != nil {
				return nil, err
			}
			if inter {
				return nil, errorx.NewDefaultError("教室该时段有答辩安排")
			}
		}

		cla = claRes
		req.ClassroomFullName = claRes.Location + " " + claRes.Name
	}
	// committee check
	if len(req.Committee) != 0 {
		for index, c := range req.Committee {
			teacherId := c.TeacherID
			teaRes, err := l.svcCtx.TeacherRpc.QueryTeacher(l.ctx, &teacher.QueryTeacherRequest{
				Id: teacherId,
			})
			if err != nil {
				return nil, errorx.NewDefaultError("非法教师ID: " + err.Error())
			}
			fmt.Println("teaRes: ", teaRes)
			isAble := false
			for _, schedule := range teaRes.Teacher.Schedules {
				inc, err := timex.IsIncludedString(schedule.StartTime, schedule.EndTime, req.StartTime, req.EndTime)
				if err != nil {
					return nil, err
				}
				isAble = isAble || inc
			}
			if !isAble {
				return nil, errorx.NewDefaultError("教师无可用时间")
			}
			timesThisWeek := 0
			for _, d := range teaRes.Teacher.Defenses {
				bSameWeek, err := timex.IsSameWeekString(d.StartTime, req.StartTime)
				if err != nil {
					return nil, err
				}
				if bSameWeek {
					timesThisWeek++
				}
				inter, err := timex.IsIntersectedString(d.StartTime, d.EndTime, req.StartTime, req.EndTime)
				if err != nil {
					return nil, err
				}
				if inter {
					return nil, errorx.NewDefaultError(("教师答辩时间冲突"))
				}
			}
			fmt.Println("defense count in this week: ", timesThisWeek+1)
			if teaRes.Teacher.MaxDefensePerWeek != 0 && timesThisWeek+1 > int(teaRes.Teacher.MaxDefensePerWeek) {
				return nil, errorx.NewDefaultError(teaRes.Teacher.Name + "老师太忙啦")
			}
			teas = append(teas, teaRes.Teacher)
			req.Committee[index].TeacherName = teaRes.Teacher.Name
		}
	}
	// student check
	if len(req.Students) != 0 {
		for index, s := range req.Students {
			studentId := s.StudentID
			stuRes, err := l.svcCtx.StudentRpc.QueryStudent(l.ctx, &student.QueryStudentRequest{
				Id: studentId,
			})
			if err != nil {
				return nil, errorx.NewDefaultError("非法学生ID: " + err.Error())
			}
			fmt.Println("stuRes: ", stuRes)
			if stuRes.Student.DefenseId != "" {
				return nil, errorx.NewDefaultError("学生已有答辩安排")
			}
			stus = append(stus, stuRes.Student)
			req.Students[index].StudentName = stuRes.Student.Name
		}
	}
	defense := &model.Defense{}
	err = typex.Convert(req, defense)
	if err != nil {
		return nil, err
	}
	defenseId, err := l.svcCtx.DefenseModel.InsertDefense(defense)
	if err != nil {
		return nil, err
	}

	//modify classroom
	cla.Defenses = append(cla.Defenses, &classroom.Defense{
		DefenseId: defenseId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	_, err = l.svcCtx.ClassroomRpc.ModifyClassroom(l.ctx, &classroomservice.ModifyClassroomRequest{
		Id:           cla.Id,
		NewClassroom: cla,
	})
	if err != nil {
		return nil, err
	}
	//modify teachers
	for _, t := range teas {
		t.Defenses = append(t.Defenses, &teacher.Defense{
			DefenseId: defenseId,
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
		})
		_, err = l.svcCtx.TeacherRpc.ModifyTeacher(l.ctx, &teacherservice.ModifyTeacherRequest{
			Id:         t.Id,
			NewTeacher: t,
		})
		if err != nil {
			return nil, err
		}
	}
	//modify students
	for _, s := range stus {
		s.DefenseId = defenseId
		_, err = l.svcCtx.StudentRpc.ModifyStudent(l.ctx, &studentservice.ModifyStudentRequest{
			Id:         s.Id,
			NewStudent: s,
		})
		if err != nil {
			return nil, err
		}
	}

	return &types.AddDefenseReply{
		Id: defenseId,
	}, err
}
