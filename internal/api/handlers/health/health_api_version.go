package health

import (
	"time"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/health"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type apiVersionImpl struct {
	version string
	githash string
	buildAt time.Time
}

// NewAPIVersionHandler builder for health.APIVersionHandler interface implementation
func NewAPIVersionHandler(version, githash string, buildAt time.Time) health.APIVersionHandler {
	return &apiVersionImpl{version: version, githash: githash, buildAt: buildAt}
}

func (av *apiVersionImpl) Handle(params health.APIVersionParams) middleware.Responder {
	return health.NewAPIVersionOK().WithPayload(&models.APIVersion{
		BuildAt: strfmt.DateTime(av.buildAt),
		Githash: av.githash,
		Version: av.version,
	})
}
