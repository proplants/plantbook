// Code generated by go-swagger; DO NOT EDIT.

package gardens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteUserGardenParams creates a new DeleteUserGardenParams object
//
// There are no default values defined in the spec.
func NewDeleteUserGardenParams() DeleteUserGardenParams {

	return DeleteUserGardenParams{}
}

// DeleteUserGardenParams contains all the bound params for the delete user garden operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteUserGarden
type DeleteUserGardenParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Numeric ID of the user's garden to delete.
	  Required: true
	  In: path
	*/
	GardenID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteUserGardenParams() beforehand.
func (o *DeleteUserGardenParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGardenID, rhkGardenID, _ := route.Params.GetOK("garden_id")
	if err := o.bindGardenID(rGardenID, rhkGardenID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGardenID binds and validates parameter GardenID from path.
func (o *DeleteUserGardenParams) bindGardenID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("garden_id", "path", "int64", raw)
	}
	o.GardenID = value

	return nil
}
