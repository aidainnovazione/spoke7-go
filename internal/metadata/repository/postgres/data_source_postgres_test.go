package postgres

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository/postgres/dao"
	"spoke7-go/pkg/authz"
)

func expectedCheckDataSource(mock sqlmock.Sqlmock, datasourceName string, user authz.User) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "data_source" WHERE name = $1 and (owner = $2 or groups && ARRAY[$3])`)).
		WithArgs(datasourceName, user.Username, user.Groups[0]).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
}

func TestDataSourcePostgresRepository_Create(t *testing.T) {
	// Setup mock DB
	repo, mock := setupMockDB(t)

	defer repo.Close()
	defer mock.ExpectClose()

	// Set up user context
	user := authz.User{
		Username: "owner",
		Groups:   []string{"spoke7/admin"},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, authz.UserCtxKey, user)

	// now := time.Now()
	// Initialize test data
	dataSource := &models.DataSource{
		Name:          "Test",
		Description:   "Test Description",
		Type:          "example",
		RoadNetworkId: nil,
		Owner:         "owner",
		Groups:        []string{"spoke7/admin"},
		ModifiedBy:    "modifier",
		// CreatedAt:     now,
		// ModifiedAt:    now,
	}

	// Mock expected data with timestamps
	dataSourceDao := dao.NewDataSourceDaoFromModel(*dataSource)

	// Set up mock expectations
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"data_source\"").
		WithArgs(
			dataSourceDao.Name,
			dataSourceDao.Description,
			dataSourceDao.Type,
			dataSourceDao.RoadNetworkId,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			dataSourceDao.Owner,
			dataSourceDao.Groups,
			dataSourceDao.ModifiedBy,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Instantiate repository and call Create

	err := repo.CreateDataSource(ctx, dataSource)

	// Assertions
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestUpdate(t *testing.T) {
	repo, mock := setupMockDB(t)

	defer repo.Close()
	defer mock.ExpectClose()

	// Set up user context
	user := authz.User{
		Username: "owner",
		Groups:   []string{"spoke7/admin"},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, authz.UserCtxKey, user)

	updatedDesc := "test update desc"

	dataSource := &models.UpdateDataSource{
		Name:        "test",
		Description: &updatedDesc,
	}
	dataSourceDao := dao.NewUpdateDataSourceDaoFromModel(*dataSource)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WithArgs(
		dataSourceDao.Name,
		*dataSource.Description,
		sqlmock.AnyArg(),
		user.Username,
		dataSourceDao.Name,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.UpdateDataSource(ctx, dataSource)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	repo, mock := setupMockDB(t)

	// Set up user context
	user := authz.User{
		Username: "owner",
		Groups:   []string{"spoke7/admin"},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, authz.UserCtxKey, user)

	defer repo.Close()
	defer mock.ExpectClose()

	name := "test"

	expectedCheckDataSource(mock, name, user)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM").WithArgs(name).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.DeleteDataSource(ctx, name)
	assert.NoError(t, err)
}

func TestDataSourcePostgresRepository_Get(t *testing.T) {
	// Initialize sqlmock
	repo, mock := setupMockDB(t)

	defer repo.Close()
	defer mock.ExpectClose()

	// Mock data
	name := "test_name"
	createdAt := time.Now()
	modifiedAt := time.Now()

	mockDataSource := dao.DataSource{
		Name:          name,
		Description:   "test description",
		Type:          "example",
		RoadNetworkId: nil,
		CreatedAt:     createdAt,
		ModifiedAt:    modifiedAt,
		Owner:         "owner",
		Groups:        []string{"spoke7/admin"},
		ModifiedBy:    "modifier",
	}

	// Set up user context
	user := authz.User{
		Username: "owner",
		Groups:   []string{"spoke7/admin"},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, authz.UserCtxKey, user)

	expectedCheckDataSource(mock, name, user)

	// Set up mock query
	rows := sqlmock.NewRows([]string{"name", "description", "type", "road_network_id", "created_at", "modified_at", "owner", "spoke7/admin", "modified_by"}).
		AddRow(mockDataSource.Name, mockDataSource.Description, mockDataSource.Type, mockDataSource.RoadNetworkId, mockDataSource.CreatedAt, mockDataSource.ModifiedAt, mockDataSource.Owner, mockDataSource.Groups, mockDataSource.ModifiedBy)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "data_source" WHERE  name = $1 ORDER BY "data_source"."name" LIMIT $2`)).
		WithArgs(name, 1).
		WillReturnRows(rows)

	// Call the repository method
	result, err := repo.GetDataSource(ctx, name, models.DataSourceGetParams{})

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockDataSource.Name, result.Name)
	assert.Equal(t, mockDataSource.Description, result.Description)
	assert.Equal(t, models.DataSourceType(mockDataSource.Type), result.Type)
	assert.Equal(t, mockDataSource.RoadNetworkId, result.RoadNetworkId)
	assert.Equal(t, mockDataSource.Owner, result.Owner)
	assert.Equal(t, mockDataSource.Groups, result.Groups)
	assert.Equal(t, mockDataSource.ModifiedBy, result.ModifiedBy)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestList(t *testing.T) {
	// Setup mock DB
	repo, mock := setupMockDB(t)

	defer repo.Close()

	user := authz.User{
		Username: "owner",
		Groups:   []string{"spoke7/admin"},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, authz.UserCtxKey, user)

	t.Run("returns data without filters", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).AddRow("test")
		mock.ExpectQuery("SELECT .* FROM \"data_source\"").WillReturnRows(rows)

		organizationName := "spoke7"
		result, err := repo.ListDataSource(ctx, models.DataSourceListParams{}, organizationName)
		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "test", result[0].Name)
	})

	t.Run("applies DetectionSections preload", func(t *testing.T) {
		// Mock source data
		dataSourceRows := sqlmock.NewRows([]string{"name"}).AddRow("test")
		mock.ExpectQuery("SELECT .* FROM \"data_source\"").WillReturnRows(dataSourceRows)

		// Mock detection sections
		detectionSectionsRows := sqlmock.NewRows([]string{"id", "data_source_name"}).AddRow("1", "test")
		mock.ExpectQuery("SELECT .* FROM \"detection_sections\"").WillReturnRows(detectionSectionsRows)

		organizationName := "spoke7"
		result, err := repo.ListDataSource(ctx, models.DataSourceListParams{DetectionSections: true}, organizationName)
		assert.NoError(t, err)
		if !assert.Len(t, result, 1) {
			return
		}

		assert.Len(t, result[0].DetectionSections, 1)
	})

	t.Run("applies DetectionPoints preload", func(t *testing.T) {
		// Mock source data
		dataSourceRows := sqlmock.NewRows([]string{"name"}).AddRow("1")
		mock.ExpectQuery("SELECT .* FROM \"data_source\"").WillReturnRows(dataSourceRows)

		// Mock detection points
		detectionPointsRows := sqlmock.NewRows([]string{"id", "data_source_name"}).AddRow("1", "1")
		mock.ExpectQuery("SELECT .* FROM \"detection_points\"").WillReturnRows(detectionPointsRows)

		organizationName := "spoke7"
		result, err := repo.ListDataSource(ctx, models.DataSourceListParams{DetectionPoints: true}, organizationName)
		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Len(t, result[0].DetectionPoints, 1)
	})

	t.Run("returns error when query fails", func(t *testing.T) {
		mock.ExpectQuery("SELECT .* FROM \"data_source\"").WillReturnError(fmt.Errorf("query error"))

		organizationName := "spoke7"
		result, err := repo.ListDataSource(ctx, models.DataSourceListParams{}, organizationName)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
