// Code generated by goa v3.15.0, DO NOT EDIT.
//
// catalog HTTP server types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package server

import (
	catalog "github.com/tektoncd/hub/api/v1/gen/catalog"
	goa "goa.design/goa/v3/pkg"
)

// ListResponseBody is the type of the "catalog" service "List" endpoint HTTP
// response body.
type ListResponseBody struct {
	Data []*CatalogResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// ListInternalErrorResponseBody is the type of the "catalog" service "List"
// endpoint HTTP response body for the "internal-error" error.
type ListInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CatalogResponseBody is used to define fields on response body types.
type CatalogResponseBody struct {
	// ID is the unique id of the catalog
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of catalog
	Name string `form:"name" json:"name" xml:"name"`
	// Type of catalog
	Type string `form:"type" json:"type" xml:"type"`
	// URL of catalog
	URL string `form:"url" json:"url" xml:"url"`
	// Provider of catalog
	Provider string `form:"provider" json:"provider" xml:"provider"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "List" endpoint of the "catalog" service.
func NewListResponseBody(res *catalog.ListResult) *ListResponseBody {
	body := &ListResponseBody{}
	if res.Data != nil {
		body.Data = make([]*CatalogResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalCatalogCatalogToCatalogResponseBody(val)
		}
	}
	return body
}

// NewListInternalErrorResponseBody builds the HTTP response body from the
// result of the "List" endpoint of the "catalog" service.
func NewListInternalErrorResponseBody(res *goa.ServiceError) *ListInternalErrorResponseBody {
	body := &ListInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}
