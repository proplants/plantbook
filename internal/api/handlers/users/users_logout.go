package users

import (
	"fmt"
	"time"

	apimiddleware "github.com/kaatinga/plantbook/internal/api/middleware"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
)

type logoutUserImpl struct {
	expDelay time.Duration
}

// NewLoginUserHandler builder for user.LogoutUserHandler interface implementation
func NewLogoutUserHandler(expDelay time.Duration) user.LogoutUserHandler {
	return &logoutUserImpl{expDelay: expDelay}
}

// Handle implementation of the user.LogoutUserHandler interface
func (lui *logoutUserImpl) Handle(params user.LogoutUserParams) middleware.Responder {
	cookie := fmt.Sprintf("%s=%s; Expires=%s; Path=/",
		apimiddleware.JWTCookieName, "user logout", time.Now().Add(-lui.expDelay).Format(timeRFC7231))
	return user.NewLogoutUserOK().WithSetCookie(cookie).WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}
