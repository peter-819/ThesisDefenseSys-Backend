package handler

import (
	"net/http"
	"strings"
	"strconv"
	
	"TDS-backend/admin/cmd/api/internal/logic"
	"TDS-backend/admin/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)
func parseRegisterInfo(r *http.Request, info *logic.RegisterInfo) {
	form := r.MultipartForm
	// info_type := reflect.TypeOf(info).Elem()
	// info_value := reflect.ValueOf(info).Elem()
	// for k:=0; k < info_type.NumField(); k++ {
	// 	f := info_type.Field(k)
	// 	v := info_value.Field(k)
	// 	if len(form.Value[f.Name]) == 0 {		
	// 		switch v.Type().Kind() {
	// 		case reflect.Int32:
	// 			v.SetInt(-1)
	// 		}
	// 		continue
	// 	} 
	// 	switch v.Type().Kind() {
	// 	case reflect.String:
	// 		v.SetString(form.Value[f.Name][0])		
	// 	case reflect.Int32:
	// 		num, err := strconv.Atoi(form.Value[f.Name][0])
	// 		if err != nil {
	// 			num = -1
	// 		}
	// 		num64:=int64(num)
	// 		v.SetInt(num64)
	// 	}
	// }
	info.Cols = make(map[string]int32)
	for k,v := range form.Value {
		if len(v) == 0 {
			continue
		}
		switch k {
		case "DefaultRole":
			info.DefaultRole = v[0]
		case "DefaultPassword":
			info.DefaultPassword = v[0]
		}
		
		temp, err := strconv.Atoi(v[0])
		var vint int32
		if err != nil {
			vint = -1
		}
		vint = int32(temp)

		if strings.HasSuffix(k, "Col") {
			info.Cols[k] = vint
			continue
		}
		switch k {
		case "RowMin":
			info.RowMin = vint
		case "RowMax":
			info.RowMax = vint
		}
	}
	
}
func registerbatchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterbatchLogic(r.Context(), svcCtx)
		parseRegisterInfo(r, &l.Info)
		uploadedFiles := r.MultipartForm.File
		for _, files := range uploadedFiles {
			l.Files = append(l.Files, files...)
		}
	
		resp, err := l.Registerbatch()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
