package logic

import (
	"context"

	"TDS-backend/user/cmd/rpc/internal/svc"
	"TDS-backend/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTeacherListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherListLogic {
	return &GetTeacherListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTeacherListLogic) GetTeacherList(in *user.GetTeacherListRequest) (*user.GetTeacherListResponse, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.TeacherModel.GetTeacherList()
	resp := &user.GetTeacherListResponse{}
	for _, t := range list {
		resp.List = append(resp.List, &user.TeacherInfo{
			Id: t.ID,
			Name: t.Name,
			IsSecretary: t.IsSecretary,
		})
	}
	return resp, err
}
