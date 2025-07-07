package postgres

import (
	"context"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	// Create a mock SQL database
	repo, mock := setupMockDB(t)

	defer repo.Close()

	ctx := context.Background()

	// Expect schema creation queries for both tables
	mock.ExpectExec("CREATE TABLE \"data_source\".*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("CREATE TABLE \"detection_sections\".*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("CREATE TABLE \"detection_points\".*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("CREATE TABLE \"detection_section_road_networks\".*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("CREATE TABLE \"lanes\".*").WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the Migrate function
	err := repo.Migrate(ctx)
	assert.NoError(t, err, "Migration should not return an error")

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
