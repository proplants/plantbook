package plants

import (
	"context"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/plant"
)

type RepoInterface interface {
	GetPlants(ctx context.Context, params plant.GetPlantsParams) ([]*models.Plant, error)
}
