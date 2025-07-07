package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger interface {
	SetLogger(interface{}) error

	Info(...interface{})
	Error(...interface{})
	Debug(...interface{})
	Warn(...interface{})
	Fatal(...interface{})

	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Debugf(string, ...interface{})
	Warnf(string, ...interface{})
	Fatalf(string, ...interface{})

	ZapInfo(msg string, fields ...zap.Field)
	ZapError(msg string, fields ...zap.Field)
}

type LoggerConfig struct {
	Level  string        `toml:"level" yaml:"level"` // debug, info, warn, error
	Output string        `toml:"output" yaml:"output"`
	File   LogFileConfig `toml:"file" yaml:"file"`
}

type LogFileConfig struct {
	Path       string `toml:"path" yaml:"path"`
	MaxSize    int    `toml:"maxSize" yaml:"maxSize"`
	MaxBackups int    `toml:"maxBackups" yaml:"maxBackups"`
	MaxAge     int    `toml:"maxAge" yaml:"maxAge"`
	Compress   bool   `toml:"compress" yaml:"compress"`
}

func (l *LoggerConfig) Validate() error {
	if l.Level == "" {
		return fmt.Errorf("Log.Level is required")
	}

	if l.Output == "" {
		return fmt.Errorf("Log.Output is required")
	}

	if l.Output == "file" {
		if l.File.Path == "" {
			return fmt.Errorf("Log.File.Path is required")
		}

	}

	return nil
}
