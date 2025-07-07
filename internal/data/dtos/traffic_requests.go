package dtos

import (
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"
)

// DETECTION POINT
func FromTrafficDataByDetectionPointGetParamsProtoToModel(dto *pb.GetTrafficDataByDetectionPointRequest) models.GetTrafficDataByDetectionPointParams {
	if dto == nil {
		return models.GetTrafficDataByDetectionPointParams{}
	}
	return models.GetTrafficDataByDetectionPointParams{
		DataSourceName:     dto.DataSourceName,
		DetectionPointIDs:  dto.DetectionPointIds,
		DetectionTimestamp: dto.DetectionTimestamp.AsTime(),
	}
}

func FromTrafficDataByDetectionPointDeleteParamsProtoToModel(dto *pb.DeleteTrafficDataByDetectionPointRequest) models.DeleteTrafficDataByDetectionPointParams {
	if dto == nil {
		return models.DeleteTrafficDataByDetectionPointParams{}
	}
	return models.DeleteTrafficDataByDetectionPointParams{
		DataSourceName:    dto.DataSourceName,
		DetectionPointIDs: dto.DetectionPointIds,
		StartTimestamp:    dto.StartTimestamp.AsTime(),
		EndTimestamp:      dto.EndTimestamp.AsTime(),
	}
}

func FromTrafficDataByDetectionPointListParamsProtoToModel(dto *pb.ListTrafficDataByDetectionPointRequest) models.ListTrafficDataByDetectionPointParams {
	if dto == nil {
		return models.ListTrafficDataByDetectionPointParams{}
	}
	return models.ListTrafficDataByDetectionPointParams{
		DataSourceName:    dto.DataSourceName,
		DetectionPointIDs: dto.DetectionPointIds,
		StartTime:         dto.StartTime.AsTime(),
		EndTime:           dto.EndTime.AsTime(),
	}
}

func FromTrafficDataByDetectionPointDownloadParamsProtoToModel(dto *pb.DownloadTrafficDataByDetectionPointRequest) models.ListTrafficDataByDetectionPointParams {
	if dto == nil {
		return models.ListTrafficDataByDetectionPointParams{}
	}
	return models.ListTrafficDataByDetectionPointParams{
		DataSourceName:    dto.DataSourceName,
		DetectionPointIDs: dto.DetectionPointIds,
		StartTime:         dto.StartTime.AsTime(),
		EndTime:           dto.EndTime.AsTime(),
	}
}

// SECTION

func FromTrafficDataByDetectionSectionGetParamsProtoToModel(dto *pb.GetTrafficDataByDetectionSectionRequest) models.GetTrafficDataByDetectionSectionParams {
	if dto == nil {
		return models.GetTrafficDataByDetectionSectionParams{}
	}
	return models.GetTrafficDataByDetectionSectionParams{
		DataSourceName:      dto.DataSourceName,
		DetectionSectionIDs: dto.DetectionSectionIds,
		DetectionTimestamp:  dto.DetectionTimestamp.AsTime(),
	}
}

func FromTrafficDataByDetectionSectionDeleteParamsProtoToModel(dto *pb.DeleteTrafficDataByDetectionSectionRequest) models.DeleteTrafficDataByDetectionSectionParams {
	if dto == nil {
		return models.DeleteTrafficDataByDetectionSectionParams{}
	}
	return models.DeleteTrafficDataByDetectionSectionParams{
		DataSourceName:      dto.DataSourceName,
		DetectionSectionIDs: dto.DetectionSectionIds,
		StartTimestamp:      dto.StartTimestamp.AsTime(),
		EndTimestamp:        dto.EndTimestamp.AsTime(),
	}
}

func FromTrafficDataByDetectionSectionListParamsProtoToModel(dto *pb.ListTrafficDataByDetectionSectionRequest) models.ListTrafficDataByDetectionSectionParams {
	if dto == nil {
		return models.ListTrafficDataByDetectionSectionParams{}
	}
	return models.ListTrafficDataByDetectionSectionParams{
		DataSourceName:      dto.DataSourceName,
		DetectionSectionIDs: dto.DetectionSectionIds,
		StartTime:           dto.StartTime.AsTime(),
		EndTime:             dto.EndTime.AsTime(),
	}
}

func FromTrafficDataByDetectionSectionDownloadParamsProtoToModel(dto *pb.DownloadTrafficDataByDetectionSectionRequest) models.ListTrafficDataByDetectionSectionParams {
	if dto == nil {
		return models.ListTrafficDataByDetectionSectionParams{}
	}
	return models.ListTrafficDataByDetectionSectionParams{
		DataSourceName:      dto.DataSourceName,
		DetectionSectionIDs: dto.DetectionSectionIds,
		StartTime:           dto.StartTime.AsTime(),
		EndTime:             dto.EndTime.AsTime(),
	}
}

// detection point by lane
func FromTrafficDataByDetectionPointByLaneGetParamsProtoToModel(dto *pb.GetTrafficDataByDetectionPointByLaneRequest) models.GetTrafficDataByDetectionPointByLaneParams {
	if dto == nil {
		return models.GetTrafficDataByDetectionPointByLaneParams{}
	}
	return models.GetTrafficDataByDetectionPointByLaneParams{
		DataSourceName:     dto.DataSourceName,
		DetectionTimestamp: dto.DetectionTimestamp.AsTime(),
		LaneIDs:            dto.LaneIds,
	}
}

func FromTrafficDataByDetectionPointByLaneDeleteParamsProtoToModel(dto *pb.DeleteTrafficDataByDetectionPointByLaneRequest) models.DeleteTrafficDataByDetectionPointByLaneParams {
	if dto == nil {
		return models.DeleteTrafficDataByDetectionPointByLaneParams{}
	}
	return models.DeleteTrafficDataByDetectionPointByLaneParams{
		DataSourceName: dto.DataSourceName,
		LaneIDs:        dto.LaneIds,
		StartTimestamp: dto.StartTimestamp.AsTime(),
		EndTimestamp:   dto.EndTimestamp.AsTime(),
	}
}

func FromTrafficDataByDetectionPointByLaneListParamsProtoToModel(dto *pb.ListTrafficDataByDetectionPointByLaneRequest) models.ListTrafficDataByDetectionPointByLaneParams {
	if dto == nil {
		return models.ListTrafficDataByDetectionPointByLaneParams{}
	}
	return models.ListTrafficDataByDetectionPointByLaneParams{
		DataSourceName: dto.DataSourceName,
		StartTime:      dto.StartTime.AsTime(),
		EndTime:        dto.EndTime.AsTime(),
		LaneIDs:        dto.LaneIds,
	}
}

func FromTrafficDataByDetectionPointByLaneDownloadParamsProtoToModel(dto *pb.DownloadTrafficDataByDetectionPointByLaneRequest) models.ListTrafficDataByDetectionPointByLaneParams {
	if dto == nil {
		return models.ListTrafficDataByDetectionPointByLaneParams{}
	}
	return models.ListTrafficDataByDetectionPointByLaneParams{
		DataSourceName: dto.DataSourceName,
		StartTime:      dto.StartTime.AsTime(),
		EndTime:        dto.EndTime.AsTime(),
		LaneIDs:        dto.LaneIds,
	}
}
