package main

import (
	"fmt"
	"log"

	"spoke7-go/pkg/configloader"
	"spoke7-go/pkg/logger"

	sumo "spoke7-go/internal/sumo-integration"
	"spoke7-go/internal/sumo-integration/config"

	"github.com/labstack/echo/v4"
)

// Version is the version of the application
// This value is set at build time using ldflags
var Version = "0.0.0"

var (
	defaultConfigPath = "res/configuration.yaml"
	defaultPort       = 8080
)

func init() {
	configLoader := configloader.NewConfigLoader(defaultConfigPath, defaultPort)

	err := configLoader.ParseFlags()
	if err != nil {
		log.Fatalf("Parse flags failed. Error:%v\n", err)
	}

	// get the global configuration instance
	err = configLoader.LoadConfig(&config.AppConfig)
	if err != nil {
		log.Fatalf("Load config failed. Error:%v\n", err)
	}

}

func main() {
	logger, err := logger.NewZapLogger(config.AppConfig.Log)
	if err != nil {
		log.Fatalf("Create logger failed. Error:%v\n", err)
	}
	defer logger.Instance.Sync()

	logger.Info("Starting Sumo Integration Server")

	e := echo.New()

	grpcDataString := fmt.Sprintf("%s:%v", config.AppConfig.GrpcDataService.Host, config.AppConfig.GrpcDataService.GrpcPort)
	grpcMetadataString := fmt.Sprintf("%s:%v", config.AppConfig.GrpcMetadataService.Host, config.AppConfig.GrpcMetadataService.GrpcPort)
	grpcStorageString := fmt.Sprintf("%s:%v", config.AppConfig.GrpcStorageService.Host, config.AppConfig.GrpcStorageService.GrpcPort)

	dataServer := sumo.NewSumoIntegrationServer(e, grpcDataString, grpcMetadataString, grpcStorageString, Version, logger)

	err = dataServer.Start()
	if err != nil {
		logger.Fatalf("Failed to start server. Error:%v\n", err)
	}

}
