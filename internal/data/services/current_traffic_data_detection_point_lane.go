package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type CurrentTrafficDataByDetectionPointByLaneService interface {
	Create(ctx context.Context, currentTrafficDataByDetectionPointByLane *models.CurrentTrafficDataByDetectionPointByLaneModel) error
	Update(ctx context.Context, currentTrafficDataByDetectionPointByLane *models.CurrentTrafficDataByDetectionPointByLaneModel) error
	Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error
	Get(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error)
	List(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error)
	BulkCreate(ctx context.Context, currentTrafficDataByDetectionPointByLanes []*models.CurrentTrafficDataByDetectionPointByLaneModel) error
	Statistics(ctx context.Context, dataSourceName string, laneIds []string, startTimestamp time.Time, endTimestamp time.Time) (models.CurrentTrafficDataByDetectionPointByLaneStatisticsByDatasource, error)
	Aggregate(ctx context.Context, dataSourceName string, laneIds []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.CurrentTrafficDataByDetectionPointByLaneAggregate, error)
	ListAggregatedByDay(ctx context.Context, dataSourceName string, laneIds []string, fromTime *time.Time, toTime *time.Time) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error)
}

type currentTrafficDataByDetectionPointByLaneService struct {
	repo                repository.DBClient
	grpcMetadataService grpc_client.GrpcMetadataClient
	organizationName    string
}

func NewCurrentTrafficDataByDetectionPointByLaneService(repo repository.DBClient, grpcMetadataService grpc_client.GrpcMetadataClient, organizationName string) CurrentTrafficDataByDetectionPointByLaneService {
	return &currentTrafficDataByDetectionPointByLaneService{repo: repo, grpcMetadataService: grpcMetadataService, organizationName: organizationName}
}

func (s *currentTrafficDataByDetectionPointByLaneService) Create(ctx context.Context, model *models.CurrentTrafficDataByDetectionPointByLaneModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.CreateCurrentTrafficDataByDetectionPointByLane(ctx, model)
}

func (s *currentTrafficDataByDetectionPointByLaneService) Update(ctx context.Context, model *models.CurrentTrafficDataByDetectionPointByLaneModel) error {
	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.UpdateCurrentTrafficDataByDetectionPointByLane(ctx, model)
}

func (s *currentTrafficDataByDetectionPointByLaneService) Delete(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.DeleteCurrentTrafficDataByDetectionPointByLane(ctx, params)
}

func (s *currentTrafficDataByDetectionPointByLaneService) Get(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.GetCurrentTrafficDataByDetectionPointByLane(ctx, params)
}

