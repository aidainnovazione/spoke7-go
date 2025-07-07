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

var REAL_TIME_TRAFFIC_BY_LANE_MICROSERVICE string = "data"
var REAL_TIME_TRAFFIC_BY_LANE_RESOURCE string = "real_time_traffic_data_detection_point_by_lane"

type realTimeTrafficDataByDetectionPointByLane struct {
	service           services.RealTimeTrafficDataByDetectionPointByLaneService
	fileUploadService services.UploadFileService
	logger            logger.Logger
	pb.UnimplementedRealTimeTrafficDataByDetectionPointByLaneServiceServer
}

func NewRealTimeTrafficDataByDetectionPointByLaneController(service services.RealTimeTrafficDataByDetectionPointByLaneService, fileUploadService services.UploadFileService, logger logger.Logger) pb.RealTimeTrafficDataByDetectionPointByLaneServiceServer {
	return &realTimeTrafficDataByDetectionPointByLane{service: service, fileUploadService: fileUploadService, logger: logger}
}

func (dc *realTimeTrafficDataByDetectionPointByLane) CreateRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, req *pb.CreateRealTimeTrafficDataByDetectionPointByLaneRequest) (*pb.CreateRealTimeTrafficDataByDetectionPointByLaneResponse, error) {

	realTimeTrafficModel, err := dtos.FromRealTimeTrafficDataByDetectionPointByLaneProtoToModel(req.RealTimeTrafficDataByDetectionPointByLane)
	if err != nil {
		return nil, err
	}
	if err := dc.service.Create(ctx, realTimeTrafficModel); err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response, err := dtos.FromRealTimeTrafficDataByDetectionPointByLaneModelToProto(realTimeTrafficModel)
	if err != nil {
		return nil, err
	}
	return &pb.CreateRealTimeTrafficDataByDetectionPointByLaneResponse{
		RealTimeTrafficDataByDetectionPointByLane: response,
	}, nil
}

func (dc *realTimeTrafficDataByDetectionPointByLane) UpdateRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, req *pb.UpdateRealTimeTrafficDataByDetectionPointByLaneRequest) (*pb.UpdateRealTimeTrafficDataByDetectionPointByLaneResponse, error) {

	realTimeTrafficModel, err := dtos.FromRealTimeTrafficDataByDetectionPointByLaneProtoToModel(req.RealTimeTrafficDataByDetectionPointByLane)
	if err != nil {
		return nil, err
	}
	if err := dc.service.Update(ctx, realTimeTrafficModel); err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response, err := dtos.FromRealTimeTrafficDataByDetectionPointByLaneModelToProto(realTimeTrafficModel)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateRealTimeTrafficDataByDetectionPointByLaneResponse{
		RealTimeTrafficDataByDetectionPointByLane: response,
	}, nil
}

