package postgres

import (
	"context"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository/postgres/dao"
	"spoke7-go/pkg/authz"
	"strings"

	"github.com/google/uuid"
)

func (repo *postgresClient) CreateDashboard(ctx context.Context, dashboard *models.Dashboard) error {
	user, err := authz.GetUserFromContext(ctx)
	if err != nil {
		return err
	}
	dashboard.ID = uuid.NewString()
	dashboard.Owner = user.Username
	dashboardDao := dao.NewDashboardDaoFromModel(*dashboard)
	return repo.db.WithContext(ctx).Create(&dashboardDao).Error
}
func (repo *postgresClient) UpdateDashboard(ctx context.Context, dashboard *models.Dashboard) error {
	dashboardDao := dao.NewDashboardDaoFromModel(*dashboard)

	return repo.db.WithContext(ctx).Model(&dao.Dashboard{}).Where("id = ?", dashboard.ID).Updates(&dashboardDao).Error

}
func (repo *postgresClient) DeleteDashboard(ctx context.Context, id string) error {
	// err := repo.CheckDashboardAccess(ctx, id)
	// if err != nil {
	// 	return err
	// }
	return repo.db.Delete(&dao.Dashboard{}, "id = ?", id).Error
}
func (repo *postgresClient) GetDashboard(ctx context.Context, id string) (*models.Dashboard, error) {
	// TODO
	// err := repo.CheckDashboardAccess(ctx, id)
	// if err != nil {
	// 	return nil, err
	// }
	dashboardDao := dao.Dashboard{}
	err := repo.db.First(&dashboardDao, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	model := dashboardDao.ToModel()

	return &model, nil
}
func (repo *postgresClient) ListDashboard(ctx context.Context, dataSourceName string, organizationName string) ([]*models.Dashboard, error) {
	user, err := authz.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var dashboardDaos []dao.Dashboard

	var groups []string
	for _, groupStr := range user.Groups {
		groupStr = strings.TrimPrefix(groupStr, organizationName+"/")
		groups = append(groups, groupStr)
	}

	query := repo.db.WithContext(ctx).Where("owner = ? or groups && ARRAY[?]", user.Username, groups)

	if dataSourceName != "" {
		query = query.Where("data_source_name = ?", dataSourceName)
	}

	err = query.Find(&dashboardDaos).Error
	if err != nil {
		return nil, err
	}

	models := make([]*models.Dashboard, 0)
	for _, dao := range dashboardDaos {
		model := dao.ToModel()
		models = append(models, &model)
	}

	return models, nil
}
