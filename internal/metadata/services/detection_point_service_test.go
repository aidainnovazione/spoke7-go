package services

import (
	"context"
	"testing"

	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository/mock"

	"github.com/stretchr/testify/assert"
)

func TestDetectionPointService_Create_Success(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		CreateDetectionPointFunc: func(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error {
			return nil
		},
	}
	service := NewDetectionPointService(mockRepo)
	ctx := context.Background()
	detectionPoint := &models.DetectionPoint{
		Id: "1",

		Description: "Test Description",
	}

	err := service.Create(ctx, "DS1", detectionPoint)
	assert.NoError(t, err)

}

func TestDetectionPointService_Create_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		CreateDetectionPointFunc: func(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error {
			return assert.AnError
		},
	}
	service := NewDetectionPointService(mockRepo)
	ctx := context.Background()
	detectionPoint := &models.DetectionPoint{}

	err := service.Create(ctx, "DS1", detectionPoint)
	assert.Error(t, err)
}

func TestDetectionPointService_Update_Success(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		UpdateDetectionPointFunc: func(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error {
			return nil
		},
	}
	service := NewDetectionPointService(mockRepo)
	ctx := context.Background()
	detectionPoint := &models.DetectionPoint{}

	err := service.Update(ctx, "DS1", detectionPoint)
	assert.NoError(t, err)

}

func TestDetectionPointService_Update_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		UpdateDetectionPointFunc: func(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error {
			return assert.AnError
		},
	}
	service := NewDetectionPointService(mockRepo)
	ctx := context.Background()
	detectionPoint := &models.DetectionPoint{}

	err := service.Update(ctx, "DS1", detectionPoint)
	assert.Error(t, err)
}

func TestDetectionPointService_Delete_Success(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		DeleteDetectionPointFunc: func(ctx context.Context, datasourceName string, id string) error {
			return nil
		},
	}
	service := NewDetectionPointService(mockRepo)
	ctx := context.Background()
	id := "test-id"

	err := service.Delete(ctx, "DS1", id)
	assert.NoError(t, err)
}

func TestDetectionPointService_Delete_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		DeleteDetectionPointFunc: func(ctx context.Context, datasourceName string, id string) error {
			return assert.AnError
		},
	}
	service := NewDetectionPointService(mockRepo)
	ctx := context.Background()
	id := "test-id"

	err := service.Delete(ctx, "DS1", id)
	assert.Error(t, err)
}

func TestDetectionPointService_Get_Success(t *testing.T) {
	expectedDetectionPoint := &models.DetectionPoint{
		Id: "1",

		Description: "Test Description",
	}

	mockRepo := &mock.MockDBClient{
		GetDetectionPointFunc: func(ctx context.Context, datasourceName string, id string) (*models.DetectionPoint, error) {
			return expectedDetectionPoint, nil
		},
	}
	service := NewDetectionPointService(mockRepo)
	ctx := context.Background()
	id := "test-id"

	result, err := service.Get(ctx, "DS1", id)
	assert.NoError(t, err)
	assert.Equal(t, expectedDetectionPoint, result)
}

func TestDetectionPointService_Get_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		GetDetectionPointFunc: func(ctx context.Context, datasourceName string, id string) (*models.DetectionPoint, error) {
			return nil, assert.AnError
		},
	}
	service := NewDetectionPointService(mockRepo)
	ctx := context.Background()
	id := "test-id"

	result, err := service.Get(ctx, "DS1", id)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestDetectionPointService_List_Success(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		ListDetectionPointFunc: func(ctx context.Context, datasourceName string) ([]*models.DetectionPoint, error) {
			return []*models.DetectionPoint{
				{
					Id: "1",

					Description: "Test Description",
				},
				{
					Id: "2",

					Description: "Test Description 2",
				},
			}, nil
		},
	}
	service := NewDetectionPointService(mockRepo)
	ctx := context.Background()

	result, err := service.List(ctx, "DS1")
	assert.NoError(t, err)
	assert.Len(t, result, 2)

}

func TestDetectionPointService_List_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		ListDetectionPointFunc: func(ctx context.Context, datasourceName string) ([]*models.DetectionPoint, error) {
			return nil, assert.AnError
		},
	}
	service := NewDetectionPointService(mockRepo)
	ctx := context.Background()

	result, err := service.List(ctx, "DS1")
	assert.Error(t, err)
	assert.Nil(t, result)
}
func TestDetectionPointService_BulkCreate_Success(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		BulkCreateDetectionPointFunc: func(ctx context.Context, datasourceName string, detectionPoints []*models.DetectionPoint) error {
			return nil
		},
	}
	service := NewDetectionPointService(mockRepo)
	ctx := context.Background()
	detectionPoints := []*models.DetectionPoint{
		{
			Id: "1",

			Description: "Test Description 1",
		},
		{
			Id: "2",

			Description: "Test Description 2",
		},
	}

	err := service.BulkCreate(ctx, "DS1", detectionPoints)
	assert.NoError(t, err)
}

func TestDetectionPointService_BulkCreate_Fail(t *testing.T) {
	mockRepo := &mock.MockDBClient{
		BulkCreateDetectionPointFunc: func(ctx context.Context, datasourceName string, detectionPoints []*models.DetectionPoint) error {
			return assert.AnError
		},
	}
	service := NewDetectionPointService(mockRepo)
	ctx := context.Background()
	detectionPoints := []*models.DetectionPoint{
		{
			Id: "1",

			Description: "Test Description 1",
		},
		{
			Id: "2",

			Description: "Test Description 2",
		},
	}

	err := service.BulkCreate(ctx, "DS1", detectionPoints)
	assert.Error(t, err)
}
