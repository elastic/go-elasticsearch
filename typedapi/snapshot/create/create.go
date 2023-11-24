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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Creates a snapshot in a repository.
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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	repository string
	snapshot   string
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

// Creates a snapshot in a repository.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
func New(tp elastictransport.Interface) *Create {
	r := &Create{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Create: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == repositoryMask|snapshotMask:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")

		path.WriteString(r.repository)
		path.WriteString("/")

		path.WriteString(r.snapshot)

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Create) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Create query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a create.Response
func (r Create) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// Header set a key, value pair in the Create headers map.
func (r *Create) Header(key, value string) *Create {
	r.headers.Set(key, value)

	return r
}

// Repository Repository for the snapshot.
// API Name: repository
func (r *Create) _repository(repository string) *Create {
	r.paramSet |= repositoryMask
	r.repository = repository

	return r
}

// Snapshot Name of the snapshot. Must be unique in the repository.
// API Name: snapshot
func (r *Create) _snapshot(snapshot string) *Create {
	r.paramSet |= snapshotMask
	r.snapshot = snapshot

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *Create) MasterTimeout(duration string) *Create {
	r.values.Set("master_timeout", duration)

	return r
}

// WaitForCompletion If `true`, the request returns a response when the snapshot is complete. If
// `false`, the request returns a response when the snapshot initializes.
// API name: wait_for_completion
func (r *Create) WaitForCompletion(waitforcompletion bool) *Create {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// FeatureStates Feature states to include in the snapshot. Each feature state includes one or
// more system indices containing related data. You can view a list of eligible
// features using the get features API. If `include_global_state` is `true`, all
// current feature states are included by default. If `include_global_state` is
// `false`, no feature states are included by default.
// API name: feature_states
func (r *Create) FeatureStates(featurestates ...string) *Create {
	r.req.FeatureStates = featurestates

	return r
}

// IgnoreUnavailable If `true`, the request ignores data streams and indices in `indices` that are
// missing or closed. If `false`, the request returns an error for any data
// stream or index that is missing or closed.
// API name: ignore_unavailable
func (r *Create) IgnoreUnavailable(ignoreunavailable bool) *Create {
	r.req.IgnoreUnavailable = &ignoreunavailable

	return r
}

// IncludeGlobalState If `true`, the current cluster state is included in the snapshot. The cluster
// state includes persistent cluster settings, composable index templates,
// legacy index templates, ingest pipelines, and ILM policies. It also includes
// data stored in system indices, such as Watches and task records (configurable
// via `feature_states`).
// API name: include_global_state
func (r *Create) IncludeGlobalState(includeglobalstate bool) *Create {
	r.req.IncludeGlobalState = &includeglobalstate

	return r
}

// Indices Data streams and indices to include in the snapshot. Supports multi-target
// syntax. Includes all data streams and indices by default.
// API name: indices
func (r *Create) Indices(indices ...string) *Create {
	r.req.Indices = indices

	return r
}

// Metadata Optional metadata for the snapshot. May have any contents. Must be less than
// 1024 bytes. This map is not automatically generated by Elasticsearch.
// API name: metadata
func (r *Create) Metadata(metadata types.Metadata) *Create {
	r.req.Metadata = metadata

	return r
}

// Partial If `true`, allows restoring a partial snapshot of indices with unavailable
// shards. Only shards that were successfully included in the snapshot will be
// restored. All missing shards will be recreated as empty. If `false`, the
// entire restore operation will fail if one or more indices included in the
// snapshot do not have all primary shards available.
// API name: partial
func (r *Create) Partial(partial bool) *Create {
	r.req.Partial = &partial

	return r
}
