package logic

import (
	"context"
	"fmt"
	"io"

	"TDS-backend/common/errorx"
	"TDS-backend/common/timex"
	"TDS-backend/excel/cmd/api/internal/svc"
	"TDS-backend/excel/cmd/api/internal/types"
	"TDS-backend/schedule/cmd/rpc/types/defense"
	"TDS-backend/student/cmd/rpc/types/student"
	"TDS-backend/teacher/cmd/rpc/types/teacher"

	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type ExportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportLogic {
	return &ExportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *ExportLogic) GetRow(rowChan chan []interface{}, errChan chan error) {
	defenseRes, err := l.svcCtx.DefenseRpc.QueryAllDefenses(l.ctx, &defense.EmptyRequest{})
	if err != nil {
		errChan <- err
		return
	}
	number := 0
	headerRow := []interface{}{"序号", "学号", "学生姓名", "导师姓名", "答辩开始时间", "答辩结束时间", "论文", "委员一", "委员二", "委员三", "委员四"}
	rowChan <- headerRow
	for _, d := range defenseRes.Defenses {
		teaRow := make([]interface{}, 0)
		for _, tea := range d.Committee {
			committeeInfo, err := l.svcCtx.TeacherRpc.QueryTeacher(l.ctx, &teacher.QueryTeacherRequest{
				Id: tea.TeacherId,
			})
			if err != nil {
				errChan <- err
				return
			}
			teaRow = append(teaRow, committeeInfo.Teacher.Name)
		}
		for _, stu := range d.Students {
			row := make([]interface{}, 0)
			stuInfo, err := l.svcCtx.StudentRpc.QueryStudent(l.ctx, &student.QueryStudentRequest{
				Id: stu.StudentId,
			})
			if err != nil {
				errChan <- err
				return
			}
			mentorInfo, err := l.svcCtx.TeacherRpc.QueryTeacher(l.ctx, &teacher.QueryTeacherRequest{
				Id: stuInfo.Student.Mentor,
			})
			if err != nil {
				errChan <- err
				return
			}

			// ID
			number++
			row = append(row, fmt.Sprint(number))
			//学号
			row = append(row, stu.StudentId)
			//学生姓名
			row = append(row, stu.StudentName)
			//导师姓名
			row = append(row, mentorInfo.Teacher.Name)
			start_time, _ := timex.GStringToTime(d.StartTime)
			end_time, _ := timex.GStringToTime(d.EndTime)
			row = append(row, start_time.Format("01-02 15:04"))
			row = append(row, end_time.Format("01-02 15:04"))
			row = append(row, stuInfo.Student.PaperTitle)
			row = append(row, teaRow...)
			rowChan <- row
		}
	}
	rowChan <- nil
}
func (l *ExportLogic) GetDefenseForm() ([][]interface{}, error) {
	rowChan := make(chan []interface{})
	errChan := make(chan error)
	go l.GetRow(rowChan, errChan)
	form := make([][]interface{}, 0)
	for {
		select {
		case err := <-errChan:
			return nil, err
		case row := <-rowChan:
			if row == nil {
				return form, nil
			}
			fmt.Println(row)
			form = append(form, row)
		}
	}
}
func setColWidth(sw *excelize.StreamWriter) error {
	//colName := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
	colWidth := []float64{6.3, 15.3, 9.3, 12.9, 13, 13, 9.3, 8.4, 8.4, 8.4, 8.4}
	for i, w := range colWidth {
		if err := sw.SetColWidth(i+1, i+1, w); err != nil {
			return err
		}
	}
	return nil
}
func (l *ExportLogic) Export(req *types.ExportExcelReq, fw io.Writer) (resp *types.ExportExcelResp, err error) {
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		return nil, errorx.NewDefaultError("创建表格文件失败")
	}
	if err = setColWidth(streamWriter); err != nil {
		return nil, errorx.NewDefaultError("修改列宽失败")
	}
	form, err := l.GetDefenseForm()
	fmt.Println("form: ", form)
	if err != nil {
		return nil, err
	}
	for rowIndex, row := range form {
		cell, _ := excelize.CoordinatesToCellName(1, rowIndex+1)

		if err := streamWriter.SetRow(cell, row); err != nil {
			return nil, errorx.NewDefaultError("写入行失败")
		}
	}
	if err := streamWriter.Flush(); err != nil {
		return nil, errorx.NewDefaultError("保存表格失败")
	}
	if _, err := file.WriteTo(fw); err != nil {
		return nil, errorx.NewDefaultError("写入http失败")
	}
	return nil, nil
}
