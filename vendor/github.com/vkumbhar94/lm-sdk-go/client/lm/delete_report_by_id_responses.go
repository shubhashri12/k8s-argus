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

// DeleteReportByIDReader is a Reader for the DeleteReportByID structure.
type DeleteReportByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReportByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteReportByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteReportByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteReportByIDOK creates a DeleteReportByIDOK with default headers values
func NewDeleteReportByIDOK() *DeleteReportByIDOK {
	return &DeleteReportByIDOK{}
}

/*DeleteReportByIDOK handles this case with default header values.

successful operation
*/
type DeleteReportByIDOK struct {
	Payload interface{}
}

func (o *DeleteReportByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /report/reports/{id}][%d] deleteReportByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteReportByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReportByIDDefault creates a DeleteReportByIDDefault with default headers values
func NewDeleteReportByIDDefault(code int) *DeleteReportByIDDefault {
	return &DeleteReportByIDDefault{
		_statusCode: code,
	}
}

/*DeleteReportByIDDefault handles this case with default header values.

Error
*/
type DeleteReportByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete report by Id default response
func (o *DeleteReportByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteReportByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /report/reports/{id}][%d] deleteReportById default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteReportByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}