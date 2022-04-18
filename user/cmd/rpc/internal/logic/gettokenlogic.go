package logic

import (
	"context"
	"time"

	"TDS-backend/user/cmd/rpc/internal/svc"
	"TDS-backend/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
	jwt "github.com/dgrijalva/jwt-go"
)

type GetTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenLogic) GetToken(in *user.TokenRequest) (*user.TokenResponse, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now + in.ExpireTime
	claims["iat"] = now
	claims["id"] = in.Id
	claims["role"] = in.Role
	claims["is_secretary"] = in.IsSecretary
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	
	tokenStr, err := token.SignedString([]byte(in.SecretKey))
	return &user.TokenResponse{
		Token: tokenStr,
	},err
}
