package postgres

import (
	"context"
	"fmt"
	"spoke7-go/internal/errors"
	"spoke7-go/internal/storage/models"
	"spoke7-go/internal/storage/repository/postgres/dao"

	"github.com/google/uuid"
)

func (repo *postgresClient) CreateStoredFile(ctx context.Context, storedFile *models.StoredFileUpload) (*models.StoredFile, error) {
	if storedFile == nil {
		return nil, errors.ErrWrongStoredFileRequest
	}

	if len(storedFile.FileContent) == 0 {
		return nil, errors.ErrMissingStoredFileContent
	}

	storedFileDao := dao.NewStoredFileUploadDaoFromModel(*storedFile)
	storedFileDao.ID = uuid.NewString()

	err := repo.db.WithContext(ctx).Create(&storedFileDao).Error
	if err != nil {
		return nil, err
	}

	model := storedFileDao.ToModel()
	return &model, nil
}

func (repo *postgresClient) UpdateStoredFile(ctx context.Context, storedFile *models.StoredFileUpdate) error {
	if storedFile == nil {
		return errors.ErrWrongStoredFileRequest
	}

	if len(storedFile.FileContent) == 0 {
		return errors.ErrMissingStoredFileContent
	}

	storedFileDao := dao.NewStoredFileUpdateDaoFromModel(*storedFile)

	return repo.db.WithContext(ctx).Model(&dao.StoredFile{}).Where("id = ?", storedFile.ID).Updates(&storedFileDao).Error
}

func (repo *postgresClient) DeleteStoredFile(ctx context.Context, id string) error {

	return repo.db.Delete(&dao.StoredFile{}, "id = ?", id).Error
}

func (repo *postgresClient) GetStoredFile(ctx context.Context, id string) (*models.StoredFile, error) {
	storedFileDao := dao.StoredFile{}
	conn := repo.db.WithContext(ctx)
	err := conn.First(&storedFileDao, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	model := storedFileDao.ToModel()

	return &model, nil
}

func (repo *postgresClient) ListStoredFile(ctx context.Context, dataSourceName string, tag string) ([]*models.StoredFile, error) {
	var storedFileDaos []dao.StoredFile

	query := repo.db.WithContext(ctx).Model(&dao.StoredFile{})

	if dataSourceName != "" {
		query = query.Where("data_source_name = ?", dataSourceName)
	}

	if tag != "" {
		query = query.Where("tag = ?", tag)
	}

	if err := query.Find(&storedFileDaos).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve stored files: %w", err)
	}

	models := make([]*models.StoredFile, 0)
	for _, dao := range storedFileDaos {
		model := dao.ToModel()
		models = append(models, &model)
	}

	return models, nil
}
