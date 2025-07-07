package dtos

import (
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func DataSourceTypeProtoFromModel(model models.DataSourceType) pb.DataSourceType {
	switch model {
	case models.Simulator:
		return pb.DataSourceType_DATA_SOURCE_TYPE_SIMULATOR
	case models.Real:
		return pb.DataSourceType_DATA_SOURCE_TYPE_REAL
	default:
		return pb.DataSourceType_DATA_SOURCE_TYPE_UNSPECIFIED
	}
}

func DataSourceTypeProtoToModel(proto pb.DataSourceType) models.DataSourceType {

	switch proto {
	case pb.DataSourceType_DATA_SOURCE_TYPE_SIMULATOR:
		return models.Simulator
	case pb.DataSourceType_DATA_SOURCE_TYPE_REAL:
		return models.Real
	default:
		return models.DataSourceType("")
	}
}

func NewDataSourceProtoFromModel(model *models.DataSource) pb.DataSource {
	rnid := ""
	if model.RoadNetworkId != nil {
		rnid = *model.RoadNetworkId
	}

	return pb.DataSource{
		Name:              model.Name,
		Description:       model.Description,
		Type:              DataSourceTypeProtoFromModel(model.Type),
		Owner:             model.Owner,
		Groups:            model.Groups,
		ModifiedBy:        model.ModifiedBy,
		CreatedAt:         timestamppb.New(model.CreatedAt),
		ModifiedAt:        timestamppb.New(model.ModifiedAt),
		RoadNetworkId:     rnid,
		DetectionSections: DetectionSectionsProtosFromModels(model.DetectionSections),
		DetectionPoints:   DetectionPointsProtosFromModels(model.DetectionPoints),
	}
}

func DataSourceProtoToModel(dto *pb.DataSource) models.DataSource {

	var networkId *string = nil
	if dto.RoadNetworkId != "" {
		networkId = &dto.RoadNetworkId
	}
	return models.DataSource{
		Name:              dto.Name,
		Description:       dto.Description,
		Type:              DataSourceTypeProtoToModel(dto.Type),
		RoadNetworkId:     networkId,
		Owner:             dto.Owner,
		Groups:            dto.Groups,
		CreatedAt:         dto.CreatedAt.AsTime(),
		ModifiedAt:        dto.ModifiedAt.AsTime(),
		DetectionSections: DetectionSectionProtosToModels(dto.DetectionSections),
		DetectionPoints:   DetectionPointProtosToModels(dto.DetectionPoints),
	}
}

func UpdateDataSourceProtoToModel(dto *pb.UpdateDataSource) models.UpdateDataSource {

	if dto == nil {
		return models.UpdateDataSource{}
	}

	dataSourceType := DataSourceTypeProtoToModel(dto.Type)

	return models.UpdateDataSource{
		Name:          dto.Name,
		Description:   &dto.Description,
		Type:          &dataSourceType,
		RoadNetworkId: &dto.RoadNetworkId,
		Groups:        dto.Groups,
		Owner:         &dto.Owner,
	}
}

func DataSourceListParamsProtoToModel(dto *pb.DataSourceListParams) models.DataSourceListParams {
	if dto == nil {
		return models.DataSourceListParams{}
	}
	return models.DataSourceListParams{
		DetectionSections: dto.DetectionSections,
		DetectionPoints:   dto.DetectionPoints,
	}
}

func DataSourceGetParamsProtoToModel(dto *pb.DataSourceGetParams) models.DataSourceGetParams {
	if dto == nil {
		return models.DataSourceGetParams{}
	}
	return models.DataSourceGetParams{
		DetectionSections: dto.DetectionSections,
		DetectionPoints:   dto.DetectionPoints,
	}
}
