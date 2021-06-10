package refplants

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/refplant"
)

type getRefPlantByIDImpl struct {
	storage RepoInterface
}

// NewGetRefPlantByIDHandler builder for refplant.GetRefPlantByIDHandler interface implementation.
func NewGetRefPlantByIDHandler(repo RepoInterface) refplant.GetRefPlantByIDHandler {
	return &getRefPlantByIDImpl{storage: repo}
}

// Handle implementation of the refplant.GetRefPlantByIDHandler interface.
func (impl *getRefPlantByIDImpl) Handle(params refplant.GetRefPlantByIDParams) middleware.Responder {
	oneRefPlant, err := impl.storage.GetRefPlantByID(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		return refplant.NewGetRefPlantByIDDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: err.Error()})
	}
	return refplant.NewGetRefPlantByIDOK().WithPayload(oneRefPlant)
}
