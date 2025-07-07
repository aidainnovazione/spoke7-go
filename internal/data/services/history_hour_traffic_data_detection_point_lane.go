package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type HistoryHourTrafficDataByDetectionPointByLaneService interface {
	Create(ctx context.Context, historyHourTrafficDataByDetectionPointByLane *models.HistoryHourTrafficDataByDetectionPointByLaneModel) error
	Update(ctx context.Context, historyHourTrafficDataByDetectionPointByLane *models.HistoryHourTrafficDataByDetectionPointByLaneModel) error
	Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error
	Get(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.HistoryHourTrafficDataByDetectionPointByLaneModel, error)
	List(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.HistoryHourTrafficDataByDetectionPointByLaneModel, error)
	BulkCreate(ctx context.Context, historyHourTrafficDataByDetectionPointByLanes []*models.HistoryHourTrafficDataByDetectionPointByLaneModel) error
	Statistics(ctx context.Context, dataSourceName string, laneIds []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionPointByLaneStatisticsByDatasource, error)
	ListAggregatedByDay(ctx context.Context, dataSourceName string, laneIds []string, fromTime *time.Time, toTime *time.Time) ([]*models.HistoryHourTrafficDataByDetectionPointByLaneModel, error)
}

type historyHourTrafficDataByDetectionPointByLaneService struct {
	repo                repository.DBClient
	grpcMetadataService grpc_client.GrpcMetadataClient
	organizationName    string
}

func NewHistoryHourTrafficDataByDetectionPointByLaneService(repo repository.DBClient, grpcMetadataService grpc_client.GrpcMetadataClient, organizationName string) HistoryHourTrafficDataByDetectionPointByLaneService {
	return &historyHourTrafficDataByDetectionPointByLaneService{repo: repo, grpcMetadataService: grpcMetadataService, organizationName: organizationName}
}

func (s *historyHourTrafficDataByDetectionPointByLaneService) Create(ctx context.Context, model *models.HistoryHourTrafficDataByDetectionPointByLaneModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.CreateHistoryHourTrafficDataByDetectionPointByLane(ctx, model)
}

func (s *historyHourTrafficDataByDetectionPointByLaneService) Update(ctx context.Context, model *models.HistoryHourTrafficDataByDetectionPointByLaneModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.UpdateHistoryHourTrafficDataByDetectionPointByLane(ctx, model)
}

func (s *historyHourTrafficDataByDetectionPointByLaneService) Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.DeleteHistoryHourTrafficDataByDetectionPointByLane(ctx, params)
}

func (s *historyHourTrafficDataByDetectionPointByLaneService) Get(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.HistoryHourTrafficDataByDetectionPointByLaneModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.GetHistoryHourTrafficDataByDetectionPointByLane(ctx, params)
}

func (s *historyHourTrafficDataByDetectionPointByLaneService) List(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.HistoryHourTrafficDataByDetectionPointByLaneModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.ListHistoryHourTrafficDataByDetectionPointByLane(ctx, params)
}

func (s *historyHourTrafficDataByDetectionPointByLaneService) BulkCreate(ctx context.Context, models []*models.HistoryHourTrafficDataByDetectionPointByLaneModel) error {
	for _, model := range models {
		if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
			return err
		}
	}
	return s.repo.BulkCreateHistoryHourTrafficDataByDetectionPointByLane(ctx, models)
}

func (s *historyHourTrafficDataByDetectionPointByLaneService) Statistics(ctx context.Context, dataSourceName string, laneIds []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionPointByLaneStatisticsByDatasource, error) {
	aggregate := models.TrafficDataByDetectionPointByLaneStatisticsByDatasource{}

	aggregatedStats, err := s.repo.ListHistoryHourTrafficDataByDetectionPointByLaneAggregatedByLane(ctx, dataSourceName, laneIds, startTimestamp, endTimestamp)
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
}

func (s *historyHourTrafficDataByDetectionPointByLaneService) ListAggregatedByDay(ctx context.Context, dataSourceName string, laneIds []string, fromTime *time.Time, toTime *time.Time) ([]*models.HistoryHourTrafficDataByDetectionPointByLaneModel, error) {
	//query db for daily basis stats
	return s.repo.ListHistoryHourTrafficDataByDetectionPointByLaneAggregatedByDay(ctx, dataSourceName, laneIds, fromTime, toTime)
}
