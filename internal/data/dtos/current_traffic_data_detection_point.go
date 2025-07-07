package dtos

import (
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromCurrentTrafficDataByDetectionPointModelsToProtos(trafficList []models.CurrentTrafficDataByDetectionPointModel) ([]*pb.CurrentTrafficDataByDetectionPoint, error) {
	protoTraffic := make([]*pb.CurrentTrafficDataByDetectionPoint, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromCurrentTrafficDataByDetectionPointModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromCurrentTrafficDataByDetectionPointModelToProto(traffic *models.CurrentTrafficDataByDetectionPointModel) (*pb.CurrentTrafficDataByDetectionPoint, error) {
	protoTraffic := &pb.CurrentTrafficDataByDetectionPoint{
		CreatedAt:                               timestamppb.New(traffic.CreatedAt),
		ModifiedAt:                              timestamppb.New(traffic.ModifiedAt),
		DataSourceName:                          traffic.DataSourceName,
		DetectionPointId:                        traffic.DetectionPointID,
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
		CountVehicleClassEquivalent:             traffic.CountVehicleClassEquivalent,
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

func FromCurrentTrafficDataByDetectionPointProtosToModels(trafficList []*pb.CurrentTrafficDataByDetectionPoint) ([]*models.CurrentTrafficDataByDetectionPointModel, error) {
	protoTraffic := make([]*models.CurrentTrafficDataByDetectionPointModel, len(trafficList))
	for i, traffic := range trafficList {
		protoTrafficList, err := FromCurrentTrafficDataByDetectionPointProtoToModel(traffic)
		if err != nil {
			return nil, err
		}
		protoTraffic[i] = protoTrafficList
	}
	return protoTraffic, nil
}

func FromCurrentTrafficDataByDetectionPointProtoToModel(traffic *pb.CurrentTrafficDataByDetectionPoint) (*models.CurrentTrafficDataByDetectionPointModel, error) {
	modelTraffic := models.CurrentTrafficDataByDetectionPointModel{
		CreatedAt:                              traffic.CreatedAt.AsTime(),
		ModifiedAt:                             traffic.ModifiedAt.AsTime(),
		DataSourceName:                         traffic.DataSourceName,
		DetectionPointID:                       traffic.DetectionPointId,
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
		CountVehicleClassEquivalent:            traffic.CountVehicleClassEquivalent,
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

func FromCurrentTrafficDataByDetectionPointStatisticsModelsToProto(list []models.CurrentTrafficDataByDetectionPointStatistics) ([]*pb.CurrentTrafficDataByDetectionPointStatistics, error) {
	protoStats := make([]*pb.CurrentTrafficDataByDetectionPointStatistics, len(list))
	for i, traffic := range list {
		protoStatsList, err := FromCurrentTrafficDataByDetectionPointStatisticsModelToProto(&traffic)
		if err != nil {
			return nil, err
		}
		protoStats[i] = protoStatsList
	}
	return protoStats, nil
}

func FromCurrentTrafficDataByDetectionPointStatisticsModelToProto(dto *models.CurrentTrafficDataByDetectionPointStatistics) (*pb.CurrentTrafficDataByDetectionPointStatistics, error) {
	return &pb.CurrentTrafficDataByDetectionPointStatistics{
		DetectionPointId:                          dto.DetectionPointID,
		RecordsCount:                              dto.RecordsCount,
		FirstRecordTimestamp:                      timestamppb.New(dto.FirstRecordTimestamp),
		LastRecordTimestamp:                       timestamppb.New(dto.LastRecordTimestamp),
		LongestDataGap:                            dto.LongestDataGap,
		LongestDataGapStartTimestamp:              timestamppb.New(dto.LongestDataGapStartTimestamp),
		LongestDataGapEndTimestamp:                timestamppb.New(dto.LongestDataGapEndTimestamp),
		MissingIntervalDetectionRate:              dto.MissingIntervalDetectionRate,
		TotalCountEquivalentVehicles:              dto.TotalCountEquivalentVehicles,
		TotalHarmonicMeanSpeedAllRecords:          dto.TotalHarmonicMeanSpeedAllRecords,
		PercentageRecordsWithEquivalentCounts:     dto.PercentageRecordsWithEquivalentCounts,
		PercentageRecordsWithPositiveAverageSpeed: dto.PercentageRecordsWithPositiveAverageSpeed,
	}, nil
}
