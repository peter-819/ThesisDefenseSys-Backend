package logic

import (
	"context"

	"TDS-backend/excel/cmd/api/internal/svc"
	"TDS-backend/excel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportClassroomsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportClassroomsLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportClassroomsLogic {
	return ImportClassroomsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportClassroomsLogic) ImportClassrooms() (resp *types.ImportExcelReply, err error) {
	// todo: add your logic here and delete this line

	return
}
