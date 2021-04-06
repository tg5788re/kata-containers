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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kata-containers/kata-containers/src/runtime/virtcontainers/pkg/firecracker/client/models"
)

// NewPatchVMParams creates a new PatchVMParams object
// with the default values initialized.
func NewPatchVMParams() *PatchVMParams {
	var ()
	return &PatchVMParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchVMParamsWithTimeout creates a new PatchVMParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchVMParamsWithTimeout(timeout time.Duration) *PatchVMParams {
	var ()
	return &PatchVMParams{

		timeout: timeout,
	}
}

// NewPatchVMParamsWithContext creates a new PatchVMParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchVMParamsWithContext(ctx context.Context) *PatchVMParams {
	var ()
	return &PatchVMParams{

		Context: ctx,
	}
}

// NewPatchVMParamsWithHTTPClient creates a new PatchVMParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchVMParamsWithHTTPClient(client *http.Client) *PatchVMParams {
	var ()
	return &PatchVMParams{
		HTTPClient: client,
	}
}

/*PatchVMParams contains all the parameters to send to the API endpoint
for the patch Vm operation typically these are written to a http.Request
*/
type PatchVMParams struct {

	/*Body
	  The microVM state

	*/
	Body *models.VM

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch Vm params
func (o *PatchVMParams) WithTimeout(timeout time.Duration) *PatchVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch Vm params
func (o *PatchVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch Vm params
func (o *PatchVMParams) WithContext(ctx context.Context) *PatchVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch Vm params
func (o *PatchVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch Vm params
func (o *PatchVMParams) WithHTTPClient(client *http.Client) *PatchVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch Vm params
func (o *PatchVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch Vm params
func (o *PatchVMParams) WithBody(body *models.VM) *PatchVMParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch Vm params
func (o *PatchVMParams) SetBody(body *models.VM) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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