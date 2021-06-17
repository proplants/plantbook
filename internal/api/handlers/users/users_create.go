package users

import (
	"net/http"

	apimiddleware "github.com/proplants/plantbook/internal/api/middleware"
	"github.com/proplants/plantbook/pkg/logging"
	"github.com/proplants/plantbook/pkg/token"

	"github.com/proplants/plantbook/internal/api/handlers"
	"github.com/proplants/plantbook/internal/api/models"
	"github.com/proplants/plantbook/internal/api/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type createUserImpl struct {
	storage RepoInterface
	tm      token.Manager
}

// NewCreateUserHandler builder for user.CreateUserHandler interface implementation.
func NewCreateUserHandler(repo RepoInterface, tm token.Manager) user.CreateUserHandler {
	return &createUserImpl{storage: repo, tm: tm}
}

// Handle implementation of the user.CreateUserHandler interface.
func (cui *createUserImpl) Handle(params user.CreateUserParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	// check cookie TODO: replace to middleware!!!
	cookie, err := params.HTTPRequest.Cookie(apimiddleware.JWTCookieName)
	if err != nil {
		log.Errorf("get cookie %s error, %s", apimiddleware.JWTCookieName, err)
		return user.NewCreateUserDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "not token cookie"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	if cookie == nil {
		return user.NewCreateUserDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "empty token cookie"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	ok, err := cui.tm.Check(params.HTTPRequest.Context(), cookie.Value)
	if err != nil {
		log.Errorf("check token %s error, %s", cookie.Value, err)
		return user.NewCreateUserDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "check token error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	if !ok {
		return user.NewCreateUserDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "token expired"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	// check admin role
	uid, uname, roleID, err := cui.tm.FindUserData(cookie.Value)
	if err != nil {
		log.Errorf("get user attributes from token %s error, %s", cookie.Value, err)
		return user.NewCreateUserDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "check permission error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	if roleID != handlers.UserRoleAdmin {
		log.With(zap.Int64("user_id", uid), zap.String("user_name", uname)).
			Warnf("create user forbidden for user.role_id=%d", roleID)
		return user.NewCreateUserDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "operation not permitted"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	// end check cookie and access rights

	// make password hash
	salt := MakeSalt(SaltLen)
	passwordHash := HashPass(salt, params.Body.Password)
	log.Debugf("found user: %+v", params.Body)
	// insert user to repo
	_user, err := cui.storage.StoreUser(params.HTTPRequest.Context(), params.Body, passwordHash)
	if err != nil {
		log.Errorf("repo.StoreUser error, %s", err)
		return user.NewCreateUserDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error happen"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	// all ok return new user
	return user.NewCreateUserCreated().WithPayload(_user).
		WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}
