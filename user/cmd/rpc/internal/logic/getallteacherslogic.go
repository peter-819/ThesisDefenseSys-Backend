package logic

import (
	"context"

	"TDS-backend/common/typex"
	"TDS-backend/user/cmd/rpc/internal/svc"
	"TDS-backend/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllTeachersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllTeachersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllTeachersLogic {
	return &GetAllTeachersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllTeachersLogic) GetAllTeachers(in *user.EmptyRequest) (*user.GetUsersResponse, error) {
	// todo: add your logic here and delete this line
	teachers, err := l.svcCtx.UserModel.GetTeachers()
	if err != nil {
		return nil, err
	}
	outTeachers := make([]*user.UserResponse, 0)
	err = typex.Convert(&teachers, &outTeachers)
	return &user.GetUsersResponse{
		Users: outTeachers,
	}, err
}
