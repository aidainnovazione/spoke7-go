package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type HistoryDayTrafficDataByDetectionPointService interface {
	Create(ctx context.Context, historyDayTrafficDataByDetectionPoint *models.HistoryDayTrafficDataByDetectionPointModel) error
	Update(ctx context.Context, historyDayTrafficDataByDetectionPoint *models.HistoryDayTrafficDataByDetectionPointModel) error
	Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointParams) error
	Get(ctx context.Context, params models.GetTrafficDataByDetectionPointParams) ([]*models.HistoryDayTrafficDataByDetectionPointModel, error)
	List(ctx context.Context, params models.ListTrafficDataByDetectionPointParams) ([]*models.HistoryDayTrafficDataByDetectionPointModel, error)
	BulkCreate(ctx context.Context, historyDayTrafficDataByDetectionPoints []*models.HistoryDayTrafficDataByDetectionPointModel) error
	Statistics(ctx context.Context, dataSourceName string, detectionPointIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionPointStatisticsByDatasource, error)
	ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionPointIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.HistoryDayTrafficDataByDetectionPointModel, error)
}

type historyDayTrafficDataByDetectionPointService struct {
	repo                repository.DBClient
	grpcMetadataService grpc_client.GrpcMetadataClient
	organizationName    string
}

func NewHistoryDayTrafficDataByDetectionPointService(repo repository.DBClient, grpcMetadataService grpc_client.GrpcMetadataClient, organizationName string) HistoryDayTrafficDataByDetectionPointService {
	return &historyDayTrafficDataByDetectionPointService{repo: repo, grpcMetadataService: grpcMetadataService, organizationName: organizationName}
}

func (s *historyDayTrafficDataByDetectionPointService) Create(ctx context.Context, model *models.HistoryDayTrafficDataByDetectionPointModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.CreateHistoryDayTrafficDataByDetectionPoint(ctx, model)
}

func (s *historyDayTrafficDataByDetectionPointService) Update(ctx context.Context, model *models.HistoryDayTrafficDataByDetectionPointModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.UpdateHistoryDayTrafficDataByDetectionPoint(ctx, model)
}

func (s *historyDayTrafficDataByDetectionPointService) Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointParams) error {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.DeleteHistoryDayTrafficDataByDetectionPoint(ctx, params)
}

func (s *historyDayTrafficDataByDetectionPointService) Get(ctx context.Context, params models.GetTrafficDataByDetectionPointParams) ([]*models.HistoryDayTrafficDataByDetectionPointModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.GetHistoryDayTrafficDataByDetectionPoint(ctx, params)
}

func (s *historyDayTrafficDataByDetectionPointService) List(ctx context.Context, params models.ListTrafficDataByDetectionPointParams) ([]*models.HistoryDayTrafficDataByDetectionPointModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.ListHistoryDayTrafficDataByDetectionPoint(ctx, params)
}

func (s *historyDayTrafficDataByDetectionPointService) BulkCreate(ctx context.Context, models []*models.HistoryDayTrafficDataByDetectionPointModel) error {
	for _, model := range models {
		if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
			return err
		}
	}
	return s.repo.BulkCreateHistoryDayTrafficDataByDetectionPoint(ctx, models)
}

func (s *historyDayTrafficDataByDetectionPointService) Statistics(ctx context.Context, dataSourceName string, detectionPointIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionPointStatisticsByDatasource, error) {
	aggregate := models.TrafficDataByDetectionPointStatisticsByDatasource{}

	aggregatedStats, err := s.repo.ListHistoryDayTrafficDataByDetectionPointAggregatedByPoint(ctx, dataSourceName, detectionPointIDs, startTimestamp, endTimestamp)
	if err != nil {
		return aggregate, fmt.Errorf("failed to retrieve aggregated traffic statistics: %w", err)
	}

	var statsList []models.TrafficStatisticsDetectionPoint
	for _, agg := range aggregatedStats {
		statsList = append(statsList, models.TrafficStatisticsDetectionPoint{
			DetectionPointID:     agg.DetectionPointID,
			RecordsCount:         agg.RecordsCount,
			FirstRecordTimestamp: agg.FirstRecordTimestamp,
			LastRecordTimestamp:  agg.LastRecordTimestamp,
		})
	}

	aggregate.StatisticsByPoint = statsList
	aggregate.RecordsCount, aggregate.FirstRecordTimestamp, aggregate.LastRecordTimestamp = AggregateStatistics(statsList)

	return aggregate, nil
}

func (s *historyDayTrafficDataByDetectionPointService) ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionPointIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.HistoryDayTrafficDataByDetectionPointModel, error) {
	//query db for daily basis stats
	return s.repo.ListHistoryDayTrafficDataByDetectionPointAggregatedByDay(ctx, dataSourceName, detectionPointIDs, fromTime, toTime)
}
