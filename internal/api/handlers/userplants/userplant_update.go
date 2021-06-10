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

type updateUserPlantImpl struct {
	storage RepoInterface
	tm      token.Manager
}

func NewUpdateUserPlantHandler(storage RepoInterface, tm token.Manager) userplant.UpdateUserPlantHandler {
	return &updateUserPlantImpl{storage: storage, tm: tm}
}

func (impl *updateUserPlantImpl) Handle(params userplant.UpdateUserPlantParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	cookie, err := params.HTTPRequest.Cookie(apimiddleware.JWTCookieName)
	if err != nil {
		log.Errorf("get cookie %s error, %s", apimiddleware.JWTCookieName, err)
		return userplant.NewUpdateUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "not token cookie"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	if cookie == nil {
		return userplant.NewUpdateUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "empty token cookie"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	ok, err := impl.tm.Check(params.HTTPRequest.Context(), cookie.Value)
	if err != nil {
		log.Errorf("check token %s error, %s", cookie.Value, err)
		return userplant.NewUpdateUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "check token error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	if !ok {
		return userplant.NewUpdateUserPlantDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "token expired"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}

	uid, _, userRoleID, err := impl.tm.FindUserData(cookie.Value)
	if err != nil {
		log.Errorf("get user attributes from token %s error, %s", cookie.Value, err)
		return userplant.NewUpdateUserPlantDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "check permission error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	existingUserPlant, err := impl.storage.GetUserPlantByID(params.HTTPRequest.Context(), params.Userplant.ID)
	if err != nil {
		log.Infof("storage.GetUserPlantByID with id=%d error, %s", params.Userplant.ID, err)
		return userplant.NewUpdateUserPlantDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error happen"})
	}
	if existingUserPlant == nil {
		log.Infof("storage.GetUserPlantByID with id=%d not found", params.Userplant.ID)
		return userplant.NewUpdateUserPlantDefault(http.StatusNotFound).
			WithPayload(&models.ErrorResponse{Message: "user's plant not found"})
	}

	isAdmin := userRoleID == handlers.UserRoleAdmin
	isOwner := existingUserPlant.UserID == uid
	if !(isAdmin || isOwner) {
		log.Errorf("userID=%d, not owner and not admin try update", uid)
		return userplant.NewDeleteUserPlantDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "forbidden"})
	}

	updatedUserPlant, err := impl.storage.UpdateUserPlant(params.HTTPRequest.Context(), params.Userplant)
	if err != nil {
		log.Errorf("Handle UpdateUserPlant error: %s", err)
		return userplant.NewUpdateUserPlantDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}

	return userplant.NewUpdateUserPlantOK().WithPayload(updatedUserPlant).
		WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}