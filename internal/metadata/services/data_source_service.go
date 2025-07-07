package services

import (
	"context"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository"
)

type DataSourceService interface {
	Create(ctx context.Context, dataSource *models.DataSource) error
	Update(ctx context.Context, dataSource *models.UpdateDataSource) (*models.DataSource, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string, params models.DataSourceGetParams) (*models.DataSource, error)
	List(ctx context.Context, params models.DataSourceListParams) ([]*models.DataSource, error)
}
type dataSourceService struct {
	repo             repository.DBClient
	organizationName string
}

func NewDataSourceService(repo repository.DBClient, _organizationName string) DataSourceService {
	return &dataSourceService{repo: repo, organizationName: _organizationName}
}

func (s *dataSourceService) Create(ctx context.Context, dataSource *models.DataSource) error {
	return s.repo.CreateDataSource(ctx, dataSource)
}

func (s *dataSourceService) Update(ctx context.Context, dataSource *models.UpdateDataSource) (*models.DataSource, error) {
	err := s.repo.UpdateDataSource(ctx, dataSource)
	if err != nil {
		return nil, err
	}

	return s.Get(ctx, dataSource.Name, models.DataSourceGetParams{})
}

func (s *dataSourceService) Delete(ctx context.Context, name string) error {
	return s.repo.DeleteDataSource(ctx, name)
}

func (s *dataSourceService) Get(ctx context.Context, name string, params models.DataSourceGetParams) (*models.DataSource, error) {
	return s.repo.GetDataSource(ctx, name, params)
}

func (s *dataSourceService) List(ctx context.Context, params models.DataSourceListParams) ([]*models.DataSource, error) {
	return s.repo.ListDataSource(ctx, params, s.organizationName)
}
