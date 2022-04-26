package handler

import (
	"net/http"

	"TDS-backend/teacher/cmd/api/internal/logic"
	"TDS-backend/teacher/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryAllTeachersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewQueryAllTeachersLogic(r.Context(), svcCtx)
		resp, err := l.QueryAllTeachers()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
