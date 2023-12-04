package response

import (
	"github.com/EasyCode-Platform/EasyCode_Backend/src/entities"
)

type AppListResponse struct {
	AppList []*entities.App
}

func (impl *AppListResponse) ExportForFeedback() interface{} {
	return impl.AppList
}

func NewAppListResponse(appList []*entities.App) *AppListResponse {
	return &AppListResponse{
		AppList: appList,
	}
}

type AppResponse struct {
	App *entities.App
}

func (impl *AppResponse) ExportForFeedback() interface{} {
	return impl.App
}

func NewAppResponse(app *entities.App) *AppResponse {
	return &AppResponse{
		App: app,
	}
}

type GetAppResponse struct {
	AppsData []entities.AppData `json:"appsdata"`
}

type CreateTableResponse struct {
	Table entities.Table `json:"table"`
}

type RenameTableResponse struct {
	Table entities.Table `json:"table"`
}

type GetTableDataResponse struct {
	TableData entities.TableData
}
