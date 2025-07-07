package main

import (
	"fmt"
	"log"

	"spoke7-go/pkg/configloader"
	"spoke7-go/pkg/logger"

	"spoke7-go/internal/managment"
	"spoke7-go/internal/managment/config"

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

	fmt.Printf("Configuration loaded: %+v\n", config.AppConfig)

}

func main() {
	logger, err := logger.NewZapLogger(config.AppConfig.Log)
	if err != nil {
		log.Fatalf("Create logger failed. Error:%v\n", err)
	}
	defer logger.Instance.Sync()

	logger.Info("Starting Managment Server")

	e := echo.New()

	//enforcer := authz.InitCasbinAndMigrate(logger, config.AppConfig.Authorization, dbconn)

	mgmtServer := managment.NewManagmentServer(e, config.AppConfig, logger, Version)

	err = mgmtServer.Start()
	if err != nil {
		logger.Fatalf("Failed to start server. Error:%v\n", err)
	}

}
