package logic

import (
	"context"
	"reflect"
	"strings"

	"TDS-backend/common/errorx"
	"TDS-backend/common/excelx"
	"TDS-backend/excel/cmd/api/internal/svc"
	"TDS-backend/excel/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/studentservice"
	"TDS-backend/teacher/cmd/rpc/teacherservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	ExcelInfo *excelx.ExcelInfo
}

func NewImportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportLogic {
	return &ImportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func getProcessFunc(usage string) func(interface{}, map[string]int, []string) error {
	FieldMap := make(map[string]string)
	switch usage {
	case "student":
		FieldMap["学号"] = "Id"
		FieldMap["姓名"] = "Name"
		FieldMap["关键字"] = "PaperKeywords"
		FieldMap["论文名"] = "PaperTitle"
		FieldMap["导师"] = "Mentor"
	case "classroom":
		FieldMap["教室名"] = "Name"
		FieldMap["位置"] = "Location"
		FieldMap["开始时间"] = "StartTime"
		FieldMap["结束时间"] = "EndTime"
	case "teacher":
		FieldMap["工号"] = "Id"
		FieldMap["姓名"] = "Name"
		FieldMap["关键字"] = "PreferKeywords"
	case "user":
		FieldMap["学工号"] = "Id"
		FieldMap["姓名"] = "Name"
		FieldMap["身份"] = "Role"
		FieldMap["密码"] = "Password"
	}

	return func(reqObj interface{}, headerMap map[string]int, row []string) error {
		for k, _ := range FieldMap {
			if _, ok := headerMap[k]; !ok {
				return errorx.NewDefaultError(k + "列错误")
			}
		}

		value := reflect.ValueOf(reqObj)
		value = value.Elem()
		for k, v := range FieldMap {
			f := value.FieldByName(v)
			if f.Kind() == reflect.String {
				f.SetString(row[headerMap[k]])
			} else if f.Kind() == reflect.Slice && f.Type().Elem().Kind() == reflect.String {
				f.Set(reflect.ValueOf(strings.Split(row[headerMap[k]], ",")))
			}
		}
		return nil
	}
}
func (l *ImportLogic) Import(req *types.ImportExcelReq) (resp *types.ImportExcelReply, err error) {
	resp = &types.ImportExcelReply{}

	data, err := l.ExcelInfo.ParseExcelData()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 || len(data[0]) == 0 {
		return nil, errorx.NewDefaultError("表格信息为空")
	}
	process := getProcessFunc(req.Usage)
	headerMap := make(map[string]int)
	for i, header := range data[0] {
		headerMap[header] = i
	}

	for _, row := range data[1:] {
		switch req.Usage {
		case "classroom":
			// cla := &classroomservice.Classroom{}
			// err := process(cla, headerMap, row)
			// if err != nil {
			// 	return nil, err
			// }
			// _, err = l.svcCtx.ClassroomRpc(l.ctx, &teacherservice.AddTeacherRequest{
			// 	Teacher: tea,
			// })
			// if err != nil {
			// 	return nil, errorx.NewDefaultError("学生信息错误: " + err.Error())
			// }
		case "student":
			stu := &studentservice.Student{}
			err := process(stu, headerMap, row)
			if err != nil {
				return nil, err
			}
			_, err = l.svcCtx.StudentRpc.AddStudent(l.ctx, &studentservice.AddStudentRequest{
				Student: stu,
			})
			if err != nil {
				return nil, errorx.NewDefaultError("学生信息错误: " + err.Error())
			}
		case "teacher":
			tea := &teacherservice.Teacher{}
			err := process(tea, headerMap, row)
			if err != nil {
				return nil, err
			}
			_, err = l.svcCtx.TeacherRpc.AddTeacher(l.ctx, &teacherservice.AddTeacherRequest{
				Teacher: tea,
			})
			if err != nil {
				return nil, errorx.NewDefaultError("教师信息错误: " + err.Error())
			}
		case "user":
		default:
			return nil, errorx.NewDefaultError("错误上传方法: " + req.Usage)
		}

	}
	return resp, nil
}
