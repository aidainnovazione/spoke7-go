package mocks

import (
	"context"
	"spoke7-go/internal/metadata/models"

	"github.com/stretchr/testify/mock"
)

// MockDataSourceService implements DataSourceService for testing purposes
type MockDataSourceService struct {
	mock.Mock
}

func (m *MockDataSourceService) Create(ctx context.Context, dataSource *models.DataSource) error {
	args := m.Called(ctx, dataSource)
	return args.Error(0)
}

func (m *MockDataSourceService) Update(ctx context.Context, dataSource *models.UpdateDataSource) (*models.DataSource, error) {
	args := m.Called(ctx, dataSource)
	if args.Get(0) != nil {
		return args.Get(0).(*models.DataSource), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataSourceService) Delete(ctx context.Context, name string) error {
	args := m.Called(ctx, name)
	return args.Error(0)
}

func (m *MockDataSourceService) Get(ctx context.Context, name string, params models.DataSourceGetParams) (*models.DataSource, error) {
	args := m.Called(ctx, name, params)
	if args.Get(0) != nil {
		return args.Get(0).(*models.DataSource), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataSourceService) List(ctx context.Context, params models.DataSourceListParams) ([]*models.DataSource, error) {
	args := m.Called(ctx, params)
	if args.Get(0) != nil {
		return args.Get(0).([]*models.DataSource), args.Error(1)
	}
	return nil, args.Error(1)
}
