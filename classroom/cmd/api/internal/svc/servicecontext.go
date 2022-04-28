package svc

import (
	"TDS-backend/classroom/cmd/api/internal/config"

	"TDS-backend/classroom/cmd/rpc/classroomservice"
	"TDS-backend/common/interceptor"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	ClassroomRpc classroomservice.ClassroomService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ClassroomRpc: classroomservice.NewClassroomService(zrpc.MustNewClient(c.ClassroomRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
	}
}
