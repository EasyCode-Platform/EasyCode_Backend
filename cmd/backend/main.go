package main

import (
	"os"

	"github.com/EasyCode-Platform/EasyCode_Backend/controller"
	"github.com/EasyCode-Platform/EasyCode_Backend/drive"
	"github.com/EasyCode-Platform/EasyCode_Backend/driver/awss3"
	"github.com/EasyCode-Platform/EasyCode_Backend/driver/postgres"
	"github.com/EasyCode-Platform/EasyCode_Backend/router"
	"github.com/EasyCode-Platform/EasyCode_Backend/storage"
	"github.com/EasyCode-Platform/EasyCode_Backend/utils/accesscontrol"
	"github.com/EasyCode-Platform/EasyCode_Backend/utils/config"
	"github.com/EasyCode-Platform/EasyCode_Backend/utils/cors"
	"github.com/EasyCode-Platform/EasyCode_Backend/utils/logger"
	"github.com/EasyCode-Platform/EasyCode_Backend/utils/recovery"
	"github.com/EasyCode-Platform/EasyCode_Backend/utils/tokenvalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	engine *gin.Engine
	router *router.Router
	logger *zap.SugaredLogger
	config *config.Config
}

func NewServer(config *config.Config, engine *gin.Engine, router *router.Router, logger *zap.SugaredLogger) *Server {
	return &Server{
		engine: engine,
		config: config,
		router: router,
		logger: logger,
	}
}

func initStorage(globalConfig *config.Config, logger *zap.SugaredLogger) *storage.Storage {
	postgresDriver, err := postgres.NewPostgresConnectionByGlobalConfig(globalConfig, logger)
	if err != nil {
		logger.Errorw("Error in startup, storage init failed.")
	}
	return storage.NewStorage(postgresDriver, logger)
}

func initDrive(globalConfig *config.Config, logger *zap.SugaredLogger) *drive.Drive {
	if globalConfig.IsAWSTypeDrive() {
		teamAWSConfig := awss3.NewTeamAwsConfigByGlobalConfig(globalConfig)
		teamDriveS3Instance := awss3.NewS3Drive(teamAWSConfig)
		return drive.NewDrive(teamDriveS3Instance, logger)
	}
	// failed
	logger.Errorw("Error in startup, drive init failed.")
	return nil
}

func initServer() (*Server, error) {
	globalConfig := config.GetInstance()
	engine := gin.New()
	sugaredLogger := logger.NewSugardLogger()

	// init validator
	validator := tokenvalidator.NewRequestTokenValidator()

	// init driver
	storage := initStorage(globalConfig, sugaredLogger)
	drive := initDrive(globalConfig, sugaredLogger)

	// init attribute group
	attrg, errInNewAttributeGroup := accesscontrol.NewRawAttributeGroup()
	if errInNewAttributeGroup != nil {
		return nil, errInNewAttributeGroup
	}

	// init controller
	c := controller.NewControllerForBackend(storage, drive, validator, attrg)
	router := router.NewRouter(c)
	server := NewServer(globalConfig, engine, router, sugaredLogger)
	return server, nil

}

func (server *Server) Start() {
	server.logger.Infow("Starting illa-builder-backend...")

	// init
	gin.SetMode(server.config.ServerMode)

	// init cors
	server.engine.Use(gin.CustomRecovery(recovery.CorsHandleRecovery))
	server.engine.Use(cors.Cors())
	server.router.RegisterRouters(server.engine)

	// run
	err := server.engine.Run(server.config.ServerHost + ":" + server.config.ServerPort)
	if err != nil {
		server.logger.Errorw("Error in startup", "err", err)
		os.Exit(2)
	}
}

func main() {
	server, err := initServer()

	if err != nil {

	}

	server.Start()
}