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

func NewGetUserPlantHandler(storage RepoInterface, tm token.Manager) userplant.GetUserPlantsHandler {
	return &getUserPlantImpl{storage: storage, tm: tm}
}

func (impl *getUserPlantImpl) Handle(params userplant.GetUserPlantsParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	cookie, err := params.HTTPRequest.Cookie(apimiddleware.JWTCookieName)
	if err != nil {
		log.Errorf("get cookie %s error, %s", apimiddleware.JWTCookieName, err)
		return userplant.NewCreateUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "not token cookie"})
	}
	if cookie == nil {
		return userplant.NewCreateUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "empty token cookie"})
	}
	ok, err := impl.tm.Check(params.HTTPRequest.Context(), cookie.Value)
	if err != nil {
		log.Errorf("check token %s error, %s", cookie.Value, err)
		return userplant.NewCreateUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "check token error"})
	}
	if !ok {
		return userplant.NewCreateUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "token expired"})
	}

	uid, _, _, err := impl.tm.FindUserData(cookie.Value)
	if err != nil {
		log.Errorf("get user attributes from token %s error, %s", cookie.Value, err)
		return userplant.NewCreateUserPlantDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "check permission error"})
	}
	userPlants, err := impl.storage.ListUserPlants(params.HTTPRequest.Context(), uid, *params.Limit, *params.Offset)
	if err != nil {
		log.Errorf("Handle StoragePlant error, %s", err)
		return userplant.NewGetUserPlantsDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error"})
	}
	return userplant.NewGetUserPlantsOK().WithPayload(userPlants).
		WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}
