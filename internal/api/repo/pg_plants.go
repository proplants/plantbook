package repo

import (
	"context"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/pkg/errors"
)

func (pg *PG) StorePlant(ctx context.Context, plant *models.Plant) (*models.Plant, error) {
	query := `insert into public.user_plants (id_user, garden_id, id_plant, planting_time,
		watering_interval, last_watering, next_watering, photo_url, name_user_plant, description
		 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10))`

	var userPlantID int64
	err := pg.db.QueryRow(ctx, query, plant.UserID, plant.GardenID, plant.PlantReferenceID, plant.Title, plant.PhotoUrls).
		Scan(&userPlantID)
	if err != nil {
		return nil, errors.WithMessage(err, "insert plant failed")
	}
	if userPlantID == 0 {
		return nil, errors.Errorf("insert plant failed, empty id")
	}
	plant.ID = userPlantID
	return plant, err
}
