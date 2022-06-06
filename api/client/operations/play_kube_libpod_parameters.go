// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewPlayKubeLibpodParams creates a new PlayKubeLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPlayKubeLibpodParams() *PlayKubeLibpodParams {
	return &PlayKubeLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPlayKubeLibpodParamsWithTimeout creates a new PlayKubeLibpodParams object
// with the ability to set a timeout on a request.
func NewPlayKubeLibpodParamsWithTimeout(timeout time.Duration) *PlayKubeLibpodParams {
	return &PlayKubeLibpodParams{
		timeout: timeout,
	}
}

// NewPlayKubeLibpodParamsWithContext creates a new PlayKubeLibpodParams object
// with the ability to set a context for a request.
func NewPlayKubeLibpodParamsWithContext(ctx context.Context) *PlayKubeLibpodParams {
	return &PlayKubeLibpodParams{
		Context: ctx,
	}
}

// NewPlayKubeLibpodParamsWithHTTPClient creates a new PlayKubeLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewPlayKubeLibpodParamsWithHTTPClient(client *http.Client) *PlayKubeLibpodParams {
	return &PlayKubeLibpodParams{
		HTTPClient: client,
	}
}

/* PlayKubeLibpodParams contains all the parameters to send to the API endpoint
   for the play kube libpod operation.

   Typically these are written to a http.Request.
*/
type PlayKubeLibpodParams struct {

	/* LogDriver.

	   Logging driver for the containers in the pod.
	*/
	LogDriver *string

	/* Network.

	   USe the network mode or specify an array of networks.
	*/
	Network []string

	/* Request.

	   Kubernetes YAML file.
	*/
	Request string

	/* Start.

	   Start the pod after creating it.

	   Default: true
	*/
	Start *bool

	/* StaticIPs.

	   Static IPs used for the pods.
	*/
	StaticIPs []string

	/* StaticMACs.

	   Static MACs used for the pods.
	*/
	StaticMACs []string

	/* TLSVerify.

	   Require HTTPS and verify signatures when contacting registries.

	   Default: true
	*/
	TLSVerify *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the play kube libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlayKubeLibpodParams) WithDefaults() *PlayKubeLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the play kube libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlayKubeLibpodParams) SetDefaults() {
	var (
		startDefault = bool(true)

		tLSVerifyDefault = bool(true)
	)

	val := PlayKubeLibpodParams{
		Start:     &startDefault,
		TLSVerify: &tLSVerifyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the play kube libpod params
func (o *PlayKubeLibpodParams) WithTimeout(timeout time.Duration) *PlayKubeLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the play kube libpod params
func (o *PlayKubeLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the play kube libpod params
func (o *PlayKubeLibpodParams) WithContext(ctx context.Context) *PlayKubeLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the play kube libpod params
func (o *PlayKubeLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the play kube libpod params
func (o *PlayKubeLibpodParams) WithHTTPClient(client *http.Client) *PlayKubeLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the play kube libpod params
func (o *PlayKubeLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogDriver adds the logDriver to the play kube libpod params
func (o *PlayKubeLibpodParams) WithLogDriver(logDriver *string) *PlayKubeLibpodParams {
	o.SetLogDriver(logDriver)
	return o
}

// SetLogDriver adds the logDriver to the play kube libpod params
func (o *PlayKubeLibpodParams) SetLogDriver(logDriver *string) {
	o.LogDriver = logDriver
}

// WithNetwork adds the network to the play kube libpod params
func (o *PlayKubeLibpodParams) WithNetwork(network []string) *PlayKubeLibpodParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the play kube libpod params
func (o *PlayKubeLibpodParams) SetNetwork(network []string) {
	o.Network = network
}

// WithRequest adds the request to the play kube libpod params
func (o *PlayKubeLibpodParams) WithRequest(request string) *PlayKubeLibpodParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the play kube libpod params
func (o *PlayKubeLibpodParams) SetRequest(request string) {
	o.Request = request
}

// WithStart adds the start to the play kube libpod params
func (o *PlayKubeLibpodParams) WithStart(start *bool) *PlayKubeLibpodParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the play kube libpod params
func (o *PlayKubeLibpodParams) SetStart(start *bool) {
	o.Start = start
}

// WithStaticIPs adds the staticIPs to the play kube libpod params
func (o *PlayKubeLibpodParams) WithStaticIPs(staticIPs []string) *PlayKubeLibpodParams {
	o.SetStaticIPs(staticIPs)
	return o
}

// SetStaticIPs adds the staticIPs to the play kube libpod params
func (o *PlayKubeLibpodParams) SetStaticIPs(staticIPs []string) {
	o.StaticIPs = staticIPs
}

// WithStaticMACs adds the staticMACs to the play kube libpod params
func (o *PlayKubeLibpodParams) WithStaticMACs(staticMACs []string) *PlayKubeLibpodParams {
	o.SetStaticMACs(staticMACs)
	return o
}

// SetStaticMACs adds the staticMACs to the play kube libpod params
func (o *PlayKubeLibpodParams) SetStaticMACs(staticMACs []string) {
	o.StaticMACs = staticMACs
}

// WithTLSVerify adds the tLSVerify to the play kube libpod params
func (o *PlayKubeLibpodParams) WithTLSVerify(tLSVerify *bool) *PlayKubeLibpodParams {
	o.SetTLSVerify(tLSVerify)
	return o
}

// SetTLSVerify adds the tlsVerify to the play kube libpod params
func (o *PlayKubeLibpodParams) SetTLSVerify(tLSVerify *bool) {
	o.TLSVerify = tLSVerify
}

// WriteToRequest writes these params to a swagger request
func (o *PlayKubeLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LogDriver != nil {

		// query param logDriver
		var qrLogDriver string

		if o.LogDriver != nil {
			qrLogDriver = *o.LogDriver
		}
		qLogDriver := qrLogDriver
		if qLogDriver != "" {

			if err := r.SetQueryParam("logDriver", qLogDriver); err != nil {
				return err
			}
		}
	}

	if o.Network != nil {

		// binding items for network
		joinedNetwork := o.bindParamNetwork(reg)

		// query array param network
		if err := r.SetQueryParam("network", joinedNetwork...); err != nil {
			return err
		}
	}
	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart bool

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatBool(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if o.StaticIPs != nil {

		// binding items for staticIPs
		joinedStaticIPs := o.bindParamStaticIPs(reg)

		// query array param staticIPs
		if err := r.SetQueryParam("staticIPs", joinedStaticIPs...); err != nil {
			return err
		}
	}

	if o.StaticMACs != nil {

		// binding items for staticMACs
		joinedStaticMACs := o.bindParamStaticMACs(reg)

		// query array param staticMACs
		if err := r.SetQueryParam("staticMACs", joinedStaticMACs...); err != nil {
			return err
		}
	}

	if o.TLSVerify != nil {

		// query param tlsVerify
		var qrTLSVerify bool

		if o.TLSVerify != nil {
			qrTLSVerify = *o.TLSVerify
		}
		qTLSVerify := swag.FormatBool(qrTLSVerify)
		if qTLSVerify != "" {

			if err := r.SetQueryParam("tlsVerify", qTLSVerify); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPlayKubeLibpod binds the parameter network
func (o *PlayKubeLibpodParams) bindParamNetwork(formats strfmt.Registry) []string {
	networkIR := o.Network

	var networkIC []string
	for _, networkIIR := range networkIR { // explode []string

		networkIIV := networkIIR // string as string
		networkIC = append(networkIC, networkIIV)
	}

	// items.CollectionFormat: ""
	networkIS := swag.JoinByFormat(networkIC, "")

	return networkIS
}

// bindParamPlayKubeLibpod binds the parameter staticIPs
func (o *PlayKubeLibpodParams) bindParamStaticIPs(formats strfmt.Registry) []string {
	staticIPsIR := o.StaticIPs

	var staticIPsIC []string
	for _, staticIPsIIR := range staticIPsIR { // explode []string

		staticIPsIIV := staticIPsIIR // string as string
		staticIPsIC = append(staticIPsIC, staticIPsIIV)
	}

	// items.CollectionFormat: ""
	staticIPsIS := swag.JoinByFormat(staticIPsIC, "")

	return staticIPsIS
}

// bindParamPlayKubeLibpod binds the parameter staticMACs
func (o *PlayKubeLibpodParams) bindParamStaticMACs(formats strfmt.Registry) []string {
	staticMACsIR := o.StaticMACs

	var staticMACsIC []string
	for _, staticMACsIIR := range staticMACsIR { // explode []string

		staticMACsIIV := staticMACsIIR // string as string
		staticMACsIC = append(staticMACsIC, staticMACsIIV)
	}

	// items.CollectionFormat: ""
	staticMACsIS := swag.JoinByFormat(staticMACsIC, "")

	return staticMACsIS
}
