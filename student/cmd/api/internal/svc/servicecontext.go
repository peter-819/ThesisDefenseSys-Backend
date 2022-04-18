package svc

import (
	"TDS-backend/student/cmd/api/internal/config"
	"TDS-backend/student/cmd/rpc/studentservice"
	"TDS-backend/excel/cmd/rpc/user"
    "github.com/zeromicro/go-zero/zrpc"
	"TDS-backend/common/interceptor"
)

type ServiceContext struct {
	Config config.Config
	StudentRpc studentservice.StudentService
	ExcelRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		StudentRpc: studentservice.NewStudentService(zrpc.MustNewClient(c.StudentRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
		ExcelRpc: user.NewUser(zrpc.MustNewClient(c.ExcelRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
	}
}
