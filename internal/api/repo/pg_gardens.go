package repo

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/kaatinga/plantbook/internal/api/models"

	"github.com/pkg/errors"
)

// StoreGarden inserts new garden to public.gardens.
func (pg *PG) StoreGarden(ctx context.Context, garden *models.Garden) (*models.Garden, error) {
	const query string = `INSERT INTO public.gardens 
		(user_id, title, description) 
		VALUES($1, $2, $3)
		returning id;`
	var gardenID int64
	err := pg.db.QueryRow(ctx, query, garden.UserID, garden.Title, garden.Description).
		Scan(&gardenID)
	if err != nil {
		return nil, errors.WithMessage(err, "insert garden failed")
	}
	if gardenID == 0 {
		return nil, errors.Errorf("insert garden failed, empty id")
	}
	garden.ID = gardenID
	return garden, nil
}


// ListGarden shows a list of user's gardens
func (pg *PG) ListGarden(ctx context.Context, garden *models.Garden) ([]models.Garden, error) {
	const query string = `SELECT (id, userId, title, description) FROM public.gardens
		WHERE user_id=$1;`
	
	gardensList, err := pg.db.Query(ctx, query, garden.UserID)
	if err != nil {
		return nil, errors.WithMessage(err, "the list of the user's gardens could not be displayed")
	}
	defer gardensList.Close()

	var gardenArr []models.Garden
	for gardensList.Next() {
		var gardenOne models.Garden
		err = gardensList.Scan(&gardenOne.ID, &gardenOne.UserID, &gardenOne.Title, &gardenOne.Description)
		if err != nil {
			return nil, errors.WithMessage(err, "can not extract the row ")
		}

	}
	return gardenArr, nil
}
// FindGardenByID extracts garden from db by specified id.
func (pg *PG) FindGardenByID(ctx context.Context, gardenID int64) (*models.Garden, error) {
	const query string = `SELECT id, user_id, title, description
		FROM public.gardens WHERE id=$1;`

	garden := &models.Garden{}
	err := pg.db.QueryRow(ctx, query, gardenID).Scan(&garden.ID, &garden.UserID, &garden.Title, &garden.Description)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, errors.WithMessage(err, "fetch garden failed")
	}
	return garden, nil
}

// DeleteGarden deletes garden from db by specified id.
func (pg *PG) DeleteGarden(ctx context.Context, gardenID int64) error {
	const query string = `DELETE FROM public.gardens WHERE id=$1 returning id;`
	var deletedGardenID int64
	err := pg.db.QueryRow(ctx, query, gardenID).Scan(&deletedGardenID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.Errorf("no rows deleted")
		}
		return errors.WithMessage(err, "delete garden failed")
	}
	if deletedGardenID == gardenID {
		return nil
	}
	return errors.Errorf("expected delete with id=%d, but deleted with id=%d", gardenID, deletedGardenID)
}
