package handler

import (
	"net/http"

	"TDS-backend/schedule/cmd/api/internal/logic"
	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func queryteacherscheduleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTScheduleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewQueryteacherscheduleLogic(r.Context(), svcCtx)
		resp, err := l.Queryteacherschedule(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
