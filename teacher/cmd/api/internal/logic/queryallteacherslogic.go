package logic

import (
	"context"

	"TDS-backend/teacher/cmd/api/internal/svc"
	"TDS-backend/teacher/cmd/api/internal/types"
	"TDS-backend/teacher/cmd/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllTeachersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAllTeachersLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryAllTeachersLogic {
	return QueryAllTeachersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAllTeachersLogic) QueryAllTeachers() (resp *types.QueryTeachersReq, err error) {
	// todo: add your logic here and delete this line
	teachers, err := l.svcCtx.TeacherModel.QueryAllTeachers()
	if err != nil {
		return nil, err
	}
	resp = &types.QueryTeachersReq{
		Teachers: []types.Teacher{},
	}
	for _, t := range teachers {
		resp.Teachers = append(resp.Teachers, *utils.TeacherBtoJ(&t))
	}
	return resp, nil
}
