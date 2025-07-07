package postgres

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"testing"
	"time"

	"spoke7-go/internal/metadata/models"
	"spoke7-go/pkg/authz"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// Mock JSONB data type for SQL mock.
type JSONBValue struct {
	Data interface{}
}

func (j JSONBValue) Value() (driver.Value, error) {
	return json.Marshal(j.Data)
}

func TestDetectionSectionPostgresRepository_Create(t *testing.T) {
	repo, mock := setupMockDB(t)
	defer repo.Close()

	detectionSection := &models.DetectionSection{
		Id:             "123",
		DataSourceName: "DS1",
		Description:    "A detection section",
		StartLatitude:  40.7128,
		StartLongitude: -74.0060,
		EndLatitude:    40.7306,
		EndLongitude:   -73.9352,
		Direction:      1,

		RoadNetworkId: "networkId",
		CreatedAt:     time.Now(),
		ModifiedAt:    time.Now(),
	}

	user := authz.User{
		Username: "owner",
		Groups:   []string{"role"},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, authz.UserCtxKey, user)

	expectedCheckDataSource(mock, "DS1", user)

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "detection_sections"`).
		WithArgs(
			detectionSection.Id,
			detectionSection.DataSourceName,
			detectionSection.Description,
			detectionSection.StartLatitude,
			detectionSection.StartLongitude,
			detectionSection.EndLatitude,
			detectionSection.EndLongitude,
			detectionSection.Direction,

			detectionSection.RoadNetworkId,
			sqlmock.AnyArg(), // CreatedAt (auto-generated)
			sqlmock.AnyArg(), // ModifiedAt (auto-generated)
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.CreateDetectionSection(ctx, detectionSection.DataSourceName, detectionSection)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDetectionSectionPostgresRepository_Get(t *testing.T) {
	repo, mock := setupMockDB(t)
	defer repo.Close()

	id := "123"
	createdAt := time.Now()
	modifiedAt := time.Now()
	user := authz.User{
		Username: "owner",
		Groups:   []string{"role"},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, authz.UserCtxKey, user)

	expectedCheckDataSource(mock, "DS1", user)

	mock.ExpectQuery(`SELECT \* FROM "detection_sections" WHERE id = \$1 ORDER BY "detection_sections"."id" LIMIT \$2`).
		WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "data_source_name", "description", "start_latitude", "start_longitude",
			"end_latitude", "end_longitude", "direction", "shape", "road_network_id", "created_at", "modified_at"}).
			AddRow(
				"123",
				"DS1",
				"A detection section",
				40.7128,
				-74.0060,
				40.7306,
				-73.9352,
				1,
				nil,
				"networkid",
				createdAt,
				modifiedAt,
			))

	result, err := repo.GetDetectionSection(ctx, "DS1", id)
	assert.NoError(t, err)
	assert.Equal(t, &models.DetectionSection{
		Id:             "123",
		DataSourceName: "DS1",
		Description:    "A detection section",
		StartLatitude:  40.7128,
		StartLongitude: -74.0060,
		EndLatitude:    40.7306,
		EndLongitude:   -73.9352,
		Direction:      1,

		RoadNetworkId:                "networkid",
		CreatedAt:                    createdAt,
		ModifiedAt:                   modifiedAt,
		DetectionSectionRoadNetworks: []models.DetectionSectionRoadNetwork{},
	}, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDetectionSectionPostgresRepository_Update(t *testing.T) {
	repo, mock := setupMockDB(t)
	defer repo.Close()

	detectionSection := &models.DetectionSection{
		Id:             "123",
		DataSourceName: "DS1",
		Description:    "Updated detection section",
		StartLatitude:  40.7128,
		StartLongitude: -74.0060,
		EndLatitude:    40.7306,
		EndLongitude:   -73.9352,
		Direction:      1,

		RoadNetworkId: "networkId",
	}

	user := authz.User{
		Username: "owner",
		Groups:   []string{"role"},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, authz.UserCtxKey, user)

	expectedCheckDataSource(mock, "DS1", user)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").
		WithArgs(
			detectionSection.Id,             // $1
			detectionSection.DataSourceName, // $2
			detectionSection.Description,    // $3
			detectionSection.StartLatitude,  // $4
			detectionSection.StartLongitude, // $5
			detectionSection.EndLatitude,    // $6
			detectionSection.EndLongitude,   // $7
			detectionSection.Direction,      // $8
			detectionSection.RoadNetworkId,  // $9

			sqlmock.AnyArg(),    // $10
			detectionSection.Id, // $11 (WHERE clause)
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err := repo.UpdateDetectionSection(ctx, detectionSection.DataSourceName, detectionSection)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDetectionSectionPostgresRepository_Delete(t *testing.T) {
	repo, mock := setupMockDB(t)
	defer repo.Close()

	id := "123"

	user := authz.User{
		Username: "owner",
		Groups:   []string{"role"},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, authz.UserCtxKey, user)

	expectedCheckDataSource(mock, "DS1", user)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "detection_sections" WHERE id = \$1`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.DeleteDetectionSection(ctx, "DS1", id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDetectionSectionPostgresRepository_List(t *testing.T) {
	repo, mock := setupMockDB(t)

	defer repo.Close()

	createdAt := time.Now()
	modifiedAt := time.Now()

	user := authz.User{
		Username: "owner",
		Groups:   []string{"role"},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, authz.UserCtxKey, user)

	expectedCheckDataSource(mock, "DS1", user)

	mock.ExpectQuery(`SELECT \* FROM "detection_sections"`).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "data_source_name", "description", "start_latitude", "start_longitude",
			"end_latitude", "end_longitude", "direction", "shape", "road_network_id", "created_at", "modified_at"}).
			AddRow(
				"123",
				"DS1",
				"A detection section",
				40.7128,
				-74.0060,
				40.7306,
				-73.9352,
				1,
				nil,
				"networkid",
				createdAt,
				modifiedAt,
			).
			AddRow(
				"124",
				"DS1",
				"Another detection section",
				41.7128,
				-75.0060,
				41.7306,
				-74.9352,
				2,
				nil,
				"networkid2",
				createdAt,
				modifiedAt,
			))

	results, err := repo.ListDetectionSection(ctx, "DS1")
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, &models.DetectionSection{
		Id:             "123",
		DataSourceName: "DS1",
		Description:    "A detection section",
		StartLatitude:  40.7128,
		StartLongitude: -74.0060,
		EndLatitude:    40.7306,
		EndLongitude:   -73.9352,
		Direction:      1,

		RoadNetworkId:                "networkid",
		CreatedAt:                    createdAt,
		ModifiedAt:                   modifiedAt,
		DetectionSectionRoadNetworks: []models.DetectionSectionRoadNetwork{},
	}, results[0])
	assert.Equal(t, &models.DetectionSection{
		Id:             "124",
		DataSourceName: "DS1",
		Description:    "Another detection section",
		StartLatitude:  41.7128,
		StartLongitude: -75.0060,
		EndLatitude:    41.7306,
		EndLongitude:   -74.9352,
		Direction:      2,

		RoadNetworkId:                "networkid2",
		CreatedAt:                    createdAt,
		ModifiedAt:                   modifiedAt,
		DetectionSectionRoadNetworks: []models.DetectionSectionRoadNetwork{},
	}, results[1])
	assert.NoError(t, mock.ExpectationsWereMet())
}
