package services

import (
	"context"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository"
)

type DetectionSectionService interface {
	Create(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error
	Update(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error
	Delete(ctx context.Context, datasourceName string, id string) error
	Get(ctx context.Context, datasourceName string, id string) (*models.DetectionSection, error)
	List(ctx context.Context, datasourceName string) ([]*models.DetectionSection, error)
	CreateMany(ctx context.Context, datasourceName string, detectionSection []*models.DetectionSection) error
}

type detectionSectionService struct {
	repo repository.DBClient
}

func NewDetectionSectionService(repo repository.DBClient) DetectionSectionService {
	return &detectionSectionService{repo: repo}
}

func (s *detectionSectionService) Create(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error {
	return s.repo.CreateDetectionSection(ctx, datasourceName, detectionSection)
}

func (s *detectionSectionService) Update(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error {
	return s.repo.UpdateDetectionSection(ctx, datasourceName, detectionSection)
}

func (s *detectionSectionService) Delete(ctx context.Context, datasourceName string, id string) error {
	return s.repo.DeleteDetectionSection(ctx, datasourceName, id)
}

func (s *detectionSectionService) Get(ctx context.Context, datasourceName string, id string) (*models.DetectionSection, error) {
	return s.repo.GetDetectionSection(ctx, datasourceName, id)
}

func (s *detectionSectionService) List(ctx context.Context, datasourceName string) ([]*models.DetectionSection, error) {
	return s.repo.ListDetectionSection(ctx, datasourceName)
}

func (s *detectionSectionService) CreateMany(ctx context.Context, datasourceName string, detectionSections []*models.DetectionSection) error {
	return s.repo.CreateManyDetectionSection(ctx, datasourceName, detectionSections)
}
