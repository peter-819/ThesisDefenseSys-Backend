package logic

import (
	"context"

	"TDS-backend/classroom/cmd/rpc/internal/svc"
	"TDS-backend/classroom/cmd/rpc/types/classroom"
	"TDS-backend/classroom/model"
	"TDS-backend/common/typex"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyClassroomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyClassroomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyClassroomLogic {
	return &ModifyClassroomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyClassroomLogic) ModifyClassroom(in *classroom.ModifyClassroomRequest) (*classroom.EmptyRequest, error) {
	// todo: add your logic here and delete this line
	cla := &model.Classroom{}
	typex.Convert(in.NewClassroom, cla)
	err := l.svcCtx.ClassroomModel.ModifyClassroom(in.Id, cla)
	if err != nil {
		return nil, err
	}
	return &classroom.EmptyRequest{}, nil
}
