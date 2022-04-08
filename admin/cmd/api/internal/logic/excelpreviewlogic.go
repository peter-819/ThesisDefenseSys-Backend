package logic

import (
	"context"

	"TDS-backend/admin/cmd/api/internal/svc"
	"TDS-backend/admin/cmd/api/internal/types"
	"TDS-backend/common/errorx"
	"TDS-backend/excel/cmd/rpc/types/excel"

	"github.com/zeromicro/go-zero/core/logx"
	"mime/multipart"
	"io"
)

type ExcelpreviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	Files []*multipart.FileHeader
}

func NewExcelpreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExcelpreviewLogic {
	return ExcelpreviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExcelpreviewLogic) Excelpreview() (resp *types.ExcelPreviewReply, err error) {
	resp = &types.ExcelPreviewReply{}
	if l.ctx.Value("role") != "Admin" {
		return resp, errorx.NewDefaultError("没有权限")
	}
	file, err := l.Files[0].Open()
	defer file.Close()
	if err != nil {
		return resp, errorx.NewDefaultError("文件打开错误")
	}
	// data := make([]byte, FileMaxLength)
	data, err := io.ReadAll(file) //TODO
	if err != nil {
		return resp, errorx.NewDefaultError("文件读取错误")
	}

	stream, err := l.svcCtx.ExcelRpc.ParseExcel(l.ctx)
	stream.Send(&excel.ExcelRequest{
		Info: &excel.ExcelInfo{
			RowMin: 0,
			RowMax: 10,
			Cols:[]int32{0,1,2,3},
		},
	})
	excel.ChunkerSrv(data).ChunkerSend(stream)
	res, err := stream.CloseAndRecv()
	if err != nil {
		return resp, errorx.NewDefaultError("rpc 接收错误: " + err.Error())
	}
	for _, row := range res.Rows{
		currentRow := []string{}
		for _, element := range row.Elements {
			currentRow = append(currentRow, element)
		}
		resp.Data = append(resp.Data, currentRow)
	}
	return resp, nil
	
}
