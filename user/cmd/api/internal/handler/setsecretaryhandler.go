package handler

import (
	"net/http"

	"TDS-backend/user/cmd/api/internal/logic"
	"TDS-backend/user/cmd/api/internal/svc"
	"TDS-backend/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetSecretaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetSecretaryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSetSecretaryLogic(r.Context(), svcCtx)
		err := l.SetSecretary(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
