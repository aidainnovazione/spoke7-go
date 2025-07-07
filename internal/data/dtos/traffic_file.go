package dtos

import (
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"
)

func FromTrafficFileProtoToModel(dto *pb.TrafficFile) models.TrafficFile {
	if dto == nil {
		return models.TrafficFile{}
	}
	return models.TrafficFile{
		Filename:    dto.Filename,
		ContentType: dto.ContentType,
		Content:     dto.Content,
	}
}
