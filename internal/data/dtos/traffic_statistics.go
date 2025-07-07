package dtos

import (
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// DETECTION POINT BY LANE
func FromTrafficDataByDetectionPointByLaneStatisticsModelsToProto(list []models.TrafficStatisticsLane) ([]*pb.TrafficStatisticsLane, error) {
	protoStats := make([]*pb.TrafficStatisticsLane, len(list))
	for i, traffic := range list {
		protoStatsList, err := FromTrafficDataByDetectionPointByLaneStatisticsModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoStats[i] = protoStatsList
	}
	return protoStats, nil
}

func FromTrafficDataByDetectionPointByLaneStatisticsModelToProto(dto *models.TrafficStatisticsLane) (*pb.TrafficStatisticsLane, error) {
	return &pb.TrafficStatisticsLane{
		LaneId:                       dto.LaneID,
		RecordsCount:                 dto.RecordsCount,
		FirstRecordTimestamp:         timestamppb.New(dto.FirstRecordTimestamp),
		LastRecordTimestamp:          timestamppb.New(dto.LastRecordTimestamp),
		LongestDataGap:               dto.LongestDataGap,
		LongestDataGapStartTimestamp: timestamppb.New(dto.LongestDataGapStartTimestamp),
		LongestDataGapEndTimestamp:   timestamppb.New(dto.LongestDataGapEndTimestamp),
	}, nil
}

// DETECTION POINT
func FromTrafficDataByDetectionPointStatisticsModelsToProto(list []models.TrafficStatisticsDetectionPoint) ([]*pb.TrafficStatisticsDetectionPoint, error) {
	protoStats := make([]*pb.TrafficStatisticsDetectionPoint, len(list))
	for i, traffic := range list {
		protoStatsList, err := FromTrafficDataByDetectionPointStatisticsModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoStats[i] = protoStatsList
	}
	return protoStats, nil
}

func FromTrafficDataByDetectionPointStatisticsModelToProto(dto *models.TrafficStatisticsDetectionPoint) (*pb.TrafficStatisticsDetectionPoint, error) {
	return &pb.TrafficStatisticsDetectionPoint{
		DetectionPointId:             dto.DetectionPointID,
		RecordsCount:                 dto.RecordsCount,
		FirstRecordTimestamp:         timestamppb.New(dto.FirstRecordTimestamp),
		LastRecordTimestamp:          timestamppb.New(dto.LastRecordTimestamp),
		LongestDataGap:               dto.LongestDataGap,
		LongestDataGapStartTimestamp: timestamppb.New(dto.LongestDataGapStartTimestamp),
		LongestDataGapEndTimestamp:   timestamppb.New(dto.LongestDataGapEndTimestamp),
	}, nil
}

// DETECTION SECTION
func FromTrafficDataByDetectionSectionStatisticsModelsToProto(list []models.TrafficStatisticsDetectionSection) ([]*pb.TrafficStatisticsDetectionSection, error) {
	protoStats := make([]*pb.TrafficStatisticsDetectionSection, len(list))
	for i, traffic := range list {
		protoStatsList, err := FromTrafficDataByDetectionSectionStatisticsModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoStats[i] = protoStatsList
	}
	return protoStats, nil
}

func FromTrafficDataByDetectionSectionStatisticsModelToProto(dto *models.TrafficStatisticsDetectionSection) (*pb.TrafficStatisticsDetectionSection, error) {
	return &pb.TrafficStatisticsDetectionSection{
		DetectionSectionId:           dto.DetectionSectionID,
		RecordsCount:                 dto.RecordsCount,
		FirstRecordTimestamp:         timestamppb.New(dto.FirstRecordTimestamp),
		LastRecordTimestamp:          timestamppb.New(dto.LastRecordTimestamp),
		LongestDataGap:               dto.LongestDataGap,
		LongestDataGapStartTimestamp: timestamppb.New(dto.LongestDataGapStartTimestamp),
		LongestDataGapEndTimestamp:   timestamppb.New(dto.LongestDataGapEndTimestamp),
	}, nil
}
