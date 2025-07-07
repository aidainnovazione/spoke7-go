package config

import (
	"spoke7-go/internal/data/repository"
	"spoke7-go/pkg/logger"
)

var AppConfig Configuration

type Configuration struct {
	Service             ServiceConfig        `toml:"service" yaml:"service"`
	Log                 logger.LoggerConfig  `toml:"log" yaml:"log"`
	Database            repository.DBConfig  `toml:"database" yaml:"database"`
	GrpcMetadataService ServiceConfig        `toml:"service" yaml:"grpcMetadataService"`
	GrpcStorageService  ServiceConfig        `toml:"service" yaml:"grpcStorageService"`
	Authentication      AuthenticationConfig `toml:"authentication" yaml:"authentication"`
}

func (c *Configuration) Validate() error {
	if err := c.Service.Validate(); err != nil {
		return err
	}

	if err := c.Log.Validate(); err != nil {
		return err
	}

	if err := c.Database.Validate(); err != nil {
		return err
	}

	if err := c.Authentication.Validate(); err != nil {
		return err
	}

	return nil
}
