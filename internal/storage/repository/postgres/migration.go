package postgres

import (
	"context"
	"spoke7-go/internal/storage/repository/postgres/dao"
)

func (pc *postgresClient) Migrate(ctx context.Context) error {
	return pc.db.WithContext(ctx).AutoMigrate(
		&dao.StoredFile{},
	)
}
