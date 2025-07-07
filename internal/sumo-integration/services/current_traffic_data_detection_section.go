package services

import (
	"context"
	"fmt"
	dataModel "spoke7-go/internal/data/models"
	"spoke7-go/internal/sumo-integration/dtos"
	"spoke7-go/pkg/grpc_client"
	"time"
)

type SumoIntegrationCurrentTrafficDataByDetectionSectionService interface {
	CreateFromXml(ctx context.Context, detector dtos.LaneDetectorModel, dataSourceName string, startTime time.Time) error
}

type currentTrafficDataByDetectionSectionService struct {
	grpcDataService     grpc_client.GrpcDataClient
	grpcMetadataService grpc_client.GrpcMetadataClient
}

func NewSumoIntegrationCurrentTrafficDataByDetectionSectionService(grpcDataService grpc_client.GrpcDataClient, grpcMetadataService grpc_client.GrpcMetadataClient) SumoIntegrationCurrentTrafficDataByDetectionSectionService {
	return &currentTrafficDataByDetectionSectionService{grpcDataService: grpcDataService, grpcMetadataService: grpcMetadataService}
}

func (s *currentTrafficDataByDetectionSectionService) CreateFromXml(ctx context.Context, detector dtos.LaneDetectorModel, dataSourceName string, startTime time.Time) error {
	datasource, err := s.grpcMetadataService.GetDataSource(ctx, dataSourceName)
	if err != nil {
		return fmt.Errorf("failed to retrieve data source section information: %w", err)
	}

	for _, loop := range detector.Intervals {
		// check interval for aggregation is 5 minutes
		interval := (loop.End - loop.Begin)
		if interval != 300 {
			return fmt.Errorf("error: the given interval is not 300 seconds")
		}

		// set timestamp
		detectionTimestamp := startTime.Add(time.Duration(loop.Begin) * time.Second)

		// initialize model
		model := dataModel.CurrentTrafficDataByDetectionSectionModel{
			DetectionTimestamp: detectionTimestamp,
			DetectionInterval:  uint32(interval),

			DetectionSectionID: loop.ID,
			DataSourceName:     dataSourceName,

			ForwardSpeed:  float32(loop.MeanSpeed),
			BackwardSpeed: float32(loop.MeanSpeed),
		}

		// retrieve information from metadata to get detction point id

		loopExists := false
		for _, detectionSection := range datasource.DetectionSections {
			if detectionSection.Id == loop.ID {
				loopExists = true
			}
		}

		if loopExists == false {
			return fmt.Errorf("failed to create data: this datasource does not have detection section with such loop id")
		}

		if _, err := s.grpcDataService.CreateCurrentTrafficDataByDetectionSection(ctx, model); err != nil {
			fmt.Printf("failed to create data: %v", err.Error())
		}
	}

	return nil
}
