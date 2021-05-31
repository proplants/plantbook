package repo

import (
	"context"

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