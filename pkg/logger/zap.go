package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger struct {
	Instance *zap.Logger
}

func NewZapLogger(cfg LoggerConfig) (*ZapLogger, error) {
	var zapLevel zapcore.Level
	switch cfg.Level {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "warn":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	case "fatal":
		zapLevel = zap.FatalLevel
	default:
		zapLevel = zap.InfoLevel
	}

	var core zapcore.Core

	// Configure output: stdout, file, or both
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	if cfg.Output == "file" {
		// Lumberjack setup for log file rotation
		lumberjackLogger := &lumberjack.Logger{
			Filename:   cfg.File.Path,
			MaxSize:    cfg.File.MaxSize,    // in megabytes
			MaxBackups: cfg.File.MaxBackups, // number of backups
			MaxAge:     cfg.File.MaxAge,     // in days
			Compress:   cfg.File.Compress,   // compress with gzip
		}

		fileWriter := zapcore.AddSync(lumberjackLogger)
		fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

		// Combine console and file
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), zapLevel),
			zapcore.NewCore(fileEncoder, fileWriter, zapLevel),
		)
	} else {
		// Console only
		core = zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapLevel)
	}

	logger := zap.New(core)

	return &ZapLogger{Instance: logger}, nil
}

func (l *ZapLogger) SetLogger(logger interface{}) error {
	newLogger, ok := logger.(*zap.Logger)
	if !ok {
		return fmt.Errorf("invalid logger type")
	}
	l.Instance = newLogger
	return nil
}

func (l *ZapLogger) Debug(msg ...interface{}) {
	l.Instance.Sugar().Debug(msg)
}

func (l *ZapLogger) Info(msg ...interface{}) {
	l.Instance.Sugar().Info(msg)
}

func (l *ZapLogger) Warn(msg ...interface{}) {
	l.Instance.Sugar().Warn(msg)
}

func (l *ZapLogger) Error(msg ...interface{}) {
	l.Instance.Sugar().Error(msg)
}

func (l *ZapLogger) Fatal(msg ...interface{}) {
	l.Instance.Sugar().Fatal(msg)
}

func (l *ZapLogger) Debugf(msg string, args ...interface{}) {
	l.Instance.Sugar().Debugf(msg, args...)
}

func (l *ZapLogger) Infof(msg string, args ...interface{}) {
	l.Instance.Sugar().Infof(msg, args...)
}

func (l *ZapLogger) Warnf(msg string, args ...interface{}) {
	l.Instance.Sugar().Warnf(msg, args...)
}

func (l *ZapLogger) Errorf(msg string, args ...interface{}) {
	l.Instance.Sugar().Errorf(msg, args...)
}

func (l *ZapLogger) Fatalf(msg string, args ...interface{}) {
	l.Instance.Sugar().Fatalf(msg, args...)
}

func (l *ZapLogger) ZapInfo(msg string, fields ...zap.Field) {
	l.Instance.Info(msg, fields...)
}

func (l *ZapLogger) ZapError(msg string, fields ...zap.Field) {
	l.Instance.Error(msg, fields...)
}
