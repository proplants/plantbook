package plants

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/plant"
)

type GetPlantByIDImpl struct {
	repo RepoInterface
}

func NewGetPlantByIDHandler(repo RepoInterface) plant.GetPlantByIDHandler {
	return &GetPlantByIDImpl{repo: repo}
}

func (impl *GetPlantByIDImpl) Handle(params plant.GetPlantByIDParams) middleware.Responder {
	onePlant, err := impl.repo.GetPlantByID(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		return plant.NewGetPlantByIDDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: err.Error()})
	}
	return plant.NewGetPlantByIDOK().WithPayload(onePlant)
}
