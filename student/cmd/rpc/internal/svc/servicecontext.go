package svc

import (
	"TDS-backend/common/mongox"
	"TDS-backend/student/cmd/rpc/internal/config"
	"TDS-backend/student/model"
)

type ServiceContext struct {
	Config       config.Config
	StudentModel model.IStudentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		StudentModel: model.NewStudentModel(mongox.ConnectMongoDB(c.Mongodb)),
	}
}
