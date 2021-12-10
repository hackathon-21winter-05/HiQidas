package heya

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/service/utils"
	"time"
)

type HeyaService interface {
	CreateHeya(c context.Context, userID uuid.UUID, title, description string) (*model.Heya, error)
	DeleteHeya(c context.Context, heyaID uuid.UUID) error
	GetHeyas(c context.Context) ([]uuid.UUID, error)
	GetHeyasByID(c context.Context, heyaID uuid.UUID) (*model.Heya, error)
}

type HeyaServiceImpl struct {
	repo repository.Repository
}

func (h *HeyaServiceImpl) GetHeyas(c context.Context) ([]uuid.UUID, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()
	heyasID, err := h.repo.GetHeyasID(ctx)
	if err != nil {
		return nil, err
	}

	return heyasID, nil
}

func (h *HeyaServiceImpl) GetHeyasByID(c context.Context, heyaID uuid.UUID) (*model.Heya, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	heya, err := h.GetHeyasByID(ctx, heyaID)
	if err != nil {
		return nil, err
	}

	return heya, nil
}
func (h *HeyaServiceImpl) DeleteHeya(c context.Context, heyaID uuid.UUID) error {
	ctx, cancel := utils.CreateTxContext(c)
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
		return err
	}

	return nil
}

func (h *HeyaServiceImpl) CreateHeya(c context.Context, userID uuid.UUID, title, description string) (*model.Heya, error) {
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
