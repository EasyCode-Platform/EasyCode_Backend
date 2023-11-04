package storage

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type ComponentStorage struct {
	logger  *zap.SugaredLogger
	storage *MongoDbStorage
}

func NewComponentStorage(logger *zap.SugaredLogger, db *mongo.Database) *ComponentStorage {
	return &ComponentStorage{
		logger:  logger,
		storage: NewMongoDb(db, "components"),
	}
}

func (ComponentStorage) UpdateComponent()
