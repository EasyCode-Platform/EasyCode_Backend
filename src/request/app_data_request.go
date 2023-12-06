package request

import (
	"github.com/EasyCode-Platform/EasyCode_Backend/src/model"
	"github.com/google/uuid"
)

// TableUpdateRequest 用于解析 PUT 请求的 JSON 载体
type TableUpdateRequest struct {
	Table struct {
		Name string `json:"name"`
	} `json:"table"`
}

type CreateAppDataRequest struct {
	Aid    uuid.UUID     `json:"aid"`
	Name   string        `json:"name"`
	Tables []model.Table `json:"tables"`
}

func NewCreateAppDataRequest() *CreateAppDataRequest {
	return &CreateAppDataRequest{}
}

func NewTableRenameRequest() *TableUpdateRequest {
	return &TableUpdateRequest{}
}
