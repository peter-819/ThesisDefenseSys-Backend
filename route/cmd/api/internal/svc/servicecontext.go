package svc

import (
	"TDS-backend/route/cmd/api/internal/config"
	"TDS-backend/route/model"
	"TDS-backend/common/mongox"
)

type ServiceContext struct {
	Config config.Config
	RouteModel model.IRouteModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RouteModel: model.NewRouteModel(mongox.ConnectMongoDB(c.Mongodb)),
	}
}
