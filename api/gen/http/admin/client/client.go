// Code generated by goa v3.15.0, DO NOT EDIT.
//
// admin client HTTP transport
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the admin service endpoint HTTP clients.
type Client struct {
	// UpdateAgent Doer is the HTTP client used to make requests to the UpdateAgent
	// endpoint.
	UpdateAgentDoer goahttp.Doer

	// RefreshConfig Doer is the HTTP client used to make requests to the
	// RefreshConfig endpoint.
	RefreshConfigDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the admin service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		UpdateAgentDoer:     doer,
		RefreshConfigDoer:   doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// UpdateAgent returns an endpoint that makes HTTP requests to the admin
// service UpdateAgent server.
func (c *Client) UpdateAgent() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateAgentRequest(c.encoder)
		decodeResponse = DecodeUpdateAgentResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUpdateAgentRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateAgentDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "UpdateAgent", err)
		}
		return decodeResponse(resp)
	}
}

// RefreshConfig returns an endpoint that makes HTTP requests to the admin
// service RefreshConfig server.
func (c *Client) RefreshConfig() goa.Endpoint {
	var (
		encodeRequest  = EncodeRefreshConfigRequest(c.encoder)
		decodeResponse = DecodeRefreshConfigResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildRefreshConfigRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RefreshConfigDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("admin", "RefreshConfig", err)
		}
		return decodeResponse(resp)
	}
}
