package handler

import (
	"net/http"

	"TDS-backend/route/cmd/api/internal/logic"
	"TDS-backend/route/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func routemenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewRoutemenuLogic(r.Context(), svcCtx)
		resp, err := l.Routemenu()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
