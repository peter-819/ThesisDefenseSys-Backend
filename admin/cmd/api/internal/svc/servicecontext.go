package svc

import (
	"TDS-backend/admin/cmd/api/internal/config"
    "github.com/zeromicro/go-zero/zrpc"
	excelrpc "TDS-backend/excel/cmd/rpc/user"
	userrpc "TDS-backend/user/cmd/rpc/user"
	"TDS-backend/common/interceptor"
)

type ServiceContext struct {
	Config config.Config
	ExcelRpc excelrpc.User
	UserRpc userrpc.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,	
		ExcelRpc: excelrpc.NewUser(zrpc.MustNewClient(c.ExcelRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
		UserRpc: userrpc.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
	}
}
