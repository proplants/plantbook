package plants

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/plant"
	"github.com/kaatinga/plantbook/pkg/logging"
)

type GetPlantsImpl struct {
	repo RepoInterface
}

func NewGetPlantsHandler(repo RepoInterface) plant.GetPlantsHandler {
	return &GetPlantsImpl{repo: repo}
}

func (impl *GetPlantsImpl) Handle(params plant.GetPlantsParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	somePlants, err := impl.repo.GetPlants(params.HTTPRequest.Context(), params)
	if err != nil {
		log.Errorf("error handle: %v", err)
		return plant.NewGetPlantsBadRequest()

	}
	return plant.NewGetPlantsOK().WithPayload(somePlants)
}
