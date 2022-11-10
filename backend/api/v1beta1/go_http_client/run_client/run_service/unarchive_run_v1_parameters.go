// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUnarchiveRunV1Params creates a new UnarchiveRunV1Params object
// with the default values initialized.
func NewUnarchiveRunV1Params() *UnarchiveRunV1Params {
	var ()
	return &UnarchiveRunV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnarchiveRunV1ParamsWithTimeout creates a new UnarchiveRunV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnarchiveRunV1ParamsWithTimeout(timeout time.Duration) *UnarchiveRunV1Params {
	var ()
	return &UnarchiveRunV1Params{

		timeout: timeout,
	}
}

// NewUnarchiveRunV1ParamsWithContext creates a new UnarchiveRunV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewUnarchiveRunV1ParamsWithContext(ctx context.Context) *UnarchiveRunV1Params {
	var ()
	return &UnarchiveRunV1Params{

		Context: ctx,
	}
}

// NewUnarchiveRunV1ParamsWithHTTPClient creates a new UnarchiveRunV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnarchiveRunV1ParamsWithHTTPClient(client *http.Client) *UnarchiveRunV1Params {
	var ()
	return &UnarchiveRunV1Params{
		HTTPClient: client,
	}
}

/*UnarchiveRunV1Params contains all the parameters to send to the API endpoint
for the unarchive run v1 operation typically these are written to a http.Request
*/
type UnarchiveRunV1Params struct {

	/*ID
	  The ID of the run to be restored.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the unarchive run v1 params
func (o *UnarchiveRunV1Params) WithTimeout(timeout time.Duration) *UnarchiveRunV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unarchive run v1 params
func (o *UnarchiveRunV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unarchive run v1 params
func (o *UnarchiveRunV1Params) WithContext(ctx context.Context) *UnarchiveRunV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unarchive run v1 params
func (o *UnarchiveRunV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unarchive run v1 params
func (o *UnarchiveRunV1Params) WithHTTPClient(client *http.Client) *UnarchiveRunV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unarchive run v1 params
func (o *UnarchiveRunV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the unarchive run v1 params
func (o *UnarchiveRunV1Params) WithID(id string) *UnarchiveRunV1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the unarchive run v1 params
func (o *UnarchiveRunV1Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UnarchiveRunV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}