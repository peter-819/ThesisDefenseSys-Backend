package handler

import (
	"net/http"

	"TDS-backend/classroom/cmd/api/internal/logic"
	"TDS-backend/classroom/cmd/api/internal/svc"
	"TDS-backend/classroom/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveClassroomHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveClassroomReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRemoveClassroomLogic(r.Context(), svcCtx)
		err := l.RemoveClassroom(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
