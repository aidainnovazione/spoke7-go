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

var REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE string = "data"
var REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_RESOURCE string = "real_time_traffic_data_detection_section"

type realTimeTrafficDataByDetectionSection struct {
	service           services.RealTimeTrafficDataByDetectionSectionService
	fileUploadService services.UploadFileService
	logger            logger.Logger
	pb.UnimplementedRealTimeTrafficDataByDetectionSectionServiceServer
}

func NewRealTimeTrafficDataByDetectionSectionController(service services.RealTimeTrafficDataByDetectionSectionService, fileUploadService services.UploadFileService, logger logger.Logger) pb.RealTimeTrafficDataByDetectionSectionServiceServer {
	return &realTimeTrafficDataByDetectionSection{service: service, fileUploadService: fileUploadService, logger: logger}
}

func (dc *realTimeTrafficDataByDetectionSection) CreateRealTimeTrafficDataByDetectionSection(ctx context.Context, req *pb.CreateRealTimeTrafficDataByDetectionSectionRequest) (*pb.CreateRealTimeTrafficDataByDetectionSectionResponse, error) {

	realTimeTrafficModel, err := dtos.FromRealTimeTrafficDataByDetectionSectionProtoToModel(req.RealTimeTrafficDataByDetectionSection)
	if err != nil {
		return nil, err
	}
	if err := dc.service.Create(ctx, realTimeTrafficModel); err != nil {

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response, err := dtos.FromRealTimeTrafficDataByDetectionSectionModelToProto(realTimeTrafficModel)
	if err != nil {
		return nil, err
	}
	return &pb.CreateRealTimeTrafficDataByDetectionSectionResponse{
		RealTimeTrafficDataByDetectionSection: response,
	}, nil
}

func (dc *realTimeTrafficDataByDetectionSection) UpdateRealTimeTrafficDataByDetectionSection(ctx context.Context, req *pb.UpdateRealTimeTrafficDataByDetectionSectionRequest) (*pb.UpdateRealTimeTrafficDataByDetectionSectionResponse, error) {

	realTimeTrafficModel, err := dtos.FromRealTimeTrafficDataByDetectionSectionProtoToModel(req.RealTimeTrafficDataByDetectionSection)
	if err != nil {
		return nil, err
	}
	if err := dc.service.Update(ctx, realTimeTrafficModel); err != nil {

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response, err := dtos.FromRealTimeTrafficDataByDetectionSectionModelToProto(realTimeTrafficModel)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateRealTimeTrafficDataByDetectionSectionResponse{
		RealTimeTrafficDataByDetectionSection: response,
	}, nil
}

func (dc *realTimeTrafficDataByDetectionSection) DeleteRealTimeTrafficDataByDetectionSection(ctx context.Context, req *pb.DeleteTrafficDataByDetectionSectionRequest) (*emptypb.Empty, error) {
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

func (dc *realTimeTrafficDataByDetectionSection) GetRealTimeTrafficDataByDetectionSection(ctx context.Context, req *pb.GetTrafficDataByDetectionSectionRequest) (*pb.GetRealTimeTrafficDataByDetectionSectionsResponse, error) {
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

	result := make([]*pb.RealTimeTrafficDataByDetectionSection, 0)
	for _, trafficModel := range trafficModels {
		dto, err := dtos.FromRealTimeTrafficDataByDetectionSectionModelToProto(trafficModel)
		if err != nil {
			return nil, err
		}
		result = append(result, dto)
	}
	return &pb.GetRealTimeTrafficDataByDetectionSectionsResponse{
		RealTimeTrafficDataByDetectionSection: result,
		TotalCount:                            uint32(len(result)),
	}, nil
}

func (dc *realTimeTrafficDataByDetectionSection) ListRealTimeTrafficDataByDetectionSections(ctx context.Context, req *pb.ListTrafficDataByDetectionSectionRequest) (*pb.ListRealTimeTrafficDataByDetectionSectionsResponse, error) {
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

	result := make([]*pb.RealTimeTrafficDataByDetectionSection, 0)
	for _, trafficModel := range trafficModels {
		dto, err := dtos.FromRealTimeTrafficDataByDetectionSectionModelToProto(trafficModel)
		if err != nil {
			return nil, err
		}
		result = append(result, dto)
	}
	return &pb.ListRealTimeTrafficDataByDetectionSectionsResponse{RealTimeTrafficDataByDetectionSection: result, TotalCount: uint32(len(result))}, nil
}

func (dc *realTimeTrafficDataByDetectionSection) BulkCreateRealTimeTrafficDataByDetectionSection(ctx context.Context, req *pb.BulkCreateTrafficDataByDetectionSectionRequest) (*pb.BulkCreateTrafficDataByDetectionSectionResponse, error) {
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

	trafficProto := []*pb.RealTimeTrafficDataByDetectionSection{}

	marshaler := &runtime.JSONPb{}
	err := marshaler.Unmarshal(req.File.Content, &trafficProto)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to parse JSON data")
	}

	trafficModels, err := dtos.FromRealTimeTrafficDataByDetectionSectionProtosToModels(trafficProto)

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

	uploadDescription := "automatic upload from BulkCreateRealTimeTrafficDataByDetectionSection"
	if err := dc.fileUploadService.UploadFile(ctx, dataSourceName, req.File.Filename, req.File.Content, uploadDescription, req.File.ContentType, models.FileTypeRealTimeTrafficByDetectionSection); err != nil {
		return &pb.BulkCreateTrafficDataByDetectionSectionResponse{
			Message: fmt.Sprintf("%v", err),
		}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.BulkCreateTrafficDataByDetectionSectionResponse{
		Message: "ok",
	}, nil

}

func (dc *realTimeTrafficDataByDetectionSection) GetRealTimeTrafficDataByDetectionSectionStatistics(ctx context.Context, req *pb.GetTrafficDataByDetectionSectionStatisticsRequest) (*pb.GetTrafficDataByDetectionSectionStatisticsResponse, error) {
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
	statProtos, err := dtos.FromTrafficDataByDetectionSectionStatisticsModelsToProto(statModels.StatisticsBySection)
	if err != nil {
		return nil, err
	}

	return &pb.GetTrafficDataByDetectionSectionStatisticsResponse{
		Statistics:           statProtos,
		RecordsCount:         statModels.RecordsCount,
		DataSourceName:       req.DataSourceName,
		FirstRecordTimestamp: timestamppb.New(statModels.FirstRecordTimestamp),
		LastRecordTimestamp:  timestamppb.New(statModels.LastRecordTimestamp),
	}, nil
}

func (dc *realTimeTrafficDataByDetectionSection) ListRealTimeTrafficDataByDetectionSectionDaily(ctx context.Context, req *pb.ListTrafficDataByDetectionSectionDailyRequest) (*pb.ListRealTimeTrafficDataByDetectionSectionDailyResponse, error) {
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

	protoStats := make([]*pb.RealTimeTrafficDataByDetectionSection, 0, len(stats))
	for _, stat := range stats {
		protoStat, err := dtos.FromRealTimeTrafficDataByDetectionSectionModelToProto(stat)
		if err != nil {
			return nil, fmt.Errorf("failed to convert stat to proto: %w", err)
		}
		protoStats = append(protoStats, protoStat)
	}

	return &pb.ListRealTimeTrafficDataByDetectionSectionDailyResponse{
		DailyStats: protoStats,
	}, nil
}

func (dc *realTimeTrafficDataByDetectionSection) DownloadRealTimeTrafficDataByDetectionSection(ctx context.Context, req *pb.DownloadTrafficDataByDetectionSectionRequest) (*pb.DownloadResponse, error) {
	logText := "Request Download Real Time Traffic By Detection Section"
	additionalFields := []zap.Field{zap.String("detection_section_ids", strings.Join(req.DetectionSectionIds, ",")), zap.String("start_timestamp", req.StartTime.String()), zap.String("end_timestamp", req.EndTime.String())}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

	params := dtos.FromTrafficDataByDetectionSectionDownloadParamsProtoToModel(req)
	if req.DataSourceName == "" {
		errorText := fmt.Sprintf("%v: %v", logText, "error: datasource name was not provided")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(req.DetectionSectionIds) == 0 {
		errorText := fmt.Sprintf("%v: %v", logText, "error: lane id was not provided")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		return nil, fmt.Errorf("error: detection section id was not provided")
	}
	trafficModels, err := dc.service.List(ctx, params)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	var trafficProtos []*pb.RealTimeTrafficDataByDetectionSection
	for _, trafficModel := range trafficModels {
		dto, err := dtos.FromRealTimeTrafficDataByDetectionSectionModelToProto(trafficModel)
		if err != nil {
			errorText := fmt.Sprintf("%v: %v", logText, err.Error())
			_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

			return nil, err
		}
		trafficProtos = append(trafficProtos, dto)
	}

	marshaler := &runtime.JSONPb{}
	jsonData, err := marshaler.Marshal(&pb.GetRealTimeTrafficDataByDetectionSectionsResponse{
		RealTimeTrafficDataByDetectionSection: trafficProtos,
	})
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		return nil, fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	logText = "Success Download Real Time Traffic By Detection Section"
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_MICROSERVICE, REAL_TIME_TRAFFIC_BY_DETECTION_SECTION_RESOURCE, "POST", req.DataSourceName, additionalFields...)

	return &pb.DownloadResponse{
		FileContent: jsonData,
		Filename:    "traffic_data_aggregated_real_time_by_detection_section.json",
		ContentType: "application/json",
	}, nil
}
