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
	storageModels "spoke7-go/internal/storage/models"
	"spoke7-go/pkg/logger"
	"spoke7-go/pkg/utils"
	"strings"

	"spoke7-go/internal/metadata/services"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var DETECTION_POINT_MICROSERVICE string = "metadata"
var DETECTION_POINT_RESOURCE string = "detection_point"

type detectionPointController struct {
	service           services.DetectionPointService
	fileUploadService services.UploadFileService
	logger            logger.Logger

	pb.UnimplementedDetectionPointServiceServer
}

func NewDetectionPointController(service services.DetectionPointService, fileUploadService services.UploadFileService, logger logger.Logger) pb.DetectionPointServiceServer {
	return &detectionPointController{service: service, fileUploadService: fileUploadService, logger: logger}
}

func (dc *detectionPointController) CreateDetectionPoints(ctx context.Context, req *pb.CreateDetectionPointRequest) (*pb.DetectionPoint, error) {
	logText := "Request Create Detection Point"
	additionalFields := []zap.Field{zap.String("detection_point_id", req.DetectionPoint.Id)}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "POST", req.DatasourceName, additionalFields...)

	detectionPointModel := dtos.DetectionPointProtoToModel(req.DetectionPoint)
	if err := dc.service.Create(ctx, req.DatasourceName, &detectionPointModel); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "POST", req.DatasourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response := dtos.DetectionPointProtoFromModel(&detectionPointModel)

	return response, nil
}

func (dc *detectionPointController) CreateManyDetectionPoints(ctx context.Context, req *pb.CreateManyDetectionPointRequest) (*pb.CreateManyDetectionPointsResponse, error) {
	logText := "Request Create Many Detection Point"
	ids := []string{}
	for _, dp := range req.DetectionPoint {
		ids = append(ids, dp.Id)
	}
	additionalFields := []zap.Field{zap.String("detection_point_ids", strings.Join(ids, ","))}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "POST", req.DatasourceName, additionalFields...)

	dpmodels := make([]*models.DetectionPoint, 0)
	for _, dp := range req.DetectionPoint {
		detectionPointModel := dtos.DetectionPointProtoToModel(dp)
		dpmodels = append(dpmodels, &detectionPointModel)
	}

	if err := dc.service.CreateMany(ctx, req.DatasourceName, dpmodels); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "POST", req.DatasourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	dp := make([]*pb.DetectionPoint, 0)
	for _, detectionPointModel := range dpmodels {
		response := dtos.DetectionPointProtoFromModel(detectionPointModel)
		dp = append(dp, response)
	}

	return &pb.CreateManyDetectionPointsResponse{
		DetectionPoints: dp,
	}, nil
}

func (dc *detectionPointController) UpdateDetectionPoint(ctx context.Context, req *pb.UpdateDetectionPointRequest) (*pb.DetectionPoint, error) {
	logText := "Request Update Detection Point"
	additionalFields := []zap.Field{zap.String("detection_point_id", req.DetectionPoint.Id)}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "PUT", req.DatasourceName, additionalFields...)

	detectionPointModel := dtos.DetectionPointProtoToModel(req.DetectionPoint)
	if err := dc.service.Update(ctx, req.DatasourceName, &detectionPointModel); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "PUT", req.DatasourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response := dtos.DetectionPointProtoFromModel(&detectionPointModel)

	return response, nil
}

