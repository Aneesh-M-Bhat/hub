// Code generated by goa v3.15.0, DO NOT EDIT.
//
// category endpoints
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package category

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "category" service endpoints.
type Endpoints struct {
	List goa.Endpoint
}

// NewEndpoints wraps the methods of the "category" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List: NewListEndpoint(s),
	}
}

// Use applies the given middleware to all the "category" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "category".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.List(ctx)
	}
}
