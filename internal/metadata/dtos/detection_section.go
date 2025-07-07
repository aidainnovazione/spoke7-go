package dtos

import (
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// type DetectionSectionDto struct {
// 	Id             string      `json:"id"`
// 	DataSourceName string      `json:"dataSourceName"`
// 	Description    string      `json:"description"`
// 	StartLatitude  float64     `json:"startLatitude"`
// 	StartLongitude float64     `json:"startLongitude"`
// 	EndLatitude    float64     `json:"endLatitude"`
// 	EndLongitude   float64     `json:"endLongitude"`
// 	Direction      int         `json:"direction"`
// 	Shape          interface{} `json:"shape"`
// 	RoadNetworkId  string      `json:"roadNetworkId"`
// 	CreatedAt      time.Time   `json:"createdAt"`
// 	ModifiedAt     time.Time   `json:"modifiedAt"`
// }

func DetectionSectionsProtosFromModels(models []models.DetectionSection) []*pb.DetectionSection {
	detectionSections := make([]*pb.DetectionSection, len(models))
	for i, model := range models {
		detectionSections[i] = DetectionSectionProtoFromModel(&model)
	}
	return detectionSections
}

func DetectionSectionProtoFromModel(model *models.DetectionSection) *pb.DetectionSection {
	return &pb.DetectionSection{
		Id:             model.Id,
		DataSourceName: model.DataSourceName,
		Description:    model.Description,
		StartLatitude:  model.StartLatitude,
		StartLongitude: model.StartLongitude,
		EndLatitude:    model.EndLatitude,
		EndLongitude:   model.EndLongitude,
		Direction:      int32(model.Direction),
		// Shape:          model.Shape,
		RoadNetworkId: model.RoadNetworkId,
		CreatedAt:     timestamppb.New(model.CreatedAt),
		ModifiedAt:    timestamppb.New(model.ModifiedAt),
	}
}

func DetectionSectionProtosToModels(dtos []*pb.DetectionSection) []models.DetectionSection {
	detectionSections := make([]models.DetectionSection, len(dtos))
	for i, dto := range dtos {
		detectionSections[i] = DetectionSectionProtoToModel(dto)
	}
	return detectionSections
}

func DetectionSectionProtoToModel(dto *pb.DetectionSection) models.DetectionSection {
	return models.DetectionSection{
		Id:             dto.Id,
		DataSourceName: dto.DataSourceName,
		Description:    dto.Description,
		StartLatitude:  dto.StartLatitude,
		StartLongitude: dto.StartLongitude,
		EndLatitude:    dto.EndLatitude,
		EndLongitude:   dto.EndLongitude,
		Direction:      int(dto.Direction),
		Shape:          dto.Shape,
		RoadNetworkId:  dto.RoadNetworkId,
		CreatedAt:      dto.CreatedAt.AsTime(),
		ModifiedAt:     dto.ModifiedAt.AsTime(),
	}
}
