package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type HistoryDayTrafficDataByDetectionSectionService interface {
	Create(ctx context.Context, historyDayTrafficDataByDetectionSection *models.HistoryTrafficDataByDetectionSectionModel) error
	Update(ctx context.Context, historyDayTrafficDataByDetectionSection *models.HistoryTrafficDataByDetectionSectionModel) error
	Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error
	Get(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error)
	List(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error)
	BulkCreate(ctx context.Context, dhistoryDayTrafficDataByDetectionSections []*models.HistoryTrafficDataByDetectionSectionModel) error
	Statistics(ctx context.Context, dataSourceName string, detectionSectionIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionSectionStatisticsByDatasource, error)
	ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.HistoryTrafficDataByDetectionSectionModel, error)
}

type historyDayTrafficDataByDetectionSectionService struct {
	repo                repository.DBClient
	grpcMetadataService grpc_client.GrpcMetadataClient
	organizationName    string
}

func NewHistoryDayTrafficDataByDetectionSectionService(repo repository.DBClient, grpcMetadataService grpc_client.GrpcMetadataClient, organizationName string) HistoryDayTrafficDataByDetectionSectionService {
	return &historyDayTrafficDataByDetectionSectionService{repo: repo, grpcMetadataService: grpcMetadataService, organizationName: organizationName}
}

func (s *historyDayTrafficDataByDetectionSectionService) Create(ctx context.Context, model *models.HistoryTrafficDataByDetectionSectionModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.CreateHistoryDayTrafficDataByDetectionSection(ctx, model)
}

func (s *historyDayTrafficDataByDetectionSectionService) Update(ctx context.Context, model *models.HistoryTrafficDataByDetectionSectionModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.UpdateHistoryDayTrafficDataByDetectionSection(ctx, model)
}

func (s *historyDayTrafficDataByDetectionSectionService) Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.DeleteHistoryDayTrafficDataByDetectionSection(ctx, params)
}

func (s *historyDayTrafficDataByDetectionSectionService) Get(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.GetHistoryDayTrafficDataByDetectionSection(ctx, params)
}

func (s *historyDayTrafficDataByDetectionSectionService) List(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.ListHistoryDayTrafficDataByDetectionSection(ctx, params)
}

func (s *historyDayTrafficDataByDetectionSectionService) BulkCreate(ctx context.Context, models []*models.HistoryTrafficDataByDetectionSectionModel) error {
	for _, model := range models {
		if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
			return err
		}
	}
	return s.repo.BulkCreateHistoryDayTrafficDataByDetectionSection(ctx, models)
}

func (s *historyDayTrafficDataByDetectionSectionService) Statistics(ctx context.Context, dataSourceName string, detectionSectionIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionSectionStatisticsByDatasource, error) {
	aggregate := models.TrafficDataByDetectionSectionStatisticsByDatasource{}

	aggregatedStats, err := s.repo.ListHistoryDayTrafficDataByDetectionSectionAggregatedBySection(ctx, dataSourceName, detectionSectionIDs, startTimestamp, endTimestamp)
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

func (s *historyDayTrafficDataByDetectionSectionService) ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.HistoryTrafficDataByDetectionSectionModel, error) {
	//query db for daily basis stats
	return s.repo.ListHistoryDayTrafficDataByDetectionSectionAggregatedByDay(ctx, dataSourceName, detectionSectionIDs, fromTime, toTime)
}
