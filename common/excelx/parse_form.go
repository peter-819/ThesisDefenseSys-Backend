package excelx
import (
	"net/http"
	"strings"
	"strconv"
	"mime/multipart"
)
type ExcelInfo struct {
	RowMin int
	RowMax int

	Cols map[string]int
	Files []*multipart.FileHeader
}

func ParseExcelInfo(r *http.Request) *ExcelInfo {
	info := &ExcelInfo{}
	info.Cols = make(map[string]int)
	form := r.MultipartForm
	for k, v := range form.Value {
		if len(v) == 0 {
			continue
		}
		temp, err := strconv.Atoi(v[0])
		if err != nil {
			continue
		}
		if strings.HasSuffix(k, "Col") {
			info.Cols[k] = temp
		} else if k == "RowMin" {
			info.RowMin = temp
		} else if k == "RowMax" {
			info.RowMax = temp
		}
	}		
	uploadedFiles := r.MultipartForm.File
	for _, files := range uploadedFiles {
		info.Files = append(info.Files, files...)
	}
	return info
}