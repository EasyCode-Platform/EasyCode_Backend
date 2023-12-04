package storage

import (
	"github.com/EasyCode-Platform/EasyCode_Backend/src/entities"
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
// @return int app.id
// @return error
func (impl *AppStorage) CreateApp(app *entities.App) (int, error) {
	if err := impl.db.Create(app).Error; err != nil {
		return 0, err
	}
	return app.ID, nil
}

// RetrieveAllApp
// @receiver impl
// @param teamId
// @return []*model.App
// @return error
func (impl *AppStorage) RetrieveAllApp(teamId int) ([]*entities.App, error) {
	var apps []*entities.App
	if err := impl.db.Where("team_id = ?", teamId).Find(&apps).Error; err != nil {
		return nil, err
	}
	return apps, nil
}

// RetrieveApp
// @receiver impl
// @param id
// @return *model.App
// @return error
func (impl *AppStorage) RetrieveAppByUid(uid int) (*entities.App, error) {
	var app *entities.App
	if err := impl.db.First(&app, uid).Error; err != nil {
		impl.logger.Errorw("Failed to retrieve app by uid", "uid", uid, "error", err)
		return nil, err
	}
	return app, nil
}

// RetrieveAppByName
// @receiver impl
// @param teamId
// @param name
// @return []*model.App
// @return error
func (impl *AppStorage) RetrieveAppByName(teamId int, name string) ([]*entities.App, error) {
	var apps []*entities.App
	if err := impl.db.Where("teamID = ? and name = ?", teamId, name).Find(&apps).Error; err != nil {
		return nil, err
	}
	return apps, nil
}

// UpdateApp
// @receiver impl
// @param app
// @return error
func (impl *AppStorage) UpdateApp(app *entities.App) error {
	if err := impl.db.Save(app).Error; err != nil {
		return err
	}
	return nil
}
