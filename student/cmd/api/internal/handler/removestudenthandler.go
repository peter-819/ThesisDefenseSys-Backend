package handler

import (
	"net/http"

	"TDS-backend/student/cmd/api/internal/logic"
	"TDS-backend/student/cmd/api/internal/svc"
	"TDS-backend/student/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveStudentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveStudentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRemoveStudentLogic(r.Context(), svcCtx)
		err := l.RemoveStudent(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
