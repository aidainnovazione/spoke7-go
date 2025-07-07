package postgres

import (
	"context"
	"errors"
	"spoke7-go/pkg/logger"
	"time"

	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// ErrRecordNotFound record not found error
var ErrRecordNotFound = errors.New("record not found")

type GormLoggerAdapter struct {
	//	logger logger.Logger
	logger logger.Logger
}

func (l *GormLoggerAdapter) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return l
}

func (l *GormLoggerAdapter) Info(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Infof(msg, data...)
}

func (l *GormLoggerAdapter) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Warnf(msg, data...)
}

func (l *GormLoggerAdapter) Error(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Errorf(msg, data...)
}

func (l *GormLoggerAdapter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && !errors.Is(err, ErrRecordNotFound):
		sql, rows := fc()
		l.logger.Errorf("%s [%.3fms] [rows:%v] [sql:%s] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql, err.Error())

	case err != nil && errors.Is(err, ErrRecordNotFound):
		sql, rows := fc()
		l.logger.Warnf("%s [%.3fms] [rows:%v] [sql:%s]", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)

	default:
		sql, rows := fc()
		l.logger.Debugf("%s [%.3fms] [rows:%v] [sql:%s]", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)

	}
}
