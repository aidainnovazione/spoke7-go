package services

import (
	"context"
	"spoke7-go/internal/storage/models"
	"spoke7-go/internal/storage/storage_interface"
	"spoke7-go/pkg/authz"
	"spoke7-go/pkg/grpc_client"
)

type StoredFileService interface {
	Create(ctx context.Context, storedFile *models.StoredFileUpload) (*models.StoredFile, error)
	Update(ctx context.Context, storedFile *models.StoredFileUpdate) (*models.StoredFile, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*models.StoredFile, error)
	List(ctx context.Context, dataSourceName string, tag string) ([]*models.StoredFile, error)
}
type storedFileService struct {
	repo                storage_interface.StorageInterface
	grpcMetadataService grpc_client.GrpcMetadataClient
	organizationName    string
}

func NewStoredFileService(repo storage_interface.StorageInterface, grpcMetadataService grpc_client.GrpcMetadataClient, organizationName string) StoredFileService {
	return &storedFileService{repo: repo, grpcMetadataService: grpcMetadataService, organizationName: organizationName}
}

func (s *storedFileService) Create(ctx context.Context, storedFile *models.StoredFileUpload) (*models.StoredFile, error) {
	if err := CheckDataSourceUserPermission(ctx, storedFile.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}

	user, _ := authz.GetUserFromContext(ctx)
	storedFile.Owner = user.Username

	if storedFile.FileSize != uint32(len(storedFile.FileContent)) {
		storedFile.FileSize = uint32(len(storedFile.FileContent))
	}

	createdStoreFile, err := s.repo.CreateStoredFile(ctx, storedFile)
	if err != nil {
		return nil, err
	}

	return s.Get(ctx, createdStoreFile.ID)
}

func (s *storedFileService) Update(ctx context.Context, storedFile *models.StoredFileUpdate) (*models.StoredFile, error) {
	if err := CheckDataSourceUserPermission(ctx, storedFile.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}
	if storedFile.FileSize != uint32(len(storedFile.FileContent)) {
		storedFile.FileSize = uint32(len(storedFile.FileContent))
	}

	err := s.repo.UpdateStoredFile(ctx, storedFile)
	if err != nil {
		return nil, err
	}

	return s.Get(ctx, storedFile.ID)
}

func (s *storedFileService) Delete(ctx context.Context, id string) error {
	storedFile, err := s.repo.GetStoredFile(ctx, id)
	if err != nil {
		return err
	}

	if err := CheckDataSourceUserPermission(ctx, storedFile.DataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return err
	}
	return s.repo.DeleteStoredFile(ctx, id)
}

func (s *storedFileService) Get(ctx context.Context, id string) (*models.StoredFile, error) {
	storedFile, err := s.repo.GetStoredFile(ctx, id)
	if err != nil {
		return nil, err
	}

	return storedFile, err
}

func (s *storedFileService) List(ctx context.Context, dataSourceName string, tag string) ([]*models.StoredFile, error) {
	if err := CheckDataSourceUserPermission(ctx, dataSourceName, s.organizationName, s.grpcMetadataService); err != nil {
		return nil, err
	}

	return s.repo.ListStoredFile(ctx, dataSourceName, tag)
}
