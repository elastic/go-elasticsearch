// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// Code generated from specification version 8.12.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
	"time"
)

func newSnapshotCleanupRepositoryFunc(t Transport) SnapshotCleanupRepository {
	return func(repository string, o ...func(*SnapshotCleanupRepositoryRequest)) (*Response, error) {
		var r = SnapshotCleanupRepositoryRequest{Repository: repository}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SnapshotCleanupRepository removes stale data from repository.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/clean-up-snapshot-repo-api.html.
type SnapshotCleanupRepository func(repository string, o ...func(*SnapshotCleanupRepositoryRequest)) (*Response, error)

// SnapshotCleanupRepositoryRequest configures the Snapshot Cleanup Repository API request.
type SnapshotCleanupRepositoryRequest struct {
	Repository string

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SnapshotCleanupRepositoryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.cleanup_repository")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_snapshot") + 1 + len(r.Repository) + 1 + len("_cleanup"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_snapshot")
	path.WriteString("/")
	path.WriteString(r.Repository)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "repository", r.Repository)
	}
	path.WriteString("/")
	path.WriteString("_cleanup")

	params = make(map[string]string)

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

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
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

	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "snapshot.cleanup_repository")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "snapshot.cleanup_repository")
	}
	if err != nil {
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
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
func (f SnapshotCleanupRepository) WithContext(v context.Context) func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
func (f SnapshotCleanupRepository) WithMasterTimeout(v time.Duration) func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
func (f SnapshotCleanupRepository) WithTimeout(v time.Duration) func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SnapshotCleanupRepository) WithPretty() func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SnapshotCleanupRepository) WithHuman() func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SnapshotCleanupRepository) WithErrorTrace() func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SnapshotCleanupRepository) WithFilterPath(v ...string) func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SnapshotCleanupRepository) WithHeader(h map[string]string) func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SnapshotCleanupRepository) WithOpaqueID(s string) func(*SnapshotCleanupRepositoryRequest) {
	return func(r *SnapshotCleanupRepositoryRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
