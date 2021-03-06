// Code generated by go-swagger; DO NOT EDIT.

package federated_lte_networks

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

// NewGetFegLTENetworkIDSubscriberConfigRuleNamesParams creates a new GetFegLTENetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized.
func NewGetFegLTENetworkIDSubscriberConfigRuleNamesParams() *GetFegLTENetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &GetFegLTENetworkIDSubscriberConfigRuleNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFegLTENetworkIDSubscriberConfigRuleNamesParamsWithTimeout creates a new GetFegLTENetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFegLTENetworkIDSubscriberConfigRuleNamesParamsWithTimeout(timeout time.Duration) *GetFegLTENetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &GetFegLTENetworkIDSubscriberConfigRuleNamesParams{

		timeout: timeout,
	}
}

// NewGetFegLTENetworkIDSubscriberConfigRuleNamesParamsWithContext creates a new GetFegLTENetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFegLTENetworkIDSubscriberConfigRuleNamesParamsWithContext(ctx context.Context) *GetFegLTENetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &GetFegLTENetworkIDSubscriberConfigRuleNamesParams{

		Context: ctx,
	}
}

// NewGetFegLTENetworkIDSubscriberConfigRuleNamesParamsWithHTTPClient creates a new GetFegLTENetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFegLTENetworkIDSubscriberConfigRuleNamesParamsWithHTTPClient(client *http.Client) *GetFegLTENetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &GetFegLTENetworkIDSubscriberConfigRuleNamesParams{
		HTTPClient: client,
	}
}

/*GetFegLTENetworkIDSubscriberConfigRuleNamesParams contains all the parameters to send to the API endpoint
for the get feg LTE network ID subscriber config rule names operation typically these are written to a http.Request
*/
type GetFegLTENetworkIDSubscriberConfigRuleNamesParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get feg LTE network ID subscriber config rule names params
func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesParams) WithTimeout(timeout time.Duration) *GetFegLTENetworkIDSubscriberConfigRuleNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get feg LTE network ID subscriber config rule names params
func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get feg LTE network ID subscriber config rule names params
func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesParams) WithContext(ctx context.Context) *GetFegLTENetworkIDSubscriberConfigRuleNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get feg LTE network ID subscriber config rule names params
func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get feg LTE network ID subscriber config rule names params
func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesParams) WithHTTPClient(client *http.Client) *GetFegLTENetworkIDSubscriberConfigRuleNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get feg LTE network ID subscriber config rule names params
func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get feg LTE network ID subscriber config rule names params
func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesParams) WithNetworkID(networkID string) *GetFegLTENetworkIDSubscriberConfigRuleNamesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get feg LTE network ID subscriber config rule names params
func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFegLTENetworkIDSubscriberConfigRuleNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
