package services

import (
	"context"
	"fmt"
	dataModel "spoke7-go/internal/data/models"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/sumo-integration/dtos"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type SumoIntegrationCurrentTrafficDataByDetectionPointByLaneService interface {
	CreateFromXml(ctx context.Context, detector dtos.InductionDetectorModel, dataSourceName string, startTime time.Time) error
}

type currentTrafficDataByDetectionPointByLaneService struct {
	grpcDataService     grpc_client.GrpcDataClient
	grpcMetadataService grpc_client.GrpcMetadataClient
}

func NewSumoIntegrationCurrentTrafficDataByDetectionPointByLaneService(grpcDataService grpc_client.GrpcDataClient, grpcMetadataService grpc_client.GrpcMetadataClient) SumoIntegrationCurrentTrafficDataByDetectionPointByLaneService {
	return &currentTrafficDataByDetectionPointByLaneService{grpcDataService: grpcDataService, grpcMetadataService: grpcMetadataService}
}

func (s *currentTrafficDataByDetectionPointByLaneService) CreateFromXml(ctx context.Context, detector dtos.InductionDetectorModel, dataSourceName string, startTime time.Time) error {
	// get datasource
	datasource, err := s.grpcMetadataService.GetDataSource(ctx, dataSourceName)
	if err != nil {
		return fmt.Errorf("failed to retrieve induction loop data source information: %w", err)
	}

	// get detection point info
	listOfDetectionPoints := []models.DetectionPoint{}
	for _, detectionPoint := range datasource.DetectionPoints {
		retrievedDetectionPoint, err := s.grpcMetadataService.GetDetectionPoint(ctx, detectionPoint.Id, dataSourceName)
		if err == nil {
			//return fmt.Errorf("failed to retrieve detection point information: %w", err)
			listOfDetectionPoints = append(listOfDetectionPoints, retrievedDetectionPoint)
		}
	}

	// iterate over intervals
	for _, loop := range detector.Intervals {
		// check interval for aggregation is 5 minutes
		interval := (loop.End - loop.Begin)
		if interval != 300 {
			return fmt.Errorf("error: the given interval is not 300 seconds")
		}

		// set timestamp
		detectionTimestamp := startTime.Add(time.Duration(loop.Begin) * time.Second)

		// initialize model
		model := dataModel.CurrentTrafficDataByDetectionPointByLaneModel{
			DetectionTimestamp: detectionTimestamp,
			DetectionInterval:  uint32(interval),

			LaneID:         loop.ID,
			DataSourceName: dataSourceName,

			AverageVehicleLength:             float32(loop.Length),
			HarmonicMeanSpeedVehicleClassAll: float32(loop.HarmonicMeanSpeed),
			HarmonicMeanSpeedVehicleClass1:   float32(loop.HarmonicMeanSpeed),
			CountVehicleClassAll:             uint32(loop.NVehContrib),
			CountVehicleClass1:               uint32(loop.NVehContrib),
		}

		// retrieve information from metadata to get detction point id
		loopExists := false
		for _, detectionPoint := range listOfDetectionPoints {
			for _, lane := range detectionPoint.Lanes {
				if lane.Id == loop.ID {
					loopExists = true
					break
				}
			}
			if loopExists {
				break
			}

		}
		if loopExists == true {
			if _, err := s.grpcDataService.CreateCurrentTrafficDataByDetectionPointByLane(ctx, model); err != nil {
				fmt.Printf("failed to create data: %v", err.Error())
			}
		}

	}

	return nil
}
