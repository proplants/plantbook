package refplants

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	apimiddleware "github.com/proplants/plantbook/internal/api/middleware"
	"github.com/proplants/plantbook/internal/api/models"
	"github.com/proplants/plantbook/internal/api/restapi/operations/refplant"
	"github.com/proplants/plantbook/pkg/logging"
)

type getRefPlantsImpl struct {
	storage RepoInterface
}

// NewGetRefPlantsHandler builder for refplant.GetRefPlantsHandler interface implementation.
func NewGetRefPlantsHandler(repo RepoInterface) refplant.GetRefPlantsHandler {
	return &getRefPlantsImpl{storage: repo}
}

// Handle implementation of the refplant.GetRefPlantsHandler interface.
func (impl *getRefPlantsImpl) Handle(params refplant.GetRefPlantsParams) middleware.Responder {
	log := logging.FromContext(params.HTTPRequest.Context())
	RefPlants, count, total, err := impl.storage.GetRefPlants(params)
	if err != nil {
		log.Errorf("Handle GetRefPlants error, %s", err)
		return refplant.NewGetRefPlantsDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: "db error"}).
			WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
	}
	resultSet := models.ResultSet{
		Limit:  params.Limit,
		Offset: params.Offset,
		Count:  count,
		Total:  total,
	}
	return refplant.NewGetRefPlantsOK().WithPayload(&models.RefPlantsResponse{Data: RefPlants, ResultSet: &resultSet}).
		WithXRequestID(apimiddleware.GetRequestID(params.HTTPRequest))
}
