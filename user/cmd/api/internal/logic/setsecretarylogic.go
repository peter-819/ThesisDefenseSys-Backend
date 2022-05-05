package logic

import (
	"context"
	"fmt"

	"TDS-backend/user/cmd/api/internal/svc"
	"TDS-backend/user/cmd/api/internal/types"
	"TDS-backend/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetSecretaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetSecretaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetSecretaryLogic {
	return &SetSecretaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetSecretaryLogic) SetSecretary(req *types.SetSecretaryReq) error {
	// todo: add your logic here and delete this line
	fmt.Println(req)
	_, err := l.svcCtx.UserRpc.SetSecretary(l.ctx, &user.SetSecretaryRequest{
		Id:          req.Id,
		IsSecretary: req.IsSecretary,
	})
	return err
}
