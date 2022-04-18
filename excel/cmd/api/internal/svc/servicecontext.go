package svc

import (
	"TDS-backend/classroom/cmd/rpc/classroomservice"
	"TDS-backend/common/interceptor"
	"TDS-backend/excel/cmd/api/internal/config"
	"TDS-backend/student/cmd/rpc/studentservice"
	"TDS-backend/teacher/cmd/rpc/teacherservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	TeacherRpc   teacherservice.TeacherService
	ClassroomRpc classroomservice.ClassroomService
	StudentRpc   studentservice.StudentService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		TeacherRpc:   teacherservice.NewTeacherService(zrpc.MustNewClient(c.TeacherRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
		ClassroomRpc: classroomservice.NewClassroomService(zrpc.MustNewClient(c.ClassroomRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
		StudentRpc:   studentservice.NewStudentService(zrpc.MustNewClient(c.StudentRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorinterceptor))),
	}
}
