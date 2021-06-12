package refplants

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/proplants/plantbook/internal/api/models"
	"github.com/proplants/plantbook/internal/api/restapi/operations/refplant"
)

type getRefPlantsImpl struct {
	storage RepoInterface
}

// NewGetRefPlantsHandler builder for refplant.GetRefPlantsHandler interface implementation.
func NewGetRefPlantsHandler(repo RepoInterface) refplant.GetRefPlantsHandler {
	return &getRefPlantsImpl{storage: repo}
}

// Handle implementation of the refplant.GetRefPlantsHandler interface.
func (impl *getRefPlantsImpl) Handle(params refplant.GetRefPlantsParams) middleware.Responder {
	someRefPlants, err := impl.storage.GetRefPlants(params.HTTPRequest.Context(), params)
	if err != nil {
		return refplant.NewGetRefPlantsDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: err.Error()})
	}
	return refplant.NewGetRefPlantsOK().WithPayload(someRefPlants)
}
