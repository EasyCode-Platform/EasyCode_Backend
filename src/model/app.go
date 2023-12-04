package model

import (
	"github.com/EasyCode-Platform/EasyCode_Backend/src/entities"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/storage"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/utils/config"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"log"
	"time"
)

func NewApp(appName string, teamID int, modifyUserID int, ComponentId string) *entities.App {
	return &entities.App{
		TeamID:      teamID,
		Name:        appName,
		ComponentId: ComponentId,
		Config:      NewAppConfig().ExportToJSONString(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}

func GetApps(logger *zap.SugaredLogger) ([]entities.AppData, error) {
	// 获取配置实例
	cfg := config.GetInstance()

	// 创建 PostgresqlStorage 实例
	storage, err := storage.NewPostgreStorage(cfg, logger)
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

func CreateNewTable(aid uuid.UUID, logger *zap.SugaredLogger) (*entities.Table, error) {
	// 假设 config 包含全局配置
	// 获取配置实例
	cfg := config.GetInstance()

	// 创建 PostgresqlStorage 实例
	storage, err := storage.NewPostgreStorage(cfg, logger)
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

func RenameTable(tid uuid.UUID, name string, logger *zap.SugaredLogger) (*entities.Table, error) {
	// 假设 config 包含全局配置
	// 获取配置实例
	cfg := config.GetInstance()

	// 创建 PostgresqlStorage 实例
	storage, err := storage.NewPostgreStorage(cfg, logger)
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
	storage, err := storage.NewPostgreStorage(cfg, logger)
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

func GetTableData(tid uuid.UUID, logger *zap.SugaredLogger) ([]entities.Field, []entities.Record, error) {
	// 假设 config 包含全局配置
	// 获取配置实例
	cfg := config.GetInstance()

	// 创建 PostgresqlStorage 实例
	storage, err := storage.NewPostgreStorage(cfg, logger)
	if err != nil {
		log.Println("Error creating postgresql entities:", err)
		return nil, nil, err
	}

	// 使用 PostgresqlStorage 获取应用数据
	field, record, err := storage.GetTableData(tid)
	if err != nil {
		log.Println("Error renaming table:", err)
		return nil, nil, err
	}

	return field, record, nil
}
