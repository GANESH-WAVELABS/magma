// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

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

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// NewPutCwfNetworkIDGatewaysGatewayIDMagmadParams creates a new PutCwfNetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized.
func NewPutCwfNetworkIDGatewaysGatewayIDMagmadParams() *PutCwfNetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutCwfNetworkIDGatewaysGatewayIDMagmadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDMagmadParamsWithTimeout creates a new PutCwfNetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCwfNetworkIDGatewaysGatewayIDMagmadParamsWithTimeout(timeout time.Duration) *PutCwfNetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutCwfNetworkIDGatewaysGatewayIDMagmadParams{

		timeout: timeout,
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDMagmadParamsWithContext creates a new PutCwfNetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCwfNetworkIDGatewaysGatewayIDMagmadParamsWithContext(ctx context.Context) *PutCwfNetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutCwfNetworkIDGatewaysGatewayIDMagmadParams{

		Context: ctx,
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDMagmadParamsWithHTTPClient creates a new PutCwfNetworkIDGatewaysGatewayIDMagmadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCwfNetworkIDGatewaysGatewayIDMagmadParamsWithHTTPClient(client *http.Client) *PutCwfNetworkIDGatewaysGatewayIDMagmadParams {
	var ()
	return &PutCwfNetworkIDGatewaysGatewayIDMagmadParams{
		HTTPClient: client,
	}
}

/*PutCwfNetworkIDGatewaysGatewayIDMagmadParams contains all the parameters to send to the API endpoint
for the put cwf network ID gateways gateway ID magmad operation typically these are written to a http.Request
*/
type PutCwfNetworkIDGatewaysGatewayIDMagmadParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*Magmad
	  New magmad configuration

	*/
	Magmad *models.MagmadGatewayConfigs
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cwf network ID gateways gateway ID magmad params
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) WithTimeout(timeout time.Duration) *PutCwfNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cwf network ID gateways gateway ID magmad params
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cwf network ID gateways gateway ID magmad params
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) WithContext(ctx context.Context) *PutCwfNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cwf network ID gateways gateway ID magmad params
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cwf network ID gateways gateway ID magmad params
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) WithHTTPClient(client *http.Client) *PutCwfNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cwf network ID gateways gateway ID magmad params
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the put cwf network ID gateways gateway ID magmad params
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) WithGatewayID(gatewayID string) *PutCwfNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put cwf network ID gateways gateway ID magmad params
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithMagmad adds the magmad to the put cwf network ID gateways gateway ID magmad params
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) WithMagmad(magmad *models.MagmadGatewayConfigs) *PutCwfNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetMagmad(magmad)
	return o
}

// SetMagmad adds the magmad to the put cwf network ID gateways gateway ID magmad params
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) SetMagmad(magmad *models.MagmadGatewayConfigs) {
	o.Magmad = magmad
}

// WithNetworkID adds the networkID to the put cwf network ID gateways gateway ID magmad params
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) WithNetworkID(networkID string) *PutCwfNetworkIDGatewaysGatewayIDMagmadParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put cwf network ID gateways gateway ID magmad params
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutCwfNetworkIDGatewaysGatewayIDMagmadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	if o.Magmad != nil {
		if err := r.SetBodyParam(o.Magmad); err != nil {
			return err
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
