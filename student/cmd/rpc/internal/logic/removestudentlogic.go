package logic

import (
	"context"

	"TDS-backend/common/errorx"
	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveStudentLogic {
	return &RemoveStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveStudentLogic) RemoveStudent(in *student.RemoveStudentRequest) (*student.EmptyResponse, error) {
	// todo: add your logic here and delete this line
	stu, err := l.svcCtx.StudentModel.QueryStudent(in.Id)
	if err != nil {
		return nil, errorx.NewDefaultError("找不到学生")
	}
	if stu.DefenseId != "" {
		return nil, errorx.NewDefaultError("请先移除该学生参加的答辩")
	}
	err = l.svcCtx.StudentModel.RemoveStudent(in.Id)
	return &student.EmptyResponse{}, err
}
