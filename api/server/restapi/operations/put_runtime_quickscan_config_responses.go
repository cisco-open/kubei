// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openclarity/kubeclarity/api/v2/server/models"
)

// PutRuntimeQuickscanConfigCreatedCode is the HTTP code returned for type PutRuntimeQuickscanConfigCreated
const PutRuntimeQuickscanConfigCreatedCode int = 201

/*PutRuntimeQuickscanConfigCreated Success

swagger:response putRuntimeQuickscanConfigCreated
*/
type PutRuntimeQuickscanConfigCreated struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPutRuntimeQuickscanConfigCreated creates PutRuntimeQuickscanConfigCreated with default headers values
func NewPutRuntimeQuickscanConfigCreated() *PutRuntimeQuickscanConfigCreated {

	return &PutRuntimeQuickscanConfigCreated{}
}

// WithPayload adds the payload to the put runtime quickscan config created response
func (o *PutRuntimeQuickscanConfigCreated) WithPayload(payload interface{}) *PutRuntimeQuickscanConfigCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put runtime quickscan config created response
func (o *PutRuntimeQuickscanConfigCreated) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutRuntimeQuickscanConfigCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PutRuntimeQuickscanConfigBadRequestCode is the HTTP code returned for type PutRuntimeQuickscanConfigBadRequest
const PutRuntimeQuickscanConfigBadRequestCode int = 400

/*PutRuntimeQuickscanConfigBadRequest Failed to set quick scan config

swagger:response putRuntimeQuickscanConfigBadRequest
*/
type PutRuntimeQuickscanConfigBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPutRuntimeQuickscanConfigBadRequest creates PutRuntimeQuickscanConfigBadRequest with default headers values
func NewPutRuntimeQuickscanConfigBadRequest() *PutRuntimeQuickscanConfigBadRequest {

	return &PutRuntimeQuickscanConfigBadRequest{}
}

// WithPayload adds the payload to the put runtime quickscan config bad request response
func (o *PutRuntimeQuickscanConfigBadRequest) WithPayload(payload string) *PutRuntimeQuickscanConfigBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put runtime quickscan config bad request response
func (o *PutRuntimeQuickscanConfigBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutRuntimeQuickscanConfigBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PutRuntimeQuickscanConfigDefault unknown error

swagger:response putRuntimeQuickscanConfigDefault
*/
type PutRuntimeQuickscanConfigDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPutRuntimeQuickscanConfigDefault creates PutRuntimeQuickscanConfigDefault with default headers values
func NewPutRuntimeQuickscanConfigDefault(code int) *PutRuntimeQuickscanConfigDefault {
	if code <= 0 {
		code = 500
	}

	return &PutRuntimeQuickscanConfigDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put runtime quickscan config default response
func (o *PutRuntimeQuickscanConfigDefault) WithStatusCode(code int) *PutRuntimeQuickscanConfigDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put runtime quickscan config default response
func (o *PutRuntimeQuickscanConfigDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put runtime quickscan config default response
func (o *PutRuntimeQuickscanConfigDefault) WithPayload(payload *models.APIResponse) *PutRuntimeQuickscanConfigDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put runtime quickscan config default response
func (o *PutRuntimeQuickscanConfigDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutRuntimeQuickscanConfigDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
