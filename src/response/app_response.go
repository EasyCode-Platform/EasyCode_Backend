package response

import (
	"github.com/EasyCode-Platform/EasyCode_Backend/src/model"
	"github.com/google/uuid"
)

type AppListResponse struct {
	AppList []*model.App
}

func (impl *AppListResponse) ExportForFeedback() interface{} {
	return impl.AppList
}

func NewAppListResponse(appList []*model.App) *AppListResponse {
	return &AppListResponse{
		AppList: appList,
	}
}

type AppResponse struct {
	App *model.App
}

func (impl *AppResponse) ExportForFeedback() interface{} {
	return impl.App
}

func NewAppResponse(app *model.App) *AppResponse {
	return &AppResponse{
		App: app,
	}
}

type GetAppResponse struct {
	AppsData []model.AppData `json:"appsdata"`
}

type CreateTableResponse struct {
	Table model.Table `json:"table"`
}

type RenameTableResponse struct {
	Table model.Table `json:"table"`
}

type GetTableDataResponse struct {
	TableData model.TableData
}

type CreateAppResponse struct {
	Aid uuid.UUID `json:"aid"`
}
