// Code generated from specification version 8.0.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
)

func newSlmDeleteLifecycleFunc(t Transport) SlmDeleteLifecycle {
	return func(o ...func(*SlmDeleteLifecycleRequest)) (*Response, error) {
		var r = SlmDeleteLifecycleRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SlmDeleteLifecycle - https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api.html
//
type SlmDeleteLifecycle func(o ...func(*SlmDeleteLifecycleRequest)) (*Response, error)

// SlmDeleteLifecycleRequest configures the Slm Delete Lifecycle API request.
//
type SlmDeleteLifecycleRequest struct {
	Policy string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SlmDeleteLifecycleRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

	path.Grow(1 + len("_slm") + 1 + len("policy") + 1 + len("{policy_id}"))
	path.WriteString("/")
	path.WriteString("_slm")
	path.WriteString("/")
	path.WriteString("policy")
	path.WriteString("/")
	path.WriteString("{policy_id}")

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), nil)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f SlmDeleteLifecycle) WithContext(v context.Context) func(*SlmDeleteLifecycleRequest) {
	return func(r *SlmDeleteLifecycleRequest) {
		r.ctx = v
	}
}

// WithPolicy - the ID of the snapshot lifecycle policy to remove.
//
func (f SlmDeleteLifecycle) WithPolicy(v string) func(*SlmDeleteLifecycleRequest) {
	return func(r *SlmDeleteLifecycleRequest) {
		r.Policy = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SlmDeleteLifecycle) WithPretty() func(*SlmDeleteLifecycleRequest) {
	return func(r *SlmDeleteLifecycleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SlmDeleteLifecycle) WithHuman() func(*SlmDeleteLifecycleRequest) {
	return func(r *SlmDeleteLifecycleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SlmDeleteLifecycle) WithErrorTrace() func(*SlmDeleteLifecycleRequest) {
	return func(r *SlmDeleteLifecycleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SlmDeleteLifecycle) WithFilterPath(v ...string) func(*SlmDeleteLifecycleRequest) {
	return func(r *SlmDeleteLifecycleRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f SlmDeleteLifecycle) WithHeader(h map[string]string) func(*SlmDeleteLifecycleRequest) {
	return func(r *SlmDeleteLifecycleRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}
