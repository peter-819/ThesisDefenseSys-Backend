package logic

import (
	"context"

	"TDS-backend/common/errorx"
	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveMemberLogic {
	return RemoveMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveMemberLogic) RemoveMember(req types.RemoveMemberReq) error {
	//修改学生
	stuRes, err := l.svcCtx.StudentRpc.QueryStudent(l.ctx, &student.QueryStudentRequest{
		Id: req.StudentId,
	})
	if err != nil {
		return err
	}
	if stuRes.Student.GroupId != req.GroupId {
		return errorx.NewDefaultError("学生不在该组中!")
	}
	stuRes.Student.GroupId = ""
	_, err = l.svcCtx.StudentRpc.ModifyStudent(l.ctx, &student.ModifyStudentRequest{
		Id:         req.StudentId,
		NewStudent: stuRes.Student,
	})
	if err != nil {
		return err
	}
	//修改组
	group, err := l.svcCtx.StudentRpc.QueryGroup(l.ctx, &student.QueryGroupRequest{
		GroupId: req.GroupId,
	})
	if err != nil {
		return err
	}
	var new_members []string
	for _, m := range group.Members {
		if m == req.StudentId {
			continue
		}
		new_members = append(new_members, m)
	}
	group.Members = new_members
	_, err = l.svcCtx.StudentRpc.ModifyGroup(l.ctx, &student.ModifyGroupRequest{
		Id:       req.GroupId,
		NewGroup: group,
	})
	return err
}
