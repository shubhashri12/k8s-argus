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

// GetRecipientGroupByIDReader is a Reader for the GetRecipientGroupByID structure.
type GetRecipientGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecipientGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRecipientGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetRecipientGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRecipientGroupByIDOK creates a GetRecipientGroupByIDOK with default headers values
func NewGetRecipientGroupByIDOK() *GetRecipientGroupByIDOK {
	return &GetRecipientGroupByIDOK{}
}

/*GetRecipientGroupByIDOK handles this case with default header values.

successful operation
*/
type GetRecipientGroupByIDOK struct {
	Payload *models.RecipientGroup
}

func (o *GetRecipientGroupByIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/recipientgroups/{id}][%d] getRecipientGroupByIdOK  %+v", 200, o.Payload)
}

func (o *GetRecipientGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecipientGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecipientGroupByIDDefault creates a GetRecipientGroupByIDDefault with default headers values
func NewGetRecipientGroupByIDDefault(code int) *GetRecipientGroupByIDDefault {
	return &GetRecipientGroupByIDDefault{
		_statusCode: code,
	}
}

/*GetRecipientGroupByIDDefault handles this case with default header values.

Error
*/
type GetRecipientGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get recipient group by Id default response
func (o *GetRecipientGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetRecipientGroupByIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/recipientgroups/{id}][%d] getRecipientGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *GetRecipientGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}