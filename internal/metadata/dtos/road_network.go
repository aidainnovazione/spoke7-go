package dtos

import (
	"encoding/json"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/pb"

	"github.com/paulmach/orb/geojson"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func RoadNetworkProtoToModel(req *pb.RoadNetwork) (*models.RoadNetwork, error) {

	//var geo geojson.FeatureCollection

	geoJson, err := req.Geometry.MarshalJSON()
	if err != nil {
		return nil, err
	}

	geo, err := geojson.UnmarshalFeatureCollection(geoJson)
	if err != nil {
		return nil, err
	}

	return &models.RoadNetwork{
		ID:         req.Id,
		Name:       req.Name,
		Geom:       geo,
		Properties: req.GetProperties(),
		Owner:      req.Owner,
		Groups:     req.Groups,
		CreatedAt:  req.CreatedAt.AsTime(),
		ModifiedAt: req.ModifiedAt.AsTime(),
	}, nil
}

func RoadNetworkModelToProto(roadNetwork *models.RoadNetwork) (*pb.RoadNetwork, error) {

	geoJsonBytes, err := json.Marshal(roadNetwork.Geom) // Usa json.Marshal invece di MarshalJSON()
	if err != nil {
		return nil, err
	}

	var geojsonMap map[string]interface{}
	if err := json.Unmarshal(geoJsonBytes, &geojsonMap); err != nil {
		return nil, err
	}

	// Converti il map[string]interface{} in google.protobuf.Struct
	geojsonStruct, err := structpb.NewStruct(geojsonMap)
	if err != nil {
		return nil, err
	}

	return &pb.RoadNetwork{
		Id:         roadNetwork.ID,
		Name:       roadNetwork.Name,
		Geometry:   geojsonStruct,
		Properties: roadNetwork.Properties,
		Owner:      roadNetwork.Owner,
		Groups:     roadNetwork.Groups,
		CreatedAt:  timestamppb.New(roadNetwork.CreatedAt),
		ModifiedAt: timestamppb.New(roadNetwork.ModifiedAt),
	}, err
}
