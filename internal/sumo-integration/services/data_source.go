package services

import (
	"context"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/pkg/grpc_client"
)

type DataSourceService interface {
	CreateRoadNetwork(ctx context.Context, network models.RoadNetwork) (*models.RoadNetwork, error)
	CreateAndAssociateRoadNetwork(ctx context.Context, datasource string, network models.RoadNetwork) (*models.RoadNetwork, error)
	GetRoadNetwork(ctx context.Context, id string) (*models.RoadNetwork, error)
	CreateManyDetectionPoints(ctx context.Context, datasource string, dp []models.DetectionPoint) ([]models.DetectionPoint, error)
}

type dataSourceService struct {
	// grpcDataService     grpc_client.GrpcDataClient
	grpcMetadataService grpc_client.GrpcMetadataClient
}

func NewDataSourceService(grpcDataService grpc_client.GrpcDataClient, grpcMetadataService grpc_client.GrpcMetadataClient) DataSourceService {
	return &dataSourceService{grpcMetadataService: grpcMetadataService}
}

func (s *dataSourceService) CreateRoadNetwork(ctx context.Context, network models.RoadNetwork) (*models.RoadNetwork, error) {

	//2. CREATE THE ROAD NETWORK
	roadNetwork, err := s.grpcMetadataService.CreateRoadNetwork(ctx, &network)
	if err != nil {
		return nil, err
	}

	//4. return the road network
	return roadNetwork, nil

}

func (s *dataSourceService) CreateAndAssociateRoadNetwork(ctx context.Context, datasourceName string, network models.RoadNetwork) (*models.RoadNetwork, error) {

	//1. check if the datasource exists
	datasource, err := s.grpcMetadataService.GetDataSource(ctx, datasourceName)
	if err != nil {
		return nil, err
	}

	//2. CREATE THE ROAD NETWORK
	roadNetwork, err := s.grpcMetadataService.CreateRoadNetwork(ctx, &network)
	if err != nil {
		return nil, err
	}

	//3. UPDATE DATASOURCE WITH THE ROAD NETWORK
	datasource.RoadNetworkId = &roadNetwork.ID

	_, err = s.grpcMetadataService.UpdateDataSource(ctx, datasource)
	if err != nil {
		return nil, err
	}
	//4. return the road network
	return roadNetwork, nil

}

func (s *dataSourceService) GetRoadNetwork(ctx context.Context, id string) (*models.RoadNetwork, error) {

	return s.grpcMetadataService.GetRoadNetwork(ctx, id)

}

func (s *dataSourceService) CreateManyDetectionPoints(ctx context.Context, datasource string, dp []models.DetectionPoint) ([]models.DetectionPoint, error) {

	return s.grpcMetadataService.CreateManyDetectionPoints(ctx, datasource, dp)

}
