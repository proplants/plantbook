package repo

import (
	"context"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/pkg/errors"
)

func (pg *PG) StorePlant(ctx context.Context, plant *models.UserPlant) (*models.UserPlant, error) {
	query := `insert into public.user_plants (user_id, ref_id, garden_id, planting_date,watering_interval, last_watering, photo_url, name, description) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id;`

	var userPlantID int64
	err := pg.db.QueryRow(ctx, query, plant.UserID, plant.PlantReferenceID, plant.GardenID, plant.PlantingDate,
		plant.WateringInterval, plant.LastWatering, plant.PhotoUrls, plant.Name, plant.Description).
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
