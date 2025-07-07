package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type RealTimeTrafficDataByDetectionPointByLaneService interface {
	Create(ctx context.Context, realTimeTrafficDataByDetectionPointByLane *models.RealTimeTrafficDataByDetectionPointByLaneModel) error
	Update(ctx context.Context, realTimeTrafficDataByDetectionPointByLane *models.RealTimeTrafficDataByDetectionPointByLaneModel) error
	Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error
	Get(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error)
	List(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error)
	BulkCreate(ctx context.Context, realTimeTrafficDataByDetectionPointByLanes []*models.RealTimeTrafficDataByDetectionPointByLaneModel) error
	Statistics(ctx context.Context, dataSourceName string, laneIds []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionPointByLaneStatisticsByDatasource, error)
	ListAggregatedByDay(ctx context.Context, dataSourceName string, laneIds []string, fromTime *time.Time, toTime *time.Time) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error)
}

type realTimeTrafficDataByDetectionPointByLaneService struct {
	repo                repository.DBClient
	grpcMetadataService grpc_client.GrpcMetadataClient
	organizationName    string
}

func NewRealTimeTrafficDataByDetectionPointByLaneService(repo repository.DBClient, grpcMetadataService grpc_client.GrpcMetadataClient, organizationName string) RealTimeTrafficDataByDetectionPointByLaneService {
	return &realTimeTrafficDataByDetectionPointByLaneService{repo: repo, grpcMetadataService: grpcMetadataService, organizationName: organizationName}
}

func (s *realTimeTrafficDataByDetectionPointByLaneService) Create(ctx context.Context, model *models.RealTimeTrafficDataByDetectionPointByLaneModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.CreateRealTimeTrafficDataByDetectionPointByLane(ctx, model)
}

func (s *realTimeTrafficDataByDetectionPointByLaneService) Update(ctx context.Context, model *models.RealTimeTrafficDataByDetectionPointByLaneModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.UpdateRealTimeTrafficDataByDetectionPointByLane(ctx, model)
}

func (s *realTimeTrafficDataByDetectionPointByLaneService) Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.DeleteRealTimeTrafficDataByDetectionPointByLane(ctx, params)
}

func (s *realTimeTrafficDataByDetectionPointByLaneService) Get(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.GetRealTimeTrafficDataByDetectionPointByLane(ctx, params)
}

func (s *realTimeTrafficDataByDetectionPointByLaneService) List(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.ListRealTimeTrafficDataByDetectionPointByLane(ctx, params)
}

func (s *realTimeTrafficDataByDetectionPointByLaneService) BulkCreate(ctx context.Context, models []*models.RealTimeTrafficDataByDetectionPointByLaneModel) error {
	for _, model := range models {
		if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
			return err
		}
	}
	return s.repo.BulkCreateRealTimeTrafficDataByDetectionPointByLane(ctx, models)
}

func (s *realTimeTrafficDataByDetectionPointByLaneService) Statistics(ctx context.Context, dataSourceName string, laneIds []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionPointByLaneStatisticsByDatasource, error) {
	aggregate := models.TrafficDataByDetectionPointByLaneStatisticsByDatasource{}

	aggregatedStats, err := s.repo.ListRealTimeTrafficDataByDetectionPointByLaneAggregatedByLane(ctx, dataSourceName, laneIds, startTimestamp, endTimestamp)
	if err != nil {
		return aggregate, fmt.Errorf("failed to retrieve aggregated traffic statistics: %w", err)
	}

	var statsList []models.TrafficStatisticsLane
	for _, agg := range aggregatedStats {
		statsList = append(statsList, models.TrafficStatisticsLane{
			LaneID:               agg.LaneID,
			RecordsCount:         agg.RecordsCount,
			FirstRecordTimestamp: agg.FirstRecordTimestamp,
			LastRecordTimestamp:  agg.LastRecordTimestamp,
		})
	}

	aggregate.StatisticsByLane = statsList
	aggregate.RecordsCount, aggregate.FirstRecordTimestamp, aggregate.LastRecordTimestamp = AggregateStatistics(statsList)

	return aggregate, nil
	return aggregate, nil
}

func (s *realTimeTrafficDataByDetectionPointByLaneService) ListAggregatedByDay(ctx context.Context, dataSourceName string, laneIds []string, fromTime *time.Time, toTime *time.Time) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error) {
	//query db for daily basis stats
	return s.repo.ListRealTimeTrafficDataByDetectionPointByLaneAggregatedByDay(ctx, dataSourceName, laneIds, fromTime, toTime)
}
