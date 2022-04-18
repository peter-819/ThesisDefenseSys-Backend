package logic

import (
	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/schedule/cmd/api/internal/utils"
	"context"
	"fmt"

	"TDS-backend/classroom/cmd/rpc/types/classroom"
	"TDS-backend/student/cmd/rpc/types/student"
	"TDS-backend/teacher/cmd/rpc/types/teacher"

	"TDS-backend/common/errorx"

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

func (l *AddDefenseLogic) AddDefense(req types.Defense) error {
	if req.GroupId == "" {
		return errorx.NewDefaultError("学生组ID不可为空")
	}
	group, err := l.svcCtx.StudentRpc.QueryGroup(l.ctx, &student.QueryGroupRequest{
		GroupId: req.GroupId,
	})
	if err != nil {
		return err
	}

	//classroom check
	classroomId := req.Classroom
	if classroomId != "" {
		claRes, err := l.svcCtx.ClassroomRpc.QueryClassroom(l.ctx, &classroom.QueryClassroomRequest{
			Id: classroomId,
		})
		fmt.Println("claRes: ", claRes)
		if err != nil {
			return errorx.NewDefaultError("非法教室ID: " + err.Error())
		}
		req.ClassroomFullName = claRes.Location + " " + claRes.Name
	}
	// committee check
	if len(req.Committee) != 0 {
		for index, c := range req.Committee {
			teacherId := c.TeacherID
			teaRes, err := l.svcCtx.TeacherRpc.QueryTeacher(l.ctx, &teacher.QueryTeacherRequest{
				Id: teacherId,
			})
			fmt.Println("teaRes: ", teaRes.Teacher)
			if err != nil {
				return errorx.NewDefaultError("非法教师ID: " + err.Error())
			}
			req.Committee[index].TeacherName = teaRes.Teacher.Name
		}
	}

	schedule, err := utils.DefenseJtoB(&req)
	if err != nil {
		return err
	}
	defenseId, err := l.svcCtx.DefenseModel.InsertDefense(schedule)
	if err != nil {
		return err
	}
	group.DefenseId = defenseId
	fmt.Println("insert defense id : ", defenseId)
	_, err = l.svcCtx.StudentRpc.ModifyGroup(l.ctx, &student.ModifyGroupRequest{
		Id:       req.GroupId,
		NewGroup: group,
	})
	return err
}
