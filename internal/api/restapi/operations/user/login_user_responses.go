// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kaatinga/plantbook/internal/api/models"
)

// LoginUserOKCode is the HTTP code returned for type LoginUserOK
const LoginUserOKCode int = 200

/*LoginUserOK successful operation

swagger:response loginUserOK
*/
type LoginUserOK struct {
	/*set cookie with jwt token value with name plantbook_token

	 */
	SetCookie string `json:"Set-Cookie"`
	/*The request id this is a response to

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewLoginUserOK creates LoginUserOK with default headers values
func NewLoginUserOK() *LoginUserOK {

	return &LoginUserOK{}
}

// WithSetCookie adds the setCookie to the login user o k response
func (o *LoginUserOK) WithSetCookie(setCookie string) *LoginUserOK {
	o.SetCookie = setCookie
	return o
}

// SetSetCookie sets the setCookie to the login user o k response
func (o *LoginUserOK) SetSetCookie(setCookie string) {
	o.SetCookie = setCookie
}

// WithXRequestID adds the xRequestId to the login user o k response
func (o *LoginUserOK) WithXRequestID(xRequestID string) *LoginUserOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the login user o k response
func (o *LoginUserOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *LoginUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Set-Cookie

	setCookie := o.SetCookie
	if setCookie != "" {
		rw.Header().Set("Set-Cookie", setCookie)
	}

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*LoginUserDefault unexpected error

swagger:response loginUserDefault
*/
type LoginUserDefault struct {
	_statusCode int
	/*The request id this is a response to

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewLoginUserDefault creates LoginUserDefault with default headers values
func NewLoginUserDefault(code int) *LoginUserDefault {
	if code <= 0 {
		code = 500
	}

	return &LoginUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the login user default response
func (o *LoginUserDefault) WithStatusCode(code int) *LoginUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the login user default response
func (o *LoginUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithXRequestID adds the xRequestId to the login user default response
func (o *LoginUserDefault) WithXRequestID(xRequestID string) *LoginUserDefault {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the login user default response
func (o *LoginUserDefault) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the login user default response
func (o *LoginUserDefault) WithPayload(payload *models.ErrorResponse) *LoginUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user default response
func (o *LoginUserDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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