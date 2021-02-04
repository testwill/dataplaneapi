// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package spoe_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models/v2"
)

// DeleteSpoeTransactionNoContentCode is the HTTP code returned for type DeleteSpoeTransactionNoContent
const DeleteSpoeTransactionNoContentCode int = 204

/*DeleteSpoeTransactionNoContent Transaction deleted

swagger:response deleteSpoeTransactionNoContent
*/
type DeleteSpoeTransactionNoContent struct {
}

// NewDeleteSpoeTransactionNoContent creates DeleteSpoeTransactionNoContent with default headers values
func NewDeleteSpoeTransactionNoContent() *DeleteSpoeTransactionNoContent {

	return &DeleteSpoeTransactionNoContent{}
}

// WriteResponse to the client
func (o *DeleteSpoeTransactionNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteSpoeTransactionNotFoundCode is the HTTP code returned for type DeleteSpoeTransactionNotFound
const DeleteSpoeTransactionNotFoundCode int = 404

/*DeleteSpoeTransactionNotFound The specified resource was not found

swagger:response deleteSpoeTransactionNotFound
*/
type DeleteSpoeTransactionNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSpoeTransactionNotFound creates DeleteSpoeTransactionNotFound with default headers values
func NewDeleteSpoeTransactionNotFound() *DeleteSpoeTransactionNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &DeleteSpoeTransactionNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the delete spoe transaction not found response
func (o *DeleteSpoeTransactionNotFound) WithConfigurationVersion(configurationVersion int64) *DeleteSpoeTransactionNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete spoe transaction not found response
func (o *DeleteSpoeTransactionNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete spoe transaction not found response
func (o *DeleteSpoeTransactionNotFound) WithPayload(payload *models.Error) *DeleteSpoeTransactionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete spoe transaction not found response
func (o *DeleteSpoeTransactionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSpoeTransactionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteSpoeTransactionDefault General Error

swagger:response deleteSpoeTransactionDefault
*/
type DeleteSpoeTransactionDefault struct {
	_statusCode int
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSpoeTransactionDefault creates DeleteSpoeTransactionDefault with default headers values
func NewDeleteSpoeTransactionDefault(code int) *DeleteSpoeTransactionDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &DeleteSpoeTransactionDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the delete spoe transaction default response
func (o *DeleteSpoeTransactionDefault) WithStatusCode(code int) *DeleteSpoeTransactionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete spoe transaction default response
func (o *DeleteSpoeTransactionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete spoe transaction default response
func (o *DeleteSpoeTransactionDefault) WithConfigurationVersion(configurationVersion int64) *DeleteSpoeTransactionDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete spoe transaction default response
func (o *DeleteSpoeTransactionDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete spoe transaction default response
func (o *DeleteSpoeTransactionDefault) WithPayload(payload *models.Error) *DeleteSpoeTransactionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete spoe transaction default response
func (o *DeleteSpoeTransactionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSpoeTransactionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}