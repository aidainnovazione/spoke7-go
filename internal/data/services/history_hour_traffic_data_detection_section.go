package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type HistoryHourTrafficDataByDetectionSectionService interface {
	Create(ctx context.Context, historyHourTrafficDataByDetectionSection *models.HistoryTrafficDataByDetectionSectionModel) error
	Update(ctx context.Context, historyHourTrafficDataByDetectionSection *models.HistoryTrafficDataByDetectionSectionModel) error
	Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error
	Get(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error)
	List(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error)
	BulkCreate(ctx context.Context, dhistoryHourTrafficDataByDetectionSections []*models.HistoryTrafficDataByDetectionSectionModel) error
	Statistics(ctx context.Context, dataSourceName string, detectionSectionIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionSectionStatisticsByDatasource, error)
	ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.HistoryTrafficDataByDetectionSectionModel, error)
}

type historyHourTrafficDataByDetectionSectionService struct {
	repo                repository.DBClient
	grpcMetadataService grpc_client.GrpcMetadataClient
	organizationName    string
}

func NewHistoryHourTrafficDataByDetectionSectionService(repo repository.DBClient, grpcMetadataService grpc_client.GrpcMetadataClient, organizationName string) HistoryHourTrafficDataByDetectionSectionService {
	return &historyHourTrafficDataByDetectionSectionService{repo: repo, grpcMetadataService: grpcMetadataService, organizationName: organizationName}
}

func (s *historyHourTrafficDataByDetectionSectionService) Create(ctx context.Context, model *models.HistoryTrafficDataByDetectionSectionModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.CreateHistoryHourTrafficDataByDetectionSection(ctx, model)
}

func (s *historyHourTrafficDataByDetectionSectionService) Update(ctx context.Context, model *models.HistoryTrafficDataByDetectionSectionModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.UpdateHistoryHourTrafficDataByDetectionSection(ctx, model)
}

func (s *historyHourTrafficDataByDetectionSectionService) Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.DeleteHistoryHourTrafficDataByDetectionSection(ctx, params)
}

func (s *historyHourTrafficDataByDetectionSectionService) Get(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.GetHistoryHourTrafficDataByDetectionSection(ctx, params)
}

func (s *historyHourTrafficDataByDetectionSectionService) List(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.ListHistoryHourTrafficDataByDetectionSection(ctx, params)
}

func (s *historyHourTrafficDataByDetectionSectionService) BulkCreate(ctx context.Context, models []*models.HistoryTrafficDataByDetectionSectionModel) error {
	for _, model := range models {
		if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
			return err
		}
	}
	return s.repo.BulkCreateHistoryHourTrafficDataByDetectionSection(ctx, models)
}

func (s *historyHourTrafficDataByDetectionSectionService) Statistics(ctx context.Context, dataSourceName string, detectionSectionIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionSectionStatisticsByDatasource, error) {
	aggregate := models.TrafficDataByDetectionSectionStatisticsByDatasource{}

	aggregatedStats, err := s.repo.ListHistoryHourTrafficDataByDetectionSectionAggregatedBySection(ctx, dataSourceName, detectionSectionIDs, startTimestamp, endTimestamp)
	if err != nil {
		return aggregate, fmt.Errorf("failed to retrieve aggregated traffic statistics: %w", err)
	}

	var statsList []models.TrafficStatisticsDetectionSection
	for _, agg := range aggregatedStats {
		statsList = append(statsList, models.TrafficStatisticsDetectionSection{
			DetectionSectionID:   agg.DetectionSectionID,
			RecordsCount:         agg.RecordsCount,
			FirstRecordTimestamp: agg.FirstRecordTimestamp,
			LastRecordTimestamp:  agg.LastRecordTimestamp,
		})
	}

	aggregate.StatisticsBySection = statsList
	aggregate.RecordsCount, aggregate.FirstRecordTimestamp, aggregate.LastRecordTimestamp = AggregateStatistics(statsList)

	return aggregate, nil
}

func (s *historyHourTrafficDataByDetectionSectionService) ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.HistoryTrafficDataByDetectionSectionModel, error) {
	//query db for daily basis stats
	return s.repo.ListHistoryHourTrafficDataByDetectionSectionAggregatedByDay(ctx, dataSourceName, detectionSectionIDs, fromTime, toTime)
}
