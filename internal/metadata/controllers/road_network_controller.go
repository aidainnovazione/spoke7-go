package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"spoke7-go/internal/metadata/dtos"
	"spoke7-go/internal/metadata/pb"
	"spoke7-go/internal/metadata/services"
	"spoke7-go/pkg/logger"
	"spoke7-go/pkg/utils"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
)

var ROADNETWORK_MICROSERVICE string = "metadata"
var ROADNETWORK_RESOURCE string = "road_network"

type roadNetworkController struct {
	service services.RoadNetworkService
	logger  logger.Logger
	pb.UnimplementedRoadNetworkServiceServer
}

func NewRoadNetworkController(service services.RoadNetworkService, logger logger.Logger) pb.RoadNetworkServiceServer {
	return &roadNetworkController{service: service, logger: logger}
}

// CreateRoadNetwork implements pb.RoadNetworkServiceServer.
func (r *roadNetworkController) CreateRoadNetwork(ctx context.Context, req *pb.RoadNetworkCreateRequest) (*pb.RoadNetwork, error) {
	logText := "Request Create Road Network"
	additionalFields := []zap.Field{zap.String("road_network_id", req.RoadNetwork.Id), zap.String("road_network_name", req.RoadNetwork.Name)}
	_ = utils.AddControllerRequestLogging(r.logger, ctx, logText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "POST", "NONE", additionalFields...)

	roadNetworkCreate, err := dtos.RoadNetworkProtoToModel(req.RoadNetwork)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "POST", "NONE", additionalFields...)

		return nil, err
	}

	if roadNetworkCreate.ID == "" {
		roadNetworkCreate.ID = uuid.NewString()
	}
	additionalFields = []zap.Field{zap.String("road_network_id", roadNetworkCreate.ID), zap.String("road_network_name", roadNetworkCreate.Name)}

	roadNetworkRes, err := r.service.Create(ctx, roadNetworkCreate)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "POST", "NONE", additionalFields...)

		return nil, err
	}

	response, err := dtos.RoadNetworkModelToProto(roadNetworkRes)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "POST", "NONE", additionalFields...)

		return nil, err
	}

	return response, nil
}

// ListRoadNetworks implements pb.RoadNetworkServiceServer.
func (r *roadNetworkController) ListRoadNetworks(ctx context.Context, req *pb.RoadNetworkListParams) (*pb.RoadNetworkListResponse, error) {
	logText := "Request List Road Network"
	_ = utils.AddControllerRequestLogging(r.logger, ctx, logText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "GET", "NONE")

	networks, err := r.service.List(ctx)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "GET", "NONE")

		return nil, err
	}

	var protoNetworks []*pb.RoadNetwork
	for _, network := range networks {
		protoNetwork, err := dtos.RoadNetworkModelToProto(network)
		if err != nil {
			errorText := fmt.Sprintf("%v: %v", logText, err.Error())
			_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "GET", "NONE")

			return nil, err
		}
		protoNetworks = append(protoNetworks, protoNetwork)
	}

	return &pb.RoadNetworkListResponse{RoadNetworks: protoNetworks}, nil
}

// DeleteRoadNetwork implements pb.RoadNetworkServiceServer.
func (r *roadNetworkController) DeleteRoadNetwork(ctx context.Context, req *pb.RoadNetworkDeleteRequest) (*emptypb.Empty, error) {
	logText := "Request Delete Road Network"
	additionalFields := []zap.Field{zap.String("road_network_id", req.Id)}
	_ = utils.AddControllerRequestLogging(r.logger, ctx, logText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "DELETE", "NONE", additionalFields...)

	err := r.service.Delete(ctx, req.Id)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "DELETE", "NONE", additionalFields...)

		return nil, err
	}
	return nil, nil
}

// GetRoadNetwork implements pb.RoadNetworkServiceServer.
func (r *roadNetworkController) GetRoadNetwork(ctx context.Context, req *pb.RoadNetworkGetRequest) (*pb.RoadNetwork, error) {
	logText := "Request Get Road Network"
	additionalFields := []zap.Field{zap.String("road_network_id", req.Id)}
	_ = utils.AddControllerRequestLogging(r.logger, ctx, logText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "GET", "NONE", additionalFields...)

	network, err := r.service.Get(ctx, req.Id)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "GET", "NONE", additionalFields...)

		return nil, err
	}

	return dtos.RoadNetworkModelToProto(network)

}

func (r *roadNetworkController) GetRoadNetworkGeoJSON(ctx context.Context, req *pb.RoadNetworkGetRequest) (*structpb.Struct, error) {
	logText := "Request Get Road Network As Geojson"
	additionalFields := []zap.Field{zap.String("road_network_id", req.Id)}
	_ = utils.AddControllerRequestLogging(r.logger, ctx, logText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "GET", "NONE", additionalFields...)

	network, err := r.service.Get(ctx, req.Id)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "GET", "NONE", additionalFields...)

		return nil, err
	}

	result, err := dtos.RoadNetworkModelToProto(network)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "GET", "NONE", additionalFields...)

		return nil, err
	}
	return result.Geometry, nil
}

