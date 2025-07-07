package postgres

import (
	"fmt"
	"sync"

	"spoke7-go/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once
var instance *postgresClient

const defaultBatchSize = 1000

type postgresClient struct {
	db     *gorm.DB
	logger logger.Logger
	conf   PostgresConf
}

func NewPostgresClient(conf PostgresConf, logger logger.Logger) (*postgresClient, error) {
	if conf.BatchSize == 0 {
		conf.BatchSize = defaultBatchSize
	}

	once.Do(func() {
		dbLogger := &dbLogger{logger: logger}

		logger.Infof("connecting to database on %s:%d", conf.Host, conf.Port)

		dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
		dsn = fmt.Sprintf(dsn, conf.Host, conf.Username, conf.Password, conf.DatabaseName, conf.Port)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 dbLogger,
		})

		if err != nil {
			logger.Fatalf("failed to connect to database on %s:%d: %s", conf.Host, conf.Port, err.Error())
		}

		logger.Infof("connected to database on %s:%d", conf.Host, conf.Port)
		instance = &postgresClient{
			db:     db,
			logger: logger,
			conf:   conf,
		}
	})

	return instance, nil
}

func (pc *postgresClient) Close() error {
	sqlDB, err := pc.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
