package dtos

import (
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromRealTimeTrafficDataByDetectionSectionModelsToProtos(trafficList []models.RealTimeTrafficDataByDetectionSectionModel) ([]*pb.RealTimeTrafficDataByDetectionSection, error) {
	protoTraffic := make([]*pb.RealTimeTrafficDataByDetectionSection, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromRealTimeTrafficDataByDetectionSectionModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromRealTimeTrafficDataByDetectionSectionModelToProto(traffic *models.RealTimeTrafficDataByDetectionSectionModel) (*pb.RealTimeTrafficDataByDetectionSection, error) {
	protoTraffic := &pb.RealTimeTrafficDataByDetectionSection{
		CreatedAt:            timestamppb.New(traffic.CreatedAt),
		ModifiedAt:           timestamppb.New(traffic.ModifiedAt),
		DataSourceName:       traffic.DataSourceName,
		DetectionSectionId:   traffic.DetectionSectionID,
		DetectionTimestamp:   timestamppb.New(traffic.DetectionTimestamp),
		DetectionSectionNode: pb.DetectionSectionNodeType(traffic.DetectionSectionNodeType),
		DetectionType:        traffic.DetectionType,
		DetectionTechnology:  traffic.DetectionTechnology,
		AnonymousDetectionId: traffic.AnonymousDetectionID,
	}

	return protoTraffic, nil
}

func FromRealTimeTrafficDataByDetectionSectionProtosToModels(trafficList []*pb.RealTimeTrafficDataByDetectionSection) ([]*models.RealTimeTrafficDataByDetectionSectionModel, error) {
	protoTraffic := make([]*models.RealTimeTrafficDataByDetectionSectionModel, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromRealTimeTrafficDataByDetectionSectionProtoToModel(traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromRealTimeTrafficDataByDetectionSectionProtoToModel(traffic *pb.RealTimeTrafficDataByDetectionSection) (*models.RealTimeTrafficDataByDetectionSectionModel, error) {
	modelTraffic := models.RealTimeTrafficDataByDetectionSectionModel{
		CreatedAt:                traffic.CreatedAt.AsTime(),
		ModifiedAt:               traffic.ModifiedAt.AsTime(),
		DataSourceName:           traffic.DataSourceName,
		DetectionSectionID:       traffic.DetectionSectionId,
		DetectionTimestamp:       traffic.DetectionTimestamp.AsTime(),
		DetectionSectionNodeType: models.DetectionSectionNodeType(traffic.DetectionSectionNode),
		DetectionType:            traffic.DetectionType,
		DetectionTechnology:      traffic.DetectionTechnology,
		AnonymousDetectionID:     traffic.AnonymousDetectionId,
	}

	return &modelTraffic, nil
}
