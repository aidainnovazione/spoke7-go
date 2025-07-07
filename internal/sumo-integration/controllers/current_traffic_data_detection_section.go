package controllers

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"

	errorsInternal "spoke7-go/internal/errors"
	"spoke7-go/internal/storage/models"
	"spoke7-go/internal/sumo-integration/dtos"
	"spoke7-go/internal/sumo-integration/pb"
	"spoke7-go/internal/sumo-integration/services"
	"spoke7-go/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type sumoIntegrationCurrentTrafficDataByDetectionSection struct {
	service           services.SumoIntegrationCurrentTrafficDataByDetectionSectionService
	fileUploadService services.UploadFileService
	logger            logger.Logger
	pb.UnimplementedSumoIntegrationCurrentTrafficDataByDetectionSectionServiceServer
}

func NewSumoIntegrationCurrentTrafficDataByDetectionSectionController(service services.SumoIntegrationCurrentTrafficDataByDetectionSectionService, fileUploadService services.UploadFileService, logger logger.Logger) pb.SumoIntegrationCurrentTrafficDataByDetectionSectionServiceServer {
	return &sumoIntegrationCurrentTrafficDataByDetectionSection{service: service, fileUploadService: fileUploadService, logger: logger}
}

func (dc *sumoIntegrationCurrentTrafficDataByDetectionSection) SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXml(ctx context.Context, req *pb.SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlRequest) (*pb.SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlResponse, error) {
	modelReq := dtos.FromSumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlRequestProtoToModel(req)
	if req.Xml == nil {
		return nil, status.Error(codes.InvalidArgument, "Missing file")
	}
	if req.Xml.Content == nil {
		return nil, status.Error(codes.InvalidArgument, "Missing file content")
	}
	//handle only xml
	if req.Xml.ContentType != "application/xml" && req.Xml.ContentType != "text/xml" {
		return nil, status.Error(codes.InvalidArgument, "Unsupported file type")
	}

	uploadDescription := "automatic upload from SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXml"
	if err := dc.fileUploadService.UploadFile(ctx, modelReq.DataSourceName, req.Xml.Filename, req.Xml.Content, uploadDescription, req.Xml.ContentType, models.FileTypeAggregatedTraffic5MinByDetectionSection); err != nil {
		return &pb.SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlResponse{
			Message: fmt.Sprintf("%v", err),
		}, status.Error(codes.InvalidArgument, err.Error())
	}

	var detector dtos.LaneDetectorModel
	xmlReader := bytes.NewReader(req.Xml.Content)
	xmlDecoder := xml.NewDecoder(xmlReader)
	if err := xmlDecoder.Decode(&detector); err != nil {
		return &pb.SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlResponse{
			Message: fmt.Sprintf("%v", err),
		}, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := dc.service.CreateFromXml(ctx, detector, modelReq.DataSourceName, modelReq.StartTime); err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return &pb.SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlResponse{
			Message: fmt.Sprintf("%v", err),
		}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlResponse{
		Message: "ok",
	}, nil
}
