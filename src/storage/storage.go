package storage

import (
	"github.com/EasyCode-Platform/EasyCode_Backend/src/driver/postgres"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/utils/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Storage struct {
	ComponentStorage *ComponentStorage
	AppStorage       *AppStorage
	// ActionStorage      *ActionStorage
	// AppSnapshotStorage *AppSnapshotStorage
	// KVStateStorage     *KVStateStorage
	// ResourceStorage    *ResourceStorage
	// SetStateStorage    *SetStateStorage
	// TreeStateStorage   *TreeStateStorage
}

// NewStorage
// @param postgresDriver
// @param mongodb
// @param logger
// @return *Storage
func NewStorage(postgresDriver *gorm.DB, mongodb *mongo.Database, logger *zap.SugaredLogger) *Storage {
	return &Storage{
		ComponentStorage: NewComponentStorage(logger, mongodb),
		AppStorage:       NewAppStorage(logger, postgresDriver),
		// ActionStorage:      NewActionStorage(logger, postgresDriver),
		// AppSnapshotStorage: NewAppSnapshotStorage(logger, postgresDriver),
		// KVStateStorage:     NewKVStateStorage(logger, postgresDriver),
		// ResourceStorage:    NewResourceStorage(logger, postgresDriver),
		// SetStateStorage:    NewSetStateStorage(logger, postgresDriver),
		// TreeStateStorage:   NewTreeStateStorage(logger, postgresDriver),
	}
}

func NewPostgreStorage(cfg *config.Config, logger *zap.SugaredLogger) (*PostgresqlStorage, error) {
	// 使用 NewPostgresConnectionByGlobalConfig 方法创建数据库连接
	db, err := postgres.NewPostgresConnectionByGlobalConfig(cfg, logger)
	if err != nil {
		logger.Errorw("Error opening database", "error", err)
		return nil, err
	}

	// 提取原生的 *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorw("Error extracting sql.DB from gorm.DB", "error", err)
		return nil, err
	}

	// 使用正确的函数创建 PostgresqlStorage 实例
	return NewPostgresqlStorage(sqlDB), nil
}
