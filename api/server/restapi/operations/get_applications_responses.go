// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openclarity/kubeclarity/api/v2/server/models"
)

// GetApplicationsOKCode is the HTTP code returned for type GetApplicationsOK
const GetApplicationsOKCode int = 200

/*GetApplicationsOK Success

swagger:response getApplicationsOK
*/
type GetApplicationsOK struct {

	/*
	  In: Body
	*/
	Payload *GetApplicationsOKBody `json:"body,omitempty"`
}

// NewGetApplicationsOK creates GetApplicationsOK with default headers values
func NewGetApplicationsOK() *GetApplicationsOK {

	return &GetApplicationsOK{}
}

// WithPayload adds the payload to the get applications o k response
func (o *GetApplicationsOK) WithPayload(payload *GetApplicationsOKBody) *GetApplicationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get applications o k response
func (o *GetApplicationsOK) SetPayload(payload *GetApplicationsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApplicationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetApplicationsDefault unknown error

swagger:response getApplicationsDefault
*/
type GetApplicationsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetApplicationsDefault creates GetApplicationsDefault with default headers values
func NewGetApplicationsDefault(code int) *GetApplicationsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetApplicationsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get applications default response
func (o *GetApplicationsDefault) WithStatusCode(code int) *GetApplicationsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get applications default response
func (o *GetApplicationsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get applications default response
func (o *GetApplicationsDefault) WithPayload(payload *models.APIResponse) *GetApplicationsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get applications default response
func (o *GetApplicationsDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApplicationsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
