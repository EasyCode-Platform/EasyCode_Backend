package model

import (
	"time"
)

func NewApp(appName string, teamID int, modifyUserID int, ComponentId string) *App {
	return &App{
		TeamID:      teamID,
		Name:        appName,
		ComponentId: ComponentId,
		Config:      NewAppConfig().ExportToJSONString(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}
