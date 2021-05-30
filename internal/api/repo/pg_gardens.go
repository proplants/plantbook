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
