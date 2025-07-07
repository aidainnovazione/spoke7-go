package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type CurrentTrafficDataByDetectionSectionService interface {
	Create(ctx context.Context, currentTrafficDataByDetectionSection *models.CurrentTrafficDataByDetectionSectionModel) error
	Update(ctx context.Context, currentTrafficDataByDetectionSection *models.CurrentTrafficDataByDetectionSectionModel) error
	Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error
	Get(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.CurrentTrafficDataByDetectionSectionModel, error)
	List(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.CurrentTrafficDataByDetectionSectionModel, error)
	BulkCreate(ctx context.Context, dcurrentTrafficDataByDetectionSections []*models.CurrentTrafficDataByDetectionSectionModel) error
	Statistics(ctx context.Context, dataSourceName string, detectionSectionIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.CurrentTrafficDataByDetectionSectionStatisticsByDatasource, error)
	ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.CurrentTrafficDataByDetectionSectionModel, error)
}

type currentTrafficDataByDetectionSectionService struct {
	repo                repository.DBClient
	grpcMetadataService grpc_client.GrpcMetadataClient
	organizationName    string
}

func NewCurrentTrafficDataByDetectionSectionService(repo repository.DBClient, grpcMetadataService grpc_client.GrpcMetadataClient, organizationName string) CurrentTrafficDataByDetectionSectionService {
	return &currentTrafficDataByDetectionSectionService{repo: repo, grpcMetadataService: grpcMetadataService, organizationName: organizationName}
}

func (s *currentTrafficDataByDetectionSectionService) Create(ctx context.Context, model *models.CurrentTrafficDataByDetectionSectionModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.CreateCurrentTrafficDataByDetectionSection(ctx, model)
}

func (s *currentTrafficDataByDetectionSectionService) Update(ctx context.Context, model *models.CurrentTrafficDataByDetectionSectionModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.UpdateCurrentTrafficDataByDetectionSection(ctx, model)
}

func (s *currentTrafficDataByDetectionSectionService) Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.DeleteCurrentTrafficDataByDetectionSection(ctx, params)
}

func (s *currentTrafficDataByDetectionSectionService) Get(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.CurrentTrafficDataByDetectionSectionModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.GetCurrentTrafficDataByDetectionSection(ctx, params)
}

func (s *currentTrafficDataByDetectionSectionService) List(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.CurrentTrafficDataByDetectionSectionModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.ListCurrentTrafficDataByDetectionSection(ctx, params)
}

func (s *currentTrafficDataByDetectionSectionService) BulkCreate(ctx context.Context, models []*models.CurrentTrafficDataByDetectionSectionModel) error {
	for _, model := range models {
		if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
			return err
		}
	}
	return s.repo.BulkCreateCurrentTrafficDataByDetectionSection(ctx, models)
}

func (s *currentTrafficDataByDetectionSectionService) Statistics(ctx context.Context, dataSourceName string, detectionSectionIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.CurrentTrafficDataByDetectionSectionStatisticsByDatasource, error) {
	aggregate := models.CurrentTrafficDataByDetectionSectionStatisticsByDatasource{}

	aggregatedStats, err := s.repo.ListCurrentTrafficDataByDetectionSectionAggregatedBySection(ctx, dataSourceName, detectionSectionIDs, startTimestamp, endTimestamp)
	if err != nil {
		return aggregate, fmt.Errorf("failed to retrieve aggregated traffic statistics: %w", err)
	}

	var statsList []models.CurrentTrafficDataByDetectionSectionStatistics
	for _, agg := range aggregatedStats {
		statsList = append(statsList, models.CurrentTrafficDataByDetectionSectionStatistics{
			DetectionSectionID:                 agg.DetectionSectionID,
			RecordsCount:                       agg.RecordsCount,
			FirstRecordTimestamp:               agg.FirstRecordTimestamp,
			LastRecordTimestamp:                agg.LastRecordTimestamp,
			TotalAverageForwardSpeed:           agg.TotalAverageForwardSpeed,
			TotalAverageBackwardSpeed:          agg.TotalAverageBackwardSpeed,
			PercentageRecordsWithForwardSpeed:  (agg.PercentageRecordsWithForwardSpeed),
			PercentageRecordsWithBackwardSpeed: (agg.PercentageRecordsWithBackwardSpeed),
		})
	}

	aggregate.StatisticsBySection = statsList
	aggregate.RecordsCount, aggregate.FirstRecordTimestamp, aggregate.LastRecordTimestamp = AggregateStatistics(statsList)

	return aggregate, nil
}

func (s *currentTrafficDataByDetectionSectionService) ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.CurrentTrafficDataByDetectionSectionModel, error) {
	//query db for daily basis stats
	return s.repo.ListCurrentTrafficDataByDetectionSectionAggregatedByDay(ctx, dataSourceName, detectionSectionIDs, fromTime, toTime)
}
