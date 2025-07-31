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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Update a cross-cluster API key.
//
// Update the attributes of an existing cross-cluster API key, which is used for
// API key based remote cluster access.
//
// To use this API, you must have at least the `manage_security` cluster
// privilege.
// Users can only update API keys that they created.
// To update another user's API key, use the `run_as` feature to submit a
// request on behalf of another user.
//
// IMPORTANT: It's not possible to use an API key as the authentication
// credential for this API.
// To update an API key, the owner user's credentials are required.
//
// It's not possible to update expired API keys, or API keys that have been
// invalidated by the invalidate API key API.
//
// This API supports updates to an API key's access scope, metadata, and
// expiration.
// The owner user's information, such as the `username` and `realm`, is also
// updated automatically on every call.
//
// NOTE: This API cannot update REST API keys, which should be updated by either
// the update API key or bulk update API keys API.
package updatecrossclusterapikey

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
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateCrossClusterApiKey struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewUpdateCrossClusterApiKey type alias for index.
type NewUpdateCrossClusterApiKey func(id string) *UpdateCrossClusterApiKey

// NewUpdateCrossClusterApiKeyFunc returns a new instance of UpdateCrossClusterApiKey with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateCrossClusterApiKeyFunc(tp elastictransport.Interface) NewUpdateCrossClusterApiKey {
	return func(id string) *UpdateCrossClusterApiKey {
		n := New(tp)

		n._id(id)

		return n
	}
}

// Update a cross-cluster API key.
//
// Update the attributes of an existing cross-cluster API key, which is used for
// API key based remote cluster access.
//
// To use this API, you must have at least the `manage_security` cluster
// privilege.
// Users can only update API keys that they created.
// To update another user's API key, use the `run_as` feature to submit a
// request on behalf of another user.
//
// IMPORTANT: It's not possible to use an API key as the authentication
// credential for this API.
// To update an API key, the owner user's credentials are required.
//
// It's not possible to update expired API keys, or API keys that have been
// invalidated by the invalidate API key API.
//
// This API supports updates to an API key's access scope, metadata, and
// expiration.
// The owner user's information, such as the `username` and `realm`, is also
// updated automatically on every call.
//
// NOTE: This API cannot update REST API keys, which should be updated by either
// the update API key or bulk update API keys API.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-cross-cluster-api-key.html
func New(tp elastictransport.Interface) *UpdateCrossClusterApiKey {
	r := &UpdateCrossClusterApiKey{
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
func (r *UpdateCrossClusterApiKey) Raw(raw io.Reader) *UpdateCrossClusterApiKey {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *UpdateCrossClusterApiKey) Request(req *Request) *UpdateCrossClusterApiKey {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *UpdateCrossClusterApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for UpdateCrossClusterApiKey: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("cross_cluster")
		path.WriteString("/")
		path.WriteString("api_key")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

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
func (r UpdateCrossClusterApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.update_cross_cluster_api_key")
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
		instrument.BeforeRequest(req, "security.update_cross_cluster_api_key")
		if reader := instrument.RecordRequestBody(ctx, "security.update_cross_cluster_api_key", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.update_cross_cluster_api_key")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the UpdateCrossClusterApiKey query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a updatecrossclusterapikey.Response
func (r UpdateCrossClusterApiKey) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.update_cross_cluster_api_key")
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

// Header set a key, value pair in the UpdateCrossClusterApiKey headers map.
func (r *UpdateCrossClusterApiKey) Header(key, value string) *UpdateCrossClusterApiKey {
	r.headers.Set(key, value)

	return r
}

// Id The ID of the cross-cluster API key to update.
// API Name: id
func (r *UpdateCrossClusterApiKey) _id(id string) *UpdateCrossClusterApiKey {
	r.paramSet |= idMask
	r.id = id

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *UpdateCrossClusterApiKey) ErrorTrace(errortrace bool) *UpdateCrossClusterApiKey {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *UpdateCrossClusterApiKey) FilterPath(filterpaths ...string) *UpdateCrossClusterApiKey {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *UpdateCrossClusterApiKey) Human(human bool) *UpdateCrossClusterApiKey {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *UpdateCrossClusterApiKey) Pretty(pretty bool) *UpdateCrossClusterApiKey {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Access The access to be granted to this API key.
// The access is composed of permissions for cross cluster search and cross
// cluster replication.
// At least one of them must be specified.
// When specified, the new access assignment fully replaces the previously
// assigned access.
// API name: access
func (r *UpdateCrossClusterApiKey) Access(access *types.Access) *UpdateCrossClusterApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Access = *access

	return r
}

// Expiration The expiration time for the API key.
// By default, API keys never expire. This property can be omitted to leave the
// value unchanged.
// API name: expiration
func (r *UpdateCrossClusterApiKey) Expiration(duration types.Duration) *UpdateCrossClusterApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Expiration = duration

	return r
}

// Metadata Arbitrary metadata that you want to associate with the API key.
// It supports nested data structure.
// Within the metadata object, keys beginning with `_` are reserved for system
// usage.
// When specified, this information fully replaces metadata previously
// associated with the API key.
// API name: metadata
func (r *UpdateCrossClusterApiKey) Metadata(metadata types.Metadata) *UpdateCrossClusterApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Metadata = metadata

	return r
}
