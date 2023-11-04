package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/EasyCode-Platform/EasyCode_Backend/src/utils/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const RETRY_TIMES = 6

type MongodbConfig struct {
	Addr          string `env:"MONGODB_ADDR" envDefault:"localhost"`
	Port          string `env:"MONGODB_PORT" envDefault:"27017"`
	Database      string `env:"MONGODB_DATABASE" envDefault:"easycode"`
	MaxCollection int64  `env:"MONGODB_MAXCOLLECTION" envDefault:"10"`
}

func NewMongodbConnectionByGlobalConfig(config *config.Config, logger *zap.SugaredLogger) (*mongo.Database, error) {
	mongodbConfig := &MongodbConfig{
		Addr:          config.GetMongodbAddr(),
		Port:          config.GetMongodbPort(),
		Database:      config.GetMongodbDatabase(),
		MaxCollection: config.GetMongodbMaxCollection(),
	}
	return NewMongodbConnection(mongodbConfig, logger)
}

func NewMongodbConnection(config *MongodbConfig, logger *zap.SugaredLogger) (*mongo.Database, error) {
	var client *mongo.Client
	var err error
	retries := RETRY_TIMES
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", config.Addr, config.Port))
	clientOptions.SetMaxPoolSize(uint64(config.MaxCollection))
	var timeout time.Duration = 10 // 设置10秒的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, err = mongo.Connect(ctx, clientOptions)
	for err != nil {
		if logger != nil {
			logger.Errorw("Failed to connect to mongodb database, %d", retries)
		}
		if retries > 1 {
			retries--
			time.Sleep(10 * time.Second)
			client, err = mongo.Connect(ctx, clientOptions)
			continue
		}
		panic(err)
	}
	err = client.Ping(ctx, nil)

	if err != nil {
		if logger != nil {
			logger.Errorw("Connection success but failed to ping mongodb database", "db", config, "err", err)
		}
		return nil, err
	}

	logger.Infow("connected to mongodb database", "db", config)

	return client.Database(config.Database), err

}
