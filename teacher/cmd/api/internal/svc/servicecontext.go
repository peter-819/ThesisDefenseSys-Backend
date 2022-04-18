package svc

import (
	"TDS-backend/teacher/cmd/api/internal/config"
	"TDS-backend/teacher/model"
	"TDS-backend/common/mongox"	
)

type ServiceContext struct {
	Config config.Config
	TeacherModel model.ITeacherModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		TeacherModel: model.NewTeacherModel(mongox.ConnectMongoDB(c.Mongodb)),
	}
}
