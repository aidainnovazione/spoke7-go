package services

import (
	"context"
	"spoke7-go/internal/storage/models"
	"spoke7-go/pkg/grpc_client"
)

type UploadFileService interface {
	UploadFile(ctx context.Context, dataSourceName string, fileName string, fileContent []byte, description string, fileFormat string, fileType models.FileType) error
}

type uploadFileService struct {
	grpcStorageService grpc_client.GrpcStorageClient
}

func NewUploadFileService(grpcStorageService grpc_client.GrpcStorageClient) UploadFileService {
	return &uploadFileService{grpcStorageService: grpcStorageService}
}

func (s *uploadFileService) UploadFile(ctx context.Context, dataSourceName string, fileName string, fileContent []byte, description string, fileFormat string, fileType models.FileType) error {
	_, err := s.grpcStorageService.UploadStoredFile(ctx, dataSourceName, fileName, fileContent, description, fileFormat, fileType)
	if err != nil {
		return err
	}

	return nil

}
