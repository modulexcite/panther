// Code generated by go-swagger; DO NOT EDIT.

// Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
// Copyright (C) 2020 Panther Labs Inc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/panther-labs/panther/api/gateway/analysis/models"
)

// GetEnabledPoliciesReader is a Reader for the GetEnabledPolicies structure.
type GetEnabledPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnabledPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEnabledPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEnabledPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEnabledPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEnabledPoliciesOK creates a GetEnabledPoliciesOK with default headers values
func NewGetEnabledPoliciesOK() *GetEnabledPoliciesOK {
	return &GetEnabledPoliciesOK{}
}

/*GetEnabledPoliciesOK handles this case with default header values.

OK
*/
type GetEnabledPoliciesOK struct {
	Payload *models.EnabledPolicies
}

func (o *GetEnabledPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /enabled][%d] getEnabledPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetEnabledPoliciesOK) GetPayload() *models.EnabledPolicies {
	return o.Payload
}

func (o *GetEnabledPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnabledPolicies)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnabledPoliciesBadRequest creates a GetEnabledPoliciesBadRequest with default headers values
func NewGetEnabledPoliciesBadRequest() *GetEnabledPoliciesBadRequest {
	return &GetEnabledPoliciesBadRequest{}
}

/*GetEnabledPoliciesBadRequest handles this case with default header values.

Bad request
*/
type GetEnabledPoliciesBadRequest struct {
	Payload *models.Error
}

func (o *GetEnabledPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /enabled][%d] getEnabledPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *GetEnabledPoliciesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEnabledPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnabledPoliciesInternalServerError creates a GetEnabledPoliciesInternalServerError with default headers values
func NewGetEnabledPoliciesInternalServerError() *GetEnabledPoliciesInternalServerError {
	return &GetEnabledPoliciesInternalServerError{}
}

/*GetEnabledPoliciesInternalServerError handles this case with default header values.

Internal server error
*/
type GetEnabledPoliciesInternalServerError struct {
}

func (o *GetEnabledPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /enabled][%d] getEnabledPoliciesInternalServerError ", 500)
}

func (o *GetEnabledPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
