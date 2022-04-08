package handler

import (
	"net/http"

	"TDS-backend/route/cmd/api/internal/logic"
	"TDS-backend/route/cmd/api/internal/svc"
	"TDS-backend/route/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func routeheaderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HeaderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRouteheaderLogic(r.Context(), svcCtx)
		resp, err := l.Routeheader(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
