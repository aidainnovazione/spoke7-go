package logger

import "fmt"

// implement mock logger
type MockLogger struct {
	Messages       []string
	FatalfMsg      string
	SetLoggerFunc  func(interface{}) error
	setLoggerCalls []struct {
		arg0 interface{}
	}

	InfoFunc  func(...interface{})
	infoCalls []struct {
		arg0 []interface{}
	}

	ErrorFunc  func(...interface{})
	errorCalls []struct {
		arg0 []interface{}
	}

	DebugFunc  func(...interface{})
	debugCalls []struct {
		arg0 []interface{}
	}

	WarnFunc  func(...interface{})
	warnCalls []struct {
		arg0 []interface{}
	}

	FatalFunc  func(...interface{})
	fatalCalls []struct {
		arg0 []interface{}
	}
}

func NewMockLogger() *MockLogger {
	return &MockLogger{}
}

func (ml *MockLogger) SetLogger(arg0 interface{}) error {
	ml.setLoggerCalls = append(ml.setLoggerCalls, struct {
		arg0 interface{}
	}{arg0})
	return ml.SetLoggerFunc(arg0)
}

func (ml *MockLogger) Info(arg0 ...interface{}) {
	ml.infoCalls = append(ml.infoCalls, struct {
		arg0 []interface{}
	}{arg0})
	ml.InfoFunc(arg0...)
}

func (ml *MockLogger) Error(arg0 ...interface{}) {
	ml.errorCalls = append(ml.errorCalls, struct {
		arg0 []interface{}
	}{arg0})
	ml.ErrorFunc(arg0...)
}

func (ml *MockLogger) Debug(arg0 ...interface{}) {
	ml.debugCalls = append(ml.debugCalls, struct {
		arg0 []interface{}
	}{arg0})
	ml.DebugFunc(arg0...)
}

func (ml *MockLogger) Warn(arg0 ...interface{}) {
	ml.warnCalls = append(ml.warnCalls, struct {
		arg0 []interface{}
	}{arg0})
	ml.WarnFunc(arg0...)
}

func (ml *MockLogger) Fatal(arg0 ...interface{}) {
	ml.fatalCalls = append(ml.fatalCalls, struct {
		arg0 []interface{}
	}{arg0})
	ml.FatalFunc(arg0...)
}

func (ml *MockLogger) Infof(format string, args ...interface{}) {
	ml.Messages = append(ml.Messages, fmt.Sprintf(format, args...))
}

func (ml *MockLogger) Errorf(format string, args ...interface{}) {
	ml.Messages = append(ml.Messages, fmt.Sprintf(format, args...))
}

func (ml *MockLogger) Debugf(format string, args ...interface{}) {
	ml.Messages = append(ml.Messages, fmt.Sprintf(format, args...))
}

func (ml *MockLogger) Warnf(format string, args ...interface{}) {
	ml.Messages = append(ml.Messages, fmt.Sprintf(format, args...))
}

func (ml *MockLogger) Fatalf(format string, args ...interface{}) {
	ml.FatalfMsg = fmt.Sprintf(format, args...)
	panic(ml.FatalfMsg) // simulate panic
}
