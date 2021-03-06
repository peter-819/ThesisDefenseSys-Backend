package logic

import (
	"context"

	"TDS-backend/common/typex"
	"TDS-backend/teacher/cmd/rpc/internal/svc"
	"TDS-backend/teacher/cmd/rpc/types/teacher"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTeacherLogic {
	return &QueryTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryTeacherLogic) QueryTeacher(in *teacher.QueryTeacherRequest) (*teacher.QueryTeacherResponse, error) {
	// todo: add your logic here and delete this line
	t, err := l.svcCtx.TeacherModel.QueryTeacher(in.Id)
	if err != nil {
		return nil, err
	}
	outTeacher := new(teacher.Teacher)
	err = typex.Convert(t, outTeacher)
	if err != nil {
		return nil, err
	}
	out := &teacher.QueryTeacherResponse{}
	out.Teacher = outTeacher
	return out, nil
}
