package handler

import (
	"net/http"

	"TDS-backend/user/cmd/api/internal/logic"
	"TDS-backend/user/cmd/api/internal/svc"
	"TDS-backend/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func setteacherinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetTeacherInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSetteacherinfoLogic(r.Context(), svcCtx)
		err := l.Setteacherinfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
