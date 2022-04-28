package logic

import (
	"context"

	"TDS-backend/common/timex"
	"TDS-backend/common/typex"
	"TDS-backend/teacher/cmd/rpc/internal/svc"
	"TDS-backend/teacher/cmd/rpc/types/teacher"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAvailableTeachersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAvailableTeachersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAvailableTeachersLogic {
	return &QueryAvailableTeachersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAvailableTeachersLogic) QueryAvailableTeachers(in *teacher.QueryAvailableTeachersRequest) (*teacher.QueryTeachersResponse, error) {
	// todo: add your logic here and delete this line
	start, err := timex.GStringToTime(in.StartTime)
	if err != nil {
		return nil, err
	}
	end, err := timex.GStringToTime(in.EndTime)
	if err != nil {
		return nil, err
	}
	resp := &teacher.QueryTeachersResponse{}
	teachers, err := l.svcCtx.TeacherModel.QueryAvailableTeachers(start, end, in.Keywords, in.Excluded)

	if err != nil {
		return nil, err
	}
	for _, t := range teachers {
		teacher := &teacher.Teacher{}
		typex.Convert(&t, teacher)
		resp.List = append(resp.List, teacher)
	}
	return resp, nil
}
