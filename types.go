package rest

import (
	"context"
	"net/http"
)

// GetPmapsParams defines parameters for GetPmaps.
type (
	GetPmapsParams struct {
		// Limit count of pmaps
		Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

		// Offset from element
		Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
	}

	// RequestEditorFn  is the function signature for the RequestEditor callback function
	RequestModifier func(ctx context.Context, req *http.Request) error

	// Doer performs HTTP requests.
	//
	// The standard http.Client implements this interface.
	HttpRequester interface {
		Do(req *http.Request) (*http.Response, error)
	}

	// ClientOption allows setting custom parameters during construction
	ClientOption func(*Client) error

	GetPmapsResponse struct {
		Body         []byte
		HTTPResponse *http.Response
	}

	GetPmapsPmapUuidResponse struct {
		Body         []byte
		HTTPResponse *http.Response
	}
)

// Status returns HTTPResponse.Status
func (r GetPmapsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPmapsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// Status returns HTTPResponse.Status
func (r GetPmapsPmapUuidResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPmapsPmapUuidResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
