package refplants

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/refplant"
)

type GetRefPlantByIDImpl struct {
	repo RepoInterface
}

func NewGetRefPlantByIDHandler(repo RepoInterface) refplant.GetRefPlantByIDHandler {
	return &GetRefPlantByIDImpl{repo: repo}
}

func (impl *GetRefPlantByIDImpl) Handle(params refplant.GetRefPlantByIDParams) middleware.Responder {
	oneRefPlant, err := impl.repo.GetRefPlantByID(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		return refplant.NewGetRefPlantByIDDefault(http.StatusInternalServerError).
			WithPayload(&models.ErrorResponse{Message: err.Error()})
	}
	return refplant.NewGetRefPlantByIDOK().WithPayload(oneRefPlant)
}
