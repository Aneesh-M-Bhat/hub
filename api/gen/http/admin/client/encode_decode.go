// Code generated by goa v3.15.0, DO NOT EDIT.
//
// admin HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	admin "github.com/tektoncd/hub/api/gen/admin"
	goahttp "goa.design/goa/v3/http"
)

// BuildUpdateAgentRequest instantiates a HTTP request object with method and
// path set to call the "admin" service "UpdateAgent" endpoint
func (c *Client) BuildUpdateAgentRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateAgentAdminPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("admin", "UpdateAgent", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateAgentRequest returns an encoder for requests sent to the admin
// UpdateAgent server.
func EncodeUpdateAgentRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*admin.UpdateAgentPayload)
		if !ok {
			return goahttp.ErrInvalidType("admin", "UpdateAgent", "*admin.UpdateAgentPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateAgentRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("admin", "UpdateAgent", err)
		}
		return nil
	}
}

// DecodeUpdateAgentResponse returns a decoder for responses returned by the
// admin UpdateAgent endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUpdateAgentResponse may return the following errors:
//   - "invalid-payload" (type *goa.ServiceError): http.StatusBadRequest
//   - "invalid-token" (type *goa.ServiceError): http.StatusUnauthorized
//   - "invalid-scopes" (type *goa.ServiceError): http.StatusForbidden
//   - "internal-error" (type *goa.ServiceError): http.StatusInternalServerError
//   - error: internal error
func DecodeUpdateAgentResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateAgentResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "UpdateAgent", err)
			}
			err = ValidateUpdateAgentResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "UpdateAgent", err)
			}
			res := NewUpdateAgentResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body UpdateAgentInvalidPayloadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "UpdateAgent", err)
			}
			err = ValidateUpdateAgentInvalidPayloadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "UpdateAgent", err)
			}
			return nil, NewUpdateAgentInvalidPayload(&body)
		case http.StatusUnauthorized:
			var (
				body UpdateAgentInvalidTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "UpdateAgent", err)
			}
			err = ValidateUpdateAgentInvalidTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "UpdateAgent", err)
			}
			return nil, NewUpdateAgentInvalidToken(&body)
		case http.StatusForbidden:
			var (
				body UpdateAgentInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "UpdateAgent", err)
			}
			err = ValidateUpdateAgentInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "UpdateAgent", err)
			}
			return nil, NewUpdateAgentInvalidScopes(&body)
		case http.StatusInternalServerError:
			var (
				body UpdateAgentInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "UpdateAgent", err)
			}
			err = ValidateUpdateAgentInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "UpdateAgent", err)
			}
			return nil, NewUpdateAgentInternalError(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("admin", "UpdateAgent", resp.StatusCode, string(body))
		}
	}
}

// BuildRefreshConfigRequest instantiates a HTTP request object with method and
// path set to call the "admin" service "RefreshConfig" endpoint
func (c *Client) BuildRefreshConfigRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RefreshConfigAdminPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("admin", "RefreshConfig", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRefreshConfigRequest returns an encoder for requests sent to the admin
// RefreshConfig server.
func EncodeRefreshConfigRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*admin.RefreshConfigPayload)
		if !ok {
			return goahttp.ErrInvalidType("admin", "RefreshConfig", "*admin.RefreshConfigPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeRefreshConfigResponse returns a decoder for responses returned by the
// admin RefreshConfig endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeRefreshConfigResponse may return the following errors:
//   - "invalid-token" (type *goa.ServiceError): http.StatusUnauthorized
//   - "invalid-scopes" (type *goa.ServiceError): http.StatusForbidden
//   - "internal-error" (type *goa.ServiceError): http.StatusInternalServerError
//   - error: internal error
func DecodeRefreshConfigResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body RefreshConfigResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "RefreshConfig", err)
			}
			err = ValidateRefreshConfigResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "RefreshConfig", err)
			}
			res := NewRefreshConfigResultOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body RefreshConfigInvalidTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "RefreshConfig", err)
			}
			err = ValidateRefreshConfigInvalidTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "RefreshConfig", err)
			}
			return nil, NewRefreshConfigInvalidToken(&body)
		case http.StatusForbidden:
			var (
				body RefreshConfigInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "RefreshConfig", err)
			}
			err = ValidateRefreshConfigInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "RefreshConfig", err)
			}
			return nil, NewRefreshConfigInvalidScopes(&body)
		case http.StatusInternalServerError:
			var (
				body RefreshConfigInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("admin", "RefreshConfig", err)
			}
			err = ValidateRefreshConfigInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("admin", "RefreshConfig", err)
			}
			return nil, NewRefreshConfigInternalError(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("admin", "RefreshConfig", resp.StatusCode, string(body))
		}
	}
}
