package mocks

import (
	"context"
	"spoke7-go/internal/metadata/models"

	"github.com/stretchr/testify/mock"
)

// MockDashboardService implements DashboardService for testing purposes
type MockDashboardService struct {
	mock.Mock
}

func (m *MockDashboardService) Create(ctx context.Context, dashboard *models.Dashboard) error {
	args := m.Called(ctx, dashboard)
	return args.Error(0)
}

func (m *MockDashboardService) Update(ctx context.Context, dashboard *models.Dashboard) (*models.Dashboard, error) {
	args := m.Called(ctx, dashboard)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Dashboard), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDashboardService) Delete(ctx context.Context, name string) error {
	args := m.Called(ctx, name)
	return args.Error(0)
}

func (m *MockDashboardService) Get(ctx context.Context, id string) (*models.Dashboard, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Dashboard), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDashboardService) List(ctx context.Context, dataSourceName string) ([]*models.Dashboard, error) {
	args := m.Called(ctx, dataSourceName)
	if args.Get(0) != nil {
		return args.Get(0).([]*models.Dashboard), args.Error(1)
	}
	return nil, args.Error(1)
}
