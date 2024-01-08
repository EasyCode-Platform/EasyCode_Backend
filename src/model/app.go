package model

import (
	"github.com/google/uuid"
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

func NewAppV2(appName string, ComponentId string) *App {
	return &App{
		Name:        appName,
		AID:         uuid.New(),
		ComponentId: ComponentId,
		Config:      NewAppConfig().ExportToJSONString(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}