func (s *currentTrafficDataByDetectionPointByLaneService) List(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error) {
	if err := CheckDataSourceUserPermission(ctx, params.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	return s.repo.ListCurrentTrafficDataByDetectionPointByLane(ctx, params)
}

func (s *currentTrafficDataByDetectionPointByLaneService) BulkCreate(ctx context.Context, models []*models.CurrentTrafficDataByDetectionPointByLaneModel) error {
	// for _, model := range models {
	// 	if err := CheckDataSourceUserPermission(ctx, model.DataSourceName,  s.organizationName, s.grpcMetadataService); err != nil {
	// 		return err
	// 	}
	// }
	return s.repo.BulkCreateCurrentTrafficDataByDetectionPointByLane(ctx, models)
}

func (s *currentTrafficDataByDetectionPointByLaneService) Aggregate(ctx context.Context, dataSourceName string, laneIds []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.CurrentTrafficDataByDetectionPointByLaneAggregate, error) {
	if err := CheckDataSourceUserPermission(ctx, dataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}

	aggregateList := []models.CurrentTrafficDataByDetectionPointByLaneAggregate{}

	// retrieve all available lane ids
	if len(laneIds) == 0 {
		datasource, err := s.grpcMetadataService.GetDataSource(ctx, dataSourceName)
		if err != nil {
			return nil, err
		}
		for _, detectionPoint := range datasource.DetectionPoints {
			detectionPointInfo, err := s.grpcMetadataService.GetDetectionPoint(ctx, detectionPoint.Id, dataSourceName)
			if err != nil {
				return nil, err
			}

			for _, lane := range detectionPointInfo.Lanes {
				laneIds = append(laneIds, lane.Id)
			}
		}
	}

	for t := startTimestamp; t.Before(endTimestamp); t = t.Add(5 * time.Minute) {
		params := models.GetTrafficDataByDetectionPointByLaneParams{
			DataSourceName:     dataSourceName,
			LaneIDs:            laneIds,
			DetectionTimestamp: t,
		}
		models, err := s.repo.GetCurrentTrafficDataByDetectionPointByLane(ctx, params)
		if err != nil {
			return nil, err
		}

		aggregate := computeLaneAggregate(models, laneIds)
		aggregate.DataSourceName = dataSourceName
		aggregate.DetectionTimestamp = t

		aggregateList = append(aggregateList, aggregate)
	}

	return aggregateList, nil
}

func computeLaneAggregate(trafficModels []*models.CurrentTrafficDataByDetectionPointByLaneModel, laneIds []string) models.CurrentTrafficDataByDetectionPointByLaneAggregate {
	aggregate := models.CurrentTrafficDataByDetectionPointByLaneAggregate{}
	// total number of records
	aggregate.RecordsCount = uint32(len(trafficModels))
	//percentage
	aggregate.AggPercentageRecords = float32(len(trafficModels)) / float32(len(laneIds)) * 100

	if len(trafficModels) > 0 {
		// compute rate
		aggregate.AggPercentageRecordsWithCounts = 0
		aggregate.AggPercentageRecordsWithPositiveAverageSpeed = 0
		for _, m := range trafficModels {
			if m.CountVehicleClassAll > 0 {
				aggregate.AggPercentageRecordsWithCounts++
			}
			if m.HarmonicMeanSpeedVehicleClassAll > 0 {
				aggregate.AggPercentageRecordsWithPositiveAverageSpeed++
			}
		}
		aggregate.AggPercentageRecordsWithCounts = (float32(aggregate.AggPercentageRecordsWithCounts) / float32(len(laneIds)) * 100)
		aggregate.AggPercentageRecordsWithPositiveAverageSpeed = (float32(aggregate.AggPercentageRecordsWithPositiveAverageSpeed) / float32(len(laneIds)) * 100)

		// average over aggregated data
		aggregate.AggCountVehicleClassAll = 0
		aggregate.AggHarmonicMeanSpeedVehicleClassAll = 0
		aggregate.AggAverageVehicleLength = 0
		aggregate.AggAverageHeadway = 0
		aggregate.AggStdHeadway = 0
		aggregate.AggAverageTimeToCollision = 0
		aggregate.AggStdTimeToCollision = 0
		for _, m := range trafficModels {
			aggregate.AggCountVehicleClassAll += float32(m.CountVehicleClassAll)
			aggregate.AggHarmonicMeanSpeedVehicleClassAll += m.HarmonicMeanSpeedVehicleClassAll
			aggregate.AggAverageVehicleLength += aggregate.AggAverageVehicleLength
			aggregate.AggAverageHeadway += aggregate.AggAverageHeadway
			aggregate.AggStdHeadway += aggregate.AggStdHeadway
			aggregate.AggAverageTimeToCollision += aggregate.AggAverageTimeToCollision
			aggregate.AggStdTimeToCollision += aggregate.AggStdTimeToCollision
		}
		aggregate.AggCountVehicleClassAll /= float32(len(trafficModels))
		aggregate.AggHarmonicMeanSpeedVehicleClassAll /= float32(len(trafficModels))
		aggregate.AggAverageVehicleLength /= float32(len(trafficModels))
		aggregate.AggAverageHeadway /= float32(len(trafficModels))
		aggregate.AggStdHeadway /= float32(len(trafficModels))
		aggregate.AggAverageTimeToCollision /= float32(len(trafficModels))
		aggregate.AggStdTimeToCollision /= float32(len(trafficModels))
	}

	return aggregate
}

func (s *currentTrafficDataByDetectionPointByLaneService) Statistics(ctx context.Context, dataSourceName string, laneIds []string, startTimestamp time.Time, endTimestamp time.Time) (models.CurrentTrafficDataByDetectionPointByLaneStatisticsByDatasource, error) {
	aggregate := models.CurrentTrafficDataByDetectionPointByLaneStatisticsByDatasource{}

	aggregatedStats, err := s.repo.ListCurrentTrafficDataByDetectionPointByLaneAggregatedByLane(ctx, dataSourceName, laneIds, startTimestamp, endTimestamp)
	if err != nil {
		return aggregate, fmt.Errorf("failed to retrieve aggregated traffic statistics: %w", err)
	}

	var statsList []models.CurrentTrafficDataByDetectionPointByLaneStatistics
	for _, agg := range aggregatedStats {
		statsList = append(statsList, models.CurrentTrafficDataByDetectionPointByLaneStatistics{
			LaneID:                                    agg.LaneID,
			RecordsCount:                              agg.RecordsCount,
			FirstRecordTimestamp:                      agg.FirstRecordTimestamp,
			LastRecordTimestamp:                       agg.LastRecordTimestamp,
			TotalCountAllVehicles:                     agg.TotalCountAllVehicles,
			TotalHarmonicMeanSpeedAllRecords:          agg.TotalHarmonicMeanSpeedAllRecords,
			PercentageRecordsWithCounts:               (agg.PercentageRecordsWithCounts),
			PercentageRecordsWithPositiveAverageSpeed: (agg.PercentageRecordsWithPositiveAverageSpeed),
		})
	}

	aggregate.StatisticsByLane = statsList
	aggregate.RecordsCount, aggregate.FirstRecordTimestamp, aggregate.LastRecordTimestamp = AggregateStatistics(statsList)

	return aggregate, nil
}

func (s *currentTrafficDataByDetectionPointByLaneService) ListAggregatedByDay(ctx context.Context, dataSourceName string, laneIds []string, fromTime *time.Time, toTime *time.Time) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error) {
	//query db for daily basis stats
	return s.repo.ListCurrentTrafficDataByDetectionPointByLaneAggregatedByDay(ctx, dataSourceName, laneIds, fromTime, toTime)
}
