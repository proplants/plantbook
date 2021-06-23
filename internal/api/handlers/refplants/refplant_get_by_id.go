package refplants

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	apimiddleware "github.com/proplants/plantbook/internal/api/middleware"
	"github.com/proplants/plantbook/internal/api/models"
	"github.com/proplants/plantbook/internal/api/restapi/operations/refplant"
	"github.com/proplants/plantbook/pkg/logging"
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
	log := logging.FromContext(params.HTTPRequest.Context())
	oneRefPlant, err := impl.storage.GetRefPlantByID(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		log.Errorf("Handle GetRefPlantsByID error, %s", err)
		return refplant.NewGetRefPlantByIDDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	return refplant.NewGetRefPlantByIDOK().WithPayload(oneRefPlant)
}
