// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// AddOpsNoteReader is a Reader for the AddOpsNote structure.
type AddOpsNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOpsNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddOpsNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddOpsNoteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddOpsNoteOK creates a AddOpsNoteOK with default headers values
func NewAddOpsNoteOK() *AddOpsNoteOK {
	return &AddOpsNoteOK{}
}

/*AddOpsNoteOK handles this case with default header values.

successful operation
*/
type AddOpsNoteOK struct {
	Payload *models.OpsNote
}

func (o *AddOpsNoteOK) Error() string {
	return fmt.Sprintf("[POST /setting/opsnotes][%d] addOpsNoteOK  %+v", 200, o.Payload)
}

func (o *AddOpsNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpsNote)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOpsNoteDefault creates a AddOpsNoteDefault with default headers values
func NewAddOpsNoteDefault(code int) *AddOpsNoteDefault {
	return &AddOpsNoteDefault{
		_statusCode: code,
	}
}

/*AddOpsNoteDefault handles this case with default header values.

Error
*/
type AddOpsNoteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add ops note default response
func (o *AddOpsNoteDefault) Code() int {
	return o._statusCode
}

func (o *AddOpsNoteDefault) Error() string {
	return fmt.Sprintf("[POST /setting/opsnotes][%d] addOpsNote default  %+v", o._statusCode, o.Payload)
}

func (o *AddOpsNoteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}