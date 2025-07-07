package dtos

import (
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromHistoryDayTrafficDataByDetectionPointByLaneModelsToProtos(trafficList []models.HistoryDayTrafficDataByDetectionPointByLaneModel) ([]*pb.HistoryDayTrafficDataByDetectionPointByLane, error) {
	protoTraffic := make([]*pb.HistoryDayTrafficDataByDetectionPointByLane, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromHistoryDayTrafficDataByDetectionPointByLaneModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromHistoryDayTrafficDataByDetectionPointByLaneModelToProto(traffic *models.HistoryDayTrafficDataByDetectionPointByLaneModel) (*pb.HistoryDayTrafficDataByDetectionPointByLane, error) {
	protoTraffic := &pb.HistoryDayTrafficDataByDetectionPointByLane{
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

		TrafficFlowParametersVehicleClass_1: FromTrafficFlowParametersByDayModelToProto(&traffic.TrafficFlowParametersVehicleClass1),
		TrafficFlowParametersVehicleClass_2: FromTrafficFlowParametersByDayModelToProto(&traffic.TrafficFlowParametersVehicleClass2),
		TrafficFlowParametersVehicleClass_3: FromTrafficFlowParametersByDayModelToProto(&traffic.TrafficFlowParametersVehicleClass3),
		TrafficFlowParametersVehicleClass_4: FromTrafficFlowParametersByDayModelToProto(&traffic.TrafficFlowParametersVehicleClass4),
		TrafficFlowParametersVehicleClass_5: FromTrafficFlowParametersByDayModelToProto(&traffic.TrafficFlowParametersVehicleClass5),
		TrafficFlowParametersVehicleClass_6: FromTrafficFlowParametersByDayModelToProto(&traffic.TrafficFlowParametersVehicleClass6),
		TrafficFlowParametersVehicleClass_7: FromTrafficFlowParametersByDayModelToProto(&traffic.TrafficFlowParametersVehicleClass7),
		TrafficFlowParametersVehicleClass_8: FromTrafficFlowParametersByDayModelToProto(&traffic.TrafficFlowParametersVehicleClass8),

		InstantaneousSpeedVelClass_1MaxFlow: traffic.InstantaneousSpeedVelClass1MaxFlow,
		InstantaneousSpeedVelClass_2MaxFlow: traffic.InstantaneousSpeedVelClass2MaxFlow,
		InstantaneousSpeedVelClass_3MaxFlow: traffic.InstantaneousSpeedVelClass3MaxFlow,
		InstantaneousSpeedVelClass_1MinFlow: traffic.InstantaneousSpeedVelClass1MinFlow,
		InstantaneousSpeedVelClass_2MinFlow: traffic.InstantaneousSpeedVelClass2MinFlow,
		InstantaneousSpeedVelClass_3MinFlow: traffic.InstantaneousSpeedVelClass3MinFlow,

		HeadwayVelClass_1MaxFlow: traffic.HeadwayVelClass1MaxFlow,
		HeadwayVelClass_2MaxFlow: traffic.HeadwayVelClass2MaxFlow,
		HeadwayVelClass_3MaxFlow: traffic.HeadwayVelClass3MaxFlow,
		HeadwayVelClass_1MinFlow: traffic.HeadwayVelClass1MinFlow,
		HeadwayVelClass_2MinFlow: traffic.HeadwayVelClass2MinFlow,
		HeadwayVelClass_3MinFlow: traffic.HeadwayVelClass3MinFlow,
	}

	return protoTraffic, nil
}

func FromHistoryDayTrafficDataByDetectionPointByLaneProtosToModels(trafficList []*pb.HistoryDayTrafficDataByDetectionPointByLane) ([]*models.HistoryDayTrafficDataByDetectionPointByLaneModel, error) {
	protoTraffic := make([]*models.HistoryDayTrafficDataByDetectionPointByLaneModel, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromHistoryDayTrafficDataByDetectionPointByLaneProtoToModel(traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromHistoryDayTrafficDataByDetectionPointByLaneProtoToModel(traffic *pb.HistoryDayTrafficDataByDetectionPointByLane) (*models.HistoryDayTrafficDataByDetectionPointByLaneModel, error) {
	modelTraffic := models.HistoryDayTrafficDataByDetectionPointByLaneModel{
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

		TrafficFlowParametersVehicleClass1: FromTrafficFlowParametersByDayProtoToModel(traffic.TrafficFlowParametersVehicleClass_1),
		TrafficFlowParametersVehicleClass2: FromTrafficFlowParametersByDayProtoToModel(traffic.TrafficFlowParametersVehicleClass_2),
		TrafficFlowParametersVehicleClass3: FromTrafficFlowParametersByDayProtoToModel(traffic.TrafficFlowParametersVehicleClass_3),
		TrafficFlowParametersVehicleClass4: FromTrafficFlowParametersByDayProtoToModel(traffic.TrafficFlowParametersVehicleClass_4),
		TrafficFlowParametersVehicleClass5: FromTrafficFlowParametersByDayProtoToModel(traffic.TrafficFlowParametersVehicleClass_5),
		TrafficFlowParametersVehicleClass6: FromTrafficFlowParametersByDayProtoToModel(traffic.TrafficFlowParametersVehicleClass_6),
		TrafficFlowParametersVehicleClass7: FromTrafficFlowParametersByDayProtoToModel(traffic.TrafficFlowParametersVehicleClass_7),
		TrafficFlowParametersVehicleClass8: FromTrafficFlowParametersByDayProtoToModel(traffic.TrafficFlowParametersVehicleClass_8),

		InstantaneousSpeedVelClass1MaxFlow: traffic.InstantaneousSpeedVelClass_1MaxFlow,
		InstantaneousSpeedVelClass2MaxFlow: traffic.InstantaneousSpeedVelClass_2MaxFlow,
		InstantaneousSpeedVelClass3MaxFlow: traffic.InstantaneousSpeedVelClass_3MaxFlow,
		InstantaneousSpeedVelClass1MinFlow: traffic.InstantaneousSpeedVelClass_1MinFlow,
		InstantaneousSpeedVelClass2MinFlow: traffic.InstantaneousSpeedVelClass_2MinFlow,
		InstantaneousSpeedVelClass3MinFlow: traffic.InstantaneousSpeedVelClass_3MinFlow,

		HeadwayVelClass1MaxFlow: traffic.HeadwayVelClass_1MaxFlow,
		HeadwayVelClass2MaxFlow: traffic.HeadwayVelClass_2MaxFlow,
		HeadwayVelClass3MaxFlow: traffic.HeadwayVelClass_3MaxFlow,
		HeadwayVelClass1MinFlow: traffic.HeadwayVelClass_1MinFlow,
		HeadwayVelClass2MinFlow: traffic.HeadwayVelClass_2MinFlow,
		HeadwayVelClass3MinFlow: traffic.HeadwayVelClass_3MinFlow,
	}

	return &modelTraffic, nil
}
