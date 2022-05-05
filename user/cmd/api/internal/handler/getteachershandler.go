package handler

import (
	"net/http"

	"TDS-backend/user/cmd/api/internal/logic"
	"TDS-backend/user/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTeachersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetTeachersLogic(r.Context(), svcCtx)
		resp, err := l.GetTeachers()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
