package logic

import (
	"context"

	"TDS-backend/common/typex"
	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryStudentLogic {
	return &QueryStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryStudentLogic) QueryStudent(req *types.QueryStudentReq) (resp *types.Student, err error) {
	// todo: add your logic here and delete this line
	stuRes, err := l.svcCtx.StudentRpc.QueryStudent(l.ctx, &student.QueryStudentRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	outStudent := &types.Student{}
	err = typex.Convert(stuRes.Student, outStudent)
	if err != nil {
		return nil, err
	}
	return outStudent, nil
}
