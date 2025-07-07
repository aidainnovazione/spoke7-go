package dtos

import (
	"spoke7-go/internal/sumo-integration/models"
	"spoke7-go/internal/sumo-integration/pb"
)

func FromSumoIntegrationCreateCurrentTrafficDataByDetectionPointByLaneFromXmlRequestProtoToModel(dto *pb.SumoIntegrationCreateCurrentTrafficDataByDetectionPointByLaneFromXmlRequest) models.SumoIntegrationCreateCurrentTrafficDataByDetectionPointByLaneFromXmlRequestModel {
	if dto == nil {
		return models.SumoIntegrationCreateCurrentTrafficDataByDetectionPointByLaneFromXmlRequestModel{}
	}
	return models.SumoIntegrationCreateCurrentTrafficDataByDetectionPointByLaneFromXmlRequestModel{
		DataSourceName: dto.DataSourceName,
		Xml:            FromSumoXmlFileProtoToModel(dto.Xml),
		StartTime:      dto.StartTime.AsTime(),
	}
}

func FromSumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlRequestProtoToModel(dto *pb.SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlRequest) models.SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlRequestModel {
	if dto == nil {
		return models.SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlRequestModel{}
	}
	return models.SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlRequestModel{
		DataSourceName: dto.DataSourceName,
		Xml:            FromSumoXmlFileProtoToModel(dto.Xml),
		StartTime:      dto.StartTime.AsTime(),
	}
}

func FromSumoXmlFileProtoToModel(dto *pb.XmlFile) models.XmlFile {
	if dto == nil {
		return models.XmlFile{}
	}
	return models.XmlFile{
		Filename:    dto.Filename,
		ContentType: dto.ContentType,
		Content:     dto.Content,
	}
}
