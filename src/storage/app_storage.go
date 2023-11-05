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

func (impl *AppStorage) RetrieveAllApp(teamId int) ([]*model.App, error) {
	var apps []*model.App
	if err := impl.db.Where("teamID = ?", teamId).Find(&apps).Error; err != nil {
		return nil, err
	}
	return apps, nil
}

// RetrieveApp
// @receiver impl
// @param id
// @return *model.App
// @return error
func (impl *AppStorage) RetrieveAppByUid(uid int) (*model.App, error) {
	var app *model.App
	if err := impl.db.First(&app, uid).Error; err != nil {
		impl.logger.Errorw("Failed to retrieve app by uid", "uid", uid, "error", err)
		return nil, err
	}
	return app, nil
}

func (impl *AppStorage) RetrieveAppByName(teamId int, name string) ([]*model.App, error) {
	var apps []*model.App
	if err := impl.db.Where("teamID = ? and name = ?", teamId, name).Find(&apps).Error; err != nil {
		return nil, err
	}
	return apps, nil
}

// UpdateApp
// @receiver impl
// @param app
// @return error
func (impl *AppStorage) UpdateApp(app *model.App) error {
	if err := impl.db.Save(app).Error; err != nil {
		return err
	}
	return nil
}
