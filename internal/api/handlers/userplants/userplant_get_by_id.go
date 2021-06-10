package userplants

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	apimiddleware "github.com/kaatinga/plantbook/internal/api/middleware"
	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/userplant"
	"github.com/kaatinga/plantbook/pkg/logging"
)

type GetUserPlantByIDImpl struct {
	storage RepoInterface
}

func NewGetUserPlantByIDHandler(storage RepoInterface) userplant.GetUserPlantByIDHandler {
	return &GetUserPlantByIDImpl{storage: storage}
}

func (impl *GetUserPlantByIDImpl) Handle(params userplant.GetUserPlantByIDParams) middleware.Responder {
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
