package grpc_client

import (
	"context"
	"fmt"
	"spoke7-go/internal/metadata/dtos"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type GrpcMetadataClient struct {
	conn                      *grpc.ClientConn
	svcDataSourceClient       pb.DataSourceServiceClient
	svcDetectionPointClient   pb.DetectionPointServiceClient
	svcDetectionSectionClient pb.DetectionSectionServiceClient
	svcRoadNetworkClient      pb.RoadNetworkServiceClient
}

func NewGrpcMetadataClient(serverAddr string) (*GrpcMetadataClient, error) {
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	clientDataSource := pb.NewDataSourceServiceClient(conn)
	clientDetectionPoint := pb.NewDetectionPointServiceClient(conn)
	clientDetectionSection := pb.NewDetectionSectionServiceClient(conn)
	clientRoadNetwork := pb.NewRoadNetworkServiceClient(conn)

	return &GrpcMetadataClient{
		conn:                      conn,
		svcDataSourceClient:       clientDataSource,
		svcDetectionPointClient:   clientDetectionPoint,
		svcDetectionSectionClient: clientDetectionSection,
		svcRoadNetworkClient:      clientRoadNetwork,
	}, nil
}
func (c *GrpcMetadataClient) ListDataSource(ctx context.Context) ([]models.DataSource, error) {
	incomingMD, _ := metadata.FromIncomingContext(ctx)
	grpcCtx := metadata.NewOutgoingContext(ctx, incomingMD)

	req := pb.DataSourceListParams{
		DetectionSections: true,
		DetectionPoints:   true,
	}
	responseProto, err := c.svcDataSourceClient.List(grpcCtx, &req)
	if err != nil {
		return []models.DataSource{}, nil
	}
	var responseModelList []models.DataSource
	for _, datasource := range responseProto.Datasource {
		responseModel := dtos.DataSourceProtoToModel(datasource)
		responseModelList = append(responseModelList, responseModel)
	}

	return responseModelList, nil
}

func (c *GrpcMetadataClient) GetDataSource(ctx context.Context, name string) (models.DataSource, error) {
	incomingMD, _ := metadata.FromIncomingContext(ctx)
	grpcCtx := metadata.NewOutgoingContext(ctx, incomingMD)

	params := pb.DataSourceGetParams{
		DetectionSections: true,
		DetectionPoints:   true,
	}
	req := pb.DataSourceGetRequest{
		Name:   name,
		Params: &params,
	}
	responseProto, err := c.svcDataSourceClient.Get(grpcCtx, &req)
	if err != nil {
		return models.DataSource{}, nil
	}
	responseModel := dtos.DataSourceProtoToModel(responseProto)

	return responseModel, nil
}

func (c *GrpcMetadataClient) UpdateDataSource(ctx context.Context, datasource models.DataSource) (models.DataSource, error) {
	incomingMD, _ := metadata.FromIncomingContext(ctx)
	grpcCtx := metadata.NewOutgoingContext(ctx, incomingMD)

	req := pb.UpdateDataSource{
		Name:          datasource.Name,
		Description:   datasource.Description,
		Type:          dtos.DataSourceTypeProtoFromModel(datasource.Type),
		RoadNetworkId: *datasource.RoadNetworkId,
		Owner:         datasource.Owner,
		Groups:        datasource.Groups,
		ModifiedBy:    datasource.ModifiedBy,
	}

	responseProto, err := c.svcDataSourceClient.Update(grpcCtx, &req)
	if err != nil {
		return models.DataSource{}, nil
	}
	responseModel := dtos.DataSourceProtoToModel(responseProto)

	return responseModel, nil
}

func (c *GrpcMetadataClient) GetDetectionPoint(ctx context.Context, id string, dataSourceName string) (models.DetectionPoint, error) {
	incomingMD, _ := metadata.FromIncomingContext(ctx)
	grpcCtx := metadata.NewOutgoingContext(ctx, incomingMD)

	req := pb.GetDetectionPointRequest{
		Id:             id,
		DatasourceName: dataSourceName,
	}

	responseProto, err := c.svcDetectionPointClient.GetDetectionPoint(grpcCtx, &req)
	if err != nil {
		return models.DetectionPoint{}, nil
	}
	responseModel := dtos.DetectionPointProtoToModel(responseProto)

	return responseModel, nil
}

func (c *GrpcMetadataClient) CreateManyDetectionPoints(ctx context.Context, dataSourceName string, detectionPoints []models.DetectionPoint) ([]models.DetectionPoint, error) {
	incomingMD, _ := metadata.FromIncomingContext(ctx)
	grpcCtx := metadata.NewOutgoingContext(ctx, incomingMD)

	req := pb.CreateManyDetectionPointRequest{
		DatasourceName: dataSourceName,
		DetectionPoint: dtos.DetectionPointsProtosFromModels(detectionPoints),
	}

	responseProto, err := c.svcDetectionPointClient.CreateManyDetectionPoints(grpcCtx, &req)
	if err != nil {
		return []models.DetectionPoint{}, nil
	}
	responseModel := dtos.DetectionPointProtosToModels(responseProto.DetectionPoints)

	return responseModel, nil
}

func (c *GrpcMetadataClient) CreateRoadNetwork(ctx context.Context, network *models.RoadNetwork) (*models.RoadNetwork, error) {
	incomingMD, _ := metadata.FromIncomingContext(ctx)
	grpcCtx := metadata.NewOutgoingContext(ctx, incomingMD)

	pbNetwork, err := dtos.RoadNetworkModelToProto(network)
	if err != nil {
		return nil, fmt.Errorf("failed to convert RoadNetwork model to proto")
	}

	req := pb.RoadNetworkCreateRequest{
		RoadNetwork: pbNetwork,
	}

	result, err := c.svcRoadNetworkClient.CreateRoadNetwork(grpcCtx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to create RoadNetwork: %v", err)
	}

	// Convert the response proto to the model
	responseModel, err := dtos.RoadNetworkProtoToModel(result)
	if err != nil {
		return nil, fmt.Errorf("failed to convert RoadNetwork proto to model: %v", err)
	}
	// Return the model
	return responseModel, nil

}

func (c *GrpcMetadataClient) GetRoadNetwork(ctx context.Context, id string) (*models.RoadNetwork, error) {
	incomingMD, _ := metadata.FromIncomingContext(ctx)
	grpcCtx := metadata.NewOutgoingContext(ctx, incomingMD)

	// pbNetwork, err := dtos.RoadNetworkModelToProto(network)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to convert RoadNetwork model to proto")
	// }

	req := pb.RoadNetworkGetRequest{
		Id: id,
	}

	result, err := c.svcRoadNetworkClient.GetRoadNetwork(grpcCtx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to create RoadNetwork: %v", err)
	}

	// Convert the response proto to the model
	responseModel, err := dtos.RoadNetworkProtoToModel(result)
	if err != nil {
		return nil, fmt.Errorf("failed to convert RoadNetwork proto to model: %v", err)
	}
	// Return the model
	return responseModel, nil

}

func (c *GrpcMetadataClient) Close() {
	c.conn.Close()
}
