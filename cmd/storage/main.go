package main

import (
	"context"
	"fmt"
	"log"
	"spoke7-go/internal/storage"
	"spoke7-go/internal/storage/config"
	"spoke7-go/internal/storage/repository"

	"spoke7-go/pkg/configloader"
	"spoke7-go/pkg/logger"

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

	e := echo.New()

	logger.Info("Starting Storage Server")

	if config.AppConfig.StorageType.Type != "database" {
		logger.Fatalf("Storage type %v not supported", config.AppConfig.StorageType.Type)
	}

	repo, err := repository.NewClient(config.AppConfig.Database, logger)
	if err != nil {
		logger.Fatal(err)
	}
	defer repo.Close()

	err = repo.Migrate(context.Background())
	if err != nil {
		logger.Fatalf("Failed to migrate database. Error:%v\n", err)
	}

	grpcMetadataString := fmt.Sprintf("%s:%v", config.AppConfig.GrpcMetadataService.Host, config.AppConfig.GrpcMetadataService.GrpcPort)

	storageServer := storage.NewStorageServer(e, repo, grpcMetadataString, Version, logger, config.AppConfig.Authentication.OrganizationName)

	err = storageServer.Start()
	if err != nil {
		logger.Fatalf("Failed to start server. Error:%v\n", err)
	}
}
