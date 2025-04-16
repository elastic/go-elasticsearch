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
// Code generated from specification version 8.19.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

func newSnapshotRepositoryVerifyIntegrityFunc(t Transport) SnapshotRepositoryVerifyIntegrity {
	return func(repository string, o ...func(*SnapshotRepositoryVerifyIntegrityRequest)) (*Response, error) {
		var r = SnapshotRepositoryVerifyIntegrityRequest{Repository: repository}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.Instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SnapshotRepositoryVerifyIntegrity verifies the integrity of the contents of a snapshot repository
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html.
type SnapshotRepositoryVerifyIntegrity func(repository string, o ...func(*SnapshotRepositoryVerifyIntegrityRequest)) (*Response, error)

// SnapshotRepositoryVerifyIntegrityRequest configures the Snapshot Repository Verify Integrity API request.
type SnapshotRepositoryVerifyIntegrityRequest struct {
	Repository string

	BlobThreadPoolConcurrency            *int
	IndexSnapshotVerificationConcurrency *int
	IndexVerificationConcurrency         *int
	MaxBytesPerSec                       string
	MaxFailedShardSnapshots              *int
	MetaThreadPoolConcurrency            *int
	SnapshotVerificationConcurrency      *int
	VerifyBlobContents                   *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SnapshotRepositoryVerifyIntegrityRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.repository_verify_integrity")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_snapshot") + 1 + len(r.Repository) + 1 + len("_verify_integrity"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_snapshot")
	path.WriteString("/")
	path.WriteString(r.Repository)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "repository", r.Repository)
	}
	path.WriteString("/")
	path.WriteString("_verify_integrity")

	params = make(map[string]string)

	if r.BlobThreadPoolConcurrency != nil {
		params["blob_thread_pool_concurrency"] = strconv.FormatInt(int64(*r.BlobThreadPoolConcurrency), 10)
	}

	if r.IndexSnapshotVerificationConcurrency != nil {
		params["index_snapshot_verification_concurrency"] = strconv.FormatInt(int64(*r.IndexSnapshotVerificationConcurrency), 10)
	}

	if r.IndexVerificationConcurrency != nil {
		params["index_verification_concurrency"] = strconv.FormatInt(int64(*r.IndexVerificationConcurrency), 10)
	}

	if r.MaxBytesPerSec != "" {
		params["max_bytes_per_sec"] = r.MaxBytesPerSec
	}

	if r.MaxFailedShardSnapshots != nil {
		params["max_failed_shard_snapshots"] = strconv.FormatInt(int64(*r.MaxFailedShardSnapshots), 10)
	}

	if r.MetaThreadPoolConcurrency != nil {
		params["meta_thread_pool_concurrency"] = strconv.FormatInt(int64(*r.MetaThreadPoolConcurrency), 10)
	}

	if r.SnapshotVerificationConcurrency != nil {
		params["snapshot_verification_concurrency"] = strconv.FormatInt(int64(*r.SnapshotVerificationConcurrency), 10)
	}

	if r.VerifyBlobContents != nil {
		params["verify_blob_contents"] = strconv.FormatBool(*r.VerifyBlobContents)
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
		if instrument, ok := r.Instrument.(Instrumentation); ok {
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

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "snapshot.repository_verify_integrity")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "snapshot.repository_verify_integrity")
	}
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
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
func (f SnapshotRepositoryVerifyIntegrity) WithContext(v context.Context) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.ctx = v
	}
}

// WithBlobThreadPoolConcurrency - number of threads to use for reading blob contents.
func (f SnapshotRepositoryVerifyIntegrity) WithBlobThreadPoolConcurrency(v int) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.BlobThreadPoolConcurrency = &v
	}
}

// WithIndexSnapshotVerificationConcurrency - number of snapshots to verify concurrently within each index.
func (f SnapshotRepositoryVerifyIntegrity) WithIndexSnapshotVerificationConcurrency(v int) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.IndexSnapshotVerificationConcurrency = &v
	}
}

// WithIndexVerificationConcurrency - number of indices to verify concurrently.
func (f SnapshotRepositoryVerifyIntegrity) WithIndexVerificationConcurrency(v int) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.IndexVerificationConcurrency = &v
	}
}

// WithMaxBytesPerSec - rate limit for individual blob verification.
func (f SnapshotRepositoryVerifyIntegrity) WithMaxBytesPerSec(v string) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.MaxBytesPerSec = v
	}
}

// WithMaxFailedShardSnapshots - maximum permitted number of failed shard snapshots.
func (f SnapshotRepositoryVerifyIntegrity) WithMaxFailedShardSnapshots(v int) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.MaxFailedShardSnapshots = &v
	}
}

// WithMetaThreadPoolConcurrency - number of threads to use for reading metadata.
func (f SnapshotRepositoryVerifyIntegrity) WithMetaThreadPoolConcurrency(v int) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.MetaThreadPoolConcurrency = &v
	}
}

// WithSnapshotVerificationConcurrency - number of snapshots to verify concurrently.
func (f SnapshotRepositoryVerifyIntegrity) WithSnapshotVerificationConcurrency(v int) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.SnapshotVerificationConcurrency = &v
	}
}

// WithVerifyBlobContents - whether to verify the contents of individual blobs.
func (f SnapshotRepositoryVerifyIntegrity) WithVerifyBlobContents(v bool) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.VerifyBlobContents = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SnapshotRepositoryVerifyIntegrity) WithPretty() func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SnapshotRepositoryVerifyIntegrity) WithHuman() func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SnapshotRepositoryVerifyIntegrity) WithErrorTrace() func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SnapshotRepositoryVerifyIntegrity) WithFilterPath(v ...string) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SnapshotRepositoryVerifyIntegrity) WithHeader(h map[string]string) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SnapshotRepositoryVerifyIntegrity) WithOpaqueID(s string) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	return func(r *SnapshotRepositoryVerifyIntegrityRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
