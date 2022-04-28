package excelx

import (
	"TDS-backend/common/errorx"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type ExcelInfo struct {
	Values map[string]int
	Files  []*multipart.FileHeader
}

const (
	DefaultMultipartMemory = 32 << 20 // 32 MB
)

func ParseExcelInfo(r *http.Request) *ExcelInfo {
	info := &ExcelInfo{}
	info.Values = make(map[string]int)
	form := r.MultipartForm
	for k, v := range form.Value {
		if len(v) == 0 {
			continue
		}
		temp, err := strconv.Atoi(v[0])
		if err != nil {
			continue
		}
		info.Values[k] = temp
	}
	uploadedFiles := r.MultipartForm.File
	for _, files := range uploadedFiles {
		info.Files = append(info.Files, files...)
	}
	return info
}

func (m *ExcelInfo) ParseExcelData() (res [][]string, err error) {
	file, err := m.Files[0].Open()
	defer file.Close()
	if err != nil {
		return nil, errorx.NewDefaultError("文件打开错误")
	}

	excelFile, err := excelize.OpenReader(file)
	defer excelFile.Close()
	if err != nil {
		return nil, errorx.NewDefaultError("open excel error")
	}
	rows, err := excelFile.GetRows("Sheet1")
	if err != nil {
		return nil, errorx.NewDefaultError("get rows error")
	}
	for i, row := range rows {
		if i >= m.Values["RowMin"] && i <= m.Values["RowMax"] {
			currentRow := []string{}
			for _, element := range row {
				currentRow = append(currentRow, element)
			}
			res = append(res, currentRow)
		}
	}
	return res, nil
}
