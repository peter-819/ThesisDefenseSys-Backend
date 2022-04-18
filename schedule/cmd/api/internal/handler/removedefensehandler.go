package handler

import (
	"net/http"

	"TDS-backend/schedule/cmd/api/internal/logic"
	"TDS-backend/schedule/cmd/api/internal/svc"
	"TDS-backend/schedule/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveDefenseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveDefenseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRemoveDefenseLogic(r.Context(), svcCtx)
		err := l.RemoveDefense(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
