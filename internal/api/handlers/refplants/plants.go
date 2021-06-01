package refplants

import (
	"context"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/refplant"
)

type RepoInterface interface {
	GetRefPlants(ctx context.Context, params refplant.GetRefPlantsParams) ([]*models.RefPlant, error)
	GetRefPlantByID(ctx context.Context, id int64) (*models.RefPlant, error)
}
