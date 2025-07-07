package repository

import (
	"errors"
	"fmt"
	"spoke7-go/internal/storage/repository/postgres"
	"spoke7-go/internal/storage/storage_interface"

	"spoke7-go/pkg/logger"
)

var (
	ErrNotFound            = errors.New("item not found")
	ErrUnsupportedDatabase = errors.New("unsupported database type! supported( postgres )")
	ErrInvalidObjectId     = errors.New("invalid object ID")
	ErrNotUnique           = errors.New("resource already exists")
	ErrCommandStillInUse   = errors.New("command is still in use by device profiles")
	ErrSlugEmpty           = errors.New("slug is nil or empty")
	ErrNameEmpty           = errors.New("name is required")
)

type DBConfig struct {
	Type         string `toml:"type" yaml:"type"`
	Host         string `toml:"host" yaml:"host"`
	Port         int    `toml:"port" yaml:"port"`
	Timeout      int    `toml:"timeout" yaml:"timeout"`
	DatabaseName string `toml:"databaseName" yaml:"databaseName"`
	Username     string `toml:"username" yaml:"username"`
	Password     string `toml:"password" yaml:"password"`
	BatchSize    int    `toml:"batchSize" yaml:"batchSize"`
	SSLMode      bool   `toml:"sslMode" yaml:"sslMode"`
	SSLRootCert  string `toml:"sslRootCert" yaml:"sslRootCert"`
}

func (c *DBConfig) Validate() error {
	if c.Type == "" {
		return fmt.Errorf("Database.Type is required")
	}
	if c.Host == "" {
		return fmt.Errorf("Database.Host is required")
	}
	if c.Port == 0 {
		return fmt.Errorf("Database.Port is required")
	}
	if c.DatabaseName == "" {
		return fmt.Errorf("Database.DatabaseName is required")
	}
	if c.Username == "" {
		return fmt.Errorf("Database.Username is required")
	}
	if c.Password == "" {
		return fmt.Errorf("Database.Password is required")
	}

	if c.SSLMode {
		if c.SSLRootCert == "" {
			return fmt.Errorf("Database.SSLRootCert is required")
		}
	}
	return nil
}

// NewDBClient creates a new database client.

func NewClient(conf DBConfig, logger logger.Logger) (storage_interface.StorageInterface, error) {
	switch conf.Type {
	case "postgres":
		postgresConf := postgres.PostgresConf{
			Host:         conf.Host,
			Port:         uint16(conf.Port),
			DatabaseName: conf.DatabaseName,
			Username:     conf.Username,
			Password:     conf.Password,
			SSLMode:      conf.SSLMode,
			SSLRootCert:  conf.SSLRootCert,
			BatchSize:    conf.BatchSize,
		}
		return postgres.NewPostgresClient(postgresConf, logger)
	default:
		return nil, ErrUnsupportedDatabase
	}
}
