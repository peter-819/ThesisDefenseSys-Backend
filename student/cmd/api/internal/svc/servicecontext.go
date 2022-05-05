package svc

import (
	"TDS-backend/common/interceptor"
	"TDS-backend/student/cmd/api/internal/config"
	"TDS-backend/student/cmd/rpc/studentservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	StudentRpc studentservice.StudentService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		StudentRpc: studentservice.NewStudentService(zrpc.MustNewClient(c.StudentRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
	}
}
