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

type deleteGardenImpl struct {
	storage RepoInterface
	tm      token.Manager
}

// NewDeleteUserGardenHandler builder for gardens.DeleteUserGardenHandler interface implementation
func NewDeleteUserGardenHandler(storage RepoInterface, tm token.Manager) gardens.DeleteUserGardenHandler {
	return &deleteGardenImpl{storage: storage, tm: tm}
}

// Handle implementation of the gardens.DeleteUserGardenHandler interface
func (dg *deleteGardenImpl) Handle(params gardens.DeleteUserGardenParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	// check cookie TODO: replace to middleware!!!
	cookie, err := params.HTTPRequest.Cookie(apimiddleware.JWTCookieName)
	if err != nil {
		log.Errorf("get cookie %s error, %s", apimiddleware.JWTCookieName, err)
		return gardens.NewCreateUserGardenDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "not token cookie"})
	}
	if cookie == nil {
		return gardens.NewCreateUserGardenDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "empty token cookie"})
	}
	ok, err := dg.tm.Check(params.HTTPRequest.Context(), cookie.Value)
	if err != nil {
		log.Errorf("check token %s error, %s", cookie.Value, err)
		return gardens.NewCreateUserGardenDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "check token error"})
	}
	if !ok {
		return gardens.NewCreateUserGardenDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "token expired"})
	}

	uid, _, userRoleID, err := dg.tm.FindUserData(cookie.Value)
	if err != nil {
		log.Errorf("get user attributes from token %s error, %s", cookie.Value, err)
		return gardens.NewCreateUserGardenDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "check permission error"})
	}

	existsGarden, err := dg.storage.FindGardenByID(params.HTTPRequest.Context(), params.GardenID)
	if err != nil {
		log.Infof("storage.FindGardenByID with id=%d error, %s", params.GardenID, err)
		return gardens.NewDeleteUserGardenDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error happen"})
	}
	if existsGarden == nil {
		log.Infof("storage.FindGardenByID with id=%d not found", params.GardenID)
		return gardens.NewDeleteUserGardenDefault(http.StatusNotFound).
			WithPayload(&models.ErrorResponse{Message: "garden not found"})
	}
	// only admin or owner can delete garden
	isAdmin := userRoleID == handlers.UserRoleAdmin
	isOwner := existsGarden.UserID == uid
	if !(isAdmin || isOwner) {
		log.Errorf("userID=%d, not owner and not admin try delete garden", uid)
		return gardens.NewCreateUserGardenDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "forbidden"})
	}

	// delete garden from storage
	err = dg.storage.DeleteGarden(params.HTTPRequest.Context(), params.GardenID)
	if err != nil {
		log.Errorf("storage.DeleteGarden error, %s", err)
		return gardens.NewCreateUserGardenDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error happen"})
	}
	// all ok return delete garden message
	return gardens.NewDeleteUserGardenOK().WithPayload(&models.Response{Message: "garden deleted"}).
		WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}
