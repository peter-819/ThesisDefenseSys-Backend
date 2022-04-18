package logic

import (
	"context"

	"TDS-backend/classroom/cmd/api/internal/svc"
	"TDS-backend/classroom/cmd/api/internal/types"
	"TDS-backend/classroom/cmd/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryClassroomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryClassroomLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryClassroomLogic {
	return QueryClassroomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryClassroomLogic) QueryClassroom(req types.QueryClassroomReq) (resp *types.Classroom, err error) {
	// todo: add your logic here and delete this line
	classroom, err := l.svcCtx.ClassroomModel.FindClassroomByIdString(req.Id)
	if err != nil {
		return nil, err
	}
	resp = utils.ClassroomBtoJ(classroom)
	return resp, nil
}
