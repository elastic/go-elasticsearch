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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Get the snapshot status.
// Get a detailed description of the current state for each shard participating
// in the snapshot.
//
// Note that this API should be used only to obtain detailed shard-level
// information for ongoing snapshots.
// If this detail is not needed or you want to obtain information about one or
// more existing snapshots, use the get snapshot API.
//
// If you omit the `<snapshot>` request path parameter, the request retrieves
// information only for currently running snapshots.
// This usage is preferred.
// If needed, you can specify `<repository>` and `<snapshot>` to retrieve
// information for specific snapshots, even if they're not currently running.
//
// WARNING: Using the API to return the status of any snapshots other than
// currently running snapshots can be expensive.
// The API requires a read from the repository for each shard in each snapshot.
// For example, if you have 100 snapshots with 1,000 shards each, an API request
// that includes all snapshots will require 100,000 reads (100 snapshots x 1,000
// shards).
//
// Depending on the latency of your storage, such requests can take an extremely
// long time to return results.
// These requests can also tax machine resources and, when using cloud storage,
// incur high processing costs.
package status

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	repositoryMask = iota + 1

	snapshotMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Status struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	repository string
	snapshot   string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewStatus type alias for index.
type NewStatus func() *Status

// NewStatusFunc returns a new instance of Status with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStatusFunc(tp elastictransport.Interface) NewStatus {
	return func() *Status {
		n := New(tp)

		return n
	}
}

// Get the snapshot status.
// Get a detailed description of the current state for each shard participating
// in the snapshot.
//
// Note that this API should be used only to obtain detailed shard-level
// information for ongoing snapshots.
// If this detail is not needed or you want to obtain information about one or
// more existing snapshots, use the get snapshot API.
//
// If you omit the `<snapshot>` request path parameter, the request retrieves
// information only for currently running snapshots.
// This usage is preferred.
// If needed, you can specify `<repository>` and `<snapshot>` to retrieve
// information for specific snapshots, even if they're not currently running.
//
// WARNING: Using the API to return the status of any snapshots other than
// currently running snapshots can be expensive.
// The API requires a read from the repository for each shard in each snapshot.
// For example, if you have 100 snapshots with 1,000 shards each, an API request
// that includes all snapshots will require 100,000 reads (100 snapshots x 1,000
// shards).
//
// Depending on the latency of your storage, such requests can take an extremely
// long time to return results.
// These requests can also tax machine resources and, when using cloud storage,
// incur high processing costs.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-snapshot-status
func New(tp elastictransport.Interface) *Status {
	r := &Status{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Status) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")
		path.WriteString("_status")

		method = http.MethodGet
	case r.paramSet == repositoryMask:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "repository", r.repository)
		}
		path.WriteString(r.repository)
		path.WriteString("/")
		path.WriteString("_status")

		method = http.MethodGet
	case r.paramSet == repositoryMask|snapshotMask:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "repository", r.repository)
		}
		path.WriteString(r.repository)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "snapshot", r.snapshot)
		}
		path.WriteString(r.snapshot)
		path.WriteString("/")
		path.WriteString("_status")

		method = http.MethodGet
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Status) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "snapshot.status")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "snapshot.status")
		if reader := instrument.RecordRequestBody(ctx, "snapshot.status", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "snapshot.status")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Status query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a status.Response
func (r Status) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.status")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r Status) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.status")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the Status query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Status headers map.
func (r *Status) Header(key, value string) *Status {
	r.headers.Set(key, value)

	return r
}

// Repository The snapshot repository name used to limit the request.
// It supports wildcards (`*`) if `<snapshot>` isn't specified.
// API Name: repository
func (r *Status) Repository(repository string) *Status {
	r.paramSet |= repositoryMask
	r.repository = repository

	return r
}

// Snapshot A comma-separated list of snapshots to retrieve status for.
// The default is currently running snapshots.
// Wildcards (`*`) are not supported.
// API Name: snapshot
func (r *Status) Snapshot(snapshot string) *Status {
	r.paramSet |= snapshotMask
	r.snapshot = snapshot

	return r
}

// IgnoreUnavailable If `false`, the request returns an error for any snapshots that are
// unavailable.
// If `true`, the request ignores snapshots that are unavailable, such as those
// that are corrupted or temporarily cannot be returned.
// API name: ignore_unavailable
func (r *Status) IgnoreUnavailable(ignoreunavailable bool) *Status {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// MasterTimeout The period to wait for the master node.
// If the master node is not available before the timeout expires, the request
// fails and returns an error.
// To indicate that the request should never timeout, set it to `-1`.
// API name: master_timeout
func (r *Status) MasterTimeout(duration string) *Status {
	r.values.Set("master_timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Status) ErrorTrace(errortrace bool) *Status {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Status) FilterPath(filterpaths ...string) *Status {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"exists_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *Status) Human(human bool) *Status {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Status) Pretty(pretty bool) *Status {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
