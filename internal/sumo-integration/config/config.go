package config

import (
	"spoke7-go/pkg/logger"
)

var AppConfig Configuration

type Configuration struct {
	Service             ServiceConfig        `toml:"service" yaml:"service"`
	GrpcDataService     ServiceConfig        `toml:"grpcDataService" yaml:"grpcDataService"`
	GrpcMetadataService ServiceConfig        `toml:"grpcDataService" yaml:"grpcMetadataService"`
	GrpcStorageService  ServiceConfig        `toml:"service" yaml:"grpcStorageService"`
	Log                 logger.LoggerConfig  `toml:"log" yaml:"log"`
	Authentication      AuthenticationConfig `toml:"authentication" yaml:"authentication"`
}

func (c *Configuration) Validate() error {
	if err := c.Service.Validate(); err != nil {
		return err
	}

	if err := c.Log.Validate(); err != nil {
		return err
	}

	if err := c.GrpcDataService.Validate(); err != nil {
		return err
	}

	if err := c.Authentication.Validate(); err != nil {
		return err
	}

	return nil
}
