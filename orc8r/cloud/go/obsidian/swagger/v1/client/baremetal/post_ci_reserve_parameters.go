// Code generated by go-swagger; DO NOT EDIT.

package baremetal

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

// NewPostCiReserveParams creates a new PostCiReserveParams object
// with the default values initialized.
func NewPostCiReserveParams() *PostCiReserveParams {
	var ()
	return &PostCiReserveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCiReserveParamsWithTimeout creates a new PostCiReserveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCiReserveParamsWithTimeout(timeout time.Duration) *PostCiReserveParams {
	var ()
	return &PostCiReserveParams{

		timeout: timeout,
	}
}

// NewPostCiReserveParamsWithContext creates a new PostCiReserveParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCiReserveParamsWithContext(ctx context.Context) *PostCiReserveParams {
	var ()
	return &PostCiReserveParams{

		Context: ctx,
	}
}

// NewPostCiReserveParamsWithHTTPClient creates a new PostCiReserveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCiReserveParamsWithHTTPClient(client *http.Client) *PostCiReserveParams {
	var ()
	return &PostCiReserveParams{
		HTTPClient: client,
	}
}

/*PostCiReserveParams contains all the parameters to send to the API endpoint
for the post ci reserve operation typically these are written to a http.Request
*/
type PostCiReserveParams struct {

	/*Tag
	  Optional tag to restrict reservation pool to

	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ci reserve params
func (o *PostCiReserveParams) WithTimeout(timeout time.Duration) *PostCiReserveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ci reserve params
func (o *PostCiReserveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ci reserve params
func (o *PostCiReserveParams) WithContext(ctx context.Context) *PostCiReserveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ci reserve params
func (o *PostCiReserveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ci reserve params
func (o *PostCiReserveParams) WithHTTPClient(client *http.Client) *PostCiReserveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ci reserve params
func (o *PostCiReserveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTag adds the tag to the post ci reserve params
func (o *PostCiReserveParams) WithTag(tag *string) *PostCiReserveParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the post ci reserve params
func (o *PostCiReserveParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *PostCiReserveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
