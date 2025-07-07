package controllers

import (
	"context"
	"errors"
	"fmt"
	"spoke7-go/internal/data/dtos"
	"spoke7-go/internal/data/pb"
	"spoke7-go/internal/data/services"
	errorsInternal "spoke7-go/internal/errors"
	"spoke7-go/internal/storage/models"
	"spoke7-go/pkg/logger"
	"spoke7-go/pkg/utils"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var CURRENT_TRAFFIC_BY_DETECTION_POINT_MICROSERVICE string = "data"
var CURRENT_TRAFFIC_BY_DETECTION_POINT_RESOURCE string = "current_traffic_data_detection_point"

type currentTrafficDataByDetectionPoint struct {
	service           services.CurrentTrafficDataByDetectionPointService
	fileUploadService services.UploadFileService
	logger            logger.Logger
	pb.UnimplementedCurrentTrafficDataByDetectionPointServiceServer
}

func NewCurrentTrafficDataByDetectionPointController(service services.CurrentTrafficDataByDetectionPointService, fileUploadService services.UploadFileService, logger logger.Logger) pb.CurrentTrafficDataByDetectionPointServiceServer {
	return &currentTrafficDataByDetectionPoint{service: service, fileUploadService: fileUploadService, logger: logger}
}

func (dc *currentTrafficDataByDetectionPoint) CreateCurrentTrafficDataByDetectionPoint(ctx context.Context, req *pb.CreateCurrentTrafficDataByDetectionPointRequest) (*pb.CreateCurrentTrafficDataByDetectionPointResponse, error) {

	currentTrafficModel, err := dtos.FromCurrentTrafficDataByDetectionPointProtoToModel(req.CurrentTrafficDataByDetectionPoint)
	if err != nil {
		return nil, err
	}
	if err := dc.service.Create(ctx, currentTrafficModel); err != nil {

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response, err := dtos.FromCurrentTrafficDataByDetectionPointModelToProto(currentTrafficModel)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCurrentTrafficDataByDetectionPointResponse{
		CurrentTrafficDataByDetectionPoint: response,
	}, nil
}

func (dc *currentTrafficDataByDetectionPoint) UpdateCurrentTrafficDataByDetectionPoint(ctx context.Context, req *pb.UpdateCurrentTrafficDataByDetectionPointRequest) (*pb.UpdateCurrentTrafficDataByDetectionPointResponse, error) {

	currentTrafficModel, err := dtos.FromCurrentTrafficDataByDetectionPointProtoToModel(req.CurrentTrafficDataByDetectionPoint)
	if err != nil {
		return nil, err
	}
	if err := dc.service.Update(ctx, currentTrafficModel); err != nil {

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response, err := dtos.FromCurrentTrafficDataByDetectionPointModelToProto(currentTrafficModel)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCurrentTrafficDataByDetectionPointResponse{
		CurrentTrafficDataByDetectionPoint: response,
	}, nil
}

func (dc *currentTrafficDataByDetectionPoint) DeleteCurrentTrafficDataByDetectionPoint(ctx context.Context, req *pb.DeleteTrafficDataByDetectionPointRequest) (*emptypb.Empty, error) {
	params := dtos.FromTrafficDataByDetectionPointDeleteParamsProtoToModel(req)
	if params.DataSourceName == "" {
		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(params.DetectionPointIDs) == 0 {
		return nil, fmt.Errorf("error: detection point id was not provided")
	}
	if err := dc.service.Delete(ctx, params); err != nil {

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (dc *currentTrafficDataByDetectionPoint) GetCurrentTrafficDataByDetectionPoint(ctx context.Context, req *pb.GetTrafficDataByDetectionPointRequest) (*pb.GetCurrentTrafficDataByDetectionPointsResponse, error) {
	params := dtos.FromTrafficDataByDetectionPointGetParamsProtoToModel(req)
	if params.DataSourceName == "" {
		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(params.DetectionPointIDs) == 0 {
		return nil, fmt.Errorf("error: detection point id was not provided")
	}
	trafficModels, err := dc.service.Get(ctx, params)
	if err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	result := make([]*pb.CurrentTrafficDataByDetectionPoint, 0)
	for _, trafficModel := range trafficModels {
		dto, err := dtos.FromCurrentTrafficDataByDetectionPointModelToProto(trafficModel)
		if err != nil {
			return nil, err
		}
		result = append(result, dto)
	}
	return &pb.GetCurrentTrafficDataByDetectionPointsResponse{
		CurrentTrafficDataByDetectionPoint: result,
		TotalCount:                         uint32(len(result)),
	}, nil
}

func (dc *currentTrafficDataByDetectionPoint) ListCurrentTrafficDataByDetectionPoints(ctx context.Context, req *pb.ListTrafficDataByDetectionPointRequest) (*pb.ListCurrentTrafficDataByDetectionPointsResponse, error) {
	params := dtos.FromTrafficDataByDetectionPointListParamsProtoToModel(req)
	if params.DataSourceName == "" {
		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(params.DetectionPointIDs) == 0 {
		return nil, fmt.Errorf("error: detection point id was not provided")
	}
	trafficModels, err := dc.service.List(ctx, params)
	if err != nil {

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	result := make([]*pb.CurrentTrafficDataByDetectionPoint, 0)
	for _, trafficModel := range trafficModels {
		dto, err := dtos.FromCurrentTrafficDataByDetectionPointModelToProto(trafficModel)
		if err != nil {
			return nil, err
		}
		result = append(result, dto)
	}
	return &pb.ListCurrentTrafficDataByDetectionPointsResponse{CurrentTrafficDataByDetectionPoints: result, TotalCount: uint32(len(result))}, nil
}

func (dc *currentTrafficDataByDetectionPoint) BulkCreateCurrentTrafficDataByDetectionPoint(ctx context.Context, req *pb.BulkCreateTrafficDataByDetectionPointRequest) (*pb.BulkCreateTrafficDataByDetectionPointResponse, error) {
	if req.File == nil {
		return nil, status.Error(codes.InvalidArgument, "Missing file")
	}
	if req.File.Content == nil {
		return nil, status.Error(codes.InvalidArgument, "Missing file content")
	}
	//handle only json
	if req.File.ContentType != "application/json" {
		return nil, status.Error(codes.InvalidArgument, "Unsupported file type")
	}

	trafficProto := []*pb.CurrentTrafficDataByDetectionPoint{}

	marshaler := &runtime.JSONPb{}
	err := marshaler.Unmarshal(req.File.Content, &trafficProto)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to parse JSON data")
	}

	trafficModels, err := dtos.FromCurrentTrafficDataByDetectionPointProtosToModels(trafficProto)

	dataSourceName := ""
	for _, trafficModel := range trafficModels {
		if dataSourceName == "" {
			dataSourceName = trafficModel.DataSourceName
		}
		if dataSourceName != trafficModel.DataSourceName {
			return &pb.BulkCreateTrafficDataByDetectionPointResponse{
				Message: fmt.Sprintf("all inserted models should have the same dataSourceName but got different dataSourceName: %s and %s", dataSourceName, trafficModel.DataSourceName),
			}, status.Error(codes.InvalidArgument, "all inserted models should have the same dataSourceName")
		}
	}
	if err := dc.service.BulkCreate(ctx, trafficModels); err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return &pb.BulkCreateTrafficDataByDetectionPointResponse{
				Message: fmt.Sprintf("%v", err),
			}, status.Error(codes.PermissionDenied, err.Error())
		}
		return &pb.BulkCreateTrafficDataByDetectionPointResponse{
			Message: fmt.Sprintf("%v", err),
		}, nil
	}

	uploadDescription := "automatic upload from BulkCreateCurrentTrafficDataByDetectionPoint"
	if err := dc.fileUploadService.UploadFile(ctx, dataSourceName, req.File.Filename, req.File.Content, uploadDescription, req.File.ContentType, models.FileTypeAggregatedTraffic5MinByDetectionPoint); err != nil {
		return &pb.BulkCreateTrafficDataByDetectionPointResponse{
			Message: fmt.Sprintf("%v", err),
		}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.BulkCreateTrafficDataByDetectionPointResponse{
		Message: "ok",
	}, nil

}

func (dc *currentTrafficDataByDetectionPoint) GetCurrentTrafficDataByDetectionPointStatistics(ctx context.Context, req *pb.GetTrafficDataByDetectionPointStatisticsRequest) (*pb.GetCurrentTrafficDataByDetectionPointStatisticsResponse, error) {
	if req.DataSourceName == "" {
		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	var from, to *time.Time
	if req.StartTimestamp != nil {
		fromTime := req.StartTimestamp.AsTime()
		from = &fromTime
	} else {
		fromTime := time.Now().AddDate(-5, 0, 0)
		from = &fromTime
	}

	if req.EndTimestamp != nil {
		toTime := req.EndTimestamp.AsTime()
		to = &toTime
	} else {
		toTime := time.Now().AddDate(5, 0, 0)
		to = &toTime
	}
	statModels, err := dc.service.Statistics(ctx, req.DataSourceName, req.DetectionPointIds, *from, *to)
	if err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}
	statProtos, err := dtos.FromCurrentTrafficDataByDetectionPointStatisticsModelsToProto(statModels.StatisticsByDetectionPoint)
	if err != nil {
		return nil, err
	}

	return &pb.GetCurrentTrafficDataByDetectionPointStatisticsResponse{
		StatisticsByDetectionPoint: statProtos,
		RecordsCount:               statModels.RecordsCount,
		DataSourceName:             req.DataSourceName,
		FirstRecordTimestamp:       timestamppb.New(statModels.FirstRecordTimestamp),
		LastRecordTimestamp:        timestamppb.New(statModels.LastRecordTimestamp),
	}, nil
}

func (dc *currentTrafficDataByDetectionPoint) ListCurrentTrafficDataByDetectionPointDaily(ctx context.Context, req *pb.ListTrafficDataByDetectionPointDailyRequest) (*pb.ListCurrentTrafficDataByDetectionPointDailyResponse, error) {
	if req.DataSourceName == "" {
		return nil, fmt.Errorf("datasource name must be provided")
	}

	var from, to *time.Time

	if req.From != nil {
		fromTime := req.From.AsTime()
		from = &fromTime
	}

	if req.To != nil {
		toTime := req.To.AsTime()
		to = &toTime
	}

	stats, err := dc.service.ListAggregatedByDay(ctx, req.DataSourceName, req.DetectionPointIds, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get daily stats: %w", err)
	}

	protoStats := make([]*pb.CurrentTrafficDataByDetectionPoint, 0, len(stats))
	for _, stat := range stats {
		protoStat, err := dtos.FromCurrentTrafficDataByDetectionPointModelToProto(stat)
		if err != nil {
			return nil, fmt.Errorf("failed to convert stat to proto: %w", err)
		}
		protoStats = append(protoStats, protoStat)
	}

	return &pb.ListCurrentTrafficDataByDetectionPointDailyResponse{
		DailyStats: protoStats,
	}, nil

}

func (dc *currentTrafficDataByDetectionPoint) DownloadCurrentTrafficDataByDetectionPoint(ctx context.Context, req *pb.DownloadTrafficDataByDetectionPointRequest) (*pb.DownloadResponse, error) {
	logText := "Request Download Current Traffic By Detection Point"
	additionalFields := []zap.Field{zap.String("detection_point_ids", strings.Join(req.DetectionPointIds, ",")), zap.String("start_timestamp", req.StartTime.String()), zap.String("end_timestamp", req.EndTime.String())}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

	params := dtos.FromTrafficDataByDetectionPointDownloadParamsProtoToModel(req)
	if req.DataSourceName == "" {
		errorText := fmt.Sprintf("%v: %v", logText, "error: datasource name was not provided")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, CURRENT_TRAFFIC_BY_DETECTION_POINT_MICROSERVICE, CURRENT_TRAFFIC_BY_DETECTION_POINT_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(req.DetectionPointIds) == 0 {
		return nil, fmt.Errorf("error: detection point id was not provided")
	}
	trafficModels, err := dc.service.List(ctx, params)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, CURRENT_TRAFFIC_BY_DETECTION_POINT_MICROSERVICE, CURRENT_TRAFFIC_BY_DETECTION_POINT_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	var trafficProtos []*pb.CurrentTrafficDataByDetectionPoint
	for _, trafficModel := range trafficModels {
		dto, err := dtos.FromCurrentTrafficDataByDetectionPointModelToProto(trafficModel)
		if err != nil {
			errorText := fmt.Sprintf("%v: %v", logText, err.Error())
			_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, CURRENT_TRAFFIC_BY_DETECTION_POINT_MICROSERVICE, CURRENT_TRAFFIC_BY_DETECTION_POINT_RESOURCE, "POST", req.DataSourceName, additionalFields...)
			return nil, err
		}
		trafficProtos = append(trafficProtos, dto)
	}

	marshaler := &runtime.JSONPb{}
	jsonData, err := marshaler.Marshal(&pb.GetCurrentTrafficDataByDetectionPointsResponse{
		CurrentTrafficDataByDetectionPoint: trafficProtos,
	})
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, CURRENT_TRAFFIC_BY_DETECTION_POINT_MICROSERVICE, CURRENT_TRAFFIC_BY_DETECTION_POINT_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		return nil, fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	logText = "Success Download Current Traffic By Detection Point"
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, CURRENT_TRAFFIC_BY_DETECTION_POINT_MICROSERVICE, CURRENT_TRAFFIC_BY_DETECTION_POINT_RESOURCE, "POST", req.DataSourceName, additionalFields...)

	return &pb.DownloadResponse{
		FileContent: jsonData,
		Filename:    "traffic_data_aggregated_5_min_by_detection_point.json",
		ContentType: "application/json",
	}, nil
}
