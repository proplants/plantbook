// Code generated by go-swagger; DO NOT EDIT.

package userplant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserPlantsHandlerFunc turns a function with the right signature into a get user plants handler
type GetUserPlantsHandlerFunc func(GetUserPlantsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserPlantsHandlerFunc) Handle(params GetUserPlantsParams) middleware.Responder {
	return fn(params)
}

// GetUserPlantsHandler interface for that can handle valid get user plants params
type GetUserPlantsHandler interface {
	Handle(GetUserPlantsParams) middleware.Responder
}

// NewGetUserPlants creates a new http.Handler for the get user plants operation
func NewGetUserPlants(ctx *middleware.Context, handler GetUserPlantsHandler) *GetUserPlants {
	return &GetUserPlants{Context: ctx, Handler: handler}
}

/* GetUserPlants swagger:route GET /api/v1/user/plants userplant getUserPlants

Find all of the user’s plants

Find all of the user’s plants

*/
type GetUserPlants struct {
	Context *middleware.Context
	Handler GetUserPlantsHandler
}

func (o *GetUserPlants) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserPlantsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}