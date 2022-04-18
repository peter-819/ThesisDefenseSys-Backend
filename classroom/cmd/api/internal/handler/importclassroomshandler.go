package handler

import (
	"net/http"
	"TDS-backend/common/excelx"

	"TDS-backend/classroom/cmd/api/internal/logic"
	"TDS-backend/classroom/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)
const (
	defaultMultipartMemory = 32 << 20 // 32 MB
)
func ImportClassroomsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewImportClassroomsLogic(r.Context(), svcCtx)
		l.Info = excelx.ParseExcelInfo(r)
		
		err := l.ImportClassrooms()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
