package postgres

import (
	"strings"
	"sync"
	"testing"

	"spoke7-go/pkg/logger"

	"github.com/DATA-DOG/go-sqlmock"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var mockInstance *postgresClient

func NewMockPostgresClient(conf PostgresConf, logger logger.Logger) (*postgresClient, error) {
	var finalErr error
	once.Do(func() {
		db, _, err := sqlmock.New()
		if err != nil {
			finalErr = err
			return
		}

		logger.Infof("connecting to database on %s:%d", conf.Host, conf.Port)

		gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

		if err != nil {
			logger.Fatalf("failed to connect to database on %s:%d: %s", conf.Host, conf.Port, err.Error())
			finalErr = err
			return
		}

		logger.Infof("connected to database on %s:%d", conf.Host, conf.Port)
		mockInstance = &postgresClient{
			db:     gdb,
			logger: logger,
		}
	})

	return mockInstance, finalErr
}
func setupMockDB(t *testing.T) (*postgresClient, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}

	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("error creating gorm database: %v", err)
	}

	repo := &postgresClient{db: gdb}

	return repo, mock
}

func setupMockLogger() *logger.MockLogger {
	return &logger.MockLogger{
		Messages:  []string{},
		FatalfMsg: "",

		SetLoggerFunc: func(interface{}) error {
			return nil
		},
		InfoFunc: func(...interface{}) {
			// Ignore info logs.
		},
		ErrorFunc: func(...interface{}) {
			// Ignore error logs.
		},
		DebugFunc: func(...interface{}) {
			// Ignore debug logs.
		},
		WarnFunc: func(...interface{}) {
			// Ignore warn logs.
		},
		FatalFunc: func(...interface{}) {
			// Ignore fatal logs.
		},
	}
}

// TestNewPostgresClient_SingletonBehavior checks if the client initializes the connection once.
// Note: Current code has a bug where subsequent clients have nil db.
func TestNewPostgresClient_SingletonBehavior(t *testing.T) {
	ml := setupMockLogger()
	conf := PostgresConf{
		Host:         "localhost",
		Port:         5432,
		Username:     "user",
		Password:     "pass",
		DatabaseName: "testdb",
	}

	// First call initializes the connection.
	client1, err := NewMockPostgresClient(conf, ml)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Second call should reuse the connection but returns nil due to the bug.
	client2, err := NewMockPostgresClient(conf, ml)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// This assertion fails due to the bug in the current code.
	if client1.db != client2.db {
		t.Error("Expected the same DB instance, but got different instances.")
	}
}

// TestNewPostgresClient_ConnectionFailure checks error handling on connection failure.
func TestNewPostgresClient_ConnectionFailure(t *testing.T) {
	// Reset the sync.Once to ensure the Do block runs.
	once = sync.Once{}

	ml := setupMockLogger()
	conf := PostgresConf{
		Host:         "invalidhost", // Unreachable host to simulate failure.
		Port:         9999,
		Username:     "user",
		Password:     "pass",
		DatabaseName: "db",
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected a panic from Fatalf, but none occurred.")
		} else {
			expectedMsg := "failed to connect to database on invalidhost:9999"
			if !strings.Contains(r.(string), expectedMsg) {
				t.Errorf("Expected panic message %q, got %q.", expectedMsg, r)
			}
		}
	}()

	_, _ = NewPostgresClient(conf, ml) // Should panic via Fatalf.
}

// TestNewPostgresClient_ConnectionParams verifies connection parameters in logs.
func TestNewPostgresClient_ConnectionParams(t *testing.T) {
	// Reset the sync.Once to ensure the Do block runs.
	once = sync.Once{}

	ml := setupMockLogger()
	conf := PostgresConf{
		Host:         "testhost",
		Port:         9999,
		Username:     "testuser",
		Password:     "testpass",
		DatabaseName: "testdb",
	}

	_, err := NewMockPostgresClient(conf, ml)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check if "connecting to database on testhost:9999" is logged.
	expectedLog := "connecting to database on testhost:9999"
	found := false
	for _, msg := range ml.Messages {
		if strings.Contains(msg, expectedLog) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected log message %q not found.", expectedLog)
	}

	// Check if "connected to database on testhost:9999" is logged.
	expectedConnectedLog := "connected to database on testhost:9999"
	foundConnected := false
	for _, msg := range ml.Messages {
		if msg == expectedConnectedLog {
			foundConnected = true
			break
		}
	}
	if !foundConnected {
		t.Errorf("Expected log message %q not found.", expectedConnectedLog)
	}
}
