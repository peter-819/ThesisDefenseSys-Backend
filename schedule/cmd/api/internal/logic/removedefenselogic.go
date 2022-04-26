package logic

import (
	"context"
	"fmt"

	"TDS-backend/common/errorx"
	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"

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

	group, err := l.svcCtx.StudentRpc.QueryGroup(l.ctx, &student.QueryGroupRequest{
		GroupId: defense.GroupId,
	})
	if err != nil {
		return err
	}
	fmt.Println(group)
	for _, member := range group.Members {
		stuRes, err := l.svcCtx.StudentRpc.QueryStudent(l.ctx, &student.QueryStudentRequest{
			Id: member,
		})
		if err != nil {
			return err
		}
		if stuRes.Student.GroupId != group.Id {
			return errorx.NewDefaultError("内部错误！")
		}
		fmt.Println("remove student: ", stuRes)
		stuRes.Student.GroupId = ""
		_, err = l.svcCtx.StudentRpc.ModifyStudent(l.ctx, &student.ModifyStudentRequest{
			Id:         member,
			NewStudent: stuRes.Student,
		})
		if err != nil {
			return err
		}
	}

	_, err = l.svcCtx.StudentRpc.RemoveGroup(l.ctx, &student.QueryGroupRequest{
		GroupId: defense.GroupId,
	})
	if err != nil {
		return err
	}
	err = l.svcCtx.DefenseModel.RemoveDefense(req.ID)
	return err
}
