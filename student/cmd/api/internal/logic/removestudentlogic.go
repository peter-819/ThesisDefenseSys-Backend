package logic

import (
	"context"

	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveStudentLogic {
	return RemoveStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveStudentLogic) RemoveStudent(req types.RemoveStudentReq) error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.StudentRpc.RemoveStudent(l.ctx, &student.RemoveStudentRequest {
		Id:req.Id,
	})
	return err
}
