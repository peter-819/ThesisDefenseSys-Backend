package svc

import (
	"TDS-backend/common/mongox"
	"TDS-backend/user/cmd/rpc/internal/config"
	"TDS-backend/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.IUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(mongox.ConnectMongoDB(c.Mongodb)),
	}
}
