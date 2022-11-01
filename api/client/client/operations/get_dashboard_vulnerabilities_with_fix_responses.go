// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openclarity/kubeclarity/api/v2/client/models"
)

// GetDashboardVulnerabilitiesWithFixReader is a Reader for the GetDashboardVulnerabilitiesWithFix structure.
type GetDashboardVulnerabilitiesWithFixReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardVulnerabilitiesWithFixReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardVulnerabilitiesWithFixOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDashboardVulnerabilitiesWithFixDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDashboardVulnerabilitiesWithFixOK creates a GetDashboardVulnerabilitiesWithFixOK with default headers values
func NewGetDashboardVulnerabilitiesWithFixOK() *GetDashboardVulnerabilitiesWithFixOK {
	return &GetDashboardVulnerabilitiesWithFixOK{}
}

/* GetDashboardVulnerabilitiesWithFixOK describes a response with status code 200, with default header values.

Success
*/
type GetDashboardVulnerabilitiesWithFixOK struct {
	Payload []*models.VulnerabilitiesWithFix
}

func (o *GetDashboardVulnerabilitiesWithFixOK) Error() string {
	return fmt.Sprintf("[GET /dashboard/vulnerabilitiesWithFix][%d] getDashboardVulnerabilitiesWithFixOK  %+v", 200, o.Payload)
}
func (o *GetDashboardVulnerabilitiesWithFixOK) GetPayload() []*models.VulnerabilitiesWithFix {
	return o.Payload
}

func (o *GetDashboardVulnerabilitiesWithFixOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVulnerabilitiesWithFixDefault creates a GetDashboardVulnerabilitiesWithFixDefault with default headers values
func NewGetDashboardVulnerabilitiesWithFixDefault(code int) *GetDashboardVulnerabilitiesWithFixDefault {
	return &GetDashboardVulnerabilitiesWithFixDefault{
		_statusCode: code,
	}
}

/* GetDashboardVulnerabilitiesWithFixDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetDashboardVulnerabilitiesWithFixDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get dashboard vulnerabilities with fix default response
func (o *GetDashboardVulnerabilitiesWithFixDefault) Code() int {
	return o._statusCode
}

func (o *GetDashboardVulnerabilitiesWithFixDefault) Error() string {
	return fmt.Sprintf("[GET /dashboard/vulnerabilitiesWithFix][%d] GetDashboardVulnerabilitiesWithFix default  %+v", o._statusCode, o.Payload)
}
func (o *GetDashboardVulnerabilitiesWithFixDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetDashboardVulnerabilitiesWithFixDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
