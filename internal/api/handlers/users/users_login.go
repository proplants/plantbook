package users

import (
	"fmt"
	"net/http"
	"time"

	apimiddleware "github.com/kaatinga/plantbook/internal/api/middleware"
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
	storage  RepoInterface
	tm       token.Manager
	expDelay time.Duration
}

// NewLoginUserHandler builder for user.LoginUserHandler interface implementation.
func NewLoginUserHandler(repo RepoInterface, tm token.Manager, expDelay time.Duration) user.LoginUserHandler {
	return &loginUserImpl{storage: repo, tm: tm, expDelay: expDelay}
}

// Handle implementation of the user.LoginUserHandler interface.
func (lui *loginUserImpl) Handle(params user.LoginUserParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	log.Debugf("login/password: %v", params.LoginPassword)
	_user, hash, err := lui.storage.FindUserByLogin(params.HTTPRequest.Context(), *params.LoginPassword.Login)
	if err != nil {
		return user.NewLoginUserDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error happen"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	ok := CheckPass(hash, *params.LoginPassword.Password)
	log.Debugf("found user: %+v hash: %s, equal pass %t", _user, hash, ok)
	if ok {
		tokenString, err := lui.tm.Create(_user.Username, _user.ID, time.Now().Add(lui.expDelay).Unix(), _user.UserRole)
		if err != nil {
			return user.NewLoginUserDefault(http.StatusInternalServerError).
				WithPayload(&models.ErrorResponse{Message: "make token error"}).
				WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
		}
		cookie := fmt.Sprintf("%s=%s; Expires=%s; Path=/",
			apimiddleware.JWTCookieName, tokenString, time.Now().Add(lui.expDelay).Format(timeRFC7231))
		return user.NewLoginUserOK().WithSetCookie(cookie).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	return user.NewLoginUserDefault(http.StatusBadRequest).
		WithPayload(&models.ErrorResponse{Message: "invalid login or password"}).
		WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}
