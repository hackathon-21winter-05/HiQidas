package heya

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/service/utils"
	"time"
)

type HeyaService interface {
	CreateHeya(userID uuid.UUID, title, description string) (*model.Heya, error)
	DeleteHeya(heyaID uuid.UUID) error
}

type HeyaServiceImpl struct {
	repo repository.Repository
}

func (h *HeyaServiceImpl) DeleteHeya(heyaID uuid.UUID) error {
	ctx, cancel := utils.CreateTxContext()
	defer cancel()

	err := h.repo.Do(ctx, nil, func(ctx context.Context) error {
		if err := h.repo.DeleteHeyaByID(ctx, heyaID); err != nil {
			return err
		}
		if err := h.repo.DeleteHistoryByHeyaID(ctx, heyaID); err != nil {
			return err
		}
		if err := h.repo.DeleteHiqidashiByHeyaID(ctx, heyaID); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if errors.Is(err, repository.ErrNoRecordDeleted) {
			return fmt.Errorf("failed to delete no record")
		}
		return err
	}

	return nil
}

func (h *HeyaServiceImpl) CreateHeya(userID uuid.UUID, title, description string) (*model.Heya, error) {
	ctx, cancel := utils.CreateTxContext()
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

	hiqidashi := &model.Hiqidashi{
		ID:           utils.GetUUID(),
		HeyaID:       heyaID,
		CreatorID:    userID,
		LastEditorID: userID,
		Title:        title,
		Description:  description,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	err := h.repo.Do(ctx, nil, func(ctx context.Context) error {
		if err := h.repo.CreateHeya(ctx, heya); err != nil {
			return err
		}

		if err := h.repo.CreateHiqidashi(ctx, hiqidashi); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return heya, nil
}

func NewHeyaServiceImpl(repo repository.Repository) *HeyaServiceImpl {
	return &HeyaServiceImpl{repo: repo}
}
