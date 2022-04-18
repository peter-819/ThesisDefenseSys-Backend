package handler

import (
	"net/http"

	"TDS-backend/classroom/cmd/api/internal/logic"
	"TDS-backend/classroom/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryAllClassroomsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewQueryAllClassroomsLogic(r.Context(), svcCtx)
		resp, err := l.QueryAllClassrooms()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
