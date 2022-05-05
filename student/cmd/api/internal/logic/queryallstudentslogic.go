package logic

import (
	"context"

	"TDS-backend/common/typex"
	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"

	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllStudentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAllStudentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryAllStudentsLogic {
	return QueryAllStudentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAllStudentsLogic) QueryAllStudents() (resp *types.QueryStudentsReply, err error) {
	// todo: add your logic here and delete this line
	stuRes, err := l.svcCtx.StudentRpc.QueryAllStudents(l.ctx, &student.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	resStu := make([]types.Student, 0)
	err = typex.Convert(&stuRes.List, &resStu)
	if err != nil {
		return nil, err
	}
	resp = &types.QueryStudentsReply{
		Students: resStu,
	}
	return resp, nil
}
