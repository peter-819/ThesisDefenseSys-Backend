// Code generated by goctl. DO NOT EDIT.
package types

type ImportExcelReq struct {
	Usage string `path:"usage"`
}

type ImportResult struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ErrorMessage string `json:"error_message"`
}

type ImportExcelReply struct {
	TotalCount   int            `json:"total_count"`
	SuccessCount int            `json:"success_count"`
	ErrorResults []ImportResult `json:"error_results"`
}

type PreviewExcelReply struct {
	Data [][]string `json:"data"`
}

type ExportExcelReq struct {
	Usage string `path:"usage"`
}

type ExportExcelResp struct {
}
