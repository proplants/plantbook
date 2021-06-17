package userplants

import (
	"context"

	"github.com/proplants/plantbook/internal/api/models"
)

type RepoInterface interface {
	StorePlant(ctx context.Context, plant *models.UserPlant) (*models.UserPlant, error)
	ListUserPlants(ctx context.Context, userID, limit, offset int64) ([]*models.UserPlant, error)
	DeleteUserPlant(ctx context.Context, userPlantID int64) error
	UpdateUserPlant(ctx context.Context, plant *models.UserPlant) (*models.UserPlant, error)
	GetUserPlantByID(ctx context.Context, userPlantID int64) (*models.UserPlant, error)
}
