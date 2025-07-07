package controllers

import (
	"context"
	"encoding/xml"
	"fmt"
	"spoke7-go/internal/metadata/models"
	storageModels "spoke7-go/internal/storage/models"
	"spoke7-go/internal/sumo-integration/dtos"
	"spoke7-go/internal/sumo-integration/pb"
	"spoke7-go/internal/sumo-integration/services"
	"spoke7-go/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type dataSourceController struct {
	service           services.DataSourceService
	fileUploadService services.UploadFileService
	logger            logger.Logger
	pb.UnimplementedSumoIntegrationDataSourceServiceServer
	// UnimplementedSumoIntegrationDataSourceServiceServer must be embedded to have forward compatible implementations.
}

func NewDataSourceController(service services.DataSourceService, fileUploadService services.UploadFileService, logger logger.Logger) pb.SumoIntegrationDataSourceServiceServer {
	return &dataSourceController{service: service, fileUploadService: fileUploadService, logger: logger}
}

func (dc *dataSourceController) ImportRoadNetworkFromXml(ctx context.Context, req *pb.RoadNetworkImportRequest) (*pb.RoadNetworkImportResponse, error) {
	// TODO implement import from xml
	//1. call the service to import the road network from xml

	if req.Xml == nil {
		return nil, status.Error(codes.InvalidArgument, "Missing file")

	}
	if req.Xml.Content == nil {
		return nil, status.Error(codes.InvalidArgument, "Missing file content")
	}
	//handle only xml file
	if req.Xml.ContentType != "application/xml" && req.Xml.ContentType != "text/xml" {
		return nil, status.Error(codes.InvalidArgument, "Unsupported file type")
	}

	var network dtos.Network
	err := xml.Unmarshal(req.Xml.Content, &network)
	if err != nil {
		fmt.Print(err.Error())
		return nil, status.Error(codes.InvalidArgument, "Failed to parse JSON data")
	}

	networkProperties := dtos.GetNetworkProperties(network)

	featureCollection := dtos.NetworkToGeoJSON(network)
	if err != nil {
		fmt.Print(err.Error())
		return nil, status.Error(codes.InvalidArgument, "Failed to convert network to GeoJSON")
	}

	md := models.RoadNetwork{
		Name:       req.Name,
		Geom:       &featureCollection,
		Properties: dtos.ToRoadNetworkPropertiesDTOs(networkProperties),
	}

	response, err := dc.service.CreateRoadNetwork(ctx, md)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to convert network to GeoJSON")
	}

	//2. return the response

	return &pb.RoadNetworkImportResponse{
		Id:      response.ID,
		Message: "ok",
	}, err

}

func (dc *dataSourceController) ImportFromXml(ctx context.Context, req *pb.ImportRequest) (*pb.ImportResponse, error) {

	//1. import XML files
	if req.NetFile == nil {
		return nil, status.Error(codes.InvalidArgument, "Missing file")

	}
	if req.NetFile.Content == nil {
		return nil, status.Error(codes.InvalidArgument, "Missing file content")
	}

	uploadDescription := "automatic upload for network from SumoImportFromXml"
	if err := dc.fileUploadService.UploadFile(ctx, req.DataSource, req.NetFile.Filename, req.NetFile.Content, uploadDescription, req.NetFile.ContentType, storageModels.FileTypeSumoNetwork); err != nil {
		return &pb.ImportResponse{
			Message: fmt.Sprintf("%v", err),
		}, status.Error(codes.InvalidArgument, err.Error())
	}

	/*handle road network */
	var network dtos.Network
	err := xml.Unmarshal(req.NetFile.Content, &network)
	if err != nil {
		fmt.Print(err.Error())
		return nil, status.Error(codes.InvalidArgument, "Failed to parse JSON data")
	}

	networkProperties := dtos.GetNetworkProperties(network)

	featureCollection := dtos.NetworkToGeoJSON(network)
	if err != nil {
		fmt.Print(err.Error())
		return nil, status.Error(codes.InvalidArgument, "Failed to convert network to GeoJSON")
	}

	md := models.RoadNetwork{
		Name:       req.NetworkName,
		Geom:       &featureCollection,
		Properties: dtos.ToRoadNetworkPropertiesDTOs(networkProperties),
	}

	_, err = dc.service.CreateAndAssociateRoadNetwork(ctx, req.DataSource, md)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to convert network to GeoJSON")
	}

	//2. Handle additional files

	//1. import XML files
	if req.AdditionsFiles == nil {
		return nil, status.Error(codes.InvalidArgument, "Missing file")

	}
	if req.AdditionsFiles.Content == nil {
		return nil, status.Error(codes.InvalidArgument, "Missing file content")
	}

	uploadAddDescription := "automatic upload for additional from SumoImportFromXml"
	if err := dc.fileUploadService.UploadFile(ctx, req.DataSource, req.AdditionsFiles.Filename, req.AdditionsFiles.Content, uploadAddDescription, req.AdditionsFiles.ContentType, storageModels.FileTypeSumoAdditional); err != nil {
		return &pb.ImportResponse{
			Message: fmt.Sprintf("%v", err),
		}, status.Error(codes.InvalidArgument, err.Error())
	}

	/*handle road network */
	var additional dtos.Additional
	err = xml.Unmarshal(req.AdditionsFiles.Content, &additional)
	if err != nil {
		fmt.Print(err.Error())
		return nil, status.Error(codes.InvalidArgument, "Failed to parse JSON data")
	}

	detectionPoint := dtos.GetInductionLoops(network, additional)

	//2. create detection points
	dc.service.CreateManyDetectionPoints(ctx, req.DataSource, detectionPoint)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to convert network to GeoJSON")
	}

	return &pb.ImportResponse{

		Message:         "ok",
		NetCount:        1,
		AdditionalCount: int64(len(detectionPoint)),
		RouteCount:      0,
	}, nil

}
