package logic

import (
	"context"

	"TDS-backend/user/cmd/rpc/internal/svc"
	"TDS-backend/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterBatchLogic {
	return &RegisterBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterBatchLogic) RegisterBatch(in *user.RegisterBatchRequest) (*user.RegisterBatchResponse, error) {
	for _, user := range in.List {
		err := l.svcCtx.UserModel.AddUser(user.Id, user.Name, user.Role, user.Password)
		if err != nil {
			return nil, err
		}
	}
	return &user.RegisterBatchResponse{}, nil
}
