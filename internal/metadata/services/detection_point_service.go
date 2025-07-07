package services

import (
	"context"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository"
)

type DetectionPointService interface {
	CreateMany(ctx context.Context, datasourceName string, detectionPoints []*models.DetectionPoint) error
	Create(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error
	Update(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error
	Delete(ctx context.Context, datasourceName string, id string) error
	DeleteAll(ctx context.Context, datasourceName string) error
	Get(ctx context.Context, datasourceName string, id string) (*models.DetectionPoint, error)
	List(ctx context.Context, datasourceName string) ([]*models.DetectionPoint, error)
	BulkCreate(ctx context.Context, datasourceName string, detectionPoints []*models.DetectionPoint) error
}

type detectionPointService struct {
	repo repository.DBClient
}

func NewDetectionPointService(repo repository.DBClient) DetectionPointService {
	return &detectionPointService{repo: repo}
}

func (s *detectionPointService) Create(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error {
	return s.repo.CreateDetectionPoint(ctx, datasourceName, detectionPoint)
}

func (s *detectionPointService) CreateMany(ctx context.Context, datasourceName string, detectionPoints []*models.DetectionPoint) error {
	return s.repo.CreateManyDetectionPoint(ctx, datasourceName, detectionPoints)
}

func (s *detectionPointService) Update(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error {
	return s.repo.UpdateDetectionPoint(ctx, datasourceName, detectionPoint)
}

func (s *detectionPointService) Delete(ctx context.Context, datasourceName string, id string) error {
	return s.repo.DeleteDetectionPoint(ctx, datasourceName, id)
}
func (s *detectionPointService) DeleteAll(ctx context.Context, datasourceName string) error {
	return s.repo.DeleteAllDetectionPointByDatasourceName(ctx, datasourceName)
}

func (s *detectionPointService) Get(ctx context.Context, datasourceName string, id string) (*models.DetectionPoint, error) {
	return s.repo.GetDetectionPoint(ctx, datasourceName, id)
}

func (s *detectionPointService) List(ctx context.Context, datasourceName string) ([]*models.DetectionPoint, error) {
	return s.repo.ListDetectionPoint(ctx, datasourceName)
}

func (s *detectionPointService) BulkCreate(ctx context.Context, datasourceName string, detectionPoints []*models.DetectionPoint) error {
	return s.repo.BulkCreateDetectionPoint(ctx, datasourceName, detectionPoints)
}
