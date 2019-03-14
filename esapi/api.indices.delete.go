// Code generated from specification version 5.6.16 (052c67e4ebe): DO NOT EDIT

package esapi

import (
	"context"
	"strings"
	"time"
)

func newIndicesDeleteFunc(t Transport) IndicesDelete {
	return func(index []string, o ...func(*IndicesDeleteRequest)) (*Response, error) {
		var r = IndicesDeleteRequest{Index: index}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// IndicesDelete deletes an index.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-delete-index.html.
//
type IndicesDelete func(index []string, o ...func(*IndicesDeleteRequest)) (*Response, error)

// IndicesDeleteRequest configures the Indices Delete API request.
//
type IndicesDeleteRequest struct {
	Index []string

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IndicesDeleteRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

	path.Grow(1 + len(strings.Join(r.Index, ",")))
	path.WriteString("/")
	path.WriteString(strings.Join(r.Index, ","))

	params = make(map[string]string)

	if r.MasterTimeout != 0 {
		params["master_timeout"] = time.Duration(r.MasterTimeout * time.Millisecond).String()
	}

	if r.Timeout != 0 {
		params["timeout"] = time.Duration(r.Timeout * time.Millisecond).String()
	}

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
func (f IndicesDelete) WithContext(v context.Context) func(*IndicesDeleteRequest) {
	return func(r *IndicesDeleteRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - specify timeout for connection to master.
//
func (f IndicesDelete) WithMasterTimeout(v time.Duration) func(*IndicesDeleteRequest) {
	return func(r *IndicesDeleteRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
//
func (f IndicesDelete) WithTimeout(v time.Duration) func(*IndicesDeleteRequest) {
	return func(r *IndicesDeleteRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IndicesDelete) WithPretty() func(*IndicesDeleteRequest) {
	return func(r *IndicesDeleteRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IndicesDelete) WithHuman() func(*IndicesDeleteRequest) {
	return func(r *IndicesDeleteRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IndicesDelete) WithErrorTrace() func(*IndicesDeleteRequest) {
	return func(r *IndicesDeleteRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IndicesDelete) WithFilterPath(v ...string) func(*IndicesDeleteRequest) {
	return func(r *IndicesDeleteRequest) {
		r.FilterPath = v
	}
}
