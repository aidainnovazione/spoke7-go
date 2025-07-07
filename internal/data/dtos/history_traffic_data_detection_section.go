package dtos

import (
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromHistoryTrafficDataByDetectionSectionModelsToProtos(trafficList []models.HistoryTrafficDataByDetectionSectionModel) ([]*pb.HistoryTrafficDataByDetectionSection, error) {
	protoTraffic := make([]*pb.HistoryTrafficDataByDetectionSection, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromHistoryTrafficDataByDetectionSectionModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromHistoryTrafficDataByDetectionSectionModelToProto(traffic *models.HistoryTrafficDataByDetectionSectionModel) (*pb.HistoryTrafficDataByDetectionSection, error) {
	protoTraffic := &pb.HistoryTrafficDataByDetectionSection{
		CreatedAt:          timestamppb.New(traffic.CreatedAt),
		ModifiedAt:         timestamppb.New(traffic.ModifiedAt),
		DataSourceName:     traffic.DataSourceName,
		DetectionSectionId: traffic.DetectionSectionID,
		DetectionTimestamp: timestamppb.New(traffic.DetectionTimestamp),
		DetectionInterval:  traffic.DetectionInterval,

		ForwardSpeedCount:         traffic.ForwardSpeedCount,
		ForwardSpeedCountMaxFlow:  traffic.ForwardSpeedCountMaxFlow,
		ForwardSpeedCountMinFlow:  traffic.ForwardSpeedCountMinFlow,
		BackwardSpeedCount:        traffic.BackwardSpeedCount,
		BackwardSpeedCountMaxFlow: traffic.BackwardSpeedCountMaxFlow,
		BackwardSpeedCountMinFlow: traffic.BackwardSpeedCountMinFlow,
	}

	return protoTraffic, nil
}

func FromHistoryTrafficDataByDetectionSectionProtosToModels(trafficList []*pb.HistoryTrafficDataByDetectionSection) ([]*models.HistoryTrafficDataByDetectionSectionModel, error) {
	protoTraffic := make([]*models.HistoryTrafficDataByDetectionSectionModel, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromHistoryTrafficDataByDetectionSectionProtoToModel(traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromHistoryTrafficDataByDetectionSectionProtoToModel(traffic *pb.HistoryTrafficDataByDetectionSection) (*models.HistoryTrafficDataByDetectionSectionModel, error) {
	modelTraffic := models.HistoryTrafficDataByDetectionSectionModel{
		CreatedAt:          traffic.CreatedAt.AsTime(),
		ModifiedAt:         traffic.ModifiedAt.AsTime(),
		DataSourceName:     traffic.DataSourceName,
		DetectionSectionID: traffic.DetectionSectionId,
		DetectionTimestamp: traffic.DetectionTimestamp.AsTime(),
		DetectionInterval:  traffic.DetectionInterval,

		ForwardSpeedCount:         traffic.ForwardSpeedCount,
		ForwardSpeedCountMaxFlow:  traffic.ForwardSpeedCountMaxFlow,
		ForwardSpeedCountMinFlow:  traffic.ForwardSpeedCountMinFlow,
		BackwardSpeedCount:        traffic.BackwardSpeedCount,
		BackwardSpeedCountMaxFlow: traffic.BackwardSpeedCountMaxFlow,
		BackwardSpeedCountMinFlow: traffic.BackwardSpeedCountMinFlow,
	}

	return &modelTraffic, nil
}