func (dc *realTimeTrafficDataByDetectionPointByLane) DeleteRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, req *pb.DeleteTrafficDataByDetectionPointByLaneRequest) (*emptypb.Empty, error) {
	params := dtos.FromTrafficDataByDetectionPointByLaneDeleteParamsProtoToModel(req)
	if params.DataSourceName == "" {
		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(params.LaneIDs) == 0 {
		return nil, fmt.Errorf("error: lane id was not provided")
	}
	if err := dc.service.Delete(ctx, params); err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (dc *realTimeTrafficDataByDetectionPointByLane) GetRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, req *pb.GetTrafficDataByDetectionPointByLaneRequest) (*pb.GetRealTimeTrafficDataByDetectionPointByLaneResponse, error) {
	params := dtos.FromTrafficDataByDetectionPointByLaneGetParamsProtoToModel(req)
	if params.DataSourceName == "" {
		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(params.LaneIDs) == 0 {
		return nil, fmt.Errorf("error: lane id was not provided")
	}
	trafficModels, err := dc.service.Get(ctx, params)
	if err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}
	result := make([]*pb.RealTimeTrafficDataByDetectionPointByLane, 0)
	for _, trafficModel := range trafficModels {
		dto, err := dtos.FromRealTimeTrafficDataByDetectionPointByLaneModelToProto(trafficModel)
		if err != nil {
			return nil, err
		}
		result = append(result, dto)
	}

	return &pb.GetRealTimeTrafficDataByDetectionPointByLaneResponse{
		RealTimeTrafficDataByDetectionPointByLane: result,
		TotalCount: uint32(len(result)),
	}, nil
}

func (dc *realTimeTrafficDataByDetectionPointByLane) ListRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, req *pb.ListTrafficDataByDetectionPointByLaneRequest) (*pb.ListRealTimeTrafficDataByDetectionPointByLaneResponse, error) {
	params := dtos.FromTrafficDataByDetectionPointByLaneListParamsProtoToModel(req)
	if params.DataSourceName == "" {
		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(params.LaneIDs) == 0 {
		return nil, fmt.Errorf("error: lane id was not provided")
	}
	trafficModels, err := dc.service.List(ctx, params)
	if err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	result := make([]*pb.RealTimeTrafficDataByDetectionPointByLane, 0)
	for _, trafficModel := range trafficModels {
		dto, err := dtos.FromRealTimeTrafficDataByDetectionPointByLaneModelToProto(trafficModel)
		if err != nil {
			return nil, err
		}
		result = append(result, dto)
	}
	return &pb.ListRealTimeTrafficDataByDetectionPointByLaneResponse{RealTimeTrafficDataByDetectionPointByLane: result, TotalCount: uint32(len(result))}, nil
}

func (dc *realTimeTrafficDataByDetectionPointByLane) BulkCreateRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, req *pb.BulkCreateTrafficDataByDetectionPointByLaneRequest) (*pb.BulkCreateTrafficDataByDetectionPointByLaneResponse, error) {
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

	trafficProto := []*pb.RealTimeTrafficDataByDetectionPointByLane{}

	marshaler := &runtime.JSONPb{}
	err := marshaler.Unmarshal(req.File.Content, &trafficProto)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to parse JSON data")
	}

	trafficModels, err := dtos.FromRealTimeTrafficDataByDetectionPointByLaneProtosToModels(trafficProto)

	dataSourceName := ""
	for _, trafficModel := range trafficModels {
		if dataSourceName == "" {
			dataSourceName = trafficModel.DataSourceName
		}
		if dataSourceName != trafficModel.DataSourceName {
			return &pb.BulkCreateTrafficDataByDetectionPointByLaneResponse{
				Message: fmt.Sprintf("all inserted models should have the same dataSourceName but got different dataSourceName: %s and %s", dataSourceName, trafficModel.DataSourceName),
			}, status.Error(codes.InvalidArgument, "all inserted models should have the same dataSourceName")
		}
	}
	if err := dc.service.BulkCreate(ctx, trafficModels); err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return &pb.BulkCreateTrafficDataByDetectionPointByLaneResponse{
				Message: fmt.Sprintf("%v", err),
			}, status.Error(codes.PermissionDenied, err.Error())
		}
		return &pb.BulkCreateTrafficDataByDetectionPointByLaneResponse{
			Message: fmt.Sprintf("%v", err),
		}, nil
	}

	uploadDescription := "automatic upload from BulkCreateRealTimeTrafficDataByDetectionPointByLane"
	if err := dc.fileUploadService.UploadFile(ctx, dataSourceName, req.File.Filename, req.File.Content, uploadDescription, req.File.ContentType, models.FileTypeRealTimeTrafficByLane); err != nil {
		return &pb.BulkCreateTrafficDataByDetectionPointByLaneResponse{
			Message: fmt.Sprintf("%v", err),
		}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.BulkCreateTrafficDataByDetectionPointByLaneResponse{
		Message: "ok",
	}, nil

}

func (dc *realTimeTrafficDataByDetectionPointByLane) GetRealTimeTrafficDataByDetectionPointByLaneStatistics(ctx context.Context, req *pb.GetTrafficDataByDetectionPointByLaneStatisticsRequest) (*pb.GetTrafficDataByDetectionPointByLaneStatisticsResponse, error) {
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
	statModels, err := dc.service.Statistics(ctx, req.DataSourceName, req.LaneIds, *from, *to)
	if err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}
	statProtos, err := dtos.FromTrafficDataByDetectionPointByLaneStatisticsModelsToProto(statModels.StatisticsByLane)
	if err != nil {
		return nil, err
	}

	return &pb.GetTrafficDataByDetectionPointByLaneStatisticsResponse{
		Statistics:           statProtos,
		RecordsCount:         statModels.RecordsCount,
		DataSourceName:       req.DataSourceName,
		FirstRecordTimestamp: timestamppb.New(statModels.FirstRecordTimestamp),
		LastRecordTimestamp:  timestamppb.New(statModels.LastRecordTimestamp),
	}, nil
}

