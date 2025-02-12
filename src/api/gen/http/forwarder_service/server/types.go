// Code generated by goa v3.19.1, DO NOT EDIT.
//
// ForwarderService HTTP server types
//
// Command:
// $ goa gen
// github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api/design
// --output ./src/api

package server

import (
	forwarderservice "github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api/gen/forwarder_service"
	goa "goa.design/goa/v3/pkg"
)

// ForwardRequestBody is the type of the "ForwarderService" service "forward"
// endpoint HTTP request body.
type ForwardRequestBody struct {
	// The type of notification
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// The name of the event
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The description of the event
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// NewForwardNotification builds a ForwarderService service forward endpoint
// payload.
func NewForwardNotification(body *ForwardRequestBody) *forwarderservice.Notification {
	v := &forwarderservice.Notification{
		Type:        *body.Type,
		Name:        *body.Name,
		Description: *body.Description,
	}

	return v
}

// ValidateForwardRequestBody runs the validations defined on ForwardRequestBody
func ValidateForwardRequestBody(body *ForwardRequestBody) (err error) {
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	return
}
