package grpc_client

import (
	"context"
	"fmt"
	"spoke7-go/internal/storage/dtos"
	"spoke7-go/internal/storage/models"
	"spoke7-go/internal/storage/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type GrpcStorageClient struct {
	conn                *grpc.ClientConn
	svcStoredFileClient pb.StoredFileServiceClient
}

func NewGrpcStorageClient(serverAddr string) (*GrpcStorageClient, error) {
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	clientStoredFile := pb.NewStoredFileServiceClient(conn)

	return &GrpcStorageClient{
		conn:                conn,
		svcStoredFileClient: clientStoredFile,
	}, nil
}

func (c *GrpcStorageClient) UploadStoredFile(ctx context.Context, dataSourceName string, fileName string, fileContent []byte, description string, fileFormat string, fileType models.FileType) (models.StoredFile, error) {
	incomingMD, _ := metadata.FromIncomingContext(ctx)
	grpcCtx := metadata.NewOutgoingContext(ctx, incomingMD)

	req := pb.StoredFileUploadRequest{
		DataSourceName: dataSourceName,
		Description:    description,
		FileName:       fileName,
		FileContent:    fileContent,
		FileFormat:     fileFormat,
		FileSize:       uint32(len(fileContent)),
		FileType:       pb.FileType(fileType),
	}

	responseProto, err := c.svcStoredFileClient.Upload(grpcCtx, &req)
	if err != nil {
		return models.StoredFile{}, nil
	}
	responseModel := dtos.StoredFileProtoToModel(responseProto)

	return responseModel, nil
}

func (c *GrpcStorageClient) Close() {
	c.conn.Close()
}
