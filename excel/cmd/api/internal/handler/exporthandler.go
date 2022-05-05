package handler

import (
	"net/http"

	"TDS-backend/excel/cmd/api/internal/logic"
	"TDS-backend/excel/cmd/api/internal/svc"
	"TDS-backend/excel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ExportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExportExcelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		l := logic.NewExportLogic(r.Context(), svcCtx)
		_, err := l.Export(&req, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
