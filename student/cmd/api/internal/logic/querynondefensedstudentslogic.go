package logic

import (
	"context"

	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryNondefensedStudentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryNondefensedStudentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryNondefensedStudentsLogic {
	return &QueryNondefensedStudentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryNondefensedStudentsLogic) QueryNondefensedStudents() (resp *types.QueryStudentsReply, err error) {
	// todo: add your logic here and delete this line

	return
}
