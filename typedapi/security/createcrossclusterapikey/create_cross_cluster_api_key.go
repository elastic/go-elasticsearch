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

// Create a cross-cluster API key.
//
// Create an API key of the `cross_cluster` type for the API key based remote
// cluster access.
// A `cross_cluster` API key cannot be used to authenticate through the REST
// interface.
//
// IMPORTANT: To authenticate this request you must use a credential that is not
// an API key. Even if you use an API key that has the required privilege, the
// API returns an error.
//
// Cross-cluster API keys are created by the Elasticsearch API key service,
// which is automatically enabled.
//
// NOTE: Unlike REST API keys, a cross-cluster API key does not capture
// permissions of the authenticated user. The API key’s effective permission is
// exactly as specified with the `access` property.
//
// A successful request returns a JSON structure that contains the API key, its
// unique ID, and its name. If applicable, it also returns expiration
// information for the API key in milliseconds.
//
// By default, API keys never expire. You can specify expiration information
// when you create the API keys.
//
// Cross-cluster API keys can only be updated with the update cross-cluster API
// key API.
// Attempting to update them with the update REST API key API or the bulk update
// REST API keys API will result in an error.
package createcrossclusterapikey

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

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CreateCrossClusterApiKey struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewCreateCrossClusterApiKey type alias for index.
type NewCreateCrossClusterApiKey func() *CreateCrossClusterApiKey

// NewCreateCrossClusterApiKeyFunc returns a new instance of CreateCrossClusterApiKey with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCreateCrossClusterApiKeyFunc(tp elastictransport.Interface) NewCreateCrossClusterApiKey {
	return func() *CreateCrossClusterApiKey {
		n := New(tp)

		return n
	}
}

// Create a cross-cluster API key.
//
// Create an API key of the `cross_cluster` type for the API key based remote
// cluster access.
// A `cross_cluster` API key cannot be used to authenticate through the REST
// interface.
//
// IMPORTANT: To authenticate this request you must use a credential that is not
// an API key. Even if you use an API key that has the required privilege, the
// API returns an error.
//
// Cross-cluster API keys are created by the Elasticsearch API key service,
// which is automatically enabled.
//
// NOTE: Unlike REST API keys, a cross-cluster API key does not capture
// permissions of the authenticated user. The API key’s effective permission is
// exactly as specified with the `access` property.
//
// A successful request returns a JSON structure that contains the API key, its
// unique ID, and its name. If applicable, it also returns expiration
// information for the API key in milliseconds.
//
// By default, API keys never expire. You can specify expiration information
// when you create the API keys.
//
// Cross-cluster API keys can only be updated with the update cross-cluster API
// key API.
// Attempting to update them with the update REST API key API or the bulk update
// REST API keys API will result in an error.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-cross-cluster-api-key.html
func New(tp elastictransport.Interface) *CreateCrossClusterApiKey {
	r := &CreateCrossClusterApiKey{
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
func (r *CreateCrossClusterApiKey) Raw(raw io.Reader) *CreateCrossClusterApiKey {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *CreateCrossClusterApiKey) Request(req *Request) *CreateCrossClusterApiKey {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *CreateCrossClusterApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for CreateCrossClusterApiKey: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("cross_cluster")
		path.WriteString("/")
		path.WriteString("api_key")

		method = http.MethodPost
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
func (r CreateCrossClusterApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.create_cross_cluster_api_key")
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
		instrument.BeforeRequest(req, "security.create_cross_cluster_api_key")
		if reader := instrument.RecordRequestBody(ctx, "security.create_cross_cluster_api_key", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.create_cross_cluster_api_key")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the CreateCrossClusterApiKey query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a createcrossclusterapikey.Response
func (r CreateCrossClusterApiKey) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.create_cross_cluster_api_key")
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

// Header set a key, value pair in the CreateCrossClusterApiKey headers map.
func (r *CreateCrossClusterApiKey) Header(key, value string) *CreateCrossClusterApiKey {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *CreateCrossClusterApiKey) ErrorTrace(errortrace bool) *CreateCrossClusterApiKey {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *CreateCrossClusterApiKey) FilterPath(filterpaths ...string) *CreateCrossClusterApiKey {
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
func (r *CreateCrossClusterApiKey) Human(human bool) *CreateCrossClusterApiKey {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *CreateCrossClusterApiKey) Pretty(pretty bool) *CreateCrossClusterApiKey {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Access The access to be granted to this API key.
// The access is composed of permissions for cross-cluster search and
// cross-cluster replication.
// At least one of them must be specified.
//
// NOTE: No explicit privileges should be specified for either search or
// replication access.
// The creation process automatically converts the access specification to a
// role descriptor which has relevant privileges assigned accordingly.
// API name: access
func (r *CreateCrossClusterApiKey) Access(access *types.Access) *CreateCrossClusterApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Access = *access

	return r
}

// Expiration Expiration time for the API key.
// By default, API keys never expire.
// API name: expiration
func (r *CreateCrossClusterApiKey) Expiration(duration types.Duration) *CreateCrossClusterApiKey {
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
// API name: metadata
func (r *CreateCrossClusterApiKey) Metadata(metadata types.Metadata) *CreateCrossClusterApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Metadata = metadata

	return r
}

// Name Specifies the name for this API key.
// API name: name
func (r *CreateCrossClusterApiKey) Name(name string) *CreateCrossClusterApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Name = name

	return r
}
