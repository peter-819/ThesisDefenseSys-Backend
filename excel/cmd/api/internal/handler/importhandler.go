package handler

import (
	"net/http"

	"TDS-backend/common/excelx"
	"TDS-backend/excel/cmd/api/internal/logic"
	"TDS-backend/excel/cmd/api/internal/svc"
	"TDS-backend/excel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportExcelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		if err := r.ParseMultipartForm(excelx.DefaultMultipartMemory); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewImportLogic(r.Context(), svcCtx)
		l.ExcelInfo = excelx.ParseExcelInfo(r)
		resp, err := l.Import(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
