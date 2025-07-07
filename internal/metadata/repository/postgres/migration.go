package postgres

import (
	"context"
	"spoke7-go/internal/metadata/repository/postgres/dao"
)

func (pc *postgresClient) Migrate(ctx context.Context) error {
	return pc.db.WithContext(ctx).AutoMigrate(
		&dao.DataSource{},
		&dao.DetectionSection{},
		&dao.DetectionPoint{},
		&dao.DetectionSectionRoadNetwork{},
		&dao.Lane{},
		&dao.RoadNetwork{},
		&dao.Dashboard{},
	)
}
