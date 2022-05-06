package handler

import (
	"fmt"
	"net/http"

	"TDS-backend/common/errorx"
	"TDS-backend/user/cmd/api/internal/logic"
	"TDS-backend/user/cmd/api/internal/svc"
	"TDS-backend/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func VerifyTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewVerifyTokenLogic(r.Context(), svcCtx)
		resp, err := l.VerifyToken(&req, r)
		if err != nil {
			fmt.Println("verify token error: ", err)
			if err.(*errorx.CodeError).Code == 401 {
				httpx.WriteJson(w, http.StatusUnauthorized, err)
			} else if err.(*errorx.CodeError).Code == 404 {
				httpx.WriteJson(w, http.StatusNotFound, err)
			} else {
				httpx.Error(w, err)
			}
		} else {
			fmt.Println("verify token resp: ", resp)
			w.Header().Set("X-User-Id", resp.UserId)
			w.Header().Set("X-User-Role", resp.UserRole)
			httpx.OkJson(w, resp)
		}
	}
}
