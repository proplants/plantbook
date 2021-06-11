package userplants

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/kaatinga/plantbook/internal/api/handlers"
	apimiddleware "github.com/kaatinga/plantbook/internal/api/middleware"
	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/userplant"
	"github.com/kaatinga/plantbook/pkg/logging"
	"github.com/kaatinga/plantbook/pkg/token"
)

type createUserPlantImpl struct {
	storage RepoInterface
	tm      token.Manager
}

// NewCreateUserPlantHandler builder for userplant.CreateUserPlantHandler interface implementation.
func NewCreateUserPlantHandler(storage RepoInterface, tm token.Manager) userplant.CreateUserPlantHandler {
	return &createUserPlantImpl{storage: storage, tm: tm}
}

// Handle implementation of the userplant.CreateUserPlantHandler interface.
func (impl *createUserPlantImpl) Handle(params userplant.CreateUserPlantParams) middleware.Responder {
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

	uid, _, roleID, err := impl.tm.FindUserData(cookie.Value)
	if err != nil {
		log.Errorf("get user attributes from token %s error, %s", cookie.Value, err)
		return userplant.NewCreateUserPlantDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "check permission error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}

	if roleID == handlers.UserRoleGardener || params.Userplant.UserID == 0 {
		params.Userplant.UserID = uid
	}
	newUserPlant, err := impl.storage.StorePlant(params.HTTPRequest.Context(), params.Userplant)
	if err != nil {
		log.Errorf("Handle StoragePlant error, %s", err)
		return userplant.NewCreateUserPlantDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}

	return userplant.NewCreateUserPlantOK().WithPayload(newUserPlant).
		WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}
