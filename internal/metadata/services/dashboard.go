package services

import (
	"context"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository"
)

type DashboardService interface {
	Create(ctx context.Context, dashboard *models.Dashboard) error
	Update(ctx context.Context, dashboard *models.Dashboard) (*models.Dashboard, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*models.Dashboard, error)
	List(ctx context.Context, dataSourceName string) ([]*models.Dashboard, error)
}
type dashboardService struct {
	repo             repository.DBClient
	organizationName string
}

func NewDashboardService(repo repository.DBClient, _organizationName string) DashboardService {
	return &dashboardService{repo: repo, organizationName: _organizationName}
}

func (s *dashboardService) Create(ctx context.Context, dashboard *models.Dashboard) error {
	return s.repo.CreateDashboard(ctx, dashboard)
}

func (s *dashboardService) Update(ctx context.Context, dashboard *models.Dashboard) (*models.Dashboard, error) {
	err := s.repo.UpdateDashboard(ctx, dashboard)
	if err != nil {
		return nil, err
	}

	return s.Get(ctx, dashboard.ID)
}

func (s *dashboardService) Delete(ctx context.Context, id string) error {
	return s.repo.DeleteDashboard(ctx, id)
}

func (s *dashboardService) Get(ctx context.Context, id string) (*models.Dashboard, error) {
	return s.repo.GetDashboard(ctx, id)
}

func (s *dashboardService) List(ctx context.Context, dataSourceName string) ([]*models.Dashboard, error) {
	return s.repo.ListDashboard(ctx, dataSourceName, s.organizationName)
}
