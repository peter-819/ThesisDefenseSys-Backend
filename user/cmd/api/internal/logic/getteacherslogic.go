package logic

import (
	"context"

	"TDS-backend/common/typex"
	"TDS-backend/user/cmd/api/internal/svc"
	"TDS-backend/user/cmd/api/internal/types"
	"TDS-backend/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeachersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTeachersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeachersLogic {
	return &GetTeachersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTeachersLogic) GetTeachers() (resp *types.QueryUsersResp, err error) {
	// todo: add your logic here and delete this line
	teachersRes, err := l.svcCtx.UserRpc.GetAllTeachers(l.ctx, &user.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	outTeachers := make([]types.UserInfo, 0)
	err = typex.Convert(&teachersRes.Users, &outTeachers)
	if err != nil {
		return nil, err
	}
	return &types.QueryUsersResp{
		Users: outTeachers,
	}, nil
}
