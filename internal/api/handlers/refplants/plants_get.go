package refplants

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/refplant"
)

type GetRefPlantsImpl struct {
	repo RepoInterface
}

func NewGetRefPlantsHandler(repo RepoInterface) refplant.GetRefPlantsHandler {
	return &GetRefPlantsImpl{repo: repo}
}

func (impl *GetRefPlantsImpl) Handle(params refplant.GetRefPlantsParams) middleware.Responder {
	someRefPlants, err := impl.repo.GetRefPlants(params.HTTPRequest.Context(), params)
	if err != nil {
		return refplant.NewGetRefPlantsDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: err.Error()})

	}
	return refplant.NewGetRefPlantsOK().WithPayload(someRefPlants)
}
