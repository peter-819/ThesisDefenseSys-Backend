package logic

import (
	"TDS-backend/common/errorx"
	"TDS-backend/common/excelx"
	"TDS-backend/excel/cmd/rpc/types/excel"
	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"TDS-backend/student/cmd/rpc/types/student"
	"context"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportExcelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	Info *excelx.ExcelInfo
}

func NewImportExcelLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportExcelLogic {
	return ImportExcelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Pair struct {
	Key   string
	Value int
}

func (l *ImportExcelLogic) ImportExcel() (resp *types.QueryStudentsReply, err error) {
	resp = &types.QueryStudentsReply{}
	file, err := l.Info.Files[0].Open()
	defer file.Close()
	if err != nil {
		return resp, errorx.NewDefaultError("文件打开错误")
	}
	data, err := io.ReadAll(file) //TODO
	if err != nil {
		return resp, errorx.NewDefaultError("文件读取错误")
	}

	stream, err := l.svcCtx.ExcelRpc.ParseExcel(l.ctx)
	pairs := []Pair{}
	cols := []int32{}
	for k, v := range l.Info.Cols {
		cols = append(cols, int32(v))
		pairs = append(pairs, Pair{Key: k, Value: v})
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})
	lastCol := -1
	current := -1
	for _, v := range pairs {
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
			Cols:   cols,
		},
	})
	excel.ChunkerSrv(data).ChunkerSend(stream)
	res, err := stream.CloseAndRecv()
	if err != nil {
		return resp, errorx.NewDefaultError("rpc 接收错误: " + err.Error())
	}

	fmt.Println(`rpcRes: `, res)
	for _, rowInfo := range res.Rows {
		stuInfo := &student.Student{}
		row := rowInfo.Elements
		if v, ok := l.Info.Cols["IdCol"]; ok && row[v] != "" {
			stuInfo.Id = row[v]
			fmt.Println("student id: ", row[v])
		} else {
			return resp, errorx.NewDefaultError("ID列号错误或有空值!")
		}

		if v, ok := l.Info.Cols["NameCol"]; ok && row[v] != "" {
			stuInfo.Name = row[v]
			fmt.Println("student name: ", row[v])
		} else {
			return resp, errorx.NewDefaultError("姓名列号错误或有空值!")
		}

		if v, ok := l.Info.Cols["PaperTitleCol"]; ok && row[v] != "" {
			stuInfo.PaperTitle = row[v]
			fmt.Println("paper id: ", row[v])
		} else {
			return resp, errorx.NewDefaultError("论文名列号错误或有空值!")
		}

		if v, ok := l.Info.Cols["PaperKeywordsCol"]; ok && row[v] != "" {
			keywords := strings.Split(row[v], ",")
			fmt.Println("paper keywords: ", keywords)
			stuInfo.PaperKeywords = keywords
		} else {
			return resp, errorx.NewDefaultError("论文关键字列号错误或有空值!")
		}
		_, err = l.svcCtx.StudentRpc.AddStudent(l.ctx, &student.AddStudentRequest{
			Student: stuInfo,
		})
		if err != nil {
			return resp, err
		}
	}
	return resp, nil
}
