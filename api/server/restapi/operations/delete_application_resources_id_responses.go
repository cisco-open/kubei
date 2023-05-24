// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openclarity/kubeclarity/api/server/models"
)

// DeleteApplicationResourcesIDNoContentCode is the HTTP code returned for type DeleteApplicationResourcesIDNoContent
const DeleteApplicationResourcesIDNoContentCode int = 204

/*DeleteApplicationResourcesIDNoContent Success

swagger:response deleteApplicationResourcesIdNoContent
*/
type DeleteApplicationResourcesIDNoContent struct {
}

// NewDeleteApplicationResourcesIDNoContent creates DeleteApplicationResourcesIDNoContent with default headers values
func NewDeleteApplicationResourcesIDNoContent() *DeleteApplicationResourcesIDNoContent {

	return &DeleteApplicationResourcesIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteApplicationResourcesIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteApplicationResourcesIDNotFoundCode is the HTTP code returned for type DeleteApplicationResourcesIDNotFound
const DeleteApplicationResourcesIDNotFoundCode int = 404

/*DeleteApplicationResourcesIDNotFound Application not found.

swagger:response deleteApplicationResourcesIdNotFound
*/
type DeleteApplicationResourcesIDNotFound struct {
}

// NewDeleteApplicationResourcesIDNotFound creates DeleteApplicationResourcesIDNotFound with default headers values
func NewDeleteApplicationResourcesIDNotFound() *DeleteApplicationResourcesIDNotFound {

	return &DeleteApplicationResourcesIDNotFound{}
}

// WriteResponse to the client
func (o *DeleteApplicationResourcesIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*DeleteApplicationResourcesIDDefault unknown error

swagger:response deleteApplicationResourcesIdDefault
*/
type DeleteApplicationResourcesIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewDeleteApplicationResourcesIDDefault creates DeleteApplicationResourcesIDDefault with default headers values
func NewDeleteApplicationResourcesIDDefault(code int) *DeleteApplicationResourcesIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteApplicationResourcesIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete application resources ID default response
func (o *DeleteApplicationResourcesIDDefault) WithStatusCode(code int) *DeleteApplicationResourcesIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete application resources ID default response
func (o *DeleteApplicationResourcesIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete application resources ID default response
func (o *DeleteApplicationResourcesIDDefault) WithPayload(payload *models.APIResponse) *DeleteApplicationResourcesIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete application resources ID default response
func (o *DeleteApplicationResourcesIDDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteApplicationResourcesIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
