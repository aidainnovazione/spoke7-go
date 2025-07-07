package config

import (
	"spoke7-go/pkg/logger"
)

var AppConfig AppConfiguration

type AppConfiguration struct {
	Service        ServiceConfig        `toml:"service" yaml:"service"`
	Log            logger.LoggerConfig  `toml:"log" yaml:"log"`
	Authentication AuthenticationConfig `toml:"authentication" yaml:"authentication"`
	Authorization  Authorization        `toml:"authorization" yaml:"authorization"`
	Database       DatabaseConfig       `toml:"database" yaml:"database"`
}

func (c *AppConfiguration) Validate() error {
	if err := c.Service.Validate(); err != nil {
		return err
	}
	if err := c.Log.Validate(); err != nil {
		return err
	}
	if err := c.Authentication.Validate(); err != nil {
		return err
	}
	if err := c.Authorization.Validate(); err != nil {
		return err
	}
	if err := c.Database.Validate(); err != nil {
		return err
	}

	return nil
}
