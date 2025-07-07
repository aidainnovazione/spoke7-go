package dtos

import (
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"
)

func FromTrafficFlowParametersByDayProtoToModel(dto *pb.TrafficFlowParametersByDay) models.TrafficFlowParametersByDay {
	if dto == nil {
		return models.TrafficFlowParametersByDay{}
	}
	return models.TrafficFlowParametersByDay{
		Min:     dto.Min,
		Max:     dto.Max,
		Average: dto.Average,
		Std:     dto.Std,
	}
}

func FromTrafficFlowParametersByDayModelToProto(dto *models.TrafficFlowParametersByDay) *pb.TrafficFlowParametersByDay {

	return &pb.TrafficFlowParametersByDay{
		Min:     dto.Min,
		Max:     dto.Max,
		Average: dto.Average,
		Std:     dto.Std,
	}
}
