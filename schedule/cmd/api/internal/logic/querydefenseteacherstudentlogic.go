package logic

import (
	"context"

	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/schedule/cmd/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDefenseTeacherStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDefenseTeacherStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryDefenseTeacherStudentLogic {
	return QueryDefenseTeacherStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDefenseTeacherStudentLogic) QueryDefenseTeacherStudent(req types.QueryDefenseReq) (resp *types.QueryDefenseReply, err error) {
	// todo: add your logic here and delete this line
	resp = &types.QueryDefenseReply{}
	
	teacherDefenses, err := l.svcCtx.DefenseModel.GetTeacherDefense(req.Id)
	if err != nil {
		return nil, err
	}
	studentDefenses, err := l.svcCtx.DefenseModel.GetStudentDefense(req.Id)
	if err != nil {
		return nil, err
	}
	if len(teacherDefenses) > 0 {
		for _, s := range teacherDefenses {
			resp.Defenses = append(resp.Defenses, *utils.DefenseBtoJ(&s))
		}
	} else if len(studentDefenses) > 0 {
		for _, s := range studentDefenses {
			resp.Defenses = append(resp.Defenses, *utils.DefenseBtoJ(&s))
		}
	} 
	return resp, err
}
