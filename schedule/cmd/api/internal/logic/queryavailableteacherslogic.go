package logic

import (
	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAvailableTeachersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAvailableTeachersLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryAvailableTeachersLogic {
	return QueryAvailableTeachersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAvailableTeachersLogic) QueryAvailableTeachers(req types.QueryAvailableTeachersReq) (resp *types.QueryAvailableTeachersResp, err error) {
	// fmt.Println("查询可用教师..")
	// defense, err := l.svcCtx.DefenseModel.QueryDefense(req.DefenseId)
	// if err != nil {
	// 	return nil, err
	// }
	// groupContent, err := l.svcCtx.StudentRpc.QueryGroupContent(l.ctx, &student.QueryGroupRequest{
	// 	GroupId: defense.GroupId,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("答辩: ", defense)
	// teachersRes, err := l.svcCtx.TeacherRpc.QueryAvailableTeachers(l.ctx, &teacher.QueryAvailableTeachersRequest{
	// 	StartTime: timex.TimeToString(defense.StartTime),
	// 	EndTime:   timex.TimeToString(defense.EndTime),
	// 	Keywords:  groupContent.Keywords,
	// 	Excluded:  groupContent.Mentors,
	// })
	// fmt.Println("RPC回复：", teachersRes)
	// if err != nil {
	// 	return nil, err
	// }
	// resp = &types.QueryAvailableTeachersResp{}
	// for _, c := range teachersRes.List {
	// 	resp.Teachers = append(resp.Teachers, types.TeacherInfo{
	// 		Name: c.Name,
	// 		Id:   c.Id,
	// 	})
	// }
	return nil, nil
}