func (dc *detectionPointController) DeleteDetectionPoint(ctx context.Context, req *pb.DeleteDetectionPointRequest) (*emptypb.Empty, error) {
	logText := "Request Delete Detection Point"
	additionalFields := []zap.Field{zap.String("detection_point_id", req.Id)}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "DELETE", req.DatasourceName, additionalFields...)

	if err := dc.service.Delete(ctx, req.DatasourceName, req.Id); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "DELETE", req.DatasourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (dc *detectionPointController) DeleteAllDetectionPoint(ctx context.Context, req *pb.DeleteDetectionPointRequest) (*emptypb.Empty, error) {
	logText := "Request Delete All Detection Point"
	additionalFields := []zap.Field{zap.String("detection_point_id", req.Id)}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "DELETE", req.DatasourceName, additionalFields...)

	if err := dc.service.DeleteAll(ctx, req.DatasourceName); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "DELETE", req.DatasourceName, additionalFields...)

		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (dc *detectionPointController) GetDetectionPoint(ctx context.Context, req *pb.GetDetectionPointRequest) (*pb.DetectionPoint, error) {
	logText := "Request Get Detection Point"
	additionalFields := []zap.Field{zap.String("detection_point_id", req.Id)}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "GET", req.DatasourceName, additionalFields...)

	detectionPointModel, err := dc.service.Get(ctx, req.DatasourceName, req.Id)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "GET", req.DatasourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response := dtos.DetectionPointProtoFromModel(detectionPointModel)
	return response, nil
}

func (dc *detectionPointController) ListDetectionPoints(ctx context.Context, req *pb.ListDetectionPointsRequest) (*pb.ListDetectionPointsResponse, error) {
	logText := "Request List Detection Point"
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "GET", req.DatasourceName)

	detectionPoints, err := dc.service.List(ctx, req.DatasourceName)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "GET", req.DatasourceName)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	result := make([]*pb.DetectionPoint, 0)
	for _, detectionPoint := range detectionPoints {
		dto := dtos.DetectionPointProtoFromModel(detectionPoint)
		result = append(result, dto)
	}
	return &pb.ListDetectionPointsResponse{DetectionPoints: result}, nil
}

func (dc *detectionPointController) BulkCreateDetectionPoints(ctx context.Context, req *pb.BulkDetectionPointRequest) (*pb.BulkDetectionPointResponse, error) {
	logText := "Request Create Bulk Detection Point"
	additionalFields := []zap.Field{zap.String("file", req.File.Filename)}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "POST", req.DatasourceName, additionalFields...)

	datasourceName := req.DatasourceName
	if datasourceName == "" {
		errorText := fmt.Sprintf("%v: %v", logText, "Missing datasource name")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "POST", req.DatasourceName)

		return nil, status.Error(codes.InvalidArgument, "Missing datasource name")
	}

	if req.File == nil {
		errorText := fmt.Sprintf("%v: %v", logText, "Missing file")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "POST", req.DatasourceName)

		return nil, status.Error(codes.InvalidArgument, "Missing file")

	}
	if req.File.Content == nil {
		errorText := fmt.Sprintf("%v: %v", logText, "Missing file content")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "POST", req.DatasourceName)

		return nil, status.Error(codes.InvalidArgument, "Missing file content")
	}
	//handle only json
	if req.File.ContentType != "application/json" {
		errorText := fmt.Sprintf("%v: %v", logText, "Unsupported file type")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "POST", req.DatasourceName)

		return nil, status.Error(codes.InvalidArgument, "Unsupported file type")
	}

	uploadDescription := "automatic upload from BulkCreateDetectionPoints"
	if err := dc.fileUploadService.UploadFile(ctx, datasourceName, req.File.Filename, req.File.Content, uploadDescription, req.File.ContentType, storageModels.FileTypeDetectionPoints); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "POST", req.DatasourceName)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Parse the JSON data
	var detectionPoints []*pb.DetectionPoint
	err := json.Unmarshal(req.File.Content, &detectionPoints)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "POST", req.DatasourceName)

		return nil, status.Error(codes.InvalidArgument, "Failed to parse JSON data")
	}
	// Convert to model
	dpmodels := make([]*models.DetectionPoint, 0)
	for _, dp := range detectionPoints {
		detectionPointModel := dtos.DetectionPointProtoToModel(dp)
		dpmodels = append(dpmodels, &detectionPointModel)
	}
	// Create detection points
	err = dc.service.CreateMany(ctx, req.DatasourceName, dpmodels)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, DETECTION_POINT_MICROSERVICE, DETECTION_POINT_RESOURCE, "POST", req.DatasourceName)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	return &pb.BulkDetectionPointResponse{DetectionPoints: detectionPoints}, nil

}
