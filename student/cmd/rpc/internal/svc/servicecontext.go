package svc

import (
	"TDS-backend/common/mongox"
	"TDS-backend/student/cmd/rpc/internal/config"
	"TDS-backend/student/model"
)

type ServiceContext struct {
	Config       config.Config
	StudentModel model.IStudentModel
	GroupModel   model.IGroupModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		StudentModel: model.NewStudentModel(mongox.ConnectMongoDB(c.Mongodb)),
		GroupModel:   model.NewGroupModel(mongox.ConnectMongoDB(c.Mongodb)),
	}
}
