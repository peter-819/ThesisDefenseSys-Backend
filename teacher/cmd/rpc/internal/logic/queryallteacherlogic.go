package logic

import (
	"context"

	"TDS-backend/teacher/cmd/rpc/internal/svc"
	"TDS-backend/teacher/cmd/rpc/types/teacher"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllTeacherLogic {
	return &QueryAllTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllTeacherLogic) QueryAllTeacher(in *teacher.QueryAllTeacherRequest) (*teacher.QueryTeachersResponse, error) {
	// todo: add your logic here and delete this line
	// l.svcCtx.TeacherModel.
	return &teacher.QueryTeachersResponse{}, nil
}
