package logic

import (
	"context"

	"TDS-backend/common/errorx"
	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveGroupLogic {
	return RemoveGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveGroupLogic) RemoveGroup(req types.RemoveGroupReq) error {
	//修改学生
	group, err := l.svcCtx.StudentRpc.QueryGroup(l.ctx, &student.QueryGroupRequest{
		GroupId: req.Id,
	})
	if err != nil {
		return err
	}
	for _, member := range group.Members {
		stuRes, err := l.svcCtx.StudentRpc.QueryStudent(l.ctx, &student.QueryStudentRequest{
			Id: member,
		})
		if err != nil {
			return err
		}
		if stuRes.Student.GroupId != req.Id {
			return errorx.NewDefaultError("内部错误！")
		}
		stuRes.Student.GroupId = ""
		_, err = l.svcCtx.StudentRpc.ModifyStudent(l.ctx, &student.ModifyStudentRequest{
			Id:         member,
			NewStudent: stuRes.Student,
		})
		if err != nil {
			return err
		}
	}
	//删除组
	_, err = l.svcCtx.StudentRpc.RemoveGroup(l.ctx, &student.QueryGroupRequest{
		GroupId: req.Id,
	})

	return err
}
