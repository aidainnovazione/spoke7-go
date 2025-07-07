package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	errorsInternal "spoke7-go/internal/errors"
	"spoke7-go/internal/metadata/dtos"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/pb"
	"spoke7-go/internal/metadata/services"
	"spoke7-go/pkg/logger"
	"spoke7-go/pkg/utils"

	storageModels "spoke7-go/internal/storage/models"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var DETECTION_SECTION_MICROSERVICE string = "metadata"
var DETECTION_SECTION_RESOURCE string = "detection_section"

type detectionSectionController struct {
	service           services.DetectionSectionService
	fileUploadService services.UploadFileService
	logger            logger.Logger

	pb.UnimplementedDetectionSectionServiceServer
}

func NewDetectionSectionController(service services.DetectionSectionService, fileUploadService services.UploadFileService, logger logger.Logger) pb.DetectionSectionServiceServer {
	return &detectionSectionController{service: service, fileUploadService: fileUploadService, logger: logger}
}

func (dc *detectionSectionController) CreateDetectionSection(ctx context.Context, req *pb.CreateDetectionSectionRequest) (*pb.DetectionSection, error) {
	logText := "Request Create Detection Section"
	additionalFields := []zap.Field{zap.String("detection_section_id", req.DetectionSection.Id)}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "POST", req.DatasourceName, additionalFields...)

	detectionSection := dtos.DetectionSectionProtoToModel(req.DetectionSection)
	if err := dc.service.Create(ctx, req.DatasourceName, &detectionSection); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "POST", req.DatasourceName, additionalFields...)

		return nil, err
	}
	response := dtos.DetectionSectionProtoFromModel(&detectionSection)

	return response, nil
}

func (dc *detectionSectionController) UpdateDetectionSection(ctx context.Context, req *pb.UpdateDetectionSectionRequest) (*pb.DetectionSection, error) {
	logText := "Request Update Detection Section"
	additionalFields := []zap.Field{zap.String("detection_section_id", req.DetectionSection.Id)}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "PUT", req.DatasourceName, additionalFields...)

	detectionSection := dtos.DetectionSectionProtoToModel(req.DetectionSection)
	if err := dc.service.Update(ctx, req.DatasourceName, &detectionSection); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "PUT", req.DatasourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response := dtos.DetectionSectionProtoFromModel(&detectionSection)

	return response, nil
}

func (dc *detectionSectionController) DeleteDetectionSection(ctx context.Context, req *pb.DeleteDetectionSectionRequest) (*emptypb.Empty, error) {
	logText := "Request Delete Detection Section"
	additionalFields := []zap.Field{zap.String("detection_section_id", req.Id)}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "DELETE", req.DatasourceName, additionalFields...)

	if err := dc.service.Delete(ctx, req.DatasourceName, req.Id); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "DELETE", req.DatasourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (dc *detectionSectionController) GetDetectionSection(ctx context.Context, req *pb.GetDetectionSectionRequest) (*pb.DetectionSection, error) {
	logText := "Request Get Detection Section"
	additionalFields := []zap.Field{zap.String("detection_section_id", req.Id)}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "GET", req.DatasourceName, additionalFields...)

	detectionSection, err := dc.service.Get(ctx, req.DatasourceName, req.Id)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "GET", req.DatasourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response := dtos.DetectionSectionProtoFromModel(detectionSection)
	return response, nil
}

func (dc *detectionSectionController) ListDetectionSections(ctx context.Context, req *pb.ListDetectionSectionsRequest) (*pb.ListDetectionSectionsResponse, error) {
	logText := "Request List Detection Section"
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "GET", req.DatasourceName)

	detectionSections, err := dc.service.List(ctx, req.DatasourceName)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "GET", req.DatasourceName)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	detectionSectionDtos := make([]*pb.DetectionSection, 0, len(detectionSections))
	for _, detectionSection := range detectionSections {
		dto := dtos.DetectionSectionProtoFromModel(detectionSection)
		detectionSectionDtos = append(detectionSectionDtos, dto)
	}

	return &pb.ListDetectionSectionsResponse{DetectionSections: detectionSectionDtos}, nil
}

func (dc *detectionSectionController) BulkCreateDetectionSections(ctx context.Context, req *pb.BulkDetectionSectionRequest) (*pb.BulkDetectionSectionResponse, error) {
	logText := "Request Create Bulk Detection Section"
	additionalFields := []zap.Field{zap.String("file", req.File.Filename)}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "POST", req.DatasourceName, additionalFields...)

	datasourceName := req.DatasourceName
	if datasourceName == "" {
		errorText := fmt.Sprintf("%v: %v", logText, "Missing datasource name")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "POST", req.DatasourceName)

		return nil, status.Error(codes.InvalidArgument, "Missing datasource name")
	}

	if req.File == nil {
		errorText := fmt.Sprintf("%v: %v", logText, "Missing file")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "POST", req.DatasourceName)

		return nil, status.Error(codes.InvalidArgument, "Missing file")

	}
	if req.File.Content == nil {
		errorText := fmt.Sprintf("%v: %v", logText, "Missing file content")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "POST", req.DatasourceName)

		return nil, status.Error(codes.InvalidArgument, "Missing file content")
	}
	//handle only json
	if req.File.ContentType != "application/json" {
		errorText := fmt.Sprintf("%v: %v", logText, "Unsupported file type")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "POST", req.DatasourceName)

		return nil, status.Error(codes.InvalidArgument, "Unsupported file type")
	}

	uploadDescription := "automatic upload from BulkCreateDetectionSections"
	if err := dc.fileUploadService.UploadFile(ctx, datasourceName, req.File.Filename, req.File.Content, uploadDescription, req.File.ContentType, storageModels.FileTypeDetectionSections); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "POST", req.DatasourceName)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Parse the JSON data
	var detectionSections []*pb.DetectionSection
	err := json.Unmarshal(req.File.Content, &detectionSections)
	if err != nil {
		fmt.Print(err.Error())
		return nil, status.Error(codes.InvalidArgument, "Failed to parse JSON data")
	}
	// Convert to model
	dpmodels := make([]*models.DetectionSection, 0)
	for _, dp := range detectionSections {
		detectionSectionModel := dtos.DetectionSectionProtoToModel(dp)
		dpmodels = append(dpmodels, &detectionSectionModel)
	}
	// Create detection points
	err = dc.service.CreateMany(ctx, req.DatasourceName, dpmodels)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_SECTION_MICROSERVICE, DETECTION_SECTION_RESOURCE, "POST", req.DatasourceName)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	return &pb.BulkDetectionSectionResponse{DetectionSections: detectionSections}, nil

}
