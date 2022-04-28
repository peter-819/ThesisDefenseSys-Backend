package logic

import (
	"context"

	"TDS-backend/common/typex"
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
	resp := &teacher.QueryTeachersResponse{}
	teachers, err := l.svcCtx.TeacherModel.QueryAllTeachers()
	if err != nil {
		return nil, err
	}
	err = typex.Convert(&teachers, &resp.List)
	return resp, err
}
