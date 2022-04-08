package logic

import (
	"context"

	"TDS-backend/excel/cmd/rpc/internal/svc"
	"TDS-backend/excel/cmd/rpc/types/excel"
	"TDS-backend/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/xuri/excelize/v2"
	"bytes"
	"io"
	"log"
)

type ParseExcelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewParseExcelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ParseExcelLogic {
	return &ParseExcelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func contains(s []int32, e int32) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func (l *ParseExcelLogic) ParseExcel(stream excel.User_ParseExcelServer) error {
	// todo: add your logic here and delete this line
	//req, err := stream.Recv()

	//chunker
	var blob []byte
	var info *excel.ExcelInfo
	for {
		c, err := stream.Recv()
		if info == nil && c.Info != nil {
			info = c.Info
		}
		if err != nil {
			if err == io.EOF {
				log.Printf("Transfer of %d bytes successful", len(blob))
				break
			}
			return errorx.NewDefaultError("rpc excel读取失败")
		}
		blob = append(blob, c.File...)
		log.Printf("接收包长:%d",len(c.File))
	}

	// if err != nil {
	// 	return errorx.NewDefaultError("rpc excel读取失败")
	// } 
	log.Printf("接收完成，包长:%d",len(blob))
	reader := bytes.NewReader(blob)
	
	excelFile, err := excelize.OpenReader(reader)
	defer excelFile.Close()
	
	if err != nil {
		return errorx.NewDefaultError("open excel error")
	}
	rows, err := excelFile.GetRows("Sheet1")
	if err != nil {
		return errorx.NewDefaultError("get rows error")
	}
	log.Printf("转换excel 长度: %d * %d",len(rows),len(rows[0]))
	res := &excel.ExcelResponse{}
	for i,row := range rows {
		if int32(i) >= info.RowMin && int32(i) <= info.RowMax{
			currentRow := &excel.RowInfo{}
			for j,element := range row {
				if contains(info.Cols, int32(j)){
					log.Println("加入元素:" + element)
					currentRow.Elements = append(currentRow.Elements, element)
				}
			}
			res.Rows = append(res.Rows, currentRow)
		}
	}
	stream.SendAndClose(res)
	return nil
}
