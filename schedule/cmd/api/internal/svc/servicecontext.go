package svc

import (
	"TDS-backend/schedule/cmd/api/internal/config"
	"TDS-backend/schedule/model"
	"TDS-backend/common/mongox"	
	"TDS-backend/teacher/cmd/rpc/teacherservice"
	"TDS-backend/classroom/cmd/rpc/classroomservice"
	"TDS-backend/student/cmd/rpc/studentservice"
    "github.com/zeromicro/go-zero/zrpc"
	"TDS-backend/common/interceptor"
)

type ServiceContext struct {
	Config config.Config
    TeacherRpc teacherservice.TeacherService
	ClassroomRpc classroomservice.ClassroomService
	StudentRpc studentservice.StudentService
	DefenseModel model.IDefenseModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DefenseModel: model.NewDefenseModel(mongox.ConnectMongoDB(c.Mongodb)),
		TeacherRpc: teacherservice.NewTeacherService(zrpc.MustNewClient(c.TeacherRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
		ClassroomRpc: classroomservice.NewClassroomService(zrpc.MustNewClient(c.ClassroomRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
		StudentRpc: studentservice.NewStudentService(zrpc.MustNewClient(c.StudentRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
	}
}
