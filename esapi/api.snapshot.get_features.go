// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 7.12.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
	"time"
)

func newSnapshotGetFeaturesFunc(t Transport) SnapshotGetFeatures {
	return func(o ...func(*SnapshotGetFeaturesRequest)) (*Response, error) {
		var r = SnapshotGetFeaturesRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SnapshotGetFeatures returns a list of features which can be snapshotted in this cluster.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html.
//
type SnapshotGetFeatures func(o ...func(*SnapshotGetFeaturesRequest)) (*Response, error)

// SnapshotGetFeaturesRequest configures the Snapshot Get Features API request.
//
type SnapshotGetFeaturesRequest struct {
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SnapshotGetFeaturesRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_snapshottable_features"))
	path.WriteString("/_snapshottable_features")

	params = make(map[string]string)

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
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
func (f SnapshotGetFeatures) WithContext(v context.Context) func(*SnapshotGetFeaturesRequest) {
	return func(r *SnapshotGetFeaturesRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
//
func (f SnapshotGetFeatures) WithMasterTimeout(v time.Duration) func(*SnapshotGetFeaturesRequest) {
	return func(r *SnapshotGetFeaturesRequest) {
		r.MasterTimeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SnapshotGetFeatures) WithPretty() func(*SnapshotGetFeaturesRequest) {
	return func(r *SnapshotGetFeaturesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SnapshotGetFeatures) WithHuman() func(*SnapshotGetFeaturesRequest) {
	return func(r *SnapshotGetFeaturesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SnapshotGetFeatures) WithErrorTrace() func(*SnapshotGetFeaturesRequest) {
	return func(r *SnapshotGetFeaturesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SnapshotGetFeatures) WithFilterPath(v ...string) func(*SnapshotGetFeaturesRequest) {
	return func(r *SnapshotGetFeaturesRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f SnapshotGetFeatures) WithHeader(h map[string]string) func(*SnapshotGetFeaturesRequest) {
	return func(r *SnapshotGetFeaturesRequest) {
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
func (f SnapshotGetFeatures) WithOpaqueID(s string) func(*SnapshotGetFeaturesRequest) {
	return func(r *SnapshotGetFeaturesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
