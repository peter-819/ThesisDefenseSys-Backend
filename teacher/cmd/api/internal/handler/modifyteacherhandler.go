package handler

import (
	"net/http"

	"TDS-backend/teacher/cmd/api/internal/logic"
	"TDS-backend/teacher/cmd/api/internal/svc"
	"TDS-backend/teacher/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ModifyTeacherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyTeacherReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewModifyTeacherLogic(r.Context(), svcCtx)
		err := l.ModifyTeacher(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
