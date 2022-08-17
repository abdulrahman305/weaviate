//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ObjectsClassPatchReader is a Reader for the ObjectsClassPatch structure.
type ObjectsClassPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObjectsClassPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewObjectsClassPatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewObjectsClassPatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewObjectsClassPatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewObjectsClassPatchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewObjectsClassPatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewObjectsClassPatchUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewObjectsClassPatchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewObjectsClassPatchNoContent creates a ObjectsClassPatchNoContent with default headers values
func NewObjectsClassPatchNoContent() *ObjectsClassPatchNoContent {
	return &ObjectsClassPatchNoContent{}
}

/*ObjectsClassPatchNoContent handles this case with default header values.

Successfully applied. No content provided.
*/
type ObjectsClassPatchNoContent struct {
}

func (o *ObjectsClassPatchNoContent) Error() string {
	return fmt.Sprintf("[PATCH /objects/{className}/{id}][%d] objectsClassPatchNoContent ", 204)
}

func (o *ObjectsClassPatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsClassPatchBadRequest creates a ObjectsClassPatchBadRequest with default headers values
func NewObjectsClassPatchBadRequest() *ObjectsClassPatchBadRequest {
	return &ObjectsClassPatchBadRequest{}
}

/*ObjectsClassPatchBadRequest handles this case with default header values.

The patch-JSON is malformed.
*/
type ObjectsClassPatchBadRequest struct {
}

func (o *ObjectsClassPatchBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /objects/{className}/{id}][%d] objectsClassPatchBadRequest ", 400)
}

func (o *ObjectsClassPatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsClassPatchUnauthorized creates a ObjectsClassPatchUnauthorized with default headers values
func NewObjectsClassPatchUnauthorized() *ObjectsClassPatchUnauthorized {
	return &ObjectsClassPatchUnauthorized{}
}

/*ObjectsClassPatchUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ObjectsClassPatchUnauthorized struct {
}

func (o *ObjectsClassPatchUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /objects/{className}/{id}][%d] objectsClassPatchUnauthorized ", 401)
}

func (o *ObjectsClassPatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsClassPatchForbidden creates a ObjectsClassPatchForbidden with default headers values
func NewObjectsClassPatchForbidden() *ObjectsClassPatchForbidden {
	return &ObjectsClassPatchForbidden{}
}

/*ObjectsClassPatchForbidden handles this case with default header values.

Forbidden
*/
type ObjectsClassPatchForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsClassPatchForbidden) Error() string {
	return fmt.Sprintf("[PATCH /objects/{className}/{id}][%d] objectsClassPatchForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsClassPatchForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsClassPatchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsClassPatchNotFound creates a ObjectsClassPatchNotFound with default headers values
func NewObjectsClassPatchNotFound() *ObjectsClassPatchNotFound {
	return &ObjectsClassPatchNotFound{}
}

/*ObjectsClassPatchNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type ObjectsClassPatchNotFound struct {
}

func (o *ObjectsClassPatchNotFound) Error() string {
	return fmt.Sprintf("[PATCH /objects/{className}/{id}][%d] objectsClassPatchNotFound ", 404)
}

func (o *ObjectsClassPatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsClassPatchUnprocessableEntity creates a ObjectsClassPatchUnprocessableEntity with default headers values
func NewObjectsClassPatchUnprocessableEntity() *ObjectsClassPatchUnprocessableEntity {
	return &ObjectsClassPatchUnprocessableEntity{}
}

/*ObjectsClassPatchUnprocessableEntity handles this case with default header values.

The patch-JSON is valid but unprocessable.
*/
type ObjectsClassPatchUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsClassPatchUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /objects/{className}/{id}][%d] objectsClassPatchUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ObjectsClassPatchUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsClassPatchUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsClassPatchInternalServerError creates a ObjectsClassPatchInternalServerError with default headers values
func NewObjectsClassPatchInternalServerError() *ObjectsClassPatchInternalServerError {
	return &ObjectsClassPatchInternalServerError{}
}

/*ObjectsClassPatchInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ObjectsClassPatchInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsClassPatchInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /objects/{className}/{id}][%d] objectsClassPatchInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsClassPatchInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsClassPatchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
