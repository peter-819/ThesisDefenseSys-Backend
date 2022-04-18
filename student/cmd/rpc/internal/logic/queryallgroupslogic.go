package logic

import (
	"context"

	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/internal/utils"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAllGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllGroupsLogic {
	return &QueryAllGroupsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryAllGroupsLogic) QueryAllGroups(in *student.EmptyRequest) (*student.QueryGroupsResponse, error) {
	// todo: add your logic here and delete this line
	groups, err := l.svcCtx.GroupModel.QueryAllGroups()
	if err != nil {
		return nil, err
	}
	res := &student.QueryGroupsResponse{}
	for _, g := range groups {
		res.Groups = append(res.Groups, utils.GroupBtoJ(&g))
	}

	return res, nil
}
