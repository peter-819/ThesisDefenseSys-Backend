package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"TDS-backend/common/typex"
	"TDS-backend/teacher/cmd/api/internal/svc"
	"TDS-backend/teacher/cmd/api/internal/types"
	"TDS-backend/teacher/cmd/rpc/teacherservice"

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
	teacher, err := l.svcCtx.TeacherRpc.QueryTeacher(l.ctx, &teacherservice.QueryTeacherRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.Teacher)
	fmt.Println("respbefore: ", resp)
	j, _ := json.Marshal(resp)
	fmt.Println("respbefore json:", string(j))
	err = typex.Convert(teacher.Teacher, resp)
	fmt.Println("respafter: ", resp)
	return resp, err
}
