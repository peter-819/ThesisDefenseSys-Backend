package logic

import (
	"TDS-backend/common/errorx"
	"TDS-backend/common/timex"
	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/schedule/model"
	"context"
	"fmt"

	"TDS-backend/classroom/cmd/rpc/types/classroom"
	"TDS-backend/teacher/cmd/rpc/types/teacher"

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
	// todo: add your logic here and delete this line
	ori, err := l.svcCtx.DefenseModel.QueryDefense(req.ModifyID)
	if err != nil {
		return err
	}
	if req.NewDefense.GroupId != "" && req.NewDefense.GroupId != ori.GroupId {
		return errorx.NewDefaultError("不支持修改学生信息")
	}
	//classroom check
	classroomId := req.NewDefense.Classroom
	if classroomId != "" {
		claRes, err := l.svcCtx.ClassroomRpc.QueryClassroom(l.ctx, &classroom.QueryClassroomRequest{
			Id: classroomId,
		})
		fmt.Println("claRes: ", claRes)
		if err != nil {
			return errorx.NewDefaultError("非法教室ID: " + err.Error())
		}
		ori.ClassroomFullName = claRes.Location + " " + claRes.Name
		ori.Classroom = classroomId
	}
	//committee check
	if len(req.NewDefense.Committee) != 0 {
		ori.Committee = []model.Member{}
		for _, c := range req.NewDefense.Committee {
			teacherId := c.TeacherID
			teaRes, err := l.svcCtx.TeacherRpc.QueryTeacher(l.ctx, &teacher.QueryTeacherRequest{
				Id: teacherId,
			})
			fmt.Println("teaRes: ", teaRes.Teacher)
			if err != nil {
				return errorx.NewDefaultError("非法教师ID: " + err.Error())
			}
			// ori.Committee[index].TeacherName = teaRes.Teacher.Name
			// ori.Committee[index].TeacherID = teacherId
			ori.Committee = append(ori.Committee, model.Member{
				TeacherID:   teacherId,
				TeacherName: teaRes.Teacher.Name,
			})
		}
	}
	// time check
	if req.NewDefense.StartTime != "" {
		ori.StartTime, err = timex.FStringToTime(req.NewDefense.StartTime)
		if err != nil {
			return err
		}
	}
	if req.NewDefense.EndTime != "" {
		ori.EndTime, err = timex.FStringToTime(req.NewDefense.EndTime)
		if err != nil {
			return err
		}
	}
	if ori.StartTime.After(ori.EndTime) {
		return errorx.NewDefaultError("时间段设置错误")
	}
	err = l.svcCtx.DefenseModel.ModifyDefense(req.ModifyID, ori)
	return err
}
