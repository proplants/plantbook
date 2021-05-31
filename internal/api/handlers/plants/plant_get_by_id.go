package plants

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/plant"
	"github.com/kaatinga/plantbook/pkg/logging"
)

type GetPlantByIdImpl struct {
	repo RepoInterface
}

func NewGetPlantByIDHandler(repo RepoInterface) plant.GetPlantByIDHandler {
	return &GetPlantByIdImpl{repo: repo}
}

func (impl *GetPlantByIdImpl) Handle(params plant.GetPlantByIDParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	onePlant, err := impl.repo.GetPlantByID(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		log.Errorf("error handle: %v", err)
		return plant.NewGetPlantByIDBadRequest()
	}
	return plant.NewGetPlantByIDOK().WithPayload(onePlant)
}
