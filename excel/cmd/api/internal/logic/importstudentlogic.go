package logic

import (
	"context"

	"TDS-backend/common/excelx"
	"TDS-backend/excel/cmd/api/internal/svc"
	"TDS-backend/excel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	Info *excelx.ExcelInfo
}

func NewImportStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportStudentLogic {
	return ImportStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Pair struct {
	Key   string
	Value int
}

func (l *ImportStudentLogic) ImportStudent() (resp *types.ImportExcelReply, err error) {
	// todo: add your logic here and delete this line

	return
}
