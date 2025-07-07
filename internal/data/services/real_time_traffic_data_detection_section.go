package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type RealTimeTrafficDataByDetectionSectionService interface {
	Create(ctx context.Context, realTimeTrafficDataByDetectionSection *models.RealTimeTrafficDataByDetectionSectionModel) error
	Update(ctx context.Context, realTimeTrafficDataByDetectionSection *models.RealTimeTrafficDataByDetectionSectionModel) error
	Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error
	Get(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.RealTimeTrafficDataByDetectionSectionModel, error)
	List(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.RealTimeTrafficDataByDetectionSectionModel, error)
	BulkCreate(ctx context.Context, realTimeTrafficDataByDetectionSections []*models.RealTimeTrafficDataByDetectionSectionModel) error
	Statistics(ctx context.Context, dataSourceName string, detectionSectionIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionSectionStatisticsByDatasource, error)
	ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.RealTimeTrafficDataByDetectionSectionModel, error)
}

type realTimeTrafficDataByDetectionSectionService struct {
	repo                repository.DBClient
	grpcMetadataService grpc_client.GrpcMetadataClient
	organizationName    string
}

func NewRealTimeTrafficDataByDetectionSectionService(repo repository.DBClient, grpcMetadataService grpc_client.GrpcMetadataClient, organizationName string) RealTimeTrafficDataByDetectionSectionService {
	return &realTimeTrafficDataByDetectionSectionService{repo: repo, grpcMetadataService: grpcMetadataService, organizationName: organizationName}
}

func (s *realTimeTrafficDataByDetectionSectionService) Create(ctx context.Context, model *models.RealTimeTrafficDataByDetectionSectionModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.CreateRealTimeTrafficDataByDetectionSection(ctx, model)
}

func (s *realTimeTrafficDataByDetectionSectionService) Update(ctx context.Context, model *models.RealTimeTrafficDataByDetectionSectionModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.UpdateRealTimeTrafficDataByDetectionSection(ctx, model)
}

func (s *realTimeTrafficDataByDetectionSectionService) Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.DeleteRealTimeTrafficDataByDetectionSection(ctx, params)
}

func (s *realTimeTrafficDataByDetectionSectionService) Get(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.RealTimeTrafficDataByDetectionSectionModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.GetRealTimeTrafficDataByDetectionSection(ctx, params)
}

func (s *realTimeTrafficDataByDetectionSectionService) List(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.RealTimeTrafficDataByDetectionSectionModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.ListRealTimeTrafficDataByDetectionSection(ctx, params)
}

func (s *realTimeTrafficDataByDetectionSectionService) BulkCreate(ctx context.Context, models []*models.RealTimeTrafficDataByDetectionSectionModel) error {
	for _, model := range models {
		if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
			return err
		}
	}
	return s.repo.BulkCreateRealTimeTrafficDataByDetectionSection(ctx, models)
}

func (s *realTimeTrafficDataByDetectionSectionService) Statistics(ctx context.Context, dataSourceName string, detectionSectionIDs []string, startTimestamp time.Time, endTimestamp time.Time) (models.TrafficDataByDetectionSectionStatisticsByDatasource, error) {
	aggregate := models.TrafficDataByDetectionSectionStatisticsByDatasource{}

	aggregatedStats, err := s.repo.ListRealTimeTrafficDataByDetectionSectionAggregatedBySection(ctx, dataSourceName, detectionSectionIDs, startTimestamp, endTimestamp)
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

func (s *realTimeTrafficDataByDetectionSectionService) ListAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, fromTime *time.Time, toTime *time.Time) ([]*models.RealTimeTrafficDataByDetectionSectionModel, error) {
	//query db for daily basis stats
	return s.repo.ListRealTimeTrafficDataByDetectionSectionAggregatedByDay(ctx, dataSourceName, detectionSectionIDs, fromTime, toTime)
}
