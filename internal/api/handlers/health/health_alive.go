package health

import (
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/health"

	"github.com/go-openapi/runtime/middleware"
)

type healthAliveImpl struct{}

// NewLoginUserHandler builder for user.LoginUserHandler interface implementation
func NewHealthAliveHandler() health.HealthAliveHandler {
	return &healthAliveImpl{}
}

func (ha *healthAliveImpl) Handle(params health.HealthAliveParams) middleware.Responder {
	return health.NewHealthAliveOK().WithPayload("OK")
}
