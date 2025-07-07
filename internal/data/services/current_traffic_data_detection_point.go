package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type CurrentTrafficDataByDetectionPointService interface {
	Create(ctx context.Context, currentTrafficDataByDetectionPoint *models.CurrentTrafficDataByDetectionPointModel) error
	Update(ctx context.Context, currentTrafficDataByDetectionPoint *models.CurrentTrafficDataByDetectionPointModel) error
	Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointParams) error
	Get(ctx context.Context, params models.GetTrafficDataByDetectionPointParams) ([]*models.CurrentTrafficDataByDetectionPointModel, error)
	List(ctx context.Context, params models.ListTrafficDataByDetectionPointParams) ([]*models.CurrentTrafficDataByDetectionPointModel, error)
	BulkCreate(ctx context.Context, dcurrentTrafficDataByDetectionPoints []*models.CurrentTrafficDataByDetectionPointModel) error
	Statistics(ctx context.Context, dataSourceName string, detectionPointIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.CurrentTrafficDataByDetectionPointStatisticsByDatasource, error)
	ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionPointIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.CurrentTrafficDataByDetectionPointModel, error)
}

type currentTrafficDataByDetectionPointService struct {
	repo                repository.DBClient
	grpcMetadataService grpc_client.GrpcMetadataClient
	organizationName    string
}

func NewCurrentTrafficDataByDetectionPointService(repo repository.DBClient, grpcMetadataService grpc_client.GrpcMetadataClient, organizationName string) CurrentTrafficDataByDetectionPointService {
	return &currentTrafficDataByDetectionPointService{repo: repo, grpcMetadataService: grpcMetadataService, organizationName: organizationName}
}

func (s *currentTrafficDataByDetectionPointService) Create(ctx context.Context, model *models.CurrentTrafficDataByDetectionPointModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.CreateCurrentTrafficDataByDetectionPoint(ctx, model)
}

func (s *currentTrafficDataByDetectionPointService) Update(ctx context.Context, model *models.CurrentTrafficDataByDetectionPointModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.UpdateCurrentTrafficDataByDetectionPoint(ctx, model)
}

func (s *currentTrafficDataByDetectionPointService) Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointParams) error {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.DeleteCurrentTrafficDataByDetectionPoint(ctx, params)
}

func (s *currentTrafficDataByDetectionPointService) Get(ctx context.Context, params models.GetTrafficDataByDetectionPointParams) ([]*models.CurrentTrafficDataByDetectionPointModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.GetCurrentTrafficDataByDetectionPoint(ctx, params)
}

func (s *currentTrafficDataByDetectionPointService) List(ctx context.Context, params models.ListTrafficDataByDetectionPointParams) ([]*models.CurrentTrafficDataByDetectionPointModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.ListCurrentTrafficDataByDetectionPoint(ctx, params)
}

func (s *currentTrafficDataByDetectionPointService) BulkCreate(ctx context.Context, models []*models.CurrentTrafficDataByDetectionPointModel) error {
	for _, model := range models {
		if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
			return err
		}
	}
	return s.repo.BulkCreateCurrentTrafficDataByDetectionPoint(ctx, models)
}

func (s *currentTrafficDataByDetectionPointService) Statistics(ctx context.Context, dataSourceName string, detectionPointIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.CurrentTrafficDataByDetectionPointStatisticsByDatasource, error) {
	aggregate := models.CurrentTrafficDataByDetectionPointStatisticsByDatasource{}

	aggregatedStats, err := s.repo.ListCurrentTrafficDataByDetectionPointAggregatedByPoint(ctx, dataSourceName, detectionPointIDs, startTimestamp, endTimestamp)
	if err != nil {
		return aggregate, fmt.Errorf("failed to retrieve aggregated traffic statistics: %w", err)
	}

	var statsList []models.CurrentTrafficDataByDetectionPointStatistics
	for _, agg := range aggregatedStats {
		statsList = append(statsList, models.CurrentTrafficDataByDetectionPointStatistics{
			DetectionPointID:                          agg.DetectionPointID,
			RecordsCount:                              agg.RecordsCount,
			FirstRecordTimestamp:                      agg.FirstRecordTimestamp,
			LastRecordTimestamp:                       agg.LastRecordTimestamp,
			TotalCountEquivalentVehicles:              agg.TotalCountEquivalentVehicles,
			TotalHarmonicMeanSpeedAllRecords:          agg.TotalHarmonicMeanSpeedAllRecords,
			PercentageRecordsWithEquivalentCounts:     (agg.PercentageRecordsWithEquivalentCounts),
			PercentageRecordsWithPositiveAverageSpeed: (agg.PercentageRecordsWithPositiveAverageSpeed),
		})
	}

	aggregate.StatisticsByDetectionPoint = statsList
	aggregate.RecordsCount, aggregate.FirstRecordTimestamp, aggregate.LastRecordTimestamp = AggregateStatistics(statsList)

	return aggregate, nil
}

func (s *currentTrafficDataByDetectionPointService) ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionPointIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.CurrentTrafficDataByDetectionPointModel, error) {
	//query db for daily basis stats
	return s.repo.ListCurrentTrafficDataByDetectionPointAggregatedByDay(ctx, dataSourceName, detectionPointIDs, fromTime, toTime)
}
