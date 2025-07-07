package postgres

import (
	"context"
	"spoke7-go/internal/errors"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository/postgres/dao"
	"spoke7-go/pkg/authz"
	"strings"
)

func (repo *postgresClient) CreateDataSource(ctx context.Context, dataSource *models.DataSource) error {

	user, err := authz.GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	dataSource.Owner = user.Username

	dataSourceDao := dao.NewDataSourceDaoFromModel(*dataSource)
	return repo.db.WithContext(ctx).Create(&dataSourceDao).Error
}
func (repo *postgresClient) UpdateDataSource(ctx context.Context, dataSource *models.UpdateDataSource) error {
	if dataSource == nil {
		return errors.ErrWrongDataSourceRequest
	}

	user, err := authz.GetUserFromContext(ctx)
	if err != nil {
		return err
	}
	dataSource.ModifiedBy = &user.Username

	if dataSource.Name == "" {
		return errors.ErrMissingDataSourceName
	}

	dataSourceDao := dao.NewUpdateDataSourceDaoFromModel(*dataSource)

	return repo.db.WithContext(ctx).Model(&dao.DataSource{}).Where("name = ?", dataSource.Name).Updates(&dataSourceDao).Error

}
func (repo *postgresClient) DeleteDataSource(ctx context.Context, name string) error {
	err := repo.CheckDataSourceAccess(ctx, name)
	if err != nil {
		return err
	}
	return repo.db.Delete(&dao.DataSource{}, "name = ?", name).Error
}
func (repo *postgresClient) GetDataSource(ctx context.Context, name string, params models.DataSourceGetParams) (*models.DataSource, error) {

	err := repo.CheckDataSourceAccess(ctx, name)
	if err != nil {
		return nil, err
	}
	dataSourceDao := dao.DataSource{}
	conn := repo.db.WithContext(ctx)
	if params.DetectionSections {
		conn = conn.Preload("DetectionSections")
	}
	if params.DetectionPoints {
		conn = conn.Preload("DetectionPoints")
	}
	err = conn.First(&dataSourceDao, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	model := dataSourceDao.ToModel()

	return &model, nil
}
func (repo *postgresClient) ListDataSource(ctx context.Context, params models.DataSourceListParams, organizationName string) ([]*models.DataSource, error) {
	user, err := authz.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var dataSourceDaos []dao.DataSource
	conn := repo.db.WithContext(ctx)
	if params.DetectionSections {
		conn = conn.Preload("DetectionSections")
	}
	if params.DetectionPoints {
		conn = conn.Preload("DetectionPoints")
	}

	var groups []string
	for _, groupStr := range user.Groups {
		groupStr = strings.TrimPrefix(groupStr, organizationName+"/")
		groups = append(groups, groupStr)
	}

	err = conn.WithContext(ctx).Where("data_source.owner = ? or data_source.groups && ARRAY[?]", user.Username, groups).Find(&dataSourceDaos).Error
	if err != nil {
		return nil, err
	}

	models := make([]*models.DataSource, 0)
	for _, dao := range dataSourceDaos {
		model := dao.ToModel()
		models = append(models, &model)
	}

	return models, nil

}

// method to check if user has access to data source
func (repo *postgresClient) CheckDataSourceAccess(ctx context.Context, dataSourceName string) error {
	if dataSourceName == "" {
		return errors.ErrMissingDataSourceName
	}

	user, err := authz.GetUserFromContext(ctx)
	if err != nil {
		repo.logger.Error("failed to get user from context", "error", err)
		return errors.ErrUnauthorized
	}

	var count int64
	err = repo.db.WithContext(ctx).Model(&dao.DataSource{}).Where("name = ? and (owner = ? or groups && array[?])", dataSourceName, user.Username, user.Groups).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.ErrForbidden
	}

	return nil
}
