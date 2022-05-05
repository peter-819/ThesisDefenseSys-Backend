package svc

import (
	"TDS-backend/common/interceptor"
	"TDS-backend/user/cmd/api/internal/config"
	"TDS-backend/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
	}
}
