package services

import (
	"context"
	"testing"

	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository/mock"

	"github.com/stretchr/testify/assert"
)

func TestDetectionSectionService_Create_Success(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		CreateDetectionSectionFunc: func(ctx context.Context, datasourceName string, DetectionSection *models.DetectionSection) error {
			return nil
		},
	}
	service := NewDetectionSectionService(mockRepo)

	err := service.Create(context.Background(), "DS1", &models.DetectionSection{})
	assert.NoError(t, err)
}

func TestDetectionSectionService_Create_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		CreateDetectionSectionFunc: func(ctx context.Context, datasourceName string, DetectionSection *models.DetectionSection) error {
			return assert.AnError
		},
	}
	service := NewDetectionSectionService(mockRepo)

	err := service.Create(context.Background(), "DS1", &models.DetectionSection{})
	assert.Error(t, err)
}

func TestDetectionSectionService_Update_Success(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		UpdateDetectionSectionFunc: func(ctx context.Context, datasourceName string, DetectionSection *models.DetectionSection) error {
			return nil
		},
	}
	service := NewDetectionSectionService(mockRepo)

	err := service.Update(context.Background(), "DS1", &models.DetectionSection{})
	assert.NoError(t, err)
}

func TestDetectionSectionService_Update_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		UpdateDetectionSectionFunc: func(ctx context.Context, datasourceName string, DetectionSection *models.DetectionSection) error {
			return assert.AnError
		},
	}
	service := NewDetectionSectionService(mockRepo)

	err := service.Update(context.Background(), "DS1", &models.DetectionSection{})
	assert.Error(t, err)
}

func TestDetectionSectionService_Delete_Success(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		DeleteDetectionSectionFunc: func(ctx context.Context, datasourceName string, name string) error {
			return nil
		},
	}
	service := NewDetectionSectionService(mockRepo)

	err := service.Delete(context.Background(), "DS1", "test")
	assert.NoError(t, err)
}

func TestDetectionSectionService_Delete_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		DeleteDetectionSectionFunc: func(ctx context.Context, datasourceName string, name string) error {
			return assert.AnError
		},
	}
	service := NewDetectionSectionService(mockRepo)

	err := service.Delete(context.Background(), "DS1", "test")
	assert.Error(t, err)
}

func TestDetectionSectionService_Get_Success(t *testing.T) {
	expectedDetectionSection := &models.DetectionSection{}
	mockRepo := &mock.MockDBClient{
		GetDetectionSectionFunc: func(ctx context.Context, datasourceName string, name string) (*models.DetectionSection, error) {
			return expectedDetectionSection, nil
		},
	}
	service := NewDetectionSectionService(mockRepo)

	DetectionSection, err := service.Get(context.Background(), "DS1", "test")
	assert.NoError(t, err)
	assert.Equal(t, expectedDetectionSection, DetectionSection)
}

func TestDetectionSectionService_Get_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		GetDetectionSectionFunc: func(ctx context.Context, datasourceName string, name string) (*models.DetectionSection, error) {
			return nil, assert.AnError
		},
	}
	service := NewDetectionSectionService(mockRepo)

	DetectionSection, err := service.Get(context.Background(), "DS1", "test")
	assert.Error(t, err)
	assert.Nil(t, DetectionSection)
}

func TestDetectionSectionService_List_Success(t *testing.T) {
	expectedDetectionSectionList := []*models.DetectionSection{}
	mockRepo := &mock.MockDBClient{
		ListDetectionSectionFunc: func(ctx context.Context, datasourceName string) ([]*models.DetectionSection, error) {
			return expectedDetectionSectionList, nil
		},
	}
	service := NewDetectionSectionService(mockRepo)

	DetectionSectionList, err := service.List(context.Background(), "DS1")
	assert.NoError(t, err)
	assert.Equal(t, expectedDetectionSectionList, DetectionSectionList)
}

func TestDetectionSectionService_List_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		ListDetectionSectionFunc: func(ctx context.Context, datasourceName string) ([]*models.DetectionSection, error) {
			return nil, assert.AnError
		},
	}
	service := NewDetectionSectionService(mockRepo)

	DetectionSectionList, err := service.List(context.Background(), "DS1")
	assert.Error(t, err)
	assert.Nil(t, DetectionSectionList)
}
