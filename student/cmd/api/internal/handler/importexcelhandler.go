package handler

import (
	"net/http"
	"TDS-backend/common/excelx"

	"TDS-backend/student/cmd/api/internal/logic"
	"TDS-backend/student/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)
const (
	defaultMultipartMemory = 32 << 20 // 32 MB
)
func ImportExcelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
			httpx.Error(w, err)
			return
		}
		
		l := logic.NewImportExcelLogic(r.Context(), svcCtx)
		l.Info = excelx.ParseExcelInfo(r)

		resp, err := l.ImportExcel()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}