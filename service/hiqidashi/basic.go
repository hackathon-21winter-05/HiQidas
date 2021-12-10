package hiqidashi

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/service/utils"
)

func (hs *HiqidashiServiceImpl) GetHiqidashisByHeyaID(c context.Context, heyaID uuid.UUID) ([]*model.Hiqidashi, error) {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	hiqidashis, err := hs.repo.GetHiqidashisByHeyaID(ctx, heyaID)
	if err != nil {
		return nil, err
	}

	return hiqidashis, nil
}

func (hs *HiqidashiServiceImpl) CreateHiqidashi(c context.Context, createrID, heyaID, parentID uuid.UUID) (*model.Hiqidashi, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	now := time.Now()

	hiqidashi := &model.Hiqidashi{
		ID:           id,
		HeyaID:       heyaID,
		CreatorID:    createrID,
		LastEditorID: createrID,
		ParentID:     utils.NullUUIDFrom(parentID),
		Title:        "",
		Description:  "",
		Drawing:      utils.NullString(),
		ColorCode:    "#9E7A7A",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	err = hs.repo.Do(ctx, nil, func(ctx context.Context) error {
		err = hs.repo.CreateHiqidashi(ctx, hiqidashi)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return hiqidashi, nil
}

func (hs *HiqidashiServiceImpl) UpdateHiqidashiByID(c context.Context, hiqidashi *model.NullHiqidashi) error {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	tempHiqidashi, err := hs.repo.GetHiqidashiByID(ctx, hiqidashi.ID)
	if err != nil {
		return err
	}
	rootHiqidashi := false
	if !tempHiqidashi.ParentID.Valid {
		rootHiqidashi = true
	}

	err = hs.repo.Do(ctx, nil, func(ctx context.Context) error {
		err := hs.repo.UpdateHiqidashiByID(ctx, hiqidashi)
		if err != nil {
			return err
		}

		if rootHiqidashi {
			err = hs.repo.UpdateHeyaByID(ctx, &model.NullHeya{
				ID:           tempHiqidashi.HeyaID,
				Title:        hiqidashi.Title,
				Description:  hiqidashi.Description,
				LastEditorID: hiqidashi.LastEditorID,
			})
			if err != nil {
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

func (hs *HiqidashiServiceImpl) DeleteHiqidashiByID(c context.Context, id uuid.UUID) error {
	ctx, cancel := utils.CreateTxContext(c)
	defer cancel()

	err := hs.repo.Do(ctx, nil, func(ctx context.Context) error {
		err := hs.repo.DeleteHiqidashiByID(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
