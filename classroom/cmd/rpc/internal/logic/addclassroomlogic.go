package logic

import (
	"context"

	"TDS-backend/classroom/cmd/rpc/internal/svc"
	"TDS-backend/classroom/cmd/rpc/types/classroom"
	"TDS-backend/classroom/model"
	"TDS-backend/common/typex"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassroomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddClassroomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassroomLogic {
	return &AddClassroomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddClassroomLogic) AddClassroom(in *classroom.Classroom) (*classroom.EmptyRequest, error) {
	// todo: add your logic here and delete this line
	cla := &model.Classroom{}
	typex.Convert(in, cla)
	err := l.svcCtx.ClassroomModel.AddClassroom(cla)
	if err != nil {
		return nil, err
	}
	return &classroom.EmptyRequest{}, nil
}
