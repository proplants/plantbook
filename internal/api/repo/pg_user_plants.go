package repo

import (
	"context"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/proplants/plantbook/internal/api/models"
)

// StorePlant create new user's plant.
func (pg *PG) StorePlant(ctx context.Context, plant *models.UserPlant) (*models.UserPlant, error) {
	query := `insert into public.user_plants (user_id, ref_id, garden_id,
		planting_date,watering_interval, last_watering, photo_url,
		title, description) 
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

// ListUserPlants get all user's plants from db.
func (pg *PG) ListUserPlants(ctx context.Context,
	userID, limit, offset int64) ([]*models.UserPlant, error) {
	query := `SELECT id, user_id, ref_id, garden_id, planting_date, watering_interval,
		last_watering, next_watering, photo_url, title, description,
		created_at, modified_at
	FROM public.user_plants
	WHERE user_id = $1
	ORDER BY title
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

// DeleteUserPlant delete user's plant by id.
func (pg *PG) DeleteUserPlant(ctx context.Context, userPlantID int64) error {
	query := `DELETE FROM public.user_plants
	WHERE id = $1
	returning id;`
	var deletedUserPlantID int64
	err := pg.db.QueryRow(ctx, query, userPlantID).Scan(&deletedUserPlantID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.Errorf("no rows deleted")
		}
		return errors.WithMessage(err, "delete user's plant failed")
	}
	if deletedUserPlantID == userPlantID {
		return nil
	}
	return errors.Errorf("expected delete with id=%d, but deleted with id=%d", userPlantID, deletedUserPlantID)
}

// UpdateUserPlant update user's plant.
func (pg *PG) UpdateUserPlant(ctx context.Context, plant *models.UserPlant) (*models.UserPlant, error) {
	query := `UPDATE public.user_plants SET (ref_id, garden_id, planting_date, watering_interval,
		last_watering, photo_url, title, description, modified_at) = 
		($2, $3, $4, $5, $6, $7, $8, $9, NOW())
	WHERE id = $1
	returning id, next_watering, modified_at;`
	var updNextWatering, updModifiedAt strfmt.DateTime
	var updUserPlantID int64

	err := pg.db.QueryRow(ctx, query, plant.ID, plant.PlantReferenceID, plant.GardenID, plant.PlantingDate,
		plant.WateringInterval, plant.LastWatering, plant.PhotoUrls, plant.Title, plant.Description).
		Scan(&updUserPlantID, &updNextWatering, &updModifiedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Errorf("no rows updated")
		}
		return nil, errors.WithMessage(err, "update plant failed")
	}
	plant.NextWatering, plant.ModifiedAt = updNextWatering, updModifiedAt

	return plant, err
}

// GetUserPlantByID get user's plant by id from db.
func (pg *PG) GetUserPlantByID(ctx context.Context, userPlantID int64) (*models.UserPlant, error) {
	query := `SELECT id, user_id, ref_id, garden_id, planting_date, watering_interval,
	last_watering, next_watering, photo_url, title, description,
	created_at, modified_at
FROM public.user_plants
WHERE id = $1`
	var interval time.Duration
	var userPlant models.UserPlant
	err := pg.db.QueryRow(ctx, query, userPlantID).Scan(&userPlant.ID, &userPlant.UserID, &userPlant.PlantReferenceID,
		&userPlant.GardenID, &userPlant.PlantingDate, &interval, &userPlant.LastWatering,
		&userPlant.NextWatering, &userPlant.PhotoUrls, &userPlant.Title, &userPlant.Description,
		&userPlant.CreatedAt, &userPlant.ModifiedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.WithMessage(err, "get user's plant error")
	}
	userPlant.WateringInterval = interval.String()
	return &userPlant, err
}
