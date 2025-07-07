package dtos

import (
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func DetectionPointsProtosFromModels(models []models.DetectionPoint) []*pb.DetectionPoint {
	detectionPoints := make([]*pb.DetectionPoint, len(models))
	for i, model := range models {
		detectionPoints[i] = DetectionPointProtoFromModel(&model)
	}
	return detectionPoints
}

func DetectionPointProtoFromModel(model *models.DetectionPoint) *pb.DetectionPoint {
	lanes := make([]*pb.Lane, len(model.Lanes))
	for i, lane := range model.Lanes {
		lanes[i] = LaneModelToProto(lane)
	}

	return &pb.DetectionPoint{
		Id:           model.Id,
		Description:  model.Description,
		CreatedAt:    timestamppb.New(model.CreatedAt),
		ModifiedAt:   timestamppb.New(model.ModifiedAt),
		Lanes:        lanes,
		Properties:   model.Properties,
		GeojsonShape: model.GeojsonShape,
		Coordinates:  CoordinateModelToProto(model.Coordinates),
		Position:     PositionModelToProto(model.Position),
	}
}

func DetectionPointProtosToModels(dtos []*pb.DetectionPoint) []models.DetectionPoint {
	detectionPoints := make([]models.DetectionPoint, len(dtos))
	for i, dto := range dtos {
		detectionPoints[i] = DetectionPointProtoToModel(dto)
	}
	return detectionPoints
}

func DetectionPointProtoToModel(dto *pb.DetectionPoint) models.DetectionPoint {
	lanes := make([]models.Lane, len(dto.Lanes))
	for i, laneDTO := range dto.Lanes {
		lanes[i] = LaneProtoToModel(laneDTO)
	}

	return models.DetectionPoint{
		Id:           dto.Id,
		Description:  dto.Description,
		CreatedAt:    dto.CreatedAt.AsTime(),
		ModifiedAt:   dto.ModifiedAt.AsTime(),
		Lanes:        lanes,
		Properties:   dto.Properties,
		GeojsonShape: dto.GeojsonShape,
		Coordinates:  CoordinateProtoToModel(dto.Coordinates),
		Position:     PositionProtoToModel(dto.Position),
	}
}

func LaneProtoToModel(lane *pb.Lane) models.Lane {
	return models.Lane{
		Id:           lane.Id,
		Description:  lane.Description,
		Index:        lane.Index,
		Properties:   lane.Properties,
		GeojsonShape: lane.GeojsonShape,
		Coordinates:  CoordinateProtoToModel(lane.Coordinates),
		Position:     PositionProtoToModel(lane.Position),
		CreatedAt:    lane.CreatedAt.AsTime(),
		ModifiedAt:   lane.ModifiedAt.AsTime(),
	}
}

func LaneModelToProto(lane models.Lane) *pb.Lane {
	return &pb.Lane{
		Id:           lane.Id,
		Description:  lane.Description,
		Index:        lane.Index,
		Properties:   lane.Properties,
		GeojsonShape: lane.GeojsonShape,
		Coordinates:  CoordinateModelToProto(lane.Coordinates),
		Position:     PositionModelToProto(lane.Position),
		CreatedAt:    timestamppb.New(lane.CreatedAt),
		ModifiedAt:   timestamppb.New(lane.ModifiedAt),
	}
}

func CoordinateProtoToModel(coord *pb.Coordinate) *models.Coordinate {
	if coord == nil {
		return nil
	}
	return &models.Coordinate{
		Latitude:  coord.Latitude,
		Longitude: coord.Longitude,
	}
}

func CoordinateModelToProto(coord *models.Coordinate) *pb.Coordinate {
	if coord == nil {
		return nil
	}
	return &pb.Coordinate{
		Latitude:  coord.Latitude,
		Longitude: coord.Longitude,
	}
}

func PositionProtoToModel(pos *pb.Position) *models.Position {
	if pos == nil {
		return nil
	}
	return &models.Position{
		X: pos.X,
		Y: pos.Y,
	}
}

func PositionModelToProto(pos *models.Position) *pb.Position {
	if pos == nil {
		return nil
	}
	return &pb.Position{
		X: pos.X,
		Y: pos.Y,
	}
}
