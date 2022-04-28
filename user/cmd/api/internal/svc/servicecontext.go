package svc

import (
	"TDS-backend/common/interceptor"
	"TDS-backend/teacher/cmd/rpc/teacherservice"
	"TDS-backend/user/cmd/api/internal/config"
	"TDS-backend/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserRpc    user.User
	TeacherRpc teacherservice.TeacherService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
		TeacherRpc: teacherservice.NewTeacherService(zrpc.MustNewClient(c.TeacherRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
	}
}
