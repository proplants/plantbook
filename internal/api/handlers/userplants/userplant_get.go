package userplants

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	apimiddleware "github.com/kaatinga/plantbook/internal/api/middleware"
	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/userplant"
	"github.com/kaatinga/plantbook/pkg/logging"
	"github.com/kaatinga/plantbook/pkg/token"
)

type getUserPlantImpl struct {
	storage RepoInterface
	tm      token.Manager
}

// NewGetUserPlantHandler builder for userplant.GetUserPlantsHandler interface implementation.
func NewGetUserPlantHandler(storage RepoInterface, tm token.Manager) userplant.GetUserPlantsHandler {
	return &getUserPlantImpl{storage: storage, tm: tm}
}

// Handle implementation of the userplant.GetUserPlantsHandler interface.
func (impl *getUserPlantImpl) Handle(params userplant.GetUserPlantsParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	cookie, err := params.HTTPRequest.Cookie(apimiddleware.JWTCookieName)
	if err != nil {
		log.Errorf("get cookie %s error, %s", apimiddleware.JWTCookieName, err)
		return userplant.NewCreateUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "not token cookie"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	if cookie == nil {
		return userplant.NewCreateUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "empty token cookie"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	ok, err := impl.tm.Check(params.HTTPRequest.Context(), cookie.Value)
	if err != nil {
		log.Errorf("check token %s error, %s", cookie.Value, err)
		return userplant.NewCreateUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "check token error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	if !ok {
		return userplant.NewCreateUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "token expired"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}

	uid, _, _, err := impl.tm.FindUserData(cookie.Value)
	if err != nil {
		log.Errorf("get user attributes from token %s error, %s", cookie.Value, err)
		return userplant.NewCreateUserPlantDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "check permission error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	userPlants, err := impl.storage.ListUserPlants(params.HTTPRequest.Context(), uid, params.Limit, params.Offset)
	if err != nil {
		log.Errorf("Handle StoragePlant error, %s", err)
		return userplant.NewGetUserPlantsDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	resultSet := models.ResultSet{
		Limit:  params.Limit,
		Offset: params.Offset,
	}
	return userplant.NewGetUserPlantsOK().WithPayload(&models.UserPlantsResponse{Data: userPlants, ResultSet: &resultSet}).
		WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}
