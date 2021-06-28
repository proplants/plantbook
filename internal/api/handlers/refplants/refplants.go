package refplants

import (
	"context"

	"github.com/proplants/plantbook/internal/api/models"
	"github.com/proplants/plantbook/internal/api/restapi/operations/refplant"
)

type RepoInterface interface {
	GetRefPlants(params refplant.GetRefPlantsParams) ([]*models.RefPlant, int64, int64, error)
	GetRefPlantByID(ctx context.Context, id int64) (*models.RefPlant, error)
}
