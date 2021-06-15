package refplants

import (
	"context"

	"github.com/proplants/plantbook/internal/api/models"
)

type RepoInterface interface {
	GetRefPlants(ctx context.Context, category, limit, offset int32, classifier, floweringTime,
		hight, kind, recommendPosition, regardToLight, regardToMoisture string) ([]*models.RefPlant, error)
	GetRefPlantByID(ctx context.Context, id int64) (*models.RefPlant, error)
}
