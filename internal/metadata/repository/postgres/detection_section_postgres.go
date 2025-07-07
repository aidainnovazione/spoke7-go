package postgres

import (
	"context"

	"spoke7-go/internal/errors"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository/postgres/dao"
)

// Create inserts a new DetectionSection record
func (repo *postgresClient) CreateDetectionSection(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error {
	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return err
	}
	detectionSectionDao := dao.NewDetectionSectionDaoFromModel(*detectionSection)
	detectionSectionDao.DataSourceName = datasourceName
	return repo.db.WithContext(ctx).Create(&detectionSectionDao).Error
}

// Create inserts a new DetectionSection record
func (repo *postgresClient) CreateManyDetectionSection(ctx context.Context, datasourceName string, detectionSection []*models.DetectionSection) error {
	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return err
	}

	listDetectionSectionDao := make([]dao.DetectionSection, 0, len(detectionSection))
	for _, ds := range detectionSection {
		detectionSectionDao := dao.NewDetectionSectionDaoFromModel(*ds)
		listDetectionSectionDao = append(listDetectionSectionDao, detectionSectionDao)
	}

	return repo.db.WithContext(ctx).Create(listDetectionSectionDao).Error
}

// Update updates an existing DetectionSection record
func (repo *postgresClient) UpdateDetectionSection(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error {
	if detectionSection == nil {
		return errors.ErrWrongDetectionSectionRequest
	}

	if detectionSection.Id == "" {
		return errors.ErrMissingDetectionSectionID
	}
	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return err
	}

	detectionSectionDao := dao.NewDetectionSectionDaoFromModel(*detectionSection)
	detectionSectionDao.DataSourceName = datasourceName
	return repo.db.WithContext(ctx).Model(&dao.DetectionSection{}).Where("id = ?", detectionSection.Id).Updates(&detectionSectionDao).Error
}

// Delete removes a DetectionSection record by ID
func (repo *postgresClient) DeleteDetectionSection(ctx context.Context, datasourceName string, id string) error {
	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return err
	}
	return repo.db.Delete(&dao.DetectionSection{}, "id = ?", id).Error
}

// Get retrieves a DetectionSection record by ID
func (repo *postgresClient) GetDetectionSection(ctx context.Context, datasourceName string, id string) (*models.DetectionSection, error) {
	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return nil, err
	}
	detectionSectionDao := dao.DetectionSection{}
	err := repo.db.WithContext(ctx).First(&detectionSectionDao, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	model := detectionSectionDao.ToModel()

	return &model, nil
}

// List retrieves all DetectionSection records
func (repo *postgresClient) ListDetectionSection(ctx context.Context, datasourceName string) ([]*models.DetectionSection, error) {
	if err := repo.CheckDataSourceAccess(ctx, datasourceName); err != nil {
		return nil, err
	}
	var detectionSectionDaos []dao.DetectionSection
	err := repo.db.WithContext(ctx).Find(&detectionSectionDaos).Error
	if err != nil {
		return nil, err
	}

	models := make([]*models.DetectionSection, 0)
	for _, detectionSectionDao := range detectionSectionDaos {
		model := detectionSectionDao.ToModel()
		models = append(models, &model)
	}

	return models, nil
}
