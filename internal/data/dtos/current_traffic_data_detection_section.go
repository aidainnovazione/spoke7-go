package dtos

import (
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromCurrentTrafficDataByDetectionSectionModelsToProtos(trafficList []models.CurrentTrafficDataByDetectionSectionModel) ([]*pb.CurrentTrafficDataByDetectionSection, error) {
	protoTraffic := make([]*pb.CurrentTrafficDataByDetectionSection, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromCurrentTrafficDataByDetectionSectionModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromCurrentTrafficDataByDetectionSectionModelToProto(traffic *models.CurrentTrafficDataByDetectionSectionModel) (*pb.CurrentTrafficDataByDetectionSection, error) {
	protoTraffic := &pb.CurrentTrafficDataByDetectionSection{
		CreatedAt:          timestamppb.New(traffic.CreatedAt),
		ModifiedAt:         timestamppb.New(traffic.ModifiedAt),
		DataSourceName:     traffic.DataSourceName,
		DetectionSectionId: traffic.DetectionSectionID,
		DetectionTimestamp: timestamppb.New(traffic.DetectionTimestamp),
		DetectionInterval:  traffic.DetectionInterval,
		ForwardSpeed:       traffic.ForwardSpeed,
		BackwardSpeed:      traffic.BackwardSpeed,
	}

	return protoTraffic, nil
}

func FromCurrentTrafficDataByDetectionSectionProtosToModels(trafficList []*pb.CurrentTrafficDataByDetectionSection) ([]*models.CurrentTrafficDataByDetectionSectionModel, error) {
	protoTraffic := make([]*models.CurrentTrafficDataByDetectionSectionModel, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromCurrentTrafficDataByDetectionSectionProtoToModel(traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromCurrentTrafficDataByDetectionSectionProtoToModel(traffic *pb.CurrentTrafficDataByDetectionSection) (*models.CurrentTrafficDataByDetectionSectionModel, error) {
	modelTraffic := models.CurrentTrafficDataByDetectionSectionModel{
		CreatedAt:          traffic.CreatedAt.AsTime(),
		ModifiedAt:         traffic.ModifiedAt.AsTime(),
		DataSourceName:     traffic.DataSourceName,
		DetectionSectionID: traffic.DetectionSectionId,
		DetectionTimestamp: traffic.DetectionTimestamp.AsTime(),
		DetectionInterval:  traffic.DetectionInterval,
		ForwardSpeed:       traffic.ForwardSpeed,
		BackwardSpeed:      traffic.BackwardSpeed,
	}

	return &modelTraffic, nil
}

func FromCurrentTrafficDataByDetectionSectionStatisticsModelsToProto(list []models.CurrentTrafficDataByDetectionSectionStatistics) ([]*pb.CurrentTrafficDataByDetectionSectionStatistics, error) {
	protoStats := make([]*pb.CurrentTrafficDataByDetectionSectionStatistics, len(list))
	for i, traffic := range list {
		protoStatsList, err := FromCurrentTrafficDataByDetectionSectionStatisticsModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoStats[i] = protoStatsList
	}
	return protoStats, nil
}

func FromCurrentTrafficDataByDetectionSectionStatisticsModelToProto(dto *models.CurrentTrafficDataByDetectionSectionStatistics) (*pb.CurrentTrafficDataByDetectionSectionStatistics, error) {
	return &pb.CurrentTrafficDataByDetectionSectionStatistics{
		DetectionSectionId:                 dto.DetectionSectionID,
		RecordsCount:                       dto.RecordsCount,
		FirstRecordTimestamp:               timestamppb.New(dto.FirstRecordTimestamp),
		LastRecordTimestamp:                timestamppb.New(dto.LastRecordTimestamp),
		LongestDataGap:                     dto.LongestDataGap,
		LongestDataGapStartTimestamp:       timestamppb.New(dto.LongestDataGapStartTimestamp),
		LongestDataGapEndTimestamp:         timestamppb.New(dto.LongestDataGapEndTimestamp),
		MissingIntervalDetectionRate:       dto.MissingIntervalDetectionRate,
		TotalAverageForwardSpeed:           dto.TotalAverageForwardSpeed,
		TotalAverageBackwardSpeed:          dto.TotalAverageBackwardSpeed,
		PercentageRecordsWithForwardSpeed:  dto.PercentageRecordsWithForwardSpeed,
		PercentageRecordsWithBackwardSpeed: dto.PercentageRecordsWithBackwardSpeed,
	}, nil
}
