package repository

import (
	"context"
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

func (repo *GormRepository) GetHeyasID(ctx context.Context) ([]uuid.UUID, error) {
	panic("implement me")
}

func (repo *GormRepository) GetHeyaByID(ctx context.Context, id uuid.UUID) (*model.Heya, error) {
	panic("implement me")
}

func (repo *GormRepository) CreateHeya(ctx context.Context, title string, description sql.NullString) (*model.Heya, error) {
	panic("implement me")
}

func (repo *GormRepository) UpdateHeya(ctx context.Context, heya *model.Heya) error {
	panic("implement me")
}

func (repo *GormRepository) DeleteHeya(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}
