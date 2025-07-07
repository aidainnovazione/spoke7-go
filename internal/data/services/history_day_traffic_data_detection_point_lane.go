package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type HistoryDayTrafficDataByDetectionPointByLaneService interface {
	Create(ctx context.Context, historyDayTrafficDataByDetectionPointByLane *models.HistoryDayTrafficDataByDetectionPointByLaneModel) error
	Update(ctx context.Context, historyDayTrafficDataByDetectionPointByLane *models.HistoryDayTrafficDataByDetectionPointByLaneModel) error
	Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error
	Get(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.HistoryDayTrafficDataByDetectionPointByLaneModel, error)
	List(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.HistoryDayTrafficDataByDetectionPointByLaneModel, error)
	BulkCreate(ctx context.Context, historyDayTrafficDataByDetectionPointByLanes []*models.HistoryDayTrafficDataByDetectionPointByLaneModel) error
	Statistics(ctx context.Context, dataSourceName string, laneIds []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionPointByLaneStatisticsByDatasource, error)
	ListAggregatedByDay(ctx context.Context, dataSourceName string, laneIds []string, fromTime *time.Time, toTime *time.Time) ([]*models.HistoryDayTrafficDataByDetectionPointByLaneModel, error)
}

type historyDayTrafficDataByDetectionPointByLaneService struct {
	repo                repository.DBClient
	grpcMetadataService grpc_client.GrpcMetadataClient
	organizationName    string
}

func NewHistoryDayTrafficDataByDetectionPointByLaneService(repo repository.DBClient, grpcMetadataService grpc_client.GrpcMetadataClient, organizationName string) HistoryDayTrafficDataByDetectionPointByLaneService {
	return &historyDayTrafficDataByDetectionPointByLaneService{repo: repo, grpcMetadataService: grpcMetadataService, organizationName: organizationName}
}

func (s *historyDayTrafficDataByDetectionPointByLaneService) Create(ctx context.Context, model *models.HistoryDayTrafficDataByDetectionPointByLaneModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.CreateHistoryDayTrafficDataByDetectionPointByLane(ctx, model)
}

func (s *historyDayTrafficDataByDetectionPointByLaneService) Update(ctx context.Context, model *models.HistoryDayTrafficDataByDetectionPointByLaneModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.UpdateHistoryDayTrafficDataByDetectionPointByLane(ctx, model)
}

func (s *historyDayTrafficDataByDetectionPointByLaneService) Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.DeleteHistoryDayTrafficDataByDetectionPointByLane(ctx, params)
}

func (s *historyDayTrafficDataByDetectionPointByLaneService) Get(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.HistoryDayTrafficDataByDetectionPointByLaneModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.GetHistoryDayTrafficDataByDetectionPointByLane(ctx, params)
}

func (s *historyDayTrafficDataByDetectionPointByLaneService) List(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.HistoryDayTrafficDataByDetectionPointByLaneModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.ListHistoryDayTrafficDataByDetectionPointByLane(ctx, params)
}

func (s *historyDayTrafficDataByDetectionPointByLaneService) BulkCreate(ctx context.Context, models []*models.HistoryDayTrafficDataByDetectionPointByLaneModel) error {
	for _, model := range models {
		if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
			return err
		}
	}
	return s.repo.BulkCreateHistoryDayTrafficDataByDetectionPointByLane(ctx, models)
}

func (s *historyDayTrafficDataByDetectionPointByLaneService) Statistics(ctx context.Context, dataSourceName string, laneIds []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionPointByLaneStatisticsByDatasource, error) {
	aggregate := models.TrafficDataByDetectionPointByLaneStatisticsByDatasource{}

	aggregatedStats, err := s.repo.ListHistoryDayTrafficDataByDetectionPointByLaneAggregatedByLane(ctx, dataSourceName, laneIds, startTimestamp, endTimestamp)
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

func (s *historyDayTrafficDataByDetectionPointByLaneService) ListAggregatedByDay(ctx context.Context, dataSourceName string, laneIds []string, fromTime *time.Time, toTime *time.Time) ([]*models.HistoryDayTrafficDataByDetectionPointByLaneModel, error) {
	//query db for daily basis stats
	return s.repo.ListHistoryDayTrafficDataByDetectionPointByLaneAggregatedByDay(ctx, dataSourceName, laneIds, fromTime, toTime)
}
