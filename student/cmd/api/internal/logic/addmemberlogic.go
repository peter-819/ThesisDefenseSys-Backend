package logic

import (
	"context"

	"TDS-backend/common/errorx"
	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddMemberLogic {
	return AddMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMemberLogic) AddMember(req types.AddMemberReq) error {
	//修改学生
	stuRes, err := l.svcCtx.StudentRpc.QueryStudent(l.ctx, &student.QueryStudentRequest{
		Id: req.StudentId,
	})
	if err != nil {
		return err
	}
	if stuRes.Student.GroupId != "" {
		return errorx.NewDefaultError("该学生已加入组!")
	}
	stuRes.Student.GroupId = req.GroupId
	_, err = l.svcCtx.StudentRpc.ModifyStudent(l.ctx, &student.ModifyStudentRequest{
		Id:         req.StudentId,
		NewStudent: stuRes.Student,
	})
	if err != nil {
		return err
	}

	//加入组
	group, err := l.svcCtx.StudentRpc.QueryGroup(l.ctx, &student.QueryGroupRequest{
		GroupId: req.GroupId,
	})
	if err != nil {
		return err
	}
	group.Members = append(group.Members, req.StudentId)
	_, err = l.svcCtx.StudentRpc.ModifyGroup(l.ctx, &student.ModifyGroupRequest{
		Id:       req.GroupId,
		NewGroup: group,
	})
	return err
}
