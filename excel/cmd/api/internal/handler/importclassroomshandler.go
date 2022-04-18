package handler

import (
	"net/http"

	"TDS-backend/excel/cmd/api/internal/logic"
	"TDS-backend/excel/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImportClassroomsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewImportClassroomsLogic(r.Context(), svcCtx)
		resp, err := l.ImportClassrooms()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
