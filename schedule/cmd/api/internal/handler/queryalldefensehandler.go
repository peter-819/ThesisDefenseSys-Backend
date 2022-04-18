package handler

import (
	"net/http"

	"TDS-backend/schedule/cmd/api/internal/logic"
	"TDS-backend/schedule/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func queryAllDefenseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewQueryAllDefenseLogic(r.Context(), svcCtx)
		resp, err := l.QueryAllDefense()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
