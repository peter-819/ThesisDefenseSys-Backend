package logic

import (
	"context"

	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAvailableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAvailableLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryAvailableLogic {
	return QueryAvailableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAvailableLogic) QueryAvailable(req types.QueryAvailableReq) (resp *types.QueryAvailableReply, err error) {
	// todo: add your logic here and delete this line
	// resp = &types.QueryAvailableReply{}
	// content, err := l.svcCtx.StudentRpc.QueryStudentsContent(l.ctx, &student.QueryStudentsContentRequest{
	// 	Ids: req.Students,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// avaTeachers, err := l.svcCtx.TeacherRpc.QueryAvailableTeachers(l.ctx, &teacher.QueryAvailableTeachersRequest{
	// 	Keywords:  content.Keywords,
	// 	Excluded:  content.Mentors,
	// 	StartTime: "",
	// 	EndTime:   "",
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// for _, t := range avaTeachers.List {
	// 	resp.Committee = append(resp.Committee, t.Id)
	// }

	return nil, nil
}
