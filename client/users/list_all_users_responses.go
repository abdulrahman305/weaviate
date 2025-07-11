//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// ListAllUsersReader is a Reader for the ListAllUsers structure.
type ListAllUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAllUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAllUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAllUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllUsersOK creates a ListAllUsersOK with default headers values
func NewListAllUsersOK() *ListAllUsersOK {
	return &ListAllUsersOK{}
}

/*
ListAllUsersOK describes a response with status code 200, with default header values.

Info about the users
*/
type ListAllUsersOK struct {
	Payload []*models.DBUserInfo
}

// IsSuccess returns true when this list all users o k response has a 2xx status code
func (o *ListAllUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list all users o k response has a 3xx status code
func (o *ListAllUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list all users o k response has a 4xx status code
func (o *ListAllUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list all users o k response has a 5xx status code
func (o *ListAllUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list all users o k response a status code equal to that given
func (o *ListAllUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list all users o k response
func (o *ListAllUsersOK) Code() int {
	return 200
}

func (o *ListAllUsersOK) Error() string {
	return fmt.Sprintf("[GET /users/db][%d] listAllUsersOK  %+v", 200, o.Payload)
}

func (o *ListAllUsersOK) String() string {
	return fmt.Sprintf("[GET /users/db][%d] listAllUsersOK  %+v", 200, o.Payload)
}

func (o *ListAllUsersOK) GetPayload() []*models.DBUserInfo {
	return o.Payload
}

func (o *ListAllUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllUsersUnauthorized creates a ListAllUsersUnauthorized with default headers values
func NewListAllUsersUnauthorized() *ListAllUsersUnauthorized {
	return &ListAllUsersUnauthorized{}
}

/*
ListAllUsersUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type ListAllUsersUnauthorized struct {
}

// IsSuccess returns true when this list all users unauthorized response has a 2xx status code
func (o *ListAllUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list all users unauthorized response has a 3xx status code
func (o *ListAllUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list all users unauthorized response has a 4xx status code
func (o *ListAllUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list all users unauthorized response has a 5xx status code
func (o *ListAllUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list all users unauthorized response a status code equal to that given
func (o *ListAllUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list all users unauthorized response
func (o *ListAllUsersUnauthorized) Code() int {
	return 401
}

func (o *ListAllUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/db][%d] listAllUsersUnauthorized ", 401)
}

func (o *ListAllUsersUnauthorized) String() string {
	return fmt.Sprintf("[GET /users/db][%d] listAllUsersUnauthorized ", 401)
}

func (o *ListAllUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAllUsersForbidden creates a ListAllUsersForbidden with default headers values
func NewListAllUsersForbidden() *ListAllUsersForbidden {
	return &ListAllUsersForbidden{}
}

/*
ListAllUsersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListAllUsersForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list all users forbidden response has a 2xx status code
func (o *ListAllUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list all users forbidden response has a 3xx status code
func (o *ListAllUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list all users forbidden response has a 4xx status code
func (o *ListAllUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list all users forbidden response has a 5xx status code
func (o *ListAllUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list all users forbidden response a status code equal to that given
func (o *ListAllUsersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list all users forbidden response
func (o *ListAllUsersForbidden) Code() int {
	return 403
}

func (o *ListAllUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /users/db][%d] listAllUsersForbidden  %+v", 403, o.Payload)
}

func (o *ListAllUsersForbidden) String() string {
	return fmt.Sprintf("[GET /users/db][%d] listAllUsersForbidden  %+v", 403, o.Payload)
}

func (o *ListAllUsersForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAllUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllUsersInternalServerError creates a ListAllUsersInternalServerError with default headers values
func NewListAllUsersInternalServerError() *ListAllUsersInternalServerError {
	return &ListAllUsersInternalServerError{}
}

/*
ListAllUsersInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ListAllUsersInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list all users internal server error response has a 2xx status code
func (o *ListAllUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list all users internal server error response has a 3xx status code
func (o *ListAllUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list all users internal server error response has a 4xx status code
func (o *ListAllUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list all users internal server error response has a 5xx status code
func (o *ListAllUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list all users internal server error response a status code equal to that given
func (o *ListAllUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list all users internal server error response
func (o *ListAllUsersInternalServerError) Code() int {
	return 500
}

func (o *ListAllUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/db][%d] listAllUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAllUsersInternalServerError) String() string {
	return fmt.Sprintf("[GET /users/db][%d] listAllUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAllUsersInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAllUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
