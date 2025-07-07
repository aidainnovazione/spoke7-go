package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type HistoryHourTrafficDataByDetectionPointService interface {
	Create(ctx context.Context, historyHourTrafficDataByDetectionPoint *models.HistoryHourTrafficDataByDetectionPointModel) error
	Update(ctx context.Context, historyHourTrafficDataByDetectionPoint *models.HistoryHourTrafficDataByDetectionPointModel) error
	Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointParams) error
	Get(ctx context.Context, params models.GetTrafficDataByDetectionPointParams) ([]*models.HistoryHourTrafficDataByDetectionPointModel, error)
	List(ctx context.Context, params models.ListTrafficDataByDetectionPointParams) ([]*models.HistoryHourTrafficDataByDetectionPointModel, error)
	BulkCreate(ctx context.Context, historyHourTrafficDataByDetectionPoints []*models.HistoryHourTrafficDataByDetectionPointModel) error
	Statistics(ctx context.Context, dataSourceName string, detectionPointIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionPointStatisticsByDatasource, error)
	ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionPointIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.HistoryHourTrafficDataByDetectionPointModel, error)
}

type historyHourTrafficDataByDetectionPointService struct {
	repo                repository.DBClient
	grpcMetadataService grpc_client.GrpcMetadataClient
	organizationName    string
}

func NewHistoryHourTrafficDataByDetectionPointService(repo repository.DBClient, grpcMetadataService grpc_client.GrpcMetadataClient, organizationName string) HistoryHourTrafficDataByDetectionPointService {
	return &historyHourTrafficDataByDetectionPointService{repo: repo, grpcMetadataService: grpcMetadataService, organizationName: organizationName}
}

func (s *historyHourTrafficDataByDetectionPointService) Create(ctx context.Context, model *models.HistoryHourTrafficDataByDetectionPointModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.CreateHistoryHourTrafficDataByDetectionPoint(ctx, model)
}

func (s *historyHourTrafficDataByDetectionPointService) Update(ctx context.Context, model *models.HistoryHourTrafficDataByDetectionPointModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.UpdateHistoryHourTrafficDataByDetectionPoint(ctx, model)
}

func (s *historyHourTrafficDataByDetectionPointService) Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointParams) error {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.DeleteHistoryHourTrafficDataByDetectionPoint(ctx, params)
}

func (s *historyHourTrafficDataByDetectionPointService) Get(ctx context.Context, params models.GetTrafficDataByDetectionPointParams) ([]*models.HistoryHourTrafficDataByDetectionPointModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.GetHistoryHourTrafficDataByDetectionPoint(ctx, params)
}

func (s *historyHourTrafficDataByDetectionPointService) List(ctx context.Context, params models.ListTrafficDataByDetectionPointParams) ([]*models.HistoryHourTrafficDataByDetectionPointModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.ListHistoryHourTrafficDataByDetectionPoint(ctx, params)
}

func (s *historyHourTrafficDataByDetectionPointService) BulkCreate(ctx context.Context, models []*models.HistoryHourTrafficDataByDetectionPointModel) error {
	for _, model := range models {
		if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
			return err
		}
	}
	return s.repo.BulkCreateHistoryHourTrafficDataByDetectionPoint(ctx, models)
}

func (s *historyHourTrafficDataByDetectionPointService) Statistics(ctx context.Context, dataSourceName string, detectionPointIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionPointStatisticsByDatasource, error) {
	aggregate := models.TrafficDataByDetectionPointStatisticsByDatasource{}

	aggregatedStats, err := s.repo.ListHistoryHourTrafficDataByDetectionPointAggregatedByPoint(ctx, dataSourceName, detectionPointIDs, startTimestamp, endTimestamp)
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

func (s *historyHourTrafficDataByDetectionPointService) ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionPointIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.HistoryHourTrafficDataByDetectionPointModel, error) {
	//query db for daily basis stats
	return s.repo.ListHistoryHourTrafficDataByDetectionPointAggregatedByDay(ctx, dataSourceName, detectionPointIDs, fromTime, toTime)
}
