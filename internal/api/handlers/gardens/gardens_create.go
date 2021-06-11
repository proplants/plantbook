package gardens

import (
	"net/http"

	apimiddleware "github.com/kaatinga/plantbook/internal/api/middleware"
	"github.com/kaatinga/plantbook/pkg/logging"
	"github.com/kaatinga/plantbook/pkg/token"

	"github.com/kaatinga/plantbook/internal/api/handlers"
	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/gardens"

	"github.com/go-openapi/runtime/middleware"
)

type createGardenImpl struct {
	storage RepoInterface
	tm      token.Manager
}

// NewCreateUserGardenHandler builder for gardens.CreateUserGardenHandler interface implementation.
func NewCreateUserGardenHandler(storage RepoInterface, tm token.Manager) gardens.CreateUserGardenHandler {
	return &createGardenImpl{storage: storage, tm: tm}
}

// Handle implementation of the user.CreateUserHandler interface.
func (cg *createGardenImpl) Handle(params gardens.CreateUserGardenParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	// check cookie TODO: replace to middleware!!!
	cookie, err := params.HTTPRequest.Cookie(apimiddleware.JWTCookieName)
	if err != nil {
		log.Errorf("get cookie %s error, %s", apimiddleware.JWTCookieName, err)
		return gardens.NewCreateUserGardenDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "not token cookie"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	if cookie == nil {
		return gardens.NewCreateUserGardenDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "empty token cookie"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	ok, err := cg.tm.Check(params.HTTPRequest.Context(), cookie.Value)
	if err != nil {
		log.Errorf("check token %s error, %s", cookie.Value, err)
		return gardens.NewCreateUserGardenDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "check token error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	if !ok {
		return gardens.NewCreateUserGardenDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "token expired"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}

	uid, _, roleID, err := cg.tm.FindUserData(cookie.Value)
	if err != nil {
		log.Errorf("get user attributes from token %s error, %s", cookie.Value, err)
		return gardens.NewCreateUserGardenDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "check permission error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}

	// fill owner_id for garden if userRole = gardener
	// if userRole = admin, miss because admin must to set OwnerID
	if roleID == handlers.UserRoleGardener || params.Garden.UserID == 0 {
		params.Garden.UserID = uid
	}

	// insert user to repo
	newGarden, err := cg.storage.StoreGarden(params.HTTPRequest.Context(), params.Garden)
	if err != nil {
		log.Errorf("storage.StoreGarden error, %s", err)
		return gardens.NewCreateUserGardenDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error happen"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	// all ok return new garden with id
	return gardens.NewCreateUserGardenCreated().WithPayload(newGarden).
		WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}
