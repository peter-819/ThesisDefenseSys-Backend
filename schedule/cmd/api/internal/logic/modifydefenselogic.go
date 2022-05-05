package logic

import (
	"TDS-backend/classroom/cmd/rpc/classroomservice"
	"TDS-backend/classroom/cmd/rpc/types/classroom"
	"TDS-backend/common/errorx"
	"TDS-backend/common/timex"
	"TDS-backend/common/typex"
	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/schedule/model"
	"TDS-backend/student/cmd/rpc/studentservice"
	"TDS-backend/student/cmd/rpc/types/student"
	"TDS-backend/teacher/cmd/rpc/teacherservice"
	"TDS-backend/teacher/cmd/rpc/types/teacher"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyDefenseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyDefenseLogic(ctx context.Context, svcCtx *svc.ServiceContext) ModifyDefenseLogic {
	return ModifyDefenseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyDefenseLogic) ModifyDefense(req types.ModifyDefenseReq) error {
	ori, err := l.svcCtx.DefenseModel.QueryDefense(req.ModifyID)
	new := &types.Defense{}
	typex.Convert(ori, new)
	fmt.Println("old defense: ", ori)
	fmt.Println("new defense: ", new)
	if err != nil {
		return err
	}
	// get new defense
	update := req.NewDefense
	if update.Classroom != "" {
		new.Classroom = update.Classroom
	}
	if update.StartTime != "" {
		new.StartTime = update.StartTime
	}
	if update.EndTime != "" {
		new.EndTime = update.EndTime
	}
	if len(update.Committee) > 0 {
		committee := []types.Member{}
		for _, t := range update.Committee {
			committee = append(committee, types.Member{
				TeacherID: t.TeacherID,
			})
		}
		new.Committee = committee
	}
	if len(update.Students) > 0 {
		students := []types.Student{}
		for _, s := range update.Students {
			students = append(students, types.Student{
				StudentID: s.StudentID,
			})
		}
		new.Students = students
	}
	// remove old
	if ori.Classroom != "" {
		claRes, err := l.svcCtx.ClassroomRpc.QueryClassroom(l.ctx, &classroom.QueryClassroomRequest{
			Id: ori.Classroom,
		})
		if err != nil {
			return err
		}
		newDefenses := []*classroom.Defense{}
		for _, d := range claRes.Defenses {
			if d.DefenseId != req.ModifyID {
				newDefenses = append(newDefenses, d)
			}
		}
		claRes.Defenses = newDefenses
		_, err = l.svcCtx.ClassroomRpc.ModifyClassroom(l.ctx, &classroom.ModifyClassroomRequest{
			Id:           ori.Classroom,
			NewClassroom: claRes,
		})
		if err != nil {
			return err
		}
	}
	for _, s := range ori.Students {
		stuRes, err := l.svcCtx.StudentRpc.QueryStudent(l.ctx, &student.QueryStudentRequest{
			Id: s.StudentID,
		})
		if err != nil {
			return err
		}
		stuRes.Student.DefenseId = ""
		_, err = l.svcCtx.StudentRpc.ModifyStudent(l.ctx, &student.ModifyStudentRequest{
			Id:         s.StudentID,
			NewStudent: stuRes.Student,
		})
		if err != nil {
			return err
		}
	}
	for _, t := range ori.Committee {
		teaRes, err := l.svcCtx.TeacherRpc.QueryTeacher(l.ctx, &teacher.QueryTeacherRequest{
			Id: t.TeacherID,
		})
		if err != nil {
			return err
		}
		newDefenses := []*teacher.Defense{}
		for _, d := range teaRes.Teacher.Defenses {
			if d.DefenseId != req.ModifyID {
				newDefenses = append(newDefenses, d)
			}
		}
		teaRes.Teacher.Defenses = newDefenses
		_, err = l.svcCtx.TeacherRpc.ModifyTeacher(l.ctx, &teacher.ModifyTeacherRequest{
			Id:         t.TeacherID,
			NewTeacher: teaRes.Teacher,
		})
		if err != nil {
			return err
		}
	}
	// add new

	var cla *classroom.Classroom
	var teas []*teacher.Teacher
	var stus []*student.Student
	//classroom check
	classroomId := new.Classroom
	if classroomId != "" {
		claRes, err := l.svcCtx.ClassroomRpc.QueryClassroom(l.ctx, &classroom.QueryClassroomRequest{
			Id: classroomId,
		})
		if err != nil {
			return errorx.NewDefaultError("非法教室ID: " + err.Error())
		}
		fmt.Println("claRes: ", claRes)
		inc, err := timex.IsIncludedString(claRes.StartTime, claRes.EndTime, new.StartTime, new.EndTime)
		if err != nil {
			return err
		}
		if !inc {
			return errorx.NewDefaultError("教室不可用")
		}
		for _, d := range claRes.Defenses {
			inter, err := timex.IsIntersectedString(new.StartTime, new.EndTime, d.StartTime, d.EndTime)
			if err != nil {
				return err
			}
			if inter {
				return errorx.NewDefaultError("教室该时段有答辩安排")
			}
		}

		cla = claRes
		new.ClassroomFullName = claRes.Location + " " + claRes.Name
	}
	// committee check
	if len(new.Committee) != 0 {
		for index, c := range new.Committee {
			teacherId := c.TeacherID
			teaRes, err := l.svcCtx.TeacherRpc.QueryTeacher(l.ctx, &teacher.QueryTeacherRequest{
				Id: teacherId,
			})
			if err != nil {
				return errorx.NewDefaultError("非法教师ID: " + err.Error())
			}
			fmt.Println("teaRes: ", teaRes)
			isAble := false
			for _, schedule := range teaRes.Teacher.Schedules {
				inc, err := timex.IsIncludedString(schedule.StartTime, schedule.EndTime, new.StartTime, new.EndTime)
				if err != nil {
					return err
				}
				isAble = isAble || inc
			}
			if !isAble {
				return errorx.NewDefaultError("教师无可用时间")
			}
			timesThisWeek := 0
			for _, d := range teaRes.Teacher.Defenses {
				bSameWeek, err := timex.IsSameWeekString(d.StartTime, new.StartTime)
				if err != nil {
					return err
				}
				if bSameWeek {
					timesThisWeek++
				}
				inter, err := timex.IsIntersectedString(d.StartTime, d.EndTime, new.StartTime, new.EndTime)
				if err != nil {
					return err
				}
				if inter {
					return errorx.NewDefaultError(("教师答辩时间冲突"))
				}
			}
			fmt.Println("defense count in this week: ", timesThisWeek+1)
			if teaRes.Teacher.MaxDefensePerWeek != 0 && timesThisWeek+1 > int(teaRes.Teacher.MaxDefensePerWeek) {
				return errorx.NewDefaultError(teaRes.Teacher.Name + "老师太忙啦")
			}
			teas = append(teas, teaRes.Teacher)
			new.Committee[index].TeacherName = teaRes.Teacher.Name
		}
	}
	// student check
	if len(new.Students) != 0 {
		for index, s := range new.Students {
			studentId := s.StudentID
			stuRes, err := l.svcCtx.StudentRpc.QueryStudent(l.ctx, &student.QueryStudentRequest{
				Id: studentId,
			})
			if err != nil {
				return errorx.NewDefaultError("非法学生ID: " + err.Error())
			}
			fmt.Println("stuRes: ", stuRes)
			if stuRes.Student.DefenseId != "" {
				return errorx.NewDefaultError("学生已有答辩安排")
			}
			stus = append(stus, stuRes.Student)
			new.Students[index].StudentName = stuRes.Student.Name
		}
	}

	newDefense := &model.Defense{}
	err = typex.Convert(new, newDefense)
	if err != nil {
		return err
	}
	err = l.svcCtx.DefenseModel.ModifyDefense(req.ModifyID, newDefense)
	if err != nil {
		return err
	}

	//modify classroom
	cla.Defenses = append(cla.Defenses, &classroom.Defense{
		DefenseId: req.ModifyID,
		StartTime: new.StartTime,
		EndTime:   new.EndTime,
	})
	_, err = l.svcCtx.ClassroomRpc.ModifyClassroom(l.ctx, &classroomservice.ModifyClassroomRequest{
		Id:           cla.Id,
		NewClassroom: cla,
	})
	if err != nil {
		return err
	}
	//modify teachers
	for _, t := range teas {
		t.Defenses = append(t.Defenses, &teacher.Defense{
			DefenseId: req.ModifyID,
			StartTime: new.StartTime,
			EndTime:   new.EndTime,
		})
		_, err = l.svcCtx.TeacherRpc.ModifyTeacher(l.ctx, &teacherservice.ModifyTeacherRequest{
			Id:         t.Id,
			NewTeacher: t,
		})
		if err != nil {
			return err
		}
	}
	//modify students
	for _, s := range stus {
		s.DefenseId = req.ModifyID
		_, err = l.svcCtx.StudentRpc.ModifyStudent(l.ctx, &studentservice.ModifyStudentRequest{
			Id:         s.Id,
			NewStudent: s,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
