package health

import (
	"net/http"

	apimiddleware "github.com/kaatinga/plantbook/internal/api/middleware"
	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/health"
	"github.com/kaatinga/plantbook/pkg/logging"

	"github.com/go-openapi/runtime/middleware"
)

type healthReadyImpl struct {
	storage RepoInterface
}

// NewHealthReadyHandler builder for user.LoginUserHandler interface implementation.
func NewHealthReadyHandler(repo RepoInterface) health.HealthReadyHandler {
	return &healthReadyImpl{storage: repo}
}

// Handle implementation of the health.HealthReadyHandler interface.
func (hr *healthReadyImpl) Handle(params health.HealthReadyParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	err := hr.storage.Health(params.HTTPRequest.Context())
	if err != nil {
		log.Errorf("repo.Health error, %s", err)
		return health.NewHealthReadyDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error happen"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	return health.NewHealthReadyOK().WithPayload("OK").
		WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}
