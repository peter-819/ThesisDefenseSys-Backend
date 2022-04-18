package logic

import (
	"context"

	"TDS-backend/teacher/cmd/api/internal/svc"
	"TDS-backend/teacher/cmd/api/internal/types"
	"TDS-backend/teacher/cmd/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryTeacherLogic {
	return QueryTeacherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryTeacherLogic) QueryTeacher(req types.QueryTeacherReq) (resp *types.Teacher, err error) {
	// todo: add your logic here and delete this line
	teacher, err := l.svcCtx.TeacherModel.QueryTeacher(req.Id)
	if err != nil {
		return nil, err
	}
	resp = utils.TeacherBtoJ(teacher)
	return resp, nil
}
