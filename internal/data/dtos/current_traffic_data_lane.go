package dtos

import (
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromCurrentTrafficDataByDetectionPointByLaneModelsToProtos(trafficList []models.CurrentTrafficDataByDetectionPointByLaneModel) ([]*pb.CurrentTrafficDataByDetectionPointByLane, error) {
	protoTraffic := make([]*pb.CurrentTrafficDataByDetectionPointByLane, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromCurrentTrafficDataByDetectionPointByLaneModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromCurrentTrafficDataByDetectionPointByLaneModelToProto(traffic *models.CurrentTrafficDataByDetectionPointByLaneModel) (*pb.CurrentTrafficDataByDetectionPointByLane, error) {
	protoTraffic := &pb.CurrentTrafficDataByDetectionPointByLane{
		CreatedAt:                               timestamppb.New(traffic.CreatedAt),
		ModifiedAt:                              timestamppb.New(traffic.ModifiedAt),
		DataSourceName:                          traffic.DataSourceName,
		LaneId:                                  traffic.LaneID,
		DetectionTimestamp:                      timestamppb.New(traffic.DetectionTimestamp),
		DetectionInterval:                       traffic.DetectionInterval,
		CountVehicleClass_1:                     traffic.CountVehicleClass1,
		CountVehicleClass_2:                     traffic.CountVehicleClass2,
		CountVehicleClass_3:                     traffic.CountVehicleClass3,
		CountVehicleClass_4:                     traffic.CountVehicleClass4,
		CountVehicleClass_5:                     traffic.CountVehicleClass5,
		CountVehicleClass_6:                     traffic.CountVehicleClass6,
		CountVehicleClass_7:                     traffic.CountVehicleClass7,
		CountVehicleClass_8:                     traffic.CountVehicleClass8,
		CountVehicleClassAll:                    traffic.CountVehicleClassAll,
		HarmonicMeanSpeedVehicleClass_1:         traffic.HarmonicMeanSpeedVehicleClass1,
		HarmonicMeanSpeedVehicleClass_2:         traffic.HarmonicMeanSpeedVehicleClass2,
		HarmonicMeanSpeedVehicleClass_3:         traffic.HarmonicMeanSpeedVehicleClass3,
		HarmonicMeanSpeedVehicleClass_4:         traffic.HarmonicMeanSpeedVehicleClass4,
		HarmonicMeanSpeedVehicleClass_5:         traffic.HarmonicMeanSpeedVehicleClass5,
		HarmonicMeanSpeedVehicleClass_6:         traffic.HarmonicMeanSpeedVehicleClass6,
		HarmonicMeanSpeedVehicleClass_7:         traffic.HarmonicMeanSpeedVehicleClass7,
		HarmonicMeanSpeedVehicleClass_8:         traffic.HarmonicMeanSpeedVehicleClass8,
		HarmonicMeanSpeedVehicleClassAll:        traffic.HarmonicMeanSpeedVehicleClassAll,
		CountDetectedSpeedVehicleUnder_50:       traffic.CountDetectedSpeedVehicleUnder50,
		CountDetectedSpeedVehicleBetween_50_100: traffic.CountDetectedSpeedVehicleBetween50_100,
		CountDetectedSpeedVehicleOver_100:       traffic.CountDetectedSpeedVehicleOver100,
		AverageVehicleLength:                    traffic.AverageVehicleLength,
		AverageHeadway:                          traffic.AverageHeadway,
		StdHeadway:                              traffic.StdHeadway,
		AverageTimeToCollision:                  traffic.AverageTimeToCollision,
		StdTimeToCollision:                      traffic.StdTimeToCollision,
	}

	return protoTraffic, nil
}

func FromCurrentTrafficDataByDetectionPointByLaneProtosToModels(trafficList []*pb.CurrentTrafficDataByDetectionPointByLane) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error) {
	protoTraffic := make([]*models.CurrentTrafficDataByDetectionPointByLaneModel, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromCurrentTrafficDataByDetectionPointByLaneProtoToModel(traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromCurrentTrafficDataByDetectionPointByLaneProtoToModel(traffic *pb.CurrentTrafficDataByDetectionPointByLane) (*models.CurrentTrafficDataByDetectionPointByLaneModel, error) {
	modelTraffic := models.CurrentTrafficDataByDetectionPointByLaneModel{
		CreatedAt:                              traffic.CreatedAt.AsTime(),
		ModifiedAt:                             traffic.ModifiedAt.AsTime(),
		DataSourceName:                         traffic.DataSourceName,
		LaneID:                                 traffic.LaneId,
		DetectionTimestamp:                     traffic.DetectionTimestamp.AsTime(),
		DetectionInterval:                      traffic.DetectionInterval,
		CountVehicleClass1:                     traffic.CountVehicleClass_1,
		CountVehicleClass2:                     traffic.CountVehicleClass_2,
		CountVehicleClass3:                     traffic.CountVehicleClass_3,
		CountVehicleClass4:                     traffic.CountVehicleClass_4,
		CountVehicleClass5:                     traffic.CountVehicleClass_5,
		CountVehicleClass6:                     traffic.CountVehicleClass_6,
		CountVehicleClass7:                     traffic.CountVehicleClass_7,
		CountVehicleClass8:                     traffic.CountVehicleClass_8,
		CountVehicleClassAll:                   traffic.CountVehicleClassAll,
		HarmonicMeanSpeedVehicleClass1:         traffic.HarmonicMeanSpeedVehicleClass_1,
		HarmonicMeanSpeedVehicleClass2:         traffic.HarmonicMeanSpeedVehicleClass_2,
		HarmonicMeanSpeedVehicleClass3:         traffic.HarmonicMeanSpeedVehicleClass_3,
		HarmonicMeanSpeedVehicleClass4:         traffic.HarmonicMeanSpeedVehicleClass_4,
		HarmonicMeanSpeedVehicleClass5:         traffic.HarmonicMeanSpeedVehicleClass_5,
		HarmonicMeanSpeedVehicleClass6:         traffic.HarmonicMeanSpeedVehicleClass_6,
		HarmonicMeanSpeedVehicleClass7:         traffic.HarmonicMeanSpeedVehicleClass_7,
		HarmonicMeanSpeedVehicleClass8:         traffic.HarmonicMeanSpeedVehicleClass_8,
		HarmonicMeanSpeedVehicleClassAll:       traffic.HarmonicMeanSpeedVehicleClassAll,
		CountDetectedSpeedVehicleUnder50:       traffic.CountDetectedSpeedVehicleUnder_50,
		CountDetectedSpeedVehicleBetween50_100: traffic.CountDetectedSpeedVehicleBetween_50_100,
		CountDetectedSpeedVehicleOver100:       traffic.CountDetectedSpeedVehicleOver_100,
		AverageVehicleLength:                   traffic.AverageVehicleLength,
		AverageHeadway:                         traffic.AverageHeadway,
		StdHeadway:                             traffic.StdHeadway,
		AverageTimeToCollision:                 traffic.AverageTimeToCollision,
		StdTimeToCollision:                     traffic.StdTimeToCollision,
	}

	return &modelTraffic, nil

}

func FromCurrentTrafficDataByDetectionPointByLaneStatisticsModelsToProto(list []models.CurrentTrafficDataByDetectionPointByLaneStatistics) ([]*pb.CurrentTrafficDataByDetectionPointByLaneStatistics, error) {
	protoStats := make([]*pb.CurrentTrafficDataByDetectionPointByLaneStatistics, len(list))
	for i, traffic := range list {
		protoStatsList, err := FromCurrentTrafficDataByDetectionPointByLaneStatisticsModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoStats[i] = protoStatsList
	}
	return protoStats, nil
}

func FromCurrentTrafficDataByDetectionPointByLaneStatisticsModelToProto(dto *models.CurrentTrafficDataByDetectionPointByLaneStatistics) (*pb.CurrentTrafficDataByDetectionPointByLaneStatistics, error) {
	return &pb.CurrentTrafficDataByDetectionPointByLaneStatistics{
		LaneId:                                    dto.LaneID,
		RecordsCount:                              dto.RecordsCount,
		FirstRecordTimestamp:                      timestamppb.New(dto.FirstRecordTimestamp),
		LastRecordTimestamp:                       timestamppb.New(dto.LastRecordTimestamp),
		LongestDataGap:                            dto.LongestDataGap,
		LongestDataGapStartTimestamp:              timestamppb.New(dto.LongestDataGapStartTimestamp),
		LongestDataGapEndTimestamp:                timestamppb.New(dto.LongestDataGapEndTimestamp),
		MissingIntervalDetectionRate:              dto.MissingIntervalDetectionRate,
		TotalCountAllVehicles:                     dto.TotalCountAllVehicles,
		TotalHarmonicMeanSpeedAllRecords:          dto.TotalHarmonicMeanSpeedAllRecords,
		PercentageRecordsWithCounts:               dto.PercentageRecordsWithCounts,
		PercentageRecordsWithPositiveAverageSpeed: dto.PercentageRecordsWithPositiveAverageSpeed,
	}, nil
}

func FromCurrentTrafficDataByDetectionPointByLaneAggregateModelsToProto(list []models.CurrentTrafficDataByDetectionPointByLaneAggregate) ([]*pb.CurrentTrafficDataByDetectionPointByLaneAggregate, error) {
	protoStats := make([]*pb.CurrentTrafficDataByDetectionPointByLaneAggregate, len(list))
	for i, traffic := range list {
		protoStatsList, err := FromCurrentTrafficDataByDetectionPointByLaneAggregateModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoStats[i] = protoStatsList
	}
	return protoStats, nil
}

func FromCurrentTrafficDataByDetectionPointByLaneAggregateModelToProto(dto *models.CurrentTrafficDataByDetectionPointByLaneAggregate) (*pb.CurrentTrafficDataByDetectionPointByLaneAggregate, error) {
	return &pb.CurrentTrafficDataByDetectionPointByLaneAggregate{
		DataSourceName:                               dto.DataSourceName,
		DetectionTimestamp:                           timestamppb.New(dto.DetectionTimestamp),
		RecordsCount:                                 dto.RecordsCount,
		AggPercentageRecords:                         dto.AggPercentageRecords,
		AggPercentageRecordsWithCounts:               dto.AggPercentageRecordsWithCounts,
		AggPercentageRecordsWithPositiveAverageSpeed: dto.AggPercentageRecordsWithPositiveAverageSpeed,
		AggCountVehicleClassAll:                      dto.AggCountVehicleClassAll,
		AggHarmonicMeanSpeedVehicleClassAll:          dto.AggHarmonicMeanSpeedVehicleClassAll,
		AggAverageVehicleLength:                      dto.AggAverageVehicleLength,
		AggAverageHeadway:                            dto.AggAverageHeadway,
		AggStdHeadway:                                dto.AggStdHeadway,
		AggAverageTimeToCollision:                    dto.AggAverageTimeToCollision,
		AggStdTimeToCollision:                        dto.AggStdTimeToCollision,
	}, nil
}
