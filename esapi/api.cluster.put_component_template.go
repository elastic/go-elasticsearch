// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 7.10.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newClusterPutComponentTemplateFunc(t Transport) ClusterPutComponentTemplate {
	return func(name string, body io.Reader, o ...func(*ClusterPutComponentTemplateRequest)) (*Response, error) {
		var r = ClusterPutComponentTemplateRequest{Name: name, Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// ClusterPutComponentTemplate creates or updates a component template
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-component-template.html.
//
type ClusterPutComponentTemplate func(name string, body io.Reader, o ...func(*ClusterPutComponentTemplateRequest)) (*Response, error)

// ClusterPutComponentTemplateRequest configures the Cluster Put Component Template API request.
//
type ClusterPutComponentTemplateRequest struct {
	Body io.Reader

	Name string

	Create        *bool
	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r ClusterPutComponentTemplateRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len("_component_template") + 1 + len(r.Name))
	path.WriteString("/")
	path.WriteString("_component_template")
	path.WriteString("/")
	path.WriteString(r.Name)

	params = make(map[string]string)

	if r.Create != nil {
		params["create"] = strconv.FormatBool(*r.Create)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
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

	req, err := newRequest(method, path.String(), r.Body)
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

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
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
func (f ClusterPutComponentTemplate) WithContext(v context.Context) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.ctx = v
	}
}

// WithCreate - whether the index template should only be added if new or can also replace an existing one.
//
func (f ClusterPutComponentTemplate) WithCreate(v bool) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.Create = &v
	}
}

// WithMasterTimeout - specify timeout for connection to master.
//
func (f ClusterPutComponentTemplate) WithMasterTimeout(v time.Duration) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
//
func (f ClusterPutComponentTemplate) WithTimeout(v time.Duration) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f ClusterPutComponentTemplate) WithPretty() func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f ClusterPutComponentTemplate) WithHuman() func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f ClusterPutComponentTemplate) WithErrorTrace() func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f ClusterPutComponentTemplate) WithFilterPath(v ...string) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f ClusterPutComponentTemplate) WithHeader(h map[string]string) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
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
func (f ClusterPutComponentTemplate) WithOpaqueID(s string) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
