package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewContainerWaitParams creates a new ContainerWaitParams object
// with the default values initialized.
func NewContainerWaitParams() *ContainerWaitParams {
	var ()
	return &ContainerWaitParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewContainerWaitParamsWithTimeout creates a new ContainerWaitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContainerWaitParamsWithTimeout(timeout time.Duration) *ContainerWaitParams {
	var ()
	return &ContainerWaitParams{

		requestTimeout: timeout,
	}
}

// NewContainerWaitParamsWithContext creates a new ContainerWaitParams object
// with the default values initialized, and the ability to set a context for a request
func NewContainerWaitParamsWithContext(ctx context.Context) *ContainerWaitParams {
	var ()
	return &ContainerWaitParams{

		Context: ctx,
	}
}

// NewContainerWaitParamsWithHTTPClient creates a new ContainerWaitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContainerWaitParamsWithHTTPClient(client *http.Client) *ContainerWaitParams {
	var ()
	return &ContainerWaitParams{
		HTTPClient: client,
	}
}

/*ContainerWaitParams contains all the parameters to send to the API endpoint
for the container wait operation typically these are written to a http.Request
*/
type ContainerWaitParams struct {

	/*OpID*/
	OpID *string
	/*ID*/
	ID string
	/*Timeout*/
	Timeout int64

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the container wait params
func (o *ContainerWaitParams) WithRequestTimeout(timeout time.Duration) *ContainerWaitParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the container wait params
func (o *ContainerWaitParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the container wait params
func (o *ContainerWaitParams) WithContext(ctx context.Context) *ContainerWaitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container wait params
func (o *ContainerWaitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container wait params
func (o *ContainerWaitParams) WithHTTPClient(client *http.Client) *ContainerWaitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container wait params
func (o *ContainerWaitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOpID adds the opID to the container wait params
func (o *ContainerWaitParams) WithOpID(opID *string) *ContainerWaitParams {
	o.SetOpID(opID)
	return o
}

// SetOpID adds the opId to the container wait params
func (o *ContainerWaitParams) SetOpID(opID *string) {
	o.OpID = opID
}

// WithID adds the id to the container wait params
func (o *ContainerWaitParams) WithID(id string) *ContainerWaitParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the container wait params
func (o *ContainerWaitParams) SetID(id string) {
	o.ID = id
}

// WithTimeout adds the timeout to the container wait params
func (o *ContainerWaitParams) WithTimeout(timeout int64) *ContainerWaitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container wait params
func (o *ContainerWaitParams) SetTimeout(timeout int64) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerWaitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.requestTimeout)
	var res []error

	if o.OpID != nil {

		// header param Op-ID
		if err := r.SetHeaderParam("Op-ID", *o.OpID); err != nil {
			return err
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param timeout
	qrTimeout := o.Timeout
	qTimeout := swag.FormatInt64(qrTimeout)
	if qTimeout != "" {
		if err := r.SetQueryParam("timeout", qTimeout); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
