package logic

import (
	"context"
	"fmt"
	
	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/internal/utils"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyStudentLogic {
	return &ModifyStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyStudentLogic) ModifyStudent(in *student.ModifyStudentRequest) (*student.EmptyResponse, error) {
	// todo: add your logic here and delete this line
	t, err := utils.StudentJtoB(in.NewStudent)
	if err != nil {
		return nil, err
	}
	fmt.Println("bson: ", t)
	err = l.svcCtx.StudentModel.ModifyStudent(in.Id, t)
	return &student.EmptyResponse{}, err
}
