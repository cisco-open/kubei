// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openclarity/kubeclarity/shared/v2/pkg/scanner/dependency_track/api/client/models"
)

// NewUpdateProjectParams creates a new UpdateProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProjectParams() *UpdateProjectParams {
	return &UpdateProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProjectParamsWithTimeout creates a new UpdateProjectParams object
// with the ability to set a timeout on a request.
func NewUpdateProjectParamsWithTimeout(timeout time.Duration) *UpdateProjectParams {
	return &UpdateProjectParams{
		timeout: timeout,
	}
}

// NewUpdateProjectParamsWithContext creates a new UpdateProjectParams object
// with the ability to set a context for a request.
func NewUpdateProjectParamsWithContext(ctx context.Context) *UpdateProjectParams {
	return &UpdateProjectParams{
		Context: ctx,
	}
}

// NewUpdateProjectParamsWithHTTPClient creates a new UpdateProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProjectParamsWithHTTPClient(client *http.Client) *UpdateProjectParams {
	return &UpdateProjectParams{
		HTTPClient: client,
	}
}

/* UpdateProjectParams contains all the parameters to send to the API endpoint
   for the update project operation.

   Typically these are written to a http.Request.
*/
type UpdateProjectParams struct {

	// Body.
	Body *models.Project

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProjectParams) WithDefaults() *UpdateProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update project params
func (o *UpdateProjectParams) WithTimeout(timeout time.Duration) *UpdateProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update project params
func (o *UpdateProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update project params
func (o *UpdateProjectParams) WithContext(ctx context.Context) *UpdateProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update project params
func (o *UpdateProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update project params
func (o *UpdateProjectParams) WithHTTPClient(client *http.Client) *UpdateProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update project params
func (o *UpdateProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update project params
func (o *UpdateProjectParams) WithBody(body *models.Project) *UpdateProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update project params
func (o *UpdateProjectParams) SetBody(body *models.Project) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
