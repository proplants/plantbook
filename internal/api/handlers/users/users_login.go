package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/user"
	"github.com/kaatinga/plantbook/pkg/logging"
	"github.com/kaatinga/plantbook/pkg/token"

	"github.com/go-openapi/runtime/middleware"
)

const (
	timeRFC7231 string = "Mon, 02 Jan 2006 15:04:05 GMT"
)

type loginUserImpl struct {
	repo     RepoInterface
	tm       token.Manager
	expDelay time.Duration
}

// NewLoginUserHandler builder for user.LoginUserHandler interface implementation
func NewLoginUserHandler(repo RepoInterface, tm token.Manager, expDelay time.Duration) user.LoginUserHandler {
	return &loginUserImpl{repo: repo, tm: tm, expDelay: expDelay}
}

// Handle implementation of the user.LoginUserHandler interface
func (lui *loginUserImpl) Handle(params user.LoginUserParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	log.Debugf("login/password: %v", params.LoginPassword)
	_user, hash, err := lui.repo.FindUserByLogin(params.HTTPRequest.Context(), *params.LoginPassword.Login)
	if err != nil {
		return user.NewLoginUserDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error happen"})
	}
	ok := CheckPass(hash, *params.LoginPassword.Password)
	log.Debugf("found user: %+v hash: %s, equal pass %t", _user, hash, ok)
	if ok {
		tokenString, err := lui.tm.Create(_user.Username, _user.ID, time.Now().Add(lui.expDelay).Unix(), _user.UserRole)
		if err != nil {
			return user.NewLoginUserDefault(http.StatusInternalServerError).
				WithPayload(&models.ErrorResponse{Message: "make token error"})
		}
		cookie := fmt.Sprintf("%s=%s; Expires=%s; Path=/",
			jwtCookieName, tokenString, time.Now().Add(lui.expDelay).Format(timeRFC7231))
		return user.NewLoginUserOK().WithSetCookie(cookie)
	}
	return user.NewLoginUserDefault(http.StatusBadRequest).
		WithPayload(&models.ErrorResponse{Message: "invalid login or password"})
}
