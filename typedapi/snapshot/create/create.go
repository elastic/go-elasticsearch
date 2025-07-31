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

// Create a snapshot.
// Take a snapshot of a cluster or of data streams and indices.
package create

import (
	gobytes "bytes"
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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	repositoryMask = iota + 1

	snapshotMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Create struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	repository string
	snapshot   string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewCreate type alias for index.
type NewCreate func(repository, snapshot string) *Create

// NewCreateFunc returns a new instance of Create with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCreateFunc(tp elastictransport.Interface) NewCreate {
	return func(repository, snapshot string) *Create {
		n := New(tp)

		n._repository(repository)

		n._snapshot(snapshot)

		return n
	}
}

// Create a snapshot.
// Take a snapshot of a cluster or of data streams and indices.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-snapshot-create
func New(tp elastictransport.Interface) *Create {
	r := &Create{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Create) Raw(raw io.Reader) *Create {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Create) Request(req *Request) *Create {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Create) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Create: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
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

		method = http.MethodPut
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

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=9")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Create) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "snapshot.create")
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
		instrument.BeforeRequest(req, "snapshot.create")
		if reader := instrument.RecordRequestBody(ctx, "snapshot.create", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "snapshot.create")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Create query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a create.Response
func (r Create) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "snapshot.create")
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

// Header set a key, value pair in the Create headers map.
func (r *Create) Header(key, value string) *Create {
	r.headers.Set(key, value)

	return r
}

// Repository The name of the repository for the snapshot.
// API Name: repository
func (r *Create) _repository(repository string) *Create {
	r.paramSet |= repositoryMask
	r.repository = repository

	return r
}

// Snapshot The name of the snapshot.
// It supportes date math.
// It must be unique in the repository.
// API Name: snapshot
func (r *Create) _snapshot(snapshot string) *Create {
	r.paramSet |= snapshotMask
	r.snapshot = snapshot

	return r
}

// MasterTimeout The period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *Create) MasterTimeout(duration string) *Create {
	r.values.Set("master_timeout", duration)

	return r
}

// WaitForCompletion If `true`, the request returns a response when the snapshot is complete.
// If `false`, the request returns a response when the snapshot initializes.
// API name: wait_for_completion
func (r *Create) WaitForCompletion(waitforcompletion bool) *Create {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Create) ErrorTrace(errortrace bool) *Create {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Create) FilterPath(filterpaths ...string) *Create {
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
func (r *Create) Human(human bool) *Create {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Create) Pretty(pretty bool) *Create {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Determines how wildcard patterns in the `indices` parameter match data
// streams and indices.
// It supports comma-separated values such as `open,hidden`.
// API name: expand_wildcards
func (r *Create) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Create {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ExpandWildcards = expandwildcards

	return r
}

// The feature states to include in the snapshot.
// Each feature state includes one or more system indices containing related
// data.
// You can view a list of eligible features using the get features API.
//
// If `include_global_state` is `true`, all current feature states are included
// by default.
// If `include_global_state` is `false`, no feature states are included by
// default.
//
// Note that specifying an empty array will result in the default behavior.
// To exclude all feature states, regardless of the `include_global_state`
// value, specify an array with only the value `none` (`["none"]`).
// API name: feature_states
func (r *Create) FeatureStates(featurestates ...string) *Create {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range featurestates {

		r.req.FeatureStates = append(r.req.FeatureStates, v)

	}
	return r
}

// If `true`, the request ignores data streams and indices in `indices` that are
// missing or closed.
// If `false`, the request returns an error for any data stream or index that is
// missing or closed.
// API name: ignore_unavailable
func (r *Create) IgnoreUnavailable(ignoreunavailable bool) *Create {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IgnoreUnavailable = &ignoreunavailable

	return r
}

// If `true`, the current cluster state is included in the snapshot.
// The cluster state includes persistent cluster settings, composable index
// templates, legacy index templates, ingest pipelines, and ILM policies.
// It also includes data stored in system indices, such as Watches and task
// records (configurable via `feature_states`).
// API name: include_global_state
func (r *Create) IncludeGlobalState(includeglobalstate bool) *Create {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IncludeGlobalState = &includeglobalstate

	return r
}

// A comma-separated list of data streams and indices to include in the
// snapshot.
// It supports a multi-target syntax.
// The default is an empty array (`[]`), which includes all regular data streams
// and regular indices.
// To exclude all data streams and indices, use `-*`.
//
// You can't use this parameter to include or exclude system indices or system
// data streams from a snapshot.
// Use `feature_states` instead.
// API name: indices
func (r *Create) Indices(indices ...string) *Create {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Indices = indices

	return r
}

// Arbitrary metadata to the snapshot, such as a record of who took the
// snapshot, why it was taken, or any other useful data.
// It can have any contents but it must be less than 1024 bytes.
// This information is not automatically generated by Elasticsearch.
// API name: metadata
func (r *Create) Metadata(metadata types.MetadataVariant) *Create {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Metadata = *metadata.MetadataCaster()

	return r
}

// If `true`, it enables you to restore a partial snapshot of indices with
// unavailable shards.
// Only shards that were successfully included in the snapshot will be restored.
// All missing shards will be recreated as empty.
//
// If `false`, the entire restore operation will fail if one or more indices
// included in the snapshot do not have all primary shards available.
// API name: partial
func (r *Create) Partial(partial bool) *Create {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Partial = &partial

	return r
}
