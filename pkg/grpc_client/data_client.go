package grpc_client

import (
	"context"
	"fmt"

	"spoke7-go/internal/data/dtos"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type GrpcDataClient struct {
	conn                       *grpc.ClientConn
	svcCurrentByDetPoint       pb.CurrentTrafficDataByDetectionPointServiceClient
	svcCurrentByDetPointByLane pb.CurrentTrafficDataByDetectionPointByLaneServiceClient
	svcCurrentByDetSection     pb.CurrentTrafficDataByDetectionSectionServiceClient
}

func NewGrpcDataClient(serverAddr string) (*GrpcDataClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	clientByPoint := pb.NewCurrentTrafficDataByDetectionPointServiceClient(conn)
	clientByLane := pb.NewCurrentTrafficDataByDetectionPointByLaneServiceClient(conn)
	clientBySection := pb.NewCurrentTrafficDataByDetectionSectionServiceClient(conn)

	return &GrpcDataClient{
		conn:                       conn,
		svcCurrentByDetPoint:       clientByPoint,
		svcCurrentByDetPointByLane: clientByLane,
		svcCurrentByDetSection:     clientBySection,
	}, nil
}

func (c *GrpcDataClient) CreateCurrentTrafficDataByDetectionSection(ctx context.Context, model models.CurrentTrafficDataByDetectionSectionModel) (models.CurrentTrafficDataByDetectionSectionModel, error) {
	proto, err := dtos.FromCurrentTrafficDataByDetectionSectionModelToProto(&model)
	if err != nil {
		return models.CurrentTrafficDataByDetectionSectionModel{}, nil
	}
	req := pb.CreateCurrentTrafficDataByDetectionSectionRequest{
		CurrentTrafficDataByDetectionSection: proto,
	}

	incomingMD, _ := metadata.FromIncomingContext(ctx)
	grpcCtx := metadata.NewOutgoingContext(ctx, incomingMD)

	responseProto, err := c.svcCurrentByDetSection.CreateCurrentTrafficDataByDetectionSection(grpcCtx, &req)
	if err != nil {
		return models.CurrentTrafficDataByDetectionSectionModel{}, nil
	}
	responseModel, err := dtos.FromCurrentTrafficDataByDetectionSectionProtoToModel(responseProto.CurrentTrafficDataByDetectionSection)
	if err != nil {
		return models.CurrentTrafficDataByDetectionSectionModel{}, nil
	}
	return *responseModel, nil
}

func (c *GrpcDataClient) CreateCurrentTrafficDataByDetectionPoint(ctx context.Context, model models.CurrentTrafficDataByDetectionPointModel) (models.CurrentTrafficDataByDetectionPointModel, error) {
	proto, err := dtos.FromCurrentTrafficDataByDetectionPointModelToProto(&model)
	if err != nil {
		return models.CurrentTrafficDataByDetectionPointModel{}, nil
	}
	req := pb.CreateCurrentTrafficDataByDetectionPointRequest{
		CurrentTrafficDataByDetectionPoint: proto,
	}

	incomingMD, _ := metadata.FromIncomingContext(ctx)
	grpcCtx := metadata.NewOutgoingContext(ctx, incomingMD)
	responseProto, err := c.svcCurrentByDetPoint.CreateCurrentTrafficDataByDetectionPoint(grpcCtx, &req)
	if err != nil {
		return models.CurrentTrafficDataByDetectionPointModel{}, nil
	}
	responseModel, err := dtos.FromCurrentTrafficDataByDetectionPointProtoToModel(responseProto.CurrentTrafficDataByDetectionPoint)
	if err != nil {
		return models.CurrentTrafficDataByDetectionPointModel{}, nil
	}
	return *responseModel, nil
}

func (c *GrpcDataClient) CreateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, model models.CurrentTrafficDataByDetectionPointByLaneModel) (models.CurrentTrafficDataByDetectionPointByLaneModel, error) {
	proto, err := dtos.FromCurrentTrafficDataByDetectionPointByLaneModelToProto(&model)
	if err != nil {
		return models.CurrentTrafficDataByDetectionPointByLaneModel{}, nil
	}
	req := pb.CreateCurrentTrafficDataByDetectionPointByLaneRequest{
		CurrentTrafficDataByDetectionPointByLane: proto,
	}
	incomingMD, _ := metadata.FromIncomingContext(ctx)
	grpcCtx := metadata.NewOutgoingContext(ctx, incomingMD)
	responseProto, err := c.svcCurrentByDetPointByLane.CreateCurrentTrafficDataByDetectionPointByLane(grpcCtx, &req)
	if err != nil {
		return models.CurrentTrafficDataByDetectionPointByLaneModel{}, err
	}
	responseModel, err := dtos.FromCurrentTrafficDataByDetectionPointByLaneProtoToModel(responseProto.CurrentTrafficDataByDetectionPointByLane)
	if err != nil {
		return models.CurrentTrafficDataByDetectionPointByLaneModel{}, nil
	}
	return *responseModel, nil
}

func (c *GrpcDataClient) Close() {
	c.conn.Close()
}
