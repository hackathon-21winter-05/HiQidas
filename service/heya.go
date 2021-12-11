package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/service/utils"
)

type HeyaService interface {
	CreateHeya(c context.Context, userID uuid.UUID, title, description string) (*model.Heya, error)
	DeleteHeya(c context.Context, heyaID uuid.UUID) error
	GetHeyas(c context.Context) ([]*model.Heya, error)
	GetHeyaByID(c context.Context, heyaID uuid.UUID) (*model.Heya, error)
	PutHeyaByID(c context.Context, heya *model.NullHeya, heyaID, userID uuid.UUID) error
	PutFavoriteByHeyaID(c context.Context, heyaID uuid.UUID, isFavorite bool) error
}

func (s *Service) CreateHeya(c context.Context, userID uuid.UUID, title, description string) (*model.Heya, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	now := time.Now()
	heyaID := utils.GetUUID()
	heya := &model.Heya{
		ID:           heyaID,
		Title:        title,
		Description:  description,
		CreatorID:    userID,
		LastEditorID: userID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	rootHiqidashi := &model.Hiqidashi{
		ID:           utils.GetUUID(),
		HeyaID:       heyaID,
		CreatorID:    userID,
		LastEditorID: userID,
		ParentID:     uuid.NullUUID{UUID: uuid.Nil, Valid: false},
		Title:        title,
		Description:  description,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	err := s.repo.Do(ctx, nil, func(ctx context.Context) error {
		if err := s.repo.CreateHeya(ctx, heya); err != nil {
			return err
		}

		if err := s.repo.CreateHiqidashi(ctx, rootHiqidashi); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return heya, nil
}

func (s *Service) DeleteHeya(c context.Context, heyaID uuid.UUID) error {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	err := s.repo.Do(ctx, nil, func(ctx context.Context) error {
		if err := s.repo.DeleteHeyaByID(ctx, heyaID); err != nil {
			return err
		}
		if err := s.repo.DeleteHistoryByHeyaID(ctx, heyaID); err != nil {
			return err
		}
		if err := s.repo.DeleteHiqidashisByHeyaID(ctx, heyaID); err != nil {
			return err
		}
		if err := s.repo.DeleteFavoriteByHeyaID(ctx, heyaID); err != nil {
			//そもそもFavoriteにない可能性があるため無かったらerrではなくnilを返す
			if errors.Is(err, repository.ErrNoRecordDeleted) {
				return nil
			}
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetHeyas(c context.Context) ([]*model.Heya, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()
	heyas, err := s.repo.GetHeyas(ctx)
	if err != nil {
		return nil, err
	}

	return heyas, nil
}

func (s *Service) GetHeyaByID(c context.Context, heyaID uuid.UUID) (*model.Heya, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	heya, err := s.repo.GetHeyaByID(ctx, heyaID)
	if err != nil {
		return nil, err
	}

	return heya, nil
}

func (s *Service) PutHeyaByID(c context.Context, heya *model.NullHeya, heyaID, userID uuid.UUID) error {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	err := s.repo.Do(ctx, nil, func(ctx context.Context) error {
		now := time.Now()
		nullHeya := model.NullHeya{
			ID:           heyaID,
			Title:        heya.Title,
			Description:  heya.Description,
			LastEditorID: userID,
			UpdatedAt:    sql.NullTime{Time: now, Valid: true},
		}
		if err := s.repo.UpdateHeyaByID(ctx, &nullHeya); err != nil {
			return err
		}
		if err := s.repo.UpdateRootHiqidashiByHeyaID(ctx, heyaID, nullHeya.Title, nullHeya.Description); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) PutFavoriteByHeyaID(c context.Context, heyaID uuid.UUID, userID uuid.UUID, isFavorite bool) error {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	err := s.repo.Do(ctx, nil, func(ctx context.Context) error {
		if isFavorite {
			favo := model.Favorite{
				UserID: userID,
				HeyaID: heyaID,
			}
			favos, err := s.repo.GetFavoritesByUserID(ctx, userID)
			if err == repository.ErrNotFound {
				return nil
			}
			//存在していたらtrueを返す
			exists := false
			for _, favorite := range favos {
				if favorite.HeyaID != heyaID {
					continue
				}
				exists = true
			}
			if exists {
				return nil
			}
			if err = s.repo.CreateFavorite(ctx, &favo); err != nil {
				return err
			}
		} else {
			if err := s.repo.DeleteFavoriteByHeyaIDAndUserID(ctx, heyaID, userID); err != nil {
				if errors.Is(err, repository.ErrNoRecordDeleted) {
					return nil
				}
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
