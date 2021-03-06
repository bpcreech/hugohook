// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bpcreech/hugohook/models"
)

// NewHandleEventParams creates a new HandleEventParams object
// no default values defined in spec.
func NewHandleEventParams() HandleEventParams {

	return HandleEventParams{}
}

// HandleEventParams contains all the bound params for the handle event operation
// typically these are obtained from a http.Request
//
// swagger:parameters handleEvent
type HandleEventParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: header
	*/
	XGitHubDelivery *strfmt.UUID
	/*
	  Required: true
	  In: header
	*/
	XGitHubEvent string
	/*
	  Max Length: 45
	  Min Length: 45
	  In: header
	*/
	XHubSignature *string
	/*
	  In: body
	*/
	Event *models.Event
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewHandleEventParams() beforehand.
func (o *HandleEventParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXGitHubDelivery(r.Header[http.CanonicalHeaderKey("X-GitHub-Delivery")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXGitHubEvent(r.Header[http.CanonicalHeaderKey("X-GitHub-Event")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXHubSignature(r.Header[http.CanonicalHeaderKey("X-Hub-Signature")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Event
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("event", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Event = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXGitHubDelivery binds and validates parameter XGitHubDelivery from header.
func (o *HandleEventParams) bindXGitHubDelivery(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("X-GitHub-Delivery", "header", "strfmt.UUID", raw)
	}
	o.XGitHubDelivery = (value.(*strfmt.UUID))

	if err := o.validateXGitHubDelivery(formats); err != nil {
		return err
	}

	return nil
}

// validateXGitHubDelivery carries on validations for parameter XGitHubDelivery
func (o *HandleEventParams) validateXGitHubDelivery(formats strfmt.Registry) error {

	if err := validate.FormatOf("X-GitHub-Delivery", "header", "uuid", o.XGitHubDelivery.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindXGitHubEvent binds and validates parameter XGitHubEvent from header.
func (o *HandleEventParams) bindXGitHubEvent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-GitHub-Event", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-GitHub-Event", "header", raw); err != nil {
		return err
	}

	o.XGitHubEvent = raw

	if err := o.validateXGitHubEvent(formats); err != nil {
		return err
	}

	return nil
}

// validateXGitHubEvent carries on validations for parameter XGitHubEvent
func (o *HandleEventParams) validateXGitHubEvent(formats strfmt.Registry) error {

	if err := validate.Enum("X-GitHub-Event", "header", o.XGitHubEvent, []interface{}{"push", "ping"}); err != nil {
		return err
	}

	return nil
}

// bindXHubSignature binds and validates parameter XHubSignature from header.
func (o *HandleEventParams) bindXHubSignature(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XHubSignature = &raw

	if err := o.validateXHubSignature(formats); err != nil {
		return err
	}

	return nil
}

// validateXHubSignature carries on validations for parameter XHubSignature
func (o *HandleEventParams) validateXHubSignature(formats strfmt.Registry) error {

	if err := validate.MinLength("X-Hub-Signature", "header", (*o.XHubSignature), 45); err != nil {
		return err
	}

	if err := validate.MaxLength("X-Hub-Signature", "header", (*o.XHubSignature), 45); err != nil {
		return err
	}

	return nil
}
