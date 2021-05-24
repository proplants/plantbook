package users

import (
	"net/http"

	"github.com/kaatinga/plantbook/pkg/logging"
	"github.com/kaatinga/plantbook/pkg/token"
	"go.uber.org/zap"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
)

type createUserImpl struct {
	repo RepoInterface
	tm   token.Manager
}

// NewCreateUserHandler builder for user.CreateUserHandler interface implementation
func NewCreateUserHandler(repo RepoInterface, tm token.Manager) user.CreateUserHandler {
	return &createUserImpl{repo: repo, tm: tm}
}

// Handle implementation of the user.CreateUserHandler interface
func (cui *createUserImpl) Handle(params user.CreateUserParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	// check cookie TODO: replace to middleware!!!
	cookie, err := params.HTTPRequest.Cookie(jwtCookieName)
	if err != nil {
		log.Errorf("get cookie %s error, %s", jwtCookieName, err)
		return user.NewCreateUserDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "not token cookie"})
	}
	if cookie == nil {
		return user.NewCreateUserDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "empty token cookie"})
	}
	ok, err := cui.tm.Check(params.HTTPRequest.Context(), cookie.Value)
	if err != nil {
		log.Errorf("check token %s error, %s", cookie.Value, err)
		return user.NewCreateUserDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "check token error"})
	}
	if !ok {
		return user.NewCreateUserDefault(http.StatusUnauthorized).
			WithPayload(&models.ErrorResponse{Message: "token expired"})
	}
	// check admin role
	uid, uname, roleID, err := cui.tm.FindUserData(cookie.Value)
	if err != nil {
		log.Errorf("get user attributes from token %s error, %s", cookie.Value, err)
		return user.NewCreateUserDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "check permission error"})
	}
	if roleID != userRoleAdmin {
		log.With(zap.Int64("user_id", uid), zap.String("user_name", uname)).
			Warnf("create user forbidden for user.role_id=%d", roleID)
		return user.NewCreateUserDefault(http.StatusForbidden).
			WithPayload(&models.ErrorResponse{Message: "operation not permitted"})
	}
	// end check cookie and access rights

	// make password hash
	salt := MakeSalt(SaltLen)
	passwordHash := HashPass(salt, params.Body.Password)
	log.Debugf("found user: %+v", params.Body)
	// insert user to repo
	_user, err := cui.repo.StoreUser(params.HTTPRequest.Context(), params.Body, passwordHash)
	if err != nil {
		log.Errorf("repo.StoreUser error, %s", err)
		return user.NewCreateUserDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error happen"})
	}
	// all ok return new user
	return user.NewCreateUserCreated().WithPayload(_user)
}
