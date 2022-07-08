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

package log_forward

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteLogForwardAcceptedCode is the HTTP code returned for type DeleteLogForwardAccepted
const DeleteLogForwardAcceptedCode int = 202

/*DeleteLogForwardAccepted Configuration change accepted and reload requested

swagger:response deleteLogForwardAccepted
*/
type DeleteLogForwardAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteLogForwardAccepted creates DeleteLogForwardAccepted with default headers values
func NewDeleteLogForwardAccepted() *DeleteLogForwardAccepted {

	return &DeleteLogForwardAccepted{}
}

// WithReloadID adds the reloadId to the delete log forward accepted response
func (o *DeleteLogForwardAccepted) WithReloadID(reloadID string) *DeleteLogForwardAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete log forward accepted response
func (o *DeleteLogForwardAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteLogForwardAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteLogForwardNoContentCode is the HTTP code returned for type DeleteLogForwardNoContent
const DeleteLogForwardNoContentCode int = 204

/*DeleteLogForwardNoContent Log Forward deleted

swagger:response deleteLogForwardNoContent
*/
type DeleteLogForwardNoContent struct {
}

// NewDeleteLogForwardNoContent creates DeleteLogForwardNoContent with default headers values
func NewDeleteLogForwardNoContent() *DeleteLogForwardNoContent {

	return &DeleteLogForwardNoContent{}
}

// WriteResponse to the client
func (o *DeleteLogForwardNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteLogForwardNotFoundCode is the HTTP code returned for type DeleteLogForwardNotFound
const DeleteLogForwardNotFoundCode int = 404

/*DeleteLogForwardNotFound The specified resource was not found

swagger:response deleteLogForwardNotFound
*/
type DeleteLogForwardNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteLogForwardNotFound creates DeleteLogForwardNotFound with default headers values
func NewDeleteLogForwardNotFound() *DeleteLogForwardNotFound {

	return &DeleteLogForwardNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete log forward not found response
func (o *DeleteLogForwardNotFound) WithConfigurationVersion(configurationVersion string) *DeleteLogForwardNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete log forward not found response
func (o *DeleteLogForwardNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete log forward not found response
func (o *DeleteLogForwardNotFound) WithPayload(payload *models.Error) *DeleteLogForwardNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete log forward not found response
func (o *DeleteLogForwardNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteLogForwardNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
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

/*DeleteLogForwardDefault General Error

swagger:response deleteLogForwardDefault
*/
type DeleteLogForwardDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteLogForwardDefault creates DeleteLogForwardDefault with default headers values
func NewDeleteLogForwardDefault(code int) *DeleteLogForwardDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteLogForwardDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete log forward default response
func (o *DeleteLogForwardDefault) WithStatusCode(code int) *DeleteLogForwardDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete log forward default response
func (o *DeleteLogForwardDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete log forward default response
func (o *DeleteLogForwardDefault) WithConfigurationVersion(configurationVersion string) *DeleteLogForwardDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete log forward default response
func (o *DeleteLogForwardDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete log forward default response
func (o *DeleteLogForwardDefault) WithPayload(payload *models.Error) *DeleteLogForwardDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete log forward default response
func (o *DeleteLogForwardDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteLogForwardDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
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
