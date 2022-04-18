package logic

import (
	"context"

	"TDS-backend/user/cmd/api/internal/svc"
	"TDS-backend/user/cmd/api/internal/types"
	
	"TDS-backend/user/cmd/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetteacherlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetteacherlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetteacherlistLogic {
	return GetteacherlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetteacherlistLogic) Getteacherlist() (resp *types.GetTeacherListReply, err error) {
	// todo: add your logic here and delete this line
	resp = &types.GetTeacherListReply{}
	res, err := l.svcCtx.UserRpc.GetTeacherList(l.ctx,&user.GetTeacherListRequest{})
	for _, t := range res.List {
		resp.List = append(resp.List, types.TeacherInfo{
			Name: t.Name,
			Id: t.Id,
			IsSecretary: t.IsSecretary,
		})
	}
	return resp, err
}
