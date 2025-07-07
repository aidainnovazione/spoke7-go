package postgres

import (
	"context"

	"spoke7-go/internal/errors"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository/postgres/dao"

	"gorm.io/gorm"
)

// Create inserts a new DetectionPoint record
func (repo *postgresClient) CreateDetectionPoint(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error {
	if detectionPoint == nil {
		return errors.ErrWrongDetectionPointRequest
	}

	detectionPointDao := dao.FromDetectionPointModelToDao(datasourceName, *detectionPoint)

	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return err
	}

	return repo.db.WithContext(ctx).Create(detectionPointDao).Error
}

// Create inserts a new DetectionPoint record
func (repo *postgresClient) CreateManyDetectionPoint(ctx context.Context, datasourceName string, detectionPoint []*models.DetectionPoint) error {
	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return err
	}

	listDetectionPointDao := make([]dao.DetectionPoint, 0, len(detectionPoint))
	for _, dp := range detectionPoint {
		detectionPointDao := dao.FromDetectionPointModelToDao(datasourceName, *dp)
		listDetectionPointDao = append(listDetectionPointDao, detectionPointDao)
	}

	return repo.db.WithContext(ctx).Create(listDetectionPointDao).Error
}

// Update updates an existing DetectionPoint record
func (repo *postgresClient) UpdateDetectionPoint(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error {
	if detectionPoint == nil {
		return errors.ErrWrongDetectionPointRequest
	}

	if detectionPoint.Id == "" {
		return errors.ErrMissingDetectionPointID
	}

	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return err
	}

	detectionPointDao := dao.FromDetectionPointModelToDao(datasourceName, *detectionPoint)

	return repo.db.WithContext(ctx).Model(&dao.DetectionPoint{}).Where("id = ?", detectionPoint.Id).Updates(&detectionPointDao).Error
}

// Delete removes a DetectionPoint record by ID
func (repo *postgresClient) DeleteDetectionPoint(ctx context.Context, datasourceName string, id string) error {
	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return err
	}
	return repo.db.Delete(&dao.DetectionPoint{}, "id = ?", id).Error
}

func (repo *postgresClient) DeleteAllDetectionPointByDatasourceName(ctx context.Context, datasourceName string) error {
	return repo.db.Where(&dao.DetectionPoint{DataSourceName: datasourceName}).Delete(&dao.DetectionPoint{}).Error

	//Delete(&dao.DetectionPoint{}, ).Error
}

// Get retrieves a DetectionPoint record by ID
func (repo *postgresClient) GetDetectionPoint(ctx context.Context, datasourceName string, id string) (*models.DetectionPoint, error) {

	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return nil, err
	}

	detectionPointDao := dao.DetectionPoint{}
	err := repo.db.WithContext(ctx).Preload("Lanes").First(&detectionPointDao, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	model := dao.FromDetectionPointDaoToModel(detectionPointDao)

	return &model, nil
}

// List retrieves all DetectionPoint records
func (repo *postgresClient) ListDetectionPoint(ctx context.Context, datasourceName string) ([]*models.DetectionPoint, error) {
	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return nil, err
	}
	var detectionPointDaos []dao.DetectionPoint
	err := repo.db.WithContext(ctx).
		Preload("Lanes").
		Where("data_source_name = ?", datasourceName).
		Find(&detectionPointDaos).Error

	if err != nil {
		return nil, err
	}

	models := make([]*models.DetectionPoint, 0)
	for _, detectionPointDao := range detectionPointDaos {
		model := dao.FromDetectionPointDaoToModel(detectionPointDao)
		models = append(models, &model)
	}

	return models, nil
}

// BulkCreateDetectionPoint inserts multiple DetectionPoint records
func (repo *postgresClient) BulkCreateDetectionPoint(ctx context.Context, datasourceName string, detectionPoints []*models.DetectionPoint) error {
	if len(detectionPoints) == 0 {
		return errors.ErrWrongDetectionPointRequest
	}

	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return err
	}

	detectionPointDaos := make([]dao.DetectionPoint, 0, len(detectionPoints))
	for _, detectionPoint := range detectionPoints {
		detectionPointDaos = append(detectionPointDaos, dao.FromDetectionPointModelToDao(datasourceName, *detectionPoint))
	}
	db := repo.db.WithContext(ctx).Session(&gorm.Session{CreateBatchSize: repo.conf.BatchSize})

	return db.Create(&detectionPointDaos).Error
}
