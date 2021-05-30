package plants

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/plant"
)

type GetPlantsImpl struct {
	repo RepoInterface
}

func NewGetPlantsHandler(repo RepoInterface) plant.GetPlantsHandler {
	return &GetPlantsImpl{repo: repo}
}

func (impl *GetPlantsImpl) Handle(params plant.GetPlantsParams) middleware.Responder {
	plants, err := impl.repo.GetPlants(params.HTTPRequest.Context(), params)
	if err != nil {
		return plant.NewGetPlantsBadRequest()
	}
	return plant.NewGetPlantsOK().WithPayload(plants)
}
