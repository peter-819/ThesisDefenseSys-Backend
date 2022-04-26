package handler

import (
	"net/http"

	"TDS-backend/student/cmd/api/internal/logic"
	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryGroupContentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryGroupReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewQueryGroupContentLogic(r.Context(), svcCtx)
		resp, err := l.QueryGroupContent(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}