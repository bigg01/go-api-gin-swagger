// Code generated by go-swagger; DO NOT EDIT.

package albums

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
	"github.com/go-openapi/swag"
)

// NewGetAlbAlbumsIDParams creates a new GetAlbAlbumsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlbAlbumsIDParams() *GetAlbAlbumsIDParams {
	return &GetAlbAlbumsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlbAlbumsIDParamsWithTimeout creates a new GetAlbAlbumsIDParams object
// with the ability to set a timeout on a request.
func NewGetAlbAlbumsIDParamsWithTimeout(timeout time.Duration) *GetAlbAlbumsIDParams {
	return &GetAlbAlbumsIDParams{
		timeout: timeout,
	}
}

// NewGetAlbAlbumsIDParamsWithContext creates a new GetAlbAlbumsIDParams object
// with the ability to set a context for a request.
func NewGetAlbAlbumsIDParamsWithContext(ctx context.Context) *GetAlbAlbumsIDParams {
	return &GetAlbAlbumsIDParams{
		Context: ctx,
	}
}

// NewGetAlbAlbumsIDParamsWithHTTPClient creates a new GetAlbAlbumsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlbAlbumsIDParamsWithHTTPClient(client *http.Client) *GetAlbAlbumsIDParams {
	return &GetAlbAlbumsIDParams{
		HTTPClient: client,
	}
}

/* GetAlbAlbumsIDParams contains all the parameters to send to the API endpoint
   for the get alb albums ID operation.

   Typically these are written to a http.Request.
*/
type GetAlbAlbumsIDParams struct {

	/* ID.

	   Album ID
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alb albums ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlbAlbumsIDParams) WithDefaults() *GetAlbAlbumsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alb albums ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlbAlbumsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get alb albums ID params
func (o *GetAlbAlbumsIDParams) WithTimeout(timeout time.Duration) *GetAlbAlbumsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alb albums ID params
func (o *GetAlbAlbumsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alb albums ID params
func (o *GetAlbAlbumsIDParams) WithContext(ctx context.Context) *GetAlbAlbumsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alb albums ID params
func (o *GetAlbAlbumsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alb albums ID params
func (o *GetAlbAlbumsIDParams) WithHTTPClient(client *http.Client) *GetAlbAlbumsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alb albums ID params
func (o *GetAlbAlbumsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get alb albums ID params
func (o *GetAlbAlbumsIDParams) WithID(id int64) *GetAlbAlbumsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get alb albums ID params
func (o *GetAlbAlbumsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlbAlbumsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
