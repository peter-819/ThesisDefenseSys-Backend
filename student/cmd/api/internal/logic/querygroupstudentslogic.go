package logic

import (
	"context"

	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGroupStudentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryGroupStudentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryGroupStudentsLogic {
	return QueryGroupStudentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryGroupStudentsLogic) QueryGroupStudents(req types.QueryGroupReq) (resp *types.QueryStudentsReply, err error) {
	// todo: add your logic here and delete this line
	group, err := l.svcCtx.StudentRpc.QueryGroup(l.ctx, &student.QueryGroupRequest{
		GroupId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	stuRes, err := l.svcCtx.StudentRpc.QueryStudentsBatch(l.ctx, &student.QueryStudentsBatchRequest{
		Ids: group.Members,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.QueryStudentsReply{}
	for _, s := range stuRes.List {
		resp.Students = append(resp.Students, types.Student{
			Id:            s.Id,
			Name:          s.Name,
			PaperKeywords: s.PaperKeywords,
			PaperTitle:    s.PaperTitle,
			Mentor:        s.Mentor,
		})
	}

	return resp, nil
}
