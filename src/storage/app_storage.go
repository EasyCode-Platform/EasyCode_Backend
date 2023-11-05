package storage

import (
	"github.com/EasyCode-Platform/EasyCode_Backend/src/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// AppStorage
type AppStorage struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

// NewAppStorage
// @param logger
// @param db
// @return *AppStorage
func NewAppStorage(logger *zap.SugaredLogger, db *gorm.DB) *AppStorage {
	return &AppStorage{logger: logger,
		db: db}
}

// CreateApp
// @receiver impl
// @param app
// @return int
// @return error
func (impl *AppStorage) CreateApp(app *model.App) (int, error) {
	if err := impl.db.Create(app).Error; err != nil {
		return 0, err
	}
	return app.ID, nil
}
