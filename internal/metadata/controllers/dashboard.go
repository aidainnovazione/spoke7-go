package controllers

import (
	"context"
	"errors"
	"fmt"
	"spoke7-go/internal/metadata/dtos"
	"spoke7-go/internal/metadata/pb"
	"spoke7-go/internal/metadata/services"
	"spoke7-go/pkg/logger"
	"spoke7-go/pkg/utils"

	errorsInternal "spoke7-go/internal/errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var DASHBOARD_MICROSERVICE string = "metadata"
var DASHBOARD_RESOURCE string = "dashboard"

type dashboardHttpController struct {
	service services.DashboardService
	logger  logger.Logger
	pb.UnimplementedDashboardServiceServer
}

func NewDashboardController(service services.DashboardService, logger logger.Logger) pb.DashboardServiceServer {
	return &dashboardHttpController{service: service, logger: logger}
}

func (sc *dashboardHttpController) List(ctx context.Context, req *pb.DashboardListParams) (*pb.DashboardListResponse, error) {
	logText := "Request List Dashboard"
	_ = utils.AddControllerRequestLogging(sc.logger, ctx, logText, DASHBOARD_MICROSERVICE, DASHBOARD_RESOURCE, "GET", req.DataSourceName)

	list, err := sc.service.List(ctx, req.DataSourceName)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(sc.logger, ctx, errorText, DASHBOARD_MICROSERVICE, DASHBOARD_RESOURCE, "GET", req.DataSourceName)

		return nil, err
	}

	response := make([]*pb.Dashboard, 0)
	for _, dashboard := range list {
		dto := dtos.NewDashboardProtoFromModel(dashboard)
		response = append(response, &dto)
	}

	return &pb.DashboardListResponse{Dashboard: response}, nil
}

func (sc *dashboardHttpController) Get(ctx context.Context, req *pb.DashboardGetRequest) (*pb.Dashboard, error) {
	logText := "Request Get Dashboard"
	additionalFields := []zap.Field{zap.String("dashboard_id", req.Id)}
	_ = utils.AddControllerRequestLogging(sc.logger, ctx, logText, DASHBOARD_MICROSERVICE, DASHBOARD_RESOURCE, "GET", "NONE", additionalFields...)

	dashboard, err := sc.service.Get(ctx, req.Id)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(sc.logger, ctx, errorText, DASHBOARD_MICROSERVICE, DASHBOARD_RESOURCE, "GET", "NONE", additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response := dtos.NewDashboardProtoFromModel(dashboard)

	return &response, nil
}

func (sc *dashboardHttpController) Create(ctx context.Context, req *pb.Dashboard) (*pb.Dashboard, error) {
	logText := "Request Create Dashboard"
	additionalFields := []zap.Field{zap.String("dashboard_id", req.Id), zap.String("dashboard_name", req.Name)}
	_ = utils.AddControllerRequestLogging(sc.logger, ctx, logText, DASHBOARD_MICROSERVICE, DASHBOARD_RESOURCE, "POST", req.DataSourceName, additionalFields...)

	dashboard := dtos.DashboardProtoToModel(req)
	if err := sc.service.Create(ctx, &dashboard); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(sc.logger, ctx, errorText, DASHBOARD_MICROSERVICE, DASHBOARD_RESOURCE, "POST", req.DataSourceName, additionalFields...)

		return nil, err
	}

	response := dtos.NewDashboardProtoFromModel(&dashboard)

	return &response, nil
}

func (sc *dashboardHttpController) Update(ctx context.Context, req *pb.Dashboard) (*pb.Dashboard, error) {
	logText := "Request Update Dashboard"
	additionalFields := []zap.Field{zap.String("dashboard_id", req.Id), zap.String("dashboard_name", req.Name)}
	_ = utils.AddControllerRequestLogging(sc.logger, ctx, logText, DASHBOARD_MICROSERVICE, DASHBOARD_RESOURCE, "PUT", req.DataSourceName, additionalFields...)

	dashboard := dtos.DashboardProtoToModel(req)
	dashboardUpdated, err := sc.service.Update(ctx, &dashboard)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(sc.logger, ctx, errorText, DASHBOARD_MICROSERVICE, DASHBOARD_RESOURCE, "PUT", req.DataSourceName, additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response := dtos.NewDashboardProtoFromModel(dashboardUpdated)

	return &response, nil
}

func (sc *dashboardHttpController) Delete(ctx context.Context, req *pb.DashboardDeleteRequest) (*emptypb.Empty, error) {
	logText := "Request Delete Dashboard"
	additionalFields := []zap.Field{zap.String("dashboard_id", req.Id)}
	_ = utils.AddControllerRequestLogging(sc.logger, ctx, logText, DASHBOARD_MICROSERVICE, DASHBOARD_RESOURCE, "DELETE", "NONE", additionalFields...)

	if err := sc.service.Delete(ctx, req.Id); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(sc.logger, ctx, errorText, DASHBOARD_MICROSERVICE, DASHBOARD_RESOURCE, "DELETE", "NONE", additionalFields...)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
