// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openclarity/kubeclarity/api/v2/server/models"
)

// GetPackagesOKCode is the HTTP code returned for type GetPackagesOK
const GetPackagesOKCode int = 200

/*GetPackagesOK Success

swagger:response getPackagesOK
*/
type GetPackagesOK struct {

	/*
	  In: Body
	*/
	Payload *GetPackagesOKBody `json:"body,omitempty"`
}

// NewGetPackagesOK creates GetPackagesOK with default headers values
func NewGetPackagesOK() *GetPackagesOK {

	return &GetPackagesOK{}
}

// WithPayload adds the payload to the get packages o k response
func (o *GetPackagesOK) WithPayload(payload *GetPackagesOKBody) *GetPackagesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get packages o k response
func (o *GetPackagesOK) SetPayload(payload *GetPackagesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPackagesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetPackagesDefault unknown error

swagger:response getPackagesDefault
*/
type GetPackagesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetPackagesDefault creates GetPackagesDefault with default headers values
func NewGetPackagesDefault(code int) *GetPackagesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetPackagesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get packages default response
func (o *GetPackagesDefault) WithStatusCode(code int) *GetPackagesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get packages default response
func (o *GetPackagesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get packages default response
func (o *GetPackagesDefault) WithPayload(payload *models.APIResponse) *GetPackagesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get packages default response
func (o *GetPackagesDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPackagesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
