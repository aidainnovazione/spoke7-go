package mocks

import (
	"context"
	"spoke7-go/internal/metadata/models"

	"github.com/stretchr/testify/mock"
)

// MockDetectionPointService implements DetectionPointService for testing purposes
type MockDetectionPointService struct {
	mock.Mock
}

func (m *MockDetectionPointService) Create(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error {
	args := m.Called(ctx, detectionPoint)
	return args.Error(0)
}

func (m *MockDetectionPointService) Update(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error {
	args := m.Called(ctx, detectionPoint)
	return args.Error(0)
}

func (m *MockDetectionPointService) Delete(ctx context.Context, datasourceName string, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockDetectionPointService) Get(ctx context.Context, datasourceName string, id string) (*models.DetectionPoint, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.DetectionPoint), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDetectionPointService) List(ctx context.Context, datasourceName string) ([]*models.DetectionPoint, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).([]*models.DetectionPoint), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDetectionPointService) BulkCreate(ctx context.Context, datasourceName string, detectionPoints []*models.DetectionPoint) error {
	args := m.Called(ctx, detectionPoints)
	return args.Error(0)
}

func (m *MockDetectionPointService) CreateMany(ctx context.Context, datasourceName string, detectionPoint []*models.DetectionPoint) error {
	args := m.Called(ctx, detectionPoint)
	return args.Error(0)
}
