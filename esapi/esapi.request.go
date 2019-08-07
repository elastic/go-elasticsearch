package esapi

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
	req := http.Request{
		Method:     method,
		URL:        &url.URL{Path: path},
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
	}

	if body != nil {
		if body, ok := body.(io.ReadCloser); !ok {
			req.Body = ioutil.NopCloser(body)
		} else {
			req.Body = body
		}

		switch b := body.(type) {
		case *bytes.Buffer:
			req.ContentLength = int64(b.Len())
			buf := b.Bytes()
			req.GetBody = func() (io.ReadCloser, error) {
				r := bytes.NewReader(buf)
				return ioutil.NopCloser(r), nil
			}
		case *bytes.Reader:
			req.ContentLength = int64(b.Len())
		case *strings.Reader:
			req.ContentLength = int64(b.Len())
		default:
		}
	}

	// if body != nil && req.GetBody == nil {
	// 	req.GetBody = func() (io.ReadCloser, error) {
	// 		return ioutil.NopCloser(body), nil
	// 	}
	// }

	return &req, nil
}
