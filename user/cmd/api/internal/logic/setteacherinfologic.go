package logic

import (
	"context"
	"fmt"
	"TDS-backend/user/cmd/api/internal/svc"
	"TDS-backend/user/cmd/api/internal/types"
	"TDS-backend/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetteacherinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetteacherinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetteacherinfoLogic {
	return SetteacherinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetteacherinfoLogic) Setteacherinfo(req types.SetTeacherInfoReq) error {
	// todo: add your logic here and delete this line
	fmt.Println("Set Teacher Info")
	fmt.Println(req)
	
	_, err := l.svcCtx.UserRpc.SetTeacherInfo(l.ctx, &user.SetTeacherInfoRequest{
		Id: req.Id,
		Info: &user.TeacherInfo{
			Id: req.Info.Id,
			Name: req.Info.Name,
			IsSecretary: req.Info.IsSecretary,
		},
	})
	
	return err
}
