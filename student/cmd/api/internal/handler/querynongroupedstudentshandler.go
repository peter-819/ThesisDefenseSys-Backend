package handler

import (
	"net/http"

	"TDS-backend/student/cmd/api/internal/logic"
	"TDS-backend/student/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryNongroupedStudentsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewQueryNongroupedStudentsLogic(r.Context(), svcCtx)
		resp, err := l.QueryNongroupedStudents()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
