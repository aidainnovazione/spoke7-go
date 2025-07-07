package services

import (
	"context"
	"testing"

	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository/mock"

	"github.com/stretchr/testify/assert"
)

var MOCK_ORGANIZATION_NAME string = "spoke7"

func TestDataSourceService_Create_Success(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		CreateDataSourceFunc: func(ctx context.Context, dataSource *models.DataSource) error {
			return nil
		},
	}
	service := NewDataSourceService(mockRepo, MOCK_ORGANIZATION_NAME)

	err := service.Create(context.Background(), &models.DataSource{})
	assert.NoError(t, err)
}

func TestDataSourceService_Create_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		CreateDataSourceFunc: func(ctx context.Context, dataSource *models.DataSource) error {
			return assert.AnError
		},
	}
	service := NewDataSourceService(mockRepo, MOCK_ORGANIZATION_NAME)

	err := service.Create(context.Background(), &models.DataSource{})
	assert.Error(t, err)
}

func TestDataSourceService_Update_Success(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		UpdateDataSourceFunc: func(ctx context.Context, dataSource *models.UpdateDataSource) error {
			return nil
		},
		GetDataSourceFunc: func(ctx context.Context, name string, params models.DataSourceGetParams) (*models.DataSource, error) {
			return &models.DataSource{}, nil
		},
	}
	service := NewDataSourceService(mockRepo, MOCK_ORGANIZATION_NAME)

	_, err := service.Update(context.Background(), &models.UpdateDataSource{})
	assert.NoError(t, err)
}

func TestDataSourceService_Update_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		UpdateDataSourceFunc: func(ctx context.Context, dataSource *models.UpdateDataSource) error {
			return assert.AnError
		},
		GetDataSourceFunc: func(ctx context.Context, name string, params models.DataSourceGetParams) (*models.DataSource, error) {
			return nil, assert.AnError
		},
	}
	service := NewDataSourceService(mockRepo, MOCK_ORGANIZATION_NAME)

	_, err := service.Update(context.Background(), &models.UpdateDataSource{})
	assert.Error(t, err)
}

func TestDataSourceService_Delete_Success(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		DeleteDataSourceFunc: func(ctx context.Context, name string) error {
			return nil
		},
	}
	service := NewDataSourceService(mockRepo, MOCK_ORGANIZATION_NAME)

	err := service.Delete(context.Background(), "test")
	assert.NoError(t, err)
}

func TestDataSourceService_Delete_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		DeleteDataSourceFunc: func(ctx context.Context, name string) error {
			return assert.AnError
		},
	}
	service := NewDataSourceService(mockRepo, MOCK_ORGANIZATION_NAME)

	err := service.Delete(context.Background(), "test")
	assert.Error(t, err)
}

func TestDataSourceService_Get_Success(t *testing.T) {
	expectedDataSource := &models.DataSource{}
	mockRepo := &mock.MockDBClient{
		GetDataSourceFunc: func(ctx context.Context, name string, params models.DataSourceGetParams) (*models.DataSource, error) {
			return expectedDataSource, nil
		},
	}
	service := NewDataSourceService(mockRepo, MOCK_ORGANIZATION_NAME)

	dataSource, err := service.Get(context.Background(), "test", models.DataSourceGetParams{})
	assert.NoError(t, err)
	assert.Equal(t, expectedDataSource, dataSource)
}

func TestDataSourceService_Get_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		GetDataSourceFunc: func(ctx context.Context, name string, params models.DataSourceGetParams) (*models.DataSource, error) {
			return nil, assert.AnError
		},
	}
	service := NewDataSourceService(mockRepo, MOCK_ORGANIZATION_NAME)

	dataSource, err := service.Get(context.Background(), "test", models.DataSourceGetParams{})
	assert.Error(t, err)
	assert.Nil(t, dataSource)
}

func TestDataSourceService_List_Success(t *testing.T) {
	expectedDataSourceList := []*models.DataSource{}
	mockRepo := &mock.MockDBClient{
		ListDataSourceFunc: func(ctx context.Context, params models.DataSourceListParams) ([]*models.DataSource, error) {
			return expectedDataSourceList, nil
		},
	}
	service := NewDataSourceService(mockRepo, MOCK_ORGANIZATION_NAME)

	dataSourceList, err := service.List(context.Background(), models.DataSourceListParams{})
	assert.NoError(t, err)
	assert.Equal(t, expectedDataSourceList, dataSourceList)
}

func TestDataSourceService_List_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		ListDataSourceFunc: func(ctx context.Context, params models.DataSourceListParams) ([]*models.DataSource, error) {
			return nil, assert.AnError
		},
	}
	service := NewDataSourceService(mockRepo, MOCK_ORGANIZATION_NAME)

	dataSourceList, err := service.List(context.Background(), models.DataSourceListParams{})
	assert.Error(t, err)
	assert.Nil(t, dataSourceList)
}
