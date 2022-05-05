package logic

import (
	"context"

	"TDS-backend/classroom/cmd/rpc/types/classroom"
	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"
	"TDS-backend/teacher/cmd/rpc/types/teacher"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveDefenseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveDefenseLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveDefenseLogic {
	return RemoveDefenseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveDefenseLogic) RemoveDefense(req types.RemoveDefenseReq) error {
	// todo: add your logic here and delete this line
	defense, err := l.svcCtx.DefenseModel.QueryDefense(req.ID)
	if err != nil {
		return err
	}

	if defense.Classroom != "" {
		claRes, err := l.svcCtx.ClassroomRpc.QueryClassroom(l.ctx, &classroom.QueryClassroomRequest{
			Id: defense.Classroom,
		})
		if err != nil {
			return err
		}
		newDefenses := []*classroom.Defense{}
		for _, d := range claRes.Defenses {
			if d.DefenseId != req.ID {
				newDefenses = append(newDefenses, d)
			}
		}
		claRes.Defenses = newDefenses
		_, err = l.svcCtx.ClassroomRpc.ModifyClassroom(l.ctx, &classroom.ModifyClassroomRequest{
			Id:           defense.Classroom,
			NewClassroom: claRes,
		})
		if err != nil {
			return err
		}
	}
	for _, s := range defense.Students {
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
	for _, t := range defense.Committee {
		teaRes, err := l.svcCtx.TeacherRpc.QueryTeacher(l.ctx, &teacher.QueryTeacherRequest{
			Id: t.TeacherID,
		})
		if err != nil {
			return err
		}
		newDefenses := []*teacher.Defense{}
		for _, d := range teaRes.Teacher.Defenses {
			if d.DefenseId != req.ID {
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

	err = l.svcCtx.DefenseModel.RemoveDefense(req.ID)
	return err
}
