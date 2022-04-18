package logic

import (
	"context"

	"TDS-backend/common/errorx"
	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryNongroupedStudentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryNongroupedStudentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryNongroupedStudentsLogic {
	return QueryNongroupedStudentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryNongroupedStudentsLogic) QueryNongroupedStudents() (resp *types.QueryStudentsReply, err error) {
	stuRes, err := l.svcCtx.StudentRpc.QueryNongroupedStudents(l.ctx, &student.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	if len(stuRes.List) == 0 {
		return nil, errorx.NewDefaultError("所有学生均分组")
	}
	resp = &types.QueryStudentsReply{}
	for _, s := range stuRes.List {
		resp.Students = append(resp.Students, types.Student{
			Id:            s.Id,
			Name:          s.Name,
			PaperKeywords: s.PaperKeywords,
			PaperTitle:    s.PaperTitle,
			GroupId:       s.GroupId,
		})
	}
	return resp, nil
}
