// Code generated by goa v3.1.1, DO NOT EDIT.
//
// calc service
//
// Command:
// $ goa gen calc/design

package calc

import (
	"context"
)

// The calc service performs operations on numbers.
type Service interface {
	// Add implements add.
	Add(context.Context, *AddPayload) (res int, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "calc"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"add"}

// AddPayload is the payload type of the calc service add method.
type AddPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}