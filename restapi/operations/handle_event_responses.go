// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// HandleEventOKCode is the HTTP code returned for type HandleEventOK
const HandleEventOKCode int = 200

/*HandleEventOK okay status

swagger:response handleEventOK
*/
type HandleEventOK struct {
}

// NewHandleEventOK creates HandleEventOK with default headers values
func NewHandleEventOK() *HandleEventOK {

	return &HandleEventOK{}
}

// WriteResponse to the client
func (o *HandleEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// HandleEventUnauthorizedCode is the HTTP code returned for type HandleEventUnauthorized
const HandleEventUnauthorizedCode int = 401

/*HandleEventUnauthorized bad secret key

swagger:response handleEventUnauthorized
*/
type HandleEventUnauthorized struct {
}

// NewHandleEventUnauthorized creates HandleEventUnauthorized with default headers values
func NewHandleEventUnauthorized() *HandleEventUnauthorized {

	return &HandleEventUnauthorized{}
}

// WriteResponse to the client
func (o *HandleEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// HandleEventInternalServerErrorCode is the HTTP code returned for type HandleEventInternalServerError
const HandleEventInternalServerErrorCode int = 500

/*HandleEventInternalServerError internal server error

swagger:response handleEventInternalServerError
*/
type HandleEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewHandleEventInternalServerError creates HandleEventInternalServerError with default headers values
func NewHandleEventInternalServerError() *HandleEventInternalServerError {

	return &HandleEventInternalServerError{}
}

// WithPayload adds the payload to the handle event internal server error response
func (o *HandleEventInternalServerError) WithPayload(payload string) *HandleEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the handle event internal server error response
func (o *HandleEventInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HandleEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}