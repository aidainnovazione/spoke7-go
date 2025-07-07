package dtos

import (
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromRealTimeTrafficDataByDetectionPointByLaneModelsToProtos(trafficList []models.RealTimeTrafficDataByDetectionPointByLaneModel) ([]*pb.RealTimeTrafficDataByDetectionPointByLane, error) {
	protoTraffic := make([]*pb.RealTimeTrafficDataByDetectionPointByLane, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromRealTimeTrafficDataByDetectionPointByLaneModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromRealTimeTrafficDataByDetectionPointByLaneModelToProto(traffic *models.RealTimeTrafficDataByDetectionPointByLaneModel) (*pb.RealTimeTrafficDataByDetectionPointByLane, error) {
	protoTraffic := &pb.RealTimeTrafficDataByDetectionPointByLane{
		CreatedAt:            timestamppb.New(traffic.CreatedAt),
		ModifiedAt:           timestamppb.New(traffic.ModifiedAt),
		DataSourceName:       traffic.DataSourceName,
		LaneId:               traffic.LaneID,
		DetectionTimestamp:   timestamppb.New(traffic.DetectionTimestamp),
		DetectionType:        traffic.DetectionType,
		DetectionTechnology:  traffic.DetectionTechnology,
		VehicleClass:         pb.VehicleClassType(traffic.VehicleClass),
		VehicleSpeed:         traffic.VehicleSpeed,
		VehicleLength:        traffic.VehicleLength,
		VehicleHeadway:       traffic.VehicleHeadway,
		QueuePresent:         traffic.QueuePresent,
		CorrectFlowDirection: traffic.CorrectFlowDirection,
	}

	return protoTraffic, nil
}

func FromRealTimeTrafficDataByDetectionPointByLaneProtosToModels(trafficList []*pb.RealTimeTrafficDataByDetectionPointByLane) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error) {
	protoTraffic := make([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromRealTimeTrafficDataByDetectionPointByLaneProtoToModel(traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromRealTimeTrafficDataByDetectionPointByLaneProtoToModel(traffic *pb.RealTimeTrafficDataByDetectionPointByLane) (*models.RealTimeTrafficDataByDetectionPointByLaneModel, error) {
	modelTraffic := models.RealTimeTrafficDataByDetectionPointByLaneModel{
		CreatedAt:            traffic.CreatedAt.AsTime(),
		ModifiedAt:           traffic.ModifiedAt.AsTime(),
		DataSourceName:       traffic.DataSourceName,
		LaneID:               traffic.LaneId,
		DetectionTimestamp:   traffic.DetectionTimestamp.AsTime(),
		DetectionType:        traffic.DetectionType,
		DetectionTechnology:  traffic.DetectionTechnology,
		VehicleClass:         models.VehicleClassType(traffic.VehicleClass),
		VehicleSpeed:         traffic.VehicleSpeed,
		VehicleLength:        traffic.VehicleLength,
		VehicleHeadway:       traffic.VehicleHeadway,
		QueuePresent:         traffic.QueuePresent,
		CorrectFlowDirection: traffic.CorrectFlowDirection,
	}

	return &modelTraffic, nil

}
