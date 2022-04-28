package svc

import (
	"TDS-backend/common/interceptor"
	"TDS-backend/teacher/cmd/api/internal/config"
	"TDS-backend/teacher/cmd/rpc/teacherservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	TeacherRpc teacherservice.TeacherService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		TeacherRpc: teacherservice.NewTeacherService(zrpc.MustNewClient(c.TeacherRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
	}
}
