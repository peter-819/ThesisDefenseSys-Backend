package logic

import (
	"context"
	"fmt"
	"net/http"

	"TDS-backend/user/cmd/api/internal/svc"
	"TDS-backend/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyTokenLogic) VerifyToken(req *types.VerifyTokenReq, r *http.Request) (resp *types.VerifyTokenResp, err error) {
	// todo: add your logic here and delete this line
	authorization := r.Header.Get("Authorization")
	realRequestPath := r.Header.Get("X-Original-Uri")
	fmt.Println("authorization: ", authorization)
	fmt.Println("realRequestPath: ", realRequestPath)
	role := l.ctx.Value("role").(string)
	fmt.Println("role: ", role)
	id := l.ctx.Value("id").(string)
	return &types.VerifyTokenResp{
		UserId: id,
		Ok:     true,
	}, nil
}
