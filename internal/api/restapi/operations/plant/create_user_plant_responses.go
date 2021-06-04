// Code generated by go-swagger; DO NOT EDIT.

package plant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kaatinga/plantbook/internal/api/models"
)

// CreateUserPlantOKCode is the HTTP code returned for type CreateUserPlantOK
const CreateUserPlantOKCode int = 200

/*CreateUserPlantOK Plant added

swagger:response createUserPlantOK
*/
type CreateUserPlantOK struct {
	/*The request id this is a response to

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Plant `json:"body,omitempty"`
}

// NewCreateUserPlantOK creates CreateUserPlantOK with default headers values
func NewCreateUserPlantOK() *CreateUserPlantOK {

	return &CreateUserPlantOK{}
}

// WithXRequestID adds the xRequestId to the create user plant o k response
func (o *CreateUserPlantOK) WithXRequestID(xRequestID string) *CreateUserPlantOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create user plant o k response
func (o *CreateUserPlantOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create user plant o k response
func (o *CreateUserPlantOK) WithPayload(payload *models.Plant) *CreateUserPlantOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user plant o k response
func (o *CreateUserPlantOK) SetPayload(payload *models.Plant) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserPlantOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateUserPlantDefault unexpected error

swagger:response createUserPlantDefault
*/
type CreateUserPlantDefault struct {
	_statusCode int
	/*The request id this is a response to

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCreateUserPlantDefault creates CreateUserPlantDefault with default headers values
func NewCreateUserPlantDefault(code int) *CreateUserPlantDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateUserPlantDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create user plant default response
func (o *CreateUserPlantDefault) WithStatusCode(code int) *CreateUserPlantDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create user plant default response
func (o *CreateUserPlantDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithXRequestID adds the xRequestId to the create user plant default response
func (o *CreateUserPlantDefault) WithXRequestID(xRequestID string) *CreateUserPlantDefault {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create user plant default response
func (o *CreateUserPlantDefault) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create user plant default response
func (o *CreateUserPlantDefault) WithPayload(payload *models.ErrorResponse) *CreateUserPlantDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user plant default response
func (o *CreateUserPlantDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserPlantDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
