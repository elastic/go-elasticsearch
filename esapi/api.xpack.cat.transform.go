// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 7.8.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

func newCatTransformFunc(t Transport) CatTransform {
	return func(o ...func(*CatTransformRequest)) (*Response, error) {
		var r = CatTransformRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// CatTransform -
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-transforms.html.
//
type CatTransform func(o ...func(*CatTransformRequest)) (*Response, error)

// CatTransformRequest configures the Cat Transform API request.
//
type CatTransformRequest struct {
	TransformID string

	AllowNoMatch *bool
	Format       string
	From         *int
	H            []string
	Help         *bool
	S            []string
	Size         *int
	Time         string
	V            *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r CatTransformRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_cat") + 1 + len("transforms") + 1 + len(r.TransformID))
	path.WriteString("/")
	path.WriteString("_cat")
	path.WriteString("/")
	path.WriteString("transforms")
	if r.TransformID != "" {
		path.WriteString("/")
		path.WriteString(r.TransformID)
	}

	params = make(map[string]string)

	if r.AllowNoMatch != nil {
		params["allow_no_match"] = strconv.FormatBool(*r.AllowNoMatch)
	}

	if r.Format != "" {
		params["format"] = r.Format
	}

	if r.From != nil {
		params["from"] = strconv.FormatInt(int64(*r.From), 10)
	}

	if len(r.H) > 0 {
		params["h"] = strings.Join(r.H, ",")
	}

	if r.Help != nil {
		params["help"] = strconv.FormatBool(*r.Help)
	}

	if len(r.S) > 0 {
		params["s"] = strings.Join(r.S, ",")
	}

	if r.Size != nil {
		params["size"] = strconv.FormatInt(int64(*r.Size), 10)
	}

	if r.Time != "" {
		params["time"] = r.Time
	}

	if r.V != nil {
		params["v"] = strconv.FormatBool(*r.V)
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

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

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
func (f CatTransform) WithContext(v context.Context) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.ctx = v
	}
}

// WithTransformID - the ID of the transform for which to get stats. '_all' or '*' implies all transforms.
//
func (f CatTransform) WithTransformID(v string) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.TransformID = v
	}
}

// WithAllowNoMatch - whether to ignore if a wildcard expression matches no transforms. (this includes `_all` string or when no transforms have been specified).
//
func (f CatTransform) WithAllowNoMatch(v bool) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.AllowNoMatch = &v
	}
}

// WithFormat - a short version of the accept header, e.g. json, yaml.
//
func (f CatTransform) WithFormat(v string) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.Format = v
	}
}

// WithFrom - skips a number of transform configs, defaults to 0.
//
func (f CatTransform) WithFrom(v int) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.From = &v
	}
}

// WithH - comma-separated list of column names to display.
//
func (f CatTransform) WithH(v ...string) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.H = v
	}
}

// WithHelp - return help information.
//
func (f CatTransform) WithHelp(v bool) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.Help = &v
	}
}

// WithS - comma-separated list of column names or column aliases to sort by.
//
func (f CatTransform) WithS(v ...string) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.S = v
	}
}

// WithSize - specifies a max number of transforms to get, defaults to 100.
//
func (f CatTransform) WithSize(v int) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.Size = &v
	}
}

// WithTime - the unit in which to display time values.
//
func (f CatTransform) WithTime(v string) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.Time = v
	}
}

// WithV - verbose mode. display column headers.
//
func (f CatTransform) WithV(v bool) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.V = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f CatTransform) WithPretty() func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f CatTransform) WithHuman() func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f CatTransform) WithErrorTrace() func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f CatTransform) WithFilterPath(v ...string) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f CatTransform) WithHeader(h map[string]string) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f CatTransform) WithOpaqueID(s string) func(*CatTransformRequest) {
	return func(r *CatTransformRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
