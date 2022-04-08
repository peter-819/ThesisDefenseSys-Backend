package svc

import (
	"TDS-backend/user/cmd/api/internal/config"
    "github.com/zeromicro/go-zero/zrpc"
	"TDS-backend/user/cmd/rpc/user"
	"TDS-backend/common/interceptor"
	
)

type ServiceContext struct {
	Config config.Config
    UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
        UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
	}
}