func (dc *realTimeTrafficDataByDetectionPointByLane) ListRealTimeTrafficDataByDetectionPointByLaneDaily(ctx context.Context, req *pb.ListTrafficDataByDetectionPointByLaneDailyRequest) (*pb.ListRealTimeTrafficDataByDetectionPointByLaneDailyResponse, error) {
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

	stats, err := dc.service.ListAggregatedByDay(ctx, req.DataSourceName, req.LaneIds, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get daily stats: %w", err)
	}

	protoStats := make([]*pb.RealTimeTrafficDataByDetectionPointByLane, 0, len(stats))
	for _, stat := range stats {
		protoStat, err := dtos.FromRealTimeTrafficDataByDetectionPointByLaneModelToProto(stat)
		if err != nil {
			return nil, fmt.Errorf("failed to convert stat to proto: %w", err)
		}
		protoStats = append(protoStats, protoStat)
	}

	return &pb.ListRealTimeTrafficDataByDetectionPointByLaneDailyResponse{
		DailyStats: protoStats,
	}, nil

}

func (dc *realTimeTrafficDataByDetectionPointByLane) DownloadRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, req *pb.DownloadTrafficDataByDetectionPointByLaneRequest) (*pb.DownloadResponse, error) {
	logText := "Request Download Real Time Traffic By Detection Point By Lane"
	additionalFields := []zap.Field{zap.String("lanes_ids", strings.Join(req.LaneIds, ",")), zap.String("start_timestamp", req.StartTime.String()), zap.String("end_timestamp", req.EndTime.String())}
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, REAL_TIME_TRAFFIC_BY_LANE_MICROSERVICE, REAL_TIME_TRAFFIC_BY_LANE_RESOURCE, "POST", req.DataSourceName, additionalFields...)

	params := dtos.FromTrafficDataByDetectionPointByLaneDownloadParamsProtoToModel(req)
	if req.DataSourceName == "" {
		errorText := fmt.Sprintf("%v: %v", logText, "error: datasource name was not provided")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, REAL_TIME_TRAFFIC_BY_LANE_MICROSERVICE, REAL_TIME_TRAFFIC_BY_LANE_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		return nil, fmt.Errorf("error: datasource name was not provided")
	}
	if len(req.LaneIds) == 0 {
		errorText := fmt.Sprintf("%v: %v", logText, "error: lane id was not provided")
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, REAL_TIME_TRAFFIC_BY_LANE_MICROSERVICE, REAL_TIME_TRAFFIC_BY_LANE_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		return nil, fmt.Errorf("error: lane id was not provided")
	}
	trafficModels, err := dc.service.List(ctx, params)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, REAL_TIME_TRAFFIC_BY_LANE_MICROSERVICE, REAL_TIME_TRAFFIC_BY_LANE_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	var trafficProtos []*pb.RealTimeTrafficDataByDetectionPointByLane
	for _, trafficModel := range trafficModels {
		dto, err := dtos.FromRealTimeTrafficDataByDetectionPointByLaneModelToProto(trafficModel)
		if err != nil {
			errorText := fmt.Sprintf("%v: %v", logText, err.Error())
			_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, REAL_TIME_TRAFFIC_BY_LANE_MICROSERVICE, REAL_TIME_TRAFFIC_BY_LANE_RESOURCE, "POST", req.DataSourceName, additionalFields...)
			return nil, err
		}
		trafficProtos = append(trafficProtos, dto)
	}

	marshaler := &runtime.JSONPb{}
	jsonData, err := marshaler.Marshal(&pb.GetRealTimeTrafficDataByDetectionPointByLaneResponse{
		RealTimeTrafficDataByDetectionPointByLane: trafficProtos,
	})
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(dc.logger, ctx, errorText, REAL_TIME_TRAFFIC_BY_LANE_MICROSERVICE, REAL_TIME_TRAFFIC_BY_LANE_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		return nil, fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	logText = "Success Download Real Time Traffic By Detection Point By Lane"
	_ = utils.AddControllerRequestLogging(dc.logger, ctx, logText, REAL_TIME_TRAFFIC_BY_LANE_MICROSERVICE, REAL_TIME_TRAFFIC_BY_LANE_RESOURCE, "POST", req.DataSourceName, additionalFields...)

	return &pb.DownloadResponse{
		FileContent: jsonData,
		Filename:    "traffic_data_aggregated_real_time_by_lane.json",
		ContentType: "application/json",
	}, nil
}
