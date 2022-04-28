package logic

import (
	"context"

	"TDS-backend/common/excelx"
	"TDS-backend/excel/cmd/api/internal/svc"
	"TDS-backend/excel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PreviewLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	ExcelInfo *excelx.ExcelInfo
}

func NewPreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreviewLogic {
	return &PreviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PreviewLogic) Preview() (resp *types.PreviewExcelReply, err error) {
	resp = &types.PreviewExcelReply{}
	resp.Data, err = l.ExcelInfo.ParseExcelData()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
