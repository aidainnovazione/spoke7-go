package repository

import (
	"context"
	"spoke7-go/pkg/logger"
	"time"

	gormLogger "gorm.io/gorm/logger"
)

type DBLogger struct {
	Logger logger.Logger
}

func (l *DBLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return l
}

func (l *DBLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.Infof(msg, data...)
}

func (l *DBLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.Warnf(msg, data...)
}

func (l *DBLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.Errorf(msg, data...)
}

func (l *DBLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if err != nil {
		l.Logger.Errorf("trace: %s", err.Error())
	}
}
