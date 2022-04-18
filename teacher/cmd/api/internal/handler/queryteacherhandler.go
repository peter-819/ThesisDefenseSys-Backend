package handler

import (
	"net/http"

	"TDS-backend/teacher/cmd/api/internal/logic"
	"TDS-backend/teacher/cmd/api/internal/svc"
	"TDS-backend/teacher/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryTeacherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryTeacherReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewQueryTeacherLogic(r.Context(), svcCtx)
		resp, err := l.QueryTeacher(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
