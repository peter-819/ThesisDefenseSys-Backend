package logic

import (
	"context"
	"io"
	"sort"
	"fmt"
	"TDS-backend/common/excelx"
	"TDS-backend/common/errorx"
	"TDS-backend/common/timex"
	"TDS-backend/excel/cmd/rpc/types/excel"
	"TDS-backend/classroom/model"

	"TDS-backend/classroom/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ImportClassroomsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	Info *excelx.ExcelInfo
}

func NewImportClassroomsLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportClassroomsLogic {
	return ImportClassroomsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
type Pair struct {
	Key string
	Value int
}
func (l *ImportClassroomsLogic) ImportClassrooms() error {
	file, err := l.Info.Files[0].Open()
	defer file.Close()
	if err != nil {
		return errorx.NewDefaultError("文件打开错误")
	}
	data, err := io.ReadAll(file) //TODO
	if err != nil {
		return errorx.NewDefaultError("文件读取错误")
	}

	stream, err := l.svcCtx.ExcelRpc.ParseExcel(l.ctx)
	if err != nil {
		return errorx.NewDefaultError("Excel RPC 错误")
	}
	pairs := []Pair{}
	cols := []int32{}
	for k,v := range l.Info.Cols {
		cols = append(cols, int32(v))
		pairs = append(pairs, Pair{Key:k,Value:v,})
	}
	sort.SliceStable(pairs, func(i,j int) bool {
		return pairs[i].Value < pairs[j].Value
	})
	lastCol := -1
	current := -1
	for _,v := range pairs {
		if v.Value != lastCol {
			lastCol = v.Value
			current++
		}
		l.Info.Cols[v.Key] = current
	}
	fmt.Println(`Cols: `, l.Info.Cols)
	stream.Send(&excel.ExcelRequest{
		Info: &excel.ExcelInfo{
			RowMin: int32(l.Info.RowMin),
			RowMax: int32(l.Info.RowMax),
			Cols: cols,
		},
	})
	excel.ChunkerSrv(data).ChunkerSend(stream)
	res, err := stream.CloseAndRecv()
	if err != nil {
		return errorx.NewDefaultError("rpc 接收错误: " + err.Error())
	}
	
	fmt.Println(`rpcRes: `, res)
	for _, rowInfo := range res.Rows{
		classInfo := &model.Classroom{}
		row := rowInfo.Elements

		if v, ok := l.Info.Cols["NameCol"]; ok && row[v] != "" { 
			classInfo.Name = row[v] 
			fmt.Println("classroom name: ", row[v])
		} else {
			return errorx.NewDefaultError("教室名 列错误或有空值!")
		} 

		if v, ok := l.Info.Cols["LocationCol"]; ok && row[v] != "" { 
			classInfo.Location = row[v]
			fmt.Println("classroom location: ", row[v]) 
		} else {
			return errorx.NewDefaultError("教师位置 列错误或有空值!")
		} 

		if v, ok := l.Info.Cols["StartTimeCol"]; ok && row[v] != "" { 
			start_time, err := timex.FStringToTime(row[v])
			if err != nil {
				return errorx.NewDefaultError("教室开始时间格式错误!")
			}
			fmt.Println("classroom start time: ", start_time)
			classInfo.StartTime = start_time
		} else {
			return errorx.NewDefaultError("教室开始时间列错误!")
		} 
		
		if v, ok := l.Info.Cols["EndTimeCol"]; ok && row[v] != "" { 
			end_time, err := timex.FStringToTime(row[v])
			if err != nil {
				return errorx.NewDefaultError("教室结束时间格式错误!")
			}
			fmt.Println("classroom end time: ", end_time)
			classInfo.EndTime = end_time
		} else {
			return errorx.NewDefaultError("教室开始时间列错误!")
		}
		err = l.svcCtx.ClassroomModel.InsertClassroom(classInfo)
		if err != nil {
			return err
		}
	}
	return nil
}
