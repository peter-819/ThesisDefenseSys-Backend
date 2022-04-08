package handler

import (
	"net/http"

	"TDS-backend/admin/cmd/api/internal/logic"
	"TDS-backend/admin/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	defaultMultipartMemory = 32 << 20 // 32 MB
)
func excelpreviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewExcelpreviewLogic(r.Context(), svcCtx)

		uploadedFiles := r.MultipartForm.File
		for _, files := range uploadedFiles {
			l.Files = append(l.Files, files...)
		}
	
		resp, err := l.Excelpreview()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
