package refplants

import (
	"context"

	"github.com/proplants/plantbook/internal/api/models"
)

type RepoInterface interface {
	GetRefPlants(ctx context.Context, category int32, limit, offset int64, classifier, floweringTime,
		hight, kind, recommendPosition, regardToLight, regardToMoisture string) ([]*models.RefPlant, int64, int64, error)
	GetRefPlantByID(ctx context.Context, id int64) (*models.RefPlant, error)
}
