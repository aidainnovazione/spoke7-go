package mocks

import (
	"context"
	"spoke7-go/internal/metadata/models"

	"github.com/stretchr/testify/mock"
)

type MockDetectionSectionService struct {
	mock.Mock
}

func (m *MockDetectionSectionService) Create(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error {
	args := m.Called(ctx, detectionSection)
	return args.Error(0)
}

func (m *MockDetectionSectionService) Update(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error {
	args := m.Called(ctx, detectionSection)
	return args.Error(0)
}

func (m *MockDetectionSectionService) Delete(ctx context.Context, datasourceName string, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockDetectionSectionService) Get(ctx context.Context, datasourceName string, id string) (*models.DetectionSection, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.DetectionSection), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDetectionSectionService) List(ctx context.Context, datasourceName string) ([]*models.DetectionSection, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).([]*models.DetectionSection), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDetectionSectionService) BulkCreate(ctx context.Context, datasourceName string, detectionSection []*models.DetectionSection) error {
	args := m.Called(ctx, datasourceName, detectionSection)
	return args.Error(0)
}

func (m *MockDetectionSectionService) CreateMany(ctx context.Context, datasourceName string, detectionSection []*models.DetectionSection) error {
	args := m.Called(ctx, detectionSection)
	return args.Error(0)
}
