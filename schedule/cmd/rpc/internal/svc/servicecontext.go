package svc

import (
	"TDS-backend/common/mongox"
	"TDS-backend/schedule/cmd/rpc/internal/config"
	"TDS-backend/schedule/model"
)

type ServiceContext struct {
	Config       config.Config
	DefenseModel model.IDefenseModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		DefenseModel: model.NewDefenseModel(mongox.ConnectMongoDB(c.Mongodb)),
	}
}
