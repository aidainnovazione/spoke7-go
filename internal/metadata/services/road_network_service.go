package services

import (
	"context"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository"
	"spoke7-go/pkg/logger"
)

type roadNetworkService struct {
	repo repository.DBClient
	log  logger.Logger
}

type RoadNetworkService interface {
	Create(ctx context.Context, roadNetwork *models.RoadNetwork) (*models.RoadNetwork, error)
	Update(ctx context.Context, RoadNetwork *models.RoadNetwork) (*models.RoadNetwork, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*models.RoadNetwork, error)
	List(ctx context.Context) ([]*models.RoadNetwork, error)
}

func NewRoadNetworkService(repo repository.DBClient, logger logger.Logger) RoadNetworkService {
	return &roadNetworkService{repo: repo, log: logger}
}

// CreateRoadNetwork implements pb.RoadNetworkServiceServer.
func (r *roadNetworkService) Create(ctx context.Context, roadNetwork *models.RoadNetwork) (*models.RoadNetwork, error) {
	return r.repo.CreateRoadNetwork(ctx, roadNetwork)
}

func (s *roadNetworkService) Update(ctx context.Context, RoadNetwork *models.RoadNetwork) (*models.RoadNetwork, error) {
	return s.repo.UpdateRoadNetwork(ctx, RoadNetwork)
}

func (s *roadNetworkService) Delete(ctx context.Context, id string) error {
	return s.repo.DeleteRoadNetwork(ctx, id)
}

func (s *roadNetworkService) Get(ctx context.Context, id string) (*models.RoadNetwork, error) {
	return s.repo.GetRoadNetworkByID(ctx, id)
}

func (s *roadNetworkService) List(ctx context.Context) ([]*models.RoadNetwork, error) {
	return s.repo.ListRoadNetworks(ctx)
}
