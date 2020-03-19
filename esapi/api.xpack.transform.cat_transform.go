// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 7.7.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

func newTransformCatTransformFunc(t Transport) TransformCatTransform {
	return func(o ...func(*TransformCatTransformRequest)) (*Response, error) {
		var r = TransformCatTransformRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// TransformCatTransform -
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-transforms.html.
//
type TransformCatTransform func(o ...func(*TransformCatTransformRequest)) (*Response, error)

// TransformCatTransformRequest configures the Transform Cat Transform API request.
//
type TransformCatTransformRequest struct {
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
func (r TransformCatTransformRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
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
func (f TransformCatTransform) WithContext(v context.Context) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.ctx = v
	}
}

// WithTransformID - the ID of the transform for which to get stats. '_all' or '*' implies all transforms.
//
func (f TransformCatTransform) WithTransformID(v string) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.TransformID = v
	}
}

// WithAllowNoMatch - whether to ignore if a wildcard expression matches no transforms. (this includes `_all` string or when no transforms have been specified).
//
func (f TransformCatTransform) WithAllowNoMatch(v bool) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.AllowNoMatch = &v
	}
}

// WithFormat - a short version of the accept header, e.g. json, yaml.
//
func (f TransformCatTransform) WithFormat(v string) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.Format = v
	}
}

// WithFrom - skips a number of transform configs, defaults to 0.
//
func (f TransformCatTransform) WithFrom(v int) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.From = &v
	}
}

// WithH - comma-separated list of column names to display.
//
func (f TransformCatTransform) WithH(v ...string) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.H = v
	}
}

// WithHelp - return help information.
//
func (f TransformCatTransform) WithHelp(v bool) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.Help = &v
	}
}

// WithS - comma-separated list of column names or column aliases to sort by.
//
func (f TransformCatTransform) WithS(v ...string) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.S = v
	}
}

// WithSize - specifies a max number of transforms to get, defaults to 100.
//
func (f TransformCatTransform) WithSize(v int) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.Size = &v
	}
}

// WithTime - the unit in which to display time values.
//
func (f TransformCatTransform) WithTime(v string) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.Time = v
	}
}

// WithV - verbose mode. display column headers.
//
func (f TransformCatTransform) WithV(v bool) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.V = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f TransformCatTransform) WithPretty() func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f TransformCatTransform) WithHuman() func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f TransformCatTransform) WithErrorTrace() func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f TransformCatTransform) WithFilterPath(v ...string) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f TransformCatTransform) WithHeader(h map[string]string) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
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
func (f TransformCatTransform) WithOpaqueID(s string) func(*TransformCatTransformRequest) {
	return func(r *TransformCatTransformRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
