package dtos

import (
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromHistoryHourTrafficDataByDetectionPointByLaneModelsToProtos(trafficList []models.HistoryHourTrafficDataByDetectionPointByLaneModel) ([]*pb.HistoryHourTrafficDataByDetectionPointByLane, error) {
	protoTraffic := make([]*pb.HistoryHourTrafficDataByDetectionPointByLane, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromHistoryHourTrafficDataByDetectionPointByLaneModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromHistoryHourTrafficDataByDetectionPointByLaneModelToProto(traffic *models.HistoryHourTrafficDataByDetectionPointByLaneModel) (*pb.HistoryHourTrafficDataByDetectionPointByLane, error) {
	protoTraffic := &pb.HistoryHourTrafficDataByDetectionPointByLane{
		CreatedAt:          timestamppb.New(traffic.CreatedAt),
		ModifiedAt:         timestamppb.New(traffic.ModifiedAt),
		DataSourceName:     traffic.DataSourceName,
		LaneId:             traffic.LaneID,
		DetectionTimestamp: timestamppb.New(traffic.DetectionTimestamp),
		DetectionInterval:  traffic.DetectionInterval,

		TrafficFlowVehicleClass_1: traffic.TrafficFlowVehicleClass1,
		TrafficFlowVehicleClass_2: traffic.TrafficFlowVehicleClass2,
		TrafficFlowVehicleClass_3: traffic.TrafficFlowVehicleClass3,
		TrafficFlowVehicleClass_4: traffic.TrafficFlowVehicleClass4,
		TrafficFlowVehicleClass_5: traffic.TrafficFlowVehicleClass5,
		TrafficFlowVehicleClass_6: traffic.TrafficFlowVehicleClass6,
		TrafficFlowVehicleClass_7: traffic.TrafficFlowVehicleClass7,
		TrafficFlowVehicleClass_8: traffic.TrafficFlowVehicleClass8,

		TrafficFlowVehicleClassEquivalent: traffic.TrafficFlowVehicleClassEquivalent,

		AverageSpeedVehicleClass_1: traffic.AverageSpeedVehicleClass1,
		AverageSpeedVehicleClass_2: traffic.AverageSpeedVehicleClass2,
		AverageSpeedVehicleClass_3: traffic.AverageSpeedVehicleClass3,
		AverageSpeedVehicleClass_4: traffic.AverageSpeedVehicleClass4,
		AverageSpeedVehicleClass_5: traffic.AverageSpeedVehicleClass5,
		AverageSpeedVehicleClass_6: traffic.AverageSpeedVehicleClass6,
		AverageSpeedVehicleClass_7: traffic.AverageSpeedVehicleClass7,
		AverageSpeedVehicleClass_8: traffic.AverageSpeedVehicleClass8,

		AverageSpeedVehicleClassAll: traffic.AverageSpeedVehicleClassAll,

		AverageVehicleLength:   traffic.AverageVehicleLength,
		AverageHeadway:         traffic.AverageHeadway,
		AverageTimeToCollision: traffic.AverageTimeToCollision,
	}

	return protoTraffic, nil
}

func FromHistoryHourTrafficDataByDetectionPointByLaneProtosToModels(trafficList []*pb.HistoryHourTrafficDataByDetectionPointByLane) ([]*models.HistoryHourTrafficDataByDetectionPointByLaneModel, error) {
	protoTraffic := make([]*models.HistoryHourTrafficDataByDetectionPointByLaneModel, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromHistoryHourTrafficDataByDetectionPointByLaneProtoToModel(traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromHistoryHourTrafficDataByDetectionPointByLaneProtoToModel(traffic *pb.HistoryHourTrafficDataByDetectionPointByLane) (*models.HistoryHourTrafficDataByDetectionPointByLaneModel, error) {
	modelTraffic := models.HistoryHourTrafficDataByDetectionPointByLaneModel{
		CreatedAt:          traffic.CreatedAt.AsTime(),
		ModifiedAt:         traffic.ModifiedAt.AsTime(),
		DataSourceName:     traffic.DataSourceName,
		LaneID:             traffic.LaneId,
		DetectionTimestamp: traffic.DetectionTimestamp.AsTime(),
		DetectionInterval:  traffic.DetectionInterval,

		TrafficFlowVehicleClass1: traffic.TrafficFlowVehicleClass_1,
		TrafficFlowVehicleClass2: traffic.TrafficFlowVehicleClass_2,
		TrafficFlowVehicleClass3: traffic.TrafficFlowVehicleClass_3,
		TrafficFlowVehicleClass4: traffic.TrafficFlowVehicleClass_4,
		TrafficFlowVehicleClass5: traffic.TrafficFlowVehicleClass_5,
		TrafficFlowVehicleClass6: traffic.TrafficFlowVehicleClass_6,
		TrafficFlowVehicleClass7: traffic.TrafficFlowVehicleClass_7,
		TrafficFlowVehicleClass8: traffic.TrafficFlowVehicleClass_8,

		TrafficFlowVehicleClassEquivalent: traffic.TrafficFlowVehicleClassEquivalent,

		AverageSpeedVehicleClass1: traffic.AverageSpeedVehicleClass_1,
		AverageSpeedVehicleClass2: traffic.AverageSpeedVehicleClass_2,
		AverageSpeedVehicleClass3: traffic.AverageSpeedVehicleClass_3,
		AverageSpeedVehicleClass4: traffic.AverageSpeedVehicleClass_4,
		AverageSpeedVehicleClass5: traffic.AverageSpeedVehicleClass_5,
		AverageSpeedVehicleClass6: traffic.AverageSpeedVehicleClass_6,
		AverageSpeedVehicleClass7: traffic.AverageSpeedVehicleClass_7,
		AverageSpeedVehicleClass8: traffic.AverageSpeedVehicleClass_8,

		AverageSpeedVehicleClassAll: traffic.AverageSpeedVehicleClassAll,

		AverageVehicleLength:   traffic.AverageVehicleLength,
		AverageHeadway:         traffic.AverageHeadway,
		AverageTimeToCollision: traffic.AverageTimeToCollision,
	}

	return &modelTraffic, nil
}
