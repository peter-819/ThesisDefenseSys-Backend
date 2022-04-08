package svc

import (
	"TDS-backend/schedule/cmd/api/internal/config"
	"TDS-backend/schedule/model"
	"TDS-backend/common/mongox"	
)

type ServiceContext struct {
	Config config.Config
	ScheduleModel model.IScheduleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ScheduleModel: model.NewScheduleModel(mongox.ConnectMongoDB(c.Mongodb)),
	}
}
