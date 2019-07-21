package esapi

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	headerContentType = "Content-Type"
)

var (
	headerContentTypeJSON = []string{"application/json"}
)

// Request defines the API request.
//
type Request interface {
	Do(ctx context.Context, transport Transport) (*Response, error)
}

// newRequest creates an HTTP request.
//
func newRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)

	if err != nil {
		return nil, err
	}

	if body != nil && req.GetBody == nil {
		req.GetBody = func() (io.ReadCloser, error) {
			return ioutil.NopCloser(body), nil
		}
	}

	return req, err
}
