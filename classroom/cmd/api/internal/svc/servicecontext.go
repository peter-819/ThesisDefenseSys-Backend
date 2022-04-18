package svc

import (
	"TDS-backend/classroom/cmd/api/internal/config"
	"TDS-backend/classroom/model"
	"TDS-backend/common/mongox"	
	
	"TDS-backend/classroom/cmd/rpc/classroomservice"
    "github.com/zeromicro/go-zero/zrpc"
	"TDS-backend/common/interceptor"
	"TDS-backend/excel/cmd/rpc/user"
)

type ServiceContext struct {
	Config config.Config
	ClassroomModel model.IClassroomModel
	ClassroomRpc classroomservice.ClassroomService
	ExcelRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ClassroomModel: model.NewClassroomModel(mongox.ConnectMongoDB(c.Mongodb)),
		ClassroomRpc: classroomservice.NewClassroomService(zrpc.MustNewClient(c.ClassroomRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
		ExcelRpc: user.NewUser(zrpc.MustNewClient(c.ExcelRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
	}
}
