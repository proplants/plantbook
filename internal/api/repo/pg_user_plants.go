package repo

import (
	"context"
	"time"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/pkg/errors"
)

func (pg *PG) StorePlant(ctx context.Context, plant *models.UserPlant) (*models.UserPlant, error) {
	query := `insert into public.user_plants (user_id, ref_id, garden_id,
		planting_date,watering_interval, last_watering, photo_url,
		name, description) 
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
	returning id, next_watering, created_at;`
	err := pg.db.QueryRow(ctx, query, plant.UserID, plant.PlantReferenceID, plant.GardenID, plant.PlantingDate,
		plant.WateringInterval, plant.LastWatering, plant.PhotoUrls, plant.Title, plant.Description).
		Scan(&plant.ID, &plant.NextWatering, &plant.CreatedAt)
	if err != nil {
		return nil, errors.WithMessage(err, "insert plant failed")
	}
	if plant.ID == 0 {
		return nil, errors.Errorf("insert plant failed, empty id")
	}
	return plant, err
}

func (pg *PG) ListUserPlants(ctx context.Context,
	userID, limit, offset int64) ([]*models.UserPlant, error) {
	query := `SELECT id, user_id, ref_id, garden_id, planting_date, watering_interval,
		last_watering, next_watering, photo_url, name, description,
		created_at, modified_at
	FROM public.user_plants
	WHERE user_id = $1
	ORDER BY name
	OFFSET $3
	LIMIT $2;`
	var userPlants []*models.UserPlant
	rows, err := pg.db.Query(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, errors.WithMessage(err, "Select rows error: ")
	}
	defer rows.Close()

	for rows.Next() {
		var userPlant models.UserPlant
		var interval time.Duration
		err = rows.Scan(&userPlant.ID, &userPlant.UserID, &userPlant.PlantReferenceID,
			&userPlant.GardenID, &userPlant.PlantingDate, &interval, &userPlant.LastWatering,
			&userPlant.NextWatering, &userPlant.PhotoUrls, &userPlant.Title, &userPlant.Description,
			&userPlant.CreatedAt, &userPlant.ModifiedAt)
		if err != nil {
			return nil, errors.WithMessage(err, "Scan rows error: ")
		}
		userPlant.WateringInterval = interval.String()
		userPlants = append(userPlants, &userPlant)
	}
	return userPlants, err
}
