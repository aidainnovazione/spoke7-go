package postgres

import (
	"context"
	"spoke7-go/pkg/logger"
	"time"

	gormLogger "gorm.io/gorm/logger"
)

type dbLogger struct {
	logger logger.Logger
}

func (l *dbLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return l
}

func (l *dbLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Infof(msg, data...)
}

func (l *dbLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Warnf(msg, data...)
}

func (l *dbLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Errorf(msg, data...)
}

func (l *dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if err != nil {
		l.logger.Errorf("trace: %s", err.Error())
	}
}
