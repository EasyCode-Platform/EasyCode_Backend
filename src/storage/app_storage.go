package storage

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppStorage struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

func NewAppStorage(logger *zap.SugaredLogger, db *gorm.DB) *AppStorage {
	return &AppStorage{logger: logger,
		db: db}
}

func (appStorage *AppStorage) CreateApp()
