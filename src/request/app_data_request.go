package request

// TableUpdateRequest 用于解析 PUT 请求的 JSON 载体
type TableUpdateRequest struct {
	Table struct {
		Name string `json:"name"`
	} `json:"table"`
}

func NewTableRenameRequest() *TableUpdateRequest {
	return &TableUpdateRequest{}
}
