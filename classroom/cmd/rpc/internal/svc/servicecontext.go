package svc

import (
	"TDS-backend/classroom/cmd/rpc/internal/config"
	"TDS-backend/classroom/model"
	"TDS-backend/common/mongox"	
)
type ServiceContext struct {
	Config config.Config
	ClassroomModel model.IClassroomModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ClassroomModel: model.NewClassroomModel(mongox.ConnectMongoDB(c.Mongodb)),
	}
}
