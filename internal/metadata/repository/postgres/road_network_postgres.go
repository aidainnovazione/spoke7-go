package postgres

import (
	"context"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository/postgres/dao"
	"spoke7-go/pkg/authz"
)

func (repo *postgresClient) CreateRoadNetwork(ctx context.Context, roadNetwork *models.RoadNetwork) (*models.RoadNetwork, error) {

	user, err := authz.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	roadNetwork.Owner = user.Username

	roadNetworkDao := dao.NewRoadNetworkDaoFromModel(roadNetwork)
	if err := repo.db.WithContext(ctx).Create(&roadNetworkDao).Error; err != nil {
		return nil, err
	}

	md, err := dao.NewRoadNetworkModelFromDao(roadNetworkDao)
	if err != nil {
		return nil, err
	}

	return md, nil

}

func (repo *postgresClient) GetRoadNetworkByID(ctx context.Context, id string) (*models.RoadNetwork, error) {
	var roadNetworkDao dao.RoadNetwork
	if err := repo.db.WithContext(ctx).First(&roadNetworkDao, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return dao.NewRoadNetworkModelFromDao(&roadNetworkDao)
}

func (repo *postgresClient) ListRoadNetworks(ctx context.Context) ([]*models.RoadNetwork, error) {

	user, err := authz.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var roadNetworksDaos []dao.RoadNetwork

	conn := repo.db.WithContext(ctx)

	//err = conn.WithContext(ctx).Omit("Geom").Where("owner = ? or groups && ARRAY[?]", user.Username, user.Groups).Find(&roadNetworksDaos).Error
	err = conn.WithContext(ctx).Omit("Geom").Where("owner = ?", user.Username).Find(&roadNetworksDaos).Error

	if err != nil {
		return nil, err
	}

	var roadNetworks []*models.RoadNetwork
	for _, daoObj := range roadNetworksDaos {
		model, err := dao.NewRoadNetworkModelFromDao(&daoObj)
		if err != nil {
			return nil, err
		}
		roadNetworks = append(roadNetworks, model)
	}
	return roadNetworks, nil
}

func (repo *postgresClient) UpdateRoadNetwork(ctx context.Context, roadNetwork *models.RoadNetwork) (*models.RoadNetwork, error) {
	roadNetworkDao := dao.NewRoadNetworkDaoFromModel(roadNetwork)
	if err := repo.db.WithContext(ctx).Omit("owner", "role").Save(&roadNetworkDao).Error; err != nil {
		return nil, err
	}
	return dao.NewRoadNetworkModelFromDao(roadNetworkDao)
}

func (repo *postgresClient) DeleteRoadNetwork(ctx context.Context, id string) error {
	if err := repo.db.WithContext(ctx).Delete(&dao.RoadNetwork{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
