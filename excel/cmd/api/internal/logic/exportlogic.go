package logic

import (
	"context"
	"io"
	"math/rand"

	"TDS-backend/common/errorx"
	"TDS-backend/excel/cmd/api/internal/svc"
	"TDS-backend/excel/cmd/api/internal/types"

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

func (l *ExportLogic) Export(req *types.ExportExcelReq, fw io.Writer) (resp *types.ExportExcelResp, err error) {
	// todo: add your logic here and delete this line
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		return nil, errorx.NewDefaultError("创建表格文件失败")
	}
	styleID, err := file.NewStyle(&excelize.Style{Font: &excelize.Font{Color: "#777777"}})
	if err != nil {
		return nil, errorx.NewDefaultError("创建样式失败")
	}
	if err := streamWriter.SetRow("A1", []interface{}{
		excelize.Cell{StyleID: styleID, Value: "Data"}}); err != nil {
		return nil, errorx.NewDefaultError("设置表头失败")
	}
	for rowID := 2; rowID <= 102; rowID++ {
		row := make([]interface{}, 50)
		for colID := 0; colID < 50; colID++ {
			row[colID] = rand.Intn(640000)
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
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
