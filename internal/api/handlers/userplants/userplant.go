package userplants

import (
	"context"

	"github.com/kaatinga/plantbook/internal/api/models"
)

type RepoInterface interface {
	StorePlant(ctx context.Context, plant *models.UserPlant) (*models.UserPlant, error)
	ListUserPlants(ctx context.Context, userID, limit, offset int64) ([]*models.UserPlant, error)
}
