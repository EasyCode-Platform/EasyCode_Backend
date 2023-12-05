package storage

import (
	"github.com/EasyCode-Platform/EasyCode_Backend/src/model"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/utils/config"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
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
func (impl *AppStorage) CreateApp(app *model.App) (int, error) {
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
func (impl *AppStorage) RetrieveAllApp(teamId int) ([]*model.App, error) {
	var apps []*model.App
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
func (impl *AppStorage) RetrieveAppByUid(uid int) (*model.App, error) {
	var app *model.App
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

func GetApps(logger *zap.SugaredLogger) ([]model.AppData, error) {
	// 获取配置实例
	cfg := config.GetInstance()

	// 创建 PostgresqlStorage 实例
	storage, err := NewPostgreStorage(cfg, logger)
	if err != nil {
		log.Println("Error creating postgresql entities:", err)
		return nil, err
	}

	// 使用 PostgresqlStorage 获取应用数据
	apps, err := storage.GetAppsData()
	if err != nil {
		log.Println("Error getting apps:", err)
		return nil, err
	}

	return apps, nil
}

func CreateNewTable(aid uuid.UUID, logger *zap.SugaredLogger) (*model.Table, error) {
	// 假设 config 包含全局配置
	// 获取配置实例
	cfg := config.GetInstance()

	// 创建 PostgresqlStorage 实例
	storage, err := NewPostgreStorage(cfg, logger)
	if err != nil {
		log.Println("Error creating postgresql entities:", err)
		return nil, err
	}

	// 使用 PostgresqlStorage 获取应用数据
	table, err := storage.CreateNewTable(aid)
	if err != nil {
		log.Println("Error creating table:", err)
		return nil, err
	}

	return table, nil
}

func RenameTable(tid uuid.UUID, name string, logger *zap.SugaredLogger) (*model.Table, error) {
	// 假设 config 包含全局配置
	// 获取配置实例
	cfg := config.GetInstance()

	// 创建 PostgresqlStorage 实例
	storage, err := NewPostgreStorage(cfg, logger)
	if err != nil {
		log.Println("Error creating postgresql entities:", err)
		return nil, err
	}

	// 使用 PostgresqlStorage 获取应用数据
	table, err := storage.RenameTable(tid, name)
	if err != nil {
		log.Println("Error renaming table:", err)
		return nil, err
	}

	return table, nil
}

func DeleteTable(tid uuid.UUID, logger *zap.SugaredLogger) error {
	// 假设 config 包含全局配置
	// 获取配置实例
	cfg := config.GetInstance()

	// 创建 PostgresqlStorage 实例
	storage, err := NewPostgreStorage(cfg, logger)
	if err != nil {
		log.Println("Error creating postgresql entities:", err)
		return err
	}

	// 使用 PostgresqlStorage 获取应用数据
	err = storage.DeleteTable(tid)
	if err != nil {
		log.Println("Error renaming table:", err)
		return err
	}

	return nil
}

func GetTableData(tid uuid.UUID, logger *zap.SugaredLogger) (model.TableData, error) {
	// 假设 config 包含全局配置
	// 获取配置实例
	cfg := config.GetInstance()

	// 创建 PostgresqlStorage 实例
	storage, err := NewPostgreStorage(cfg, logger)
	if err != nil {
		log.Println("Error creating postgresql entities:", err)
		return model.TableData{}, err
	}

	// 使用 PostgresqlStorage 获取应用数据
	tableData, err := storage.GetTableData(tid)
	if err != nil {
		log.Println("Error renaming table:", err)
		return model.TableData{}, err
	}

	return tableData, nil
}
