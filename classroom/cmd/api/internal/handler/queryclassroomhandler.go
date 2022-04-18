package handler

import (
	"net/http"

	"TDS-backend/classroom/cmd/api/internal/logic"
	"TDS-backend/classroom/cmd/api/internal/svc"
	"TDS-backend/classroom/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryClassroomHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryClassroomReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewQueryClassroomLogic(r.Context(), svcCtx)
		resp, err := l.QueryClassroom(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