// UpdateRoadNetwork implements pb.RoadNetworkServiceServer.
func (r *roadNetworkController) UpdateRoadNetwork(ctx context.Context, req *pb.RoadNetworkUpdateRequest) (*pb.RoadNetwork, error) {
	logText := "Request Update Road Network"
	additionalFields := []zap.Field{zap.String("road_network_id", req.Id)}
	_ = utils.AddControllerRequestLogging(r.logger, ctx, logText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

	if req.Id == "" || req.Id != req.RoadNetwork.Id {
		errorText := fmt.Sprintf("%v: %v", "Missing or invalid network id")
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, status.Error(codes.InvalidArgument, "Missing or invalid network id")
	}

	roadNetworkCreate, err := dtos.RoadNetworkProtoToModel(req.RoadNetwork)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)
		return nil, err
	}

	roadNetworkRes, err := r.service.Create(ctx, roadNetworkCreate)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)
		return nil, err
	}

	response, err := dtos.RoadNetworkModelToProto(roadNetworkRes)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, err
	}

	return response, nil
}

// UploadRoadNetwork implements pb.RoadNetworkServiceServer.
func (r *roadNetworkController) UploadRoadNetwork(ctx context.Context, req *pb.RoadNetworkFile) (*pb.RoadNetwork, error) {
	logText := "Request Upload Road Network"
	additionalFields := []zap.Field{zap.String("file", req.File.Filename)}
	_ = utils.AddControllerRequestLogging(r.logger, ctx, logText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

	if req.File == nil {
		errorText := fmt.Sprintf("%v: %v", "Missing file")
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, status.Error(codes.InvalidArgument, "Missing file")
	}

	if req.File.Content == nil {
		errorText := fmt.Sprintf("%v: %v", "Missing file content")
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, status.Error(codes.InvalidArgument, "Missing file content")
	}
	//handle only json
	if req.File.ContentType != "application/json" {
		errorText := fmt.Sprintf("%v: %v", "Unsupported file type")
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, status.Error(codes.InvalidArgument, "Unsupported file type")
	}
	// Parse the JSON data
	roadNetwork := &pb.RoadNetwork{
		Name: req.File.Filename,
	}
	err := json.Unmarshal(req.File.Content, &roadNetwork.Geometry)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, status.Error(codes.InvalidArgument, "Failed to parse JSON data")
	}

	roadNetworkCreate, err := dtos.RoadNetworkProtoToModel(roadNetwork)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, err
	}

	if roadNetworkCreate.ID == "" {
		roadNetworkCreate.ID = uuid.NewString()
	}
	additionalFields = []zap.Field{zap.String("file", req.File.Filename), zap.String("road_network_id", roadNetworkCreate.ID)}

	roadNetworkRes, err := r.service.Create(ctx, roadNetworkCreate)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, err
	}
	response, err := dtos.RoadNetworkModelToProto(roadNetworkRes)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, err
	}

	return response, nil

}

func (r *roadNetworkController) UpdateRoadNetworkByImport(ctx context.Context, req *pb.RoadNetworkUpdateByImportRequest) (*pb.RoadNetwork, error) {
	logText := "Request Update Road Network By Import"
	additionalFields := []zap.Field{zap.String("file", req.File.Filename), zap.String("road_network_id", req.Id)}

	_ = utils.AddControllerRequestLogging(r.logger, ctx, logText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

	if req.File == nil {
		errorText := fmt.Sprintf("%v: %v", "Missing file")
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, status.Error(codes.InvalidArgument, "Missing file")

	}
	if req.File.Content == nil {
		errorText := fmt.Sprintf("%v: %v", "Missing file content")
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, status.Error(codes.InvalidArgument, "Missing file content")
	}
	//handle only json
	if req.File.ContentType != "application/json" {
		errorText := fmt.Sprintf("%v: %v", "Unsupported file type")
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, status.Error(codes.InvalidArgument, "Unsupported file type")
	}
	// Parse the JSON data
	roadNetwork := &pb.RoadNetwork{
		Id: req.Id,
	}
	err := json.Unmarshal(req.File.Content, &roadNetwork.Geometry)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to parse JSON data")
	}

	roadNetworkCreate, err := dtos.RoadNetworkProtoToModel(roadNetwork)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, err
	}

	if roadNetworkCreate.ID == "" {
		roadNetworkCreate.ID = uuid.NewString()
	}

	roadNetworkRes, err := r.service.Update(ctx, roadNetworkCreate)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, err
	}

	response, err := dtos.RoadNetworkModelToProto(roadNetworkRes)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", err.Error())
		_ = utils.AddControllerErrorLogging(r.logger, ctx, errorText, ROADNETWORK_MICROSERVICE, ROADNETWORK_RESOURCE, "PUT", "NONE", additionalFields...)

		return nil, err
	}

	return response, nil
}
