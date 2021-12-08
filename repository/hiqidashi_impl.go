package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

func (repo *GormRepository) GetHiqidashisByHeyaID(ctx context.Context, heyaID uuid.UUID) ([]*model.Hiqidashi, error) {
	panic("implement me")
}

func (repo *GormRepository) GetHiqidashisByParentID(ctx context.Context, parentID uuid.UUID) ([]*model.Hiqidashi, error) {
	panic("implement me")
}

func (repo *GormRepository) CreateHiqidashi(ctx context.Context, hiqidashi *model.Hiqidashi) error {
	panic("implement me")
}

func (repo *GormRepository) DeleteHiqidashi(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}

func (repo *GormRepository) UpdateHiqidashi(ctx context.Context, hiqidashi *model.Hiqidashi) error {
	panic("implement me")
}
