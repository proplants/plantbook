package userplants

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	apimiddleware "github.com/kaatinga/plantbook/internal/api/middleware"
	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/userplant"
	"github.com/kaatinga/plantbook/pkg/logging"
)

type getUserPlantByIDImpl struct {
	storage RepoInterface
}

// NewGetUserPlantByIDHandler builder for userplant.GetUserPlantByIDHandler interface implementation.
func NewGetUserPlantByIDHandler(storage RepoInterface) userplant.GetUserPlantByIDHandler {
	return &getUserPlantByIDImpl{storage: storage}
}

// Handle implementation of the userplant.GetUserPlantByIDHandler interface.
func (impl *getUserPlantByIDImpl) Handle(params userplant.GetUserPlantByIDParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	plant, err := impl.storage.GetUserPlantByID(params.HTTPRequest.Context(), params.UserplantID)
	if err != nil {
		log.Errorf("Handle GetUserPlantByID error, %s", err)
		return userplant.NewGetUserPlantByIDDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	return userplant.NewGetUserPlantByIDOK().WithPayload(plant).
		WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}
