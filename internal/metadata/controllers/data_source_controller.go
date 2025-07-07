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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var DATASOURCE_MICROSERVICE string = "metadata"
var DATASOURCE_RESOURCE string = "datasource"

type dataSourceHttpController struct {
	service services.DataSourceService
	logger  logger.Logger
	pb.UnimplementedDataSourceServiceServer
}

func NewDataSourceController(service services.DataSourceService, logger logger.Logger) pb.DataSourceServiceServer {
	return &dataSourceHttpController{service: service, logger: logger}
}

func (sc *dataSourceHttpController) List(ctx context.Context, req *pb.DataSourceListParams) (*pb.DataSourceListResponse, error) {
	logText := "Request List Data Source"
	_ = utils.AddControllerRequestLogging(sc.logger, ctx, logText, DATASOURCE_MICROSERVICE, DATASOURCE_RESOURCE, "GET", "NONE")

	params := dtos.DataSourceListParamsProtoToModel(req)

	list, err := sc.service.List(ctx, params)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(sc.logger, ctx, errorText, DATASOURCE_MICROSERVICE, DATASOURCE_RESOURCE, "GET", "NONE")
		return nil, err
	}

	response := make([]*pb.DataSource, 0)
	for _, dataSource := range list {
		dto := dtos.NewDataSourceProtoFromModel(dataSource)
		response = append(response, &dto)
	}

	return &pb.DataSourceListResponse{Datasource: response}, nil
}

func (sc *dataSourceHttpController) Get(ctx context.Context, req *pb.DataSourceGetRequest) (*pb.DataSource, error) {
	logText := "Request Get Data Source"
	_ = utils.AddControllerRequestLogging(sc.logger, ctx, logText, DATASOURCE_MICROSERVICE, DATASOURCE_RESOURCE, "GET", req.Name)

	name := req.Name
	params := dtos.DataSourceGetParamsProtoToModel(req.Params)

	dataSource, err := sc.service.Get(ctx, name, params)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(sc.logger, ctx, errorText, DATASOURCE_MICROSERVICE, DATASOURCE_RESOURCE, "GET", req.Name)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response := dtos.NewDataSourceProtoFromModel(dataSource)

	return &response, nil
}

func (sc *dataSourceHttpController) Create(ctx context.Context, req *pb.DataSource) (*pb.DataSource, error) {
	logText := "Request Create Data Source"
	_ = utils.AddControllerRequestLogging(sc.logger, ctx, logText, DATASOURCE_MICROSERVICE, DATASOURCE_RESOURCE, "POST", req.Name)

	dataSource := dtos.DataSourceProtoToModel(req)
	if err := sc.service.Create(ctx, &dataSource); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(sc.logger, ctx, errorText, DATASOURCE_MICROSERVICE, DATASOURCE_RESOURCE, "POST", req.Name)

		return nil, err
	}

	response := dtos.NewDataSourceProtoFromModel(&dataSource)

	return &response, nil
}

func (sc *dataSourceHttpController) Update(ctx context.Context, req *pb.UpdateDataSource) (*pb.DataSource, error) {
	logText := "Request Update Data Source"
	_ = utils.AddControllerRequestLogging(sc.logger, ctx, logText, DATASOURCE_MICROSERVICE, DATASOURCE_RESOURCE, "PUT", req.Name)

	dataSource := dtos.UpdateDataSourceProtoToModel(req)
	dataSourceUpdated, err := sc.service.Update(ctx, &dataSource)
	if err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(sc.logger, ctx, errorText, DATASOURCE_MICROSERVICE, DATASOURCE_RESOURCE, "PUT", req.Name)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response := dtos.NewDataSourceProtoFromModel(dataSourceUpdated)

	return &response, nil
}

func (sc *dataSourceHttpController) Delete(ctx context.Context, req *pb.DataSourceDeleteRequest) (*emptypb.Empty, error) {
	logText := "Request Delete Data Source"
	_ = utils.AddControllerRequestLogging(sc.logger, ctx, logText, DATASOURCE_MICROSERVICE, DATASOURCE_RESOURCE, "DELETE", req.Name)

	if err := sc.service.Delete(ctx, req.Name); err != nil {
		errorText := fmt.Sprintf("%v: %v", logText, err.Error())
		_ = utils.AddControllerErrorLogging(sc.logger, ctx, errorText, DATASOURCE_MICROSERVICE, DATASOURCE_RESOURCE, "DELETE", req.Name)

		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
