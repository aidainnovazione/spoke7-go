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

var CURRENT_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE string = "data"
var CURRENT_TRAFFIC_BY_DETECTION_SECTION_RESOURCE string = "current_traffic_data_detection_section"

type currentTrafficDataByDetectionSection struct {
	service           services.CurrentTrafficDataByDetectionSectionService
	fileUploadService services.UploadFileService
	logger            logger.Logger
	pb.UnimplementedCurrentTrafficDataByDetectionSectionServiceServer
}

func NewCurrentTrafficDataByDetectionSectionController(service services.CurrentTrafficDataByDetectionSectionService, fileUploadService services.UploadFileService, logger logger.Logger) pb.CurrentTrafficDataByDetectionSectionServiceServer {
	return &currentTrafficDataByDetectionSection{service: service, fileUploadService: fileUploadService, logger: logger}
}

func (dc *currentTrafficDataByDetectionSection) CreateCurrentTrafficDataByDetectionSection(ctx context.Context, req *pb.CreateCurrentTrafficDataByDetectionSectionRequest) (*pb.CreateCurrentTrafficDataByDetectionSectionResponse, error) {

	currentTrafficModel, err := dtos.FromCurrentTrafficDataByDetectionSectionProtoToModel(req.CurrentTrafficDataByDetectionSection)
	if err != nil {
		return nil, err
	}
	if err := dc.service.Create(ctx, currentTrafficModel); err != nil {

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response, err := dtos.FromCurrentTrafficDataByDetectionSectionModelToProto(currentTrafficModel)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCurrentTrafficDataByDetectionSectionResponse{
		CurrentTrafficDataByDetectionSection: response,
	}, nil
}

func (dc *currentTrafficDataByDetectionSection) UpdateCurrentTrafficDataByDetectionSection(ctx context.Context, req *pb.UpdateCurrentTrafficDataByDetectionSectionRequest) (*pb.UpdateCurrentTrafficDataByDetectionSectionResponse, error) {

	currentTrafficModel, err := dtos.FromCurrentTrafficDataByDetectionSectionProtoToModel(req.CurrentTrafficDataByDetectionSection)
	if err != nil {
		return nil, err
	}
	if err := dc.service.Update(ctx, currentTrafficModel); err != nil {

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response, err := dtos.FromCurrentTrafficDataByDetectionSectionModelToProto(currentTrafficModel)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCurrentTrafficDataByDetectionSectionResponse{
		CurrentTrafficDataByDetectionSection: response,
	}, nil
}

func (dc *currentTrafficDataByDetectionSection) DeleteCurrentTrafficDataByDetectionSection(ctx context.Context, req *pb.DeleteTrafficDataByDetectionSectionRequest) (*emptypb.Empty, error) {
	params := dtos.FromTrafficDataByDetectionSectionDeleteParamsProtoToModel(req)
	if params.DataSourceName == "" {
		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(params.DetectionSectionIDs) == 0 {
		return nil, fmt.Errorf("error: detection section id was not provided")
	}
	if err := dc.service.Delete(ctx, params); err != nil {

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (dc *currentTrafficDataByDetectionSection) GetCurrentTrafficDataByDetectionSection(ctx context.Context, req *pb.GetTrafficDataByDetectionSectionRequest) (*pb.GetCurrentTrafficDataByDetectionSectionsResponse, error) {
	params := dtos.FromTrafficDataByDetectionSectionGetParamsProtoToModel(req)
	if params.DataSourceName == "" {
		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(params.DetectionSectionIDs) == 0 {
		return nil, fmt.Errorf("error: detection section id was not provided")
	}
	trafficModels, err := dc.service.Get(ctx, params)
	if err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	result := make([]*pb.CurrentTrafficDataByDetectionSection, 0)
	for _, trafficModel := range trafficModels {
		dto, err := dtos.FromCurrentTrafficDataByDetectionSectionModelToProto(trafficModel)
		if err != nil {
			return nil, err
		}
		result = append(result, dto)
	}
	return &pb.GetCurrentTrafficDataByDetectionSectionsResponse{
		CurrentTrafficDataByDetectionSection: result,
		TotalCount:                           uint32(len(result)),
	}, nil
}

func (dc *currentTrafficDataByDetectionSection) ListCurrentTrafficDataByDetectionSections(ctx context.Context, req *pb.ListTrafficDataByDetectionSectionRequest) (*pb.ListCurrentTrafficDataByDetectionSectionsResponse, error) {
	params := dtos.FromTrafficDataByDetectionSectionListParamsProtoToModel(req)
	if params.DataSourceName == "" {
		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(params.DetectionSectionIDs) == 0 {
		return nil, fmt.Errorf("error: detection section id was not provided")
	}
	trafficModels, err := dc.service.List(ctx, params)
	if err != nil {

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	result := make([]*pb.CurrentTrafficDataByDetectionSection, 0)
	for _, trafficModel := range trafficModels {
		dto, err := dtos.FromCurrentTrafficDataByDetectionSectionModelToProto(trafficModel)
		if err != nil {
			return nil, err
		}
		result = append(result, dto)
	}
	return &pb.ListCurrentTrafficDataByDetectionSectionsResponse{CurrentTrafficDataByDetectionSection: result, TotalCount: uint32(len(result))}, nil
}

func (dc *currentTrafficDataByDetectionSection) BulkCreateCurrentTrafficDataByDetectionSection(ctx context.Context, req *pb.BulkCreateTrafficDataByDetectionSectionRequest) (*pb.BulkCreateTrafficDataByDetectionSectionResponse, error) {
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

	trafficProto := []*pb.CurrentTrafficDataByDetectionSection{}

	marshaler := &runtime.JSONPb{}
	err := marshaler.Unmarshal(req.File.Content, &trafficProto)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to parse JSON data")
	}

	trafficModels, err := dtos.FromCurrentTrafficDataByDetectionSectionProtosToModels(trafficProto)

	dataSourceName := ""
	for _, trafficModel := range trafficModels {
		if dataSourceName == "" {
			dataSourceName = trafficModel.DataSourceName
		}
		if dataSourceName != trafficModel.DataSourceName {
			return &pb.BulkCreateTrafficDataByDetectionSectionResponse{
				Message: fmt.Sprintf("all inserted models should have the same dataSourceName but got different dataSourceName: %s and %s", dataSourceName, trafficModel.DataSourceName),
			}, status.Error(codes.InvalidArgument, "all inserted models should have the same dataSourceName")
		}
	}
	if err := dc.service.BulkCreate(ctx, trafficModels); err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return &pb.BulkCreateTrafficDataByDetectionSectionResponse{
				Message: fmt.Sprintf("%v", err),
			}, status.Error(codes.PermissionDenied, err.Error())
		}
		return &pb.BulkCreateTrafficDataByDetectionSectionResponse{
			Message: fmt.Sprintf("%v", err),
		}, nil
	}

	uploadDescription := "automatic upload from BulkCreateCurrentTrafficDataByDetectionSection"
	if err := dc.fileUploadService.UploadFile(ctx, dataSourceName, req.File.Filename, req.File.Content, uploadDescription, req.File.ContentType, models.FileTypeAggregatedTraffic5MinByDetectionSection); err != nil {
		return &pb.BulkCreateTrafficDataByDetectionSectionResponse{
			Message: fmt.Sprintf("%v", err),
		}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.BulkCreateTrafficDataByDetectionSectionResponse{
		Message: "ok",
	}, nil

}

func (dc *currentTrafficDataByDetectionSection) GetCurrentTrafficDataByDetectionSectionStatistics(ctx context.Context, req *pb.GetTrafficDataByDetectionSectionStatisticsRequest) (*pb.GetCurrentTrafficDataByDetectionSectionStatisticsResponse, error) {
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
	statModels, err := dc.service.Statistics(ctx, req.DataSourceName, req.DetectionSectionIds, *from, *to)
	if err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}
	statProtos, err := dtos.FromCurrentTrafficDataByDetectionSectionStatisticsModelsToProto(statModels.StatisticsBySection)
	if err != nil {
		return nil, err
	}

	return &pb.GetCurrentTrafficDataByDetectionSectionStatisticsResponse{
		StatisticsByDetectionSection: statProtos,
		RecordsCount:                 statModels.RecordsCount,
		DataSourceName:               req.DataSourceName,
		FirstRecordTimestamp:         timestamppb.New(statModels.FirstRecordTimestamp),
		LastRecordTimestamp:          timestamppb.New(statModels.LastRecordTimestamp),
	}, nil
}

func (dc *currentTrafficDataByDetectionSection) ListCurrentTrafficDataByDetectionSectionDaily(ctx context.Context, req *pb.ListTrafficDataByDetectionSectionDailyRequest) (*pb.ListCurrentTrafficDataByDetectionSectionDailyResponse, error) {
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

	stats, err := dc.service.ListAggregatedByDay(ctx, req.DataSourceName, req.DetectionSectionIds, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get daily stats: %w", err)
	}

	protoStats := make([]*pb.CurrentTrafficDataByDetectionSection, 0, len(stats))
	for _, stat := range stats {
		protoStat, err := dtos.FromCurrentTrafficDataByDetectionSectionModelToProto(stat)
		if err != nil {
			return nil, fmt.Errorf("failed to convert stat to proto: %w", err)
		}
		protoStats = append(protoStats, protoStat)
	}

	return &pb.ListCurrentTrafficDataByDetectionSectionDailyResponse{
		DailyStats: protoStats,
	}, nil
}

func (dc *currentTrafficDataByDetectionSection) DownloadCurrentTrafficDataByDetectionSection(ctx context.Context, req *pb.DownloadTrafficDataByDetectionSectionRequest) (*pb.DownloadResponse, error) {
	logText := "Request Download Current Traffic By Detection Section"
	additionalFields := []zap.Field{zap.String("detection_section_ids", strings.Join(req.DetectionSectionIds, ",")), zap.String("start_timestamp", req.StartTime.String()), zap.String("end_timestamp", req.EndTime.String())}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, CURRENT_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, CURRENT_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

	params := dtos.FromTrafficDataByDetectionSectionDownloadParamsProtoToModel(req)
	if req.DataSourceName == "" {
		errorText := fmt.Sprintf("%v: %v", logText, "error: datasource name was not provided")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, CURRENT_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, CURRENT_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(req.DetectionSectionIds) == 0 {
		return nil, fmt.Errorf("error: detection section id was not provided")
	}
	trafficModels, err := dc.service.List(ctx, params)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, CURRENT_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, CURRENT_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	var trafficProtos []*pb.CurrentTrafficDataByDetectionSection
	for _, trafficModel := range trafficModels {
		dto, err := dtos.FromCurrentTrafficDataByDetectionSectionModelToProto(trafficModel)
		if err != nil {
			errorText := fmt.Sprintf("%v: %v", logText, err.Error())
			_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, CURRENT_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, CURRENT_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)
			return nil, err
		}
		trafficProtos = append(trafficProtos, dto)
	}

	marshaler := &runtime.JSONPb{}
	jsonData, err := marshaler.Marshal(&pb.GetCurrentTrafficDataByDetectionSectionsResponse{
		CurrentTrafficDataByDetectionSection: trafficProtos,
	})
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, CURRENT_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, CURRENT_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		return nil, fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	logText = "Success Download Current Traffic By Detection Section"
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, CURRENT_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, CURRENT_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

	return &pb.DownloadResponse{
		FileContent: jsonData,
		Filename:    "traffic_data_aggregated_5_min_by_detection_section.json",
		ContentType: "application/json",
	}, nil
}
