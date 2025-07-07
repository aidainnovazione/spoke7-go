package storage_interface

import (
	"context"
	"spoke7-go/internal/storage/models"
)

// StorageInterface defines the interface for interacting with the storage.
// It provides methods to create, update, delete, retrieve, and list source data.
type StorageInterface interface {
	Close() error
	Migrate(ctx context.Context) error

	// CreateStoredFile inserts new source data into the database.
	CreateStoredFile(ctx context.Context, storedFile *models.StoredFileUpload) (*models.StoredFile, error)
	// UpdateStoredFile updates existing source data in the database.
	UpdateStoredFile(ctx context.Context, storedFile *models.StoredFileUpdate) error
	// DeleteStoredFile removes source data from the database by name.
	DeleteStoredFile(ctx context.Context, id string) error
	// GetStoredFile retrieves source data from the database by name and optional params.
	GetStoredFile(ctx context.Context, id string) (*models.StoredFile, error)
	// ListStoredFile lists all source data in the database that match the optional params.
	ListStoredFile(ctx context.Context, dataSourceName string, tag string) ([]*models.StoredFile, error)
}
