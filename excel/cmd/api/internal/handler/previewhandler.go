package handler

import (
	"net/http"

	"TDS-backend/common/excelx"
	"TDS-backend/excel/cmd/api/internal/logic"
	"TDS-backend/excel/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PreviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(excelx.DefaultMultipartMemory); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewPreviewLogic(r.Context(), svcCtx)
		l.ExcelInfo = excelx.ParseExcelInfo(r)
		resp, err := l.Preview()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
