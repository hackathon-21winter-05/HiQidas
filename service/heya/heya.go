package heya

import (
	"context"
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/service/utils"
	"time"
)

type HeyaService interface {
	CreateHeya(c context.Context, userID uuid.UUID, title, description string) (*model.Heya, error)
	DeleteHeya(c context.Context, heyaID uuid.UUID) error
	GetHeyas(c context.Context) ([]model.Heya, error)
	GetHeyaByID(c context.Context, heyaID uuid.UUID) (*model.Heya, error)
	GetUsersByHeyaID(c context.Context, heyaID uuid.UUID) ([]uuid.UUID, error)
	PutHeyaByID(c context.Context, heya *model.NullHeya, heyaID, userID uuid.UUID) error
}

type HeyaServiceImpl struct {
	repo repository.Repository
}

func (h *HeyaServiceImpl) GetHeyas(c context.Context) ([]model.Heya, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()
	heyasID, err := h.repo.GetHeyas(ctx)
	if err != nil {
		return nil, err
	}

	return heyasID, nil
}

func (h *HeyaServiceImpl) GetHeyaByID(c context.Context, heyaID uuid.UUID) (*model.Heya, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	heya, err := h.repo.GetHeyaByID(ctx, heyaID)
	if err != nil {
		return nil, err
	}

	return heya, nil
}

func (h *HeyaServiceImpl) GetUsersByHeyaID(c context.Context, heyaID uuid.UUID) ([]uuid.UUID, error) {
	//TODO: どこからUserがそのヘヤにいるのかを持ってくる
	/*ctx,cancel := utils.CreateTxContext(c)
	defer cancel()*/
	return nil, nil
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
		if err := h.repo.DeleteHiqidashisByHeyaID(ctx, heyaID); err != nil {
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

	err := h.repo.Do(ctx, nil, func(ctx context.Context) error {
		if err := h.repo.CreateHeya(ctx, heya); err != nil {
			return err
		}

		if err := h.repo.CreateHiqidashi(ctx, rootHiqidashi); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return heya, nil
}

func (h *HeyaServiceImpl) PutHeyaByID(c context.Context, heya *model.NullHeya, heyaID, userID uuid.UUID) error {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	err := h.repo.Do(ctx, nil, func(ctx context.Context) error {
		now := time.Now()
		nullHeya := model.NullHeya{
			ID:           heyaID,
			Title:        heya.Title,
			Description:  heya.Description,
			LastEditorID: userID,
			UpdatedAt:    sql.NullTime{Time: now, Valid: true},
		}
		if err := h.repo.UpdateHeyaByID(ctx, &nullHeya); err != nil {
			return err
		}
		if err := h.repo.UpdateRootHiqidashiByHeyaID(ctx, heyaID, nullHeya.Title, nullHeya.Description); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func NewHeyaServiceImpl(repo repository.Repository) *HeyaServiceImpl {
	return &HeyaServiceImpl{repo: repo}
}
