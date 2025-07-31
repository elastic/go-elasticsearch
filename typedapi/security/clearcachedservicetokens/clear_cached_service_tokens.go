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

// Clear service account token caches.
//
// Evict a subset of all entries from the service account token caches.
// Two separate caches exist for service account tokens: one cache for tokens
// backed by the `service_tokens` file, and another for tokens backed by the
// `.security` index.
// This API clears matching entries from both caches.
//
// The cache for service account tokens backed by the `.security` index is
// cleared automatically on state changes of the security index.
// The cache for tokens backed by the `service_tokens` file is cleared
// automatically on file changes.
package clearcachedservicetokens

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	namespaceMask = iota + 1

	serviceMask

	nameMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ClearCachedServiceTokens struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	namespace string
	service   string
	name      string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewClearCachedServiceTokens type alias for index.
type NewClearCachedServiceTokens func(namespace, service, name string) *ClearCachedServiceTokens

// NewClearCachedServiceTokensFunc returns a new instance of ClearCachedServiceTokens with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewClearCachedServiceTokensFunc(tp elastictransport.Interface) NewClearCachedServiceTokens {
	return func(namespace, service, name string) *ClearCachedServiceTokens {
		n := New(tp)

		n._namespace(namespace)

		n._service(service)

		n._name(name)

		return n
	}
}

// Clear service account token caches.
//
// Evict a subset of all entries from the service account token caches.
// Two separate caches exist for service account tokens: one cache for tokens
// backed by the `service_tokens` file, and another for tokens backed by the
// `.security` index.
// This API clears matching entries from both caches.
//
// The cache for service account tokens backed by the `.security` index is
// cleared automatically on state changes of the security index.
// The cache for tokens backed by the `service_tokens` file is cleared
// automatically on file changes.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-clear-service-token-caches.html
func New(tp elastictransport.Interface) *ClearCachedServiceTokens {
	r := &ClearCachedServiceTokens{
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
func (r *ClearCachedServiceTokens) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == namespaceMask|serviceMask|nameMask:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("service")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "namespace", r.namespace)
		}
		path.WriteString(r.namespace)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "service", r.service)
		}
		path.WriteString(r.service)
		path.WriteString("/")
		path.WriteString("credential")
		path.WriteString("/")
		path.WriteString("token")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)
		path.WriteString("/")
		path.WriteString("_clear_cache")

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r ClearCachedServiceTokens) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.clear_cached_service_tokens")
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
		instrument.BeforeRequest(req, "security.clear_cached_service_tokens")
		if reader := instrument.RecordRequestBody(ctx, "security.clear_cached_service_tokens", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.clear_cached_service_tokens")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ClearCachedServiceTokens query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a clearcachedservicetokens.Response
func (r ClearCachedServiceTokens) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.clear_cached_service_tokens")
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
func (r ClearCachedServiceTokens) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.clear_cached_service_tokens")
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
		err := fmt.Errorf("an error happened during the ClearCachedServiceTokens query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the ClearCachedServiceTokens headers map.
func (r *ClearCachedServiceTokens) Header(key, value string) *ClearCachedServiceTokens {
	r.headers.Set(key, value)

	return r
}

// Namespace The namespace, which is a top-level grouping of service accounts.
// API Name: namespace
func (r *ClearCachedServiceTokens) _namespace(namespace string) *ClearCachedServiceTokens {
	r.paramSet |= namespaceMask
	r.namespace = namespace

	return r
}

// Service The name of the service, which must be unique within its namespace.
// API Name: service
func (r *ClearCachedServiceTokens) _service(service string) *ClearCachedServiceTokens {
	r.paramSet |= serviceMask
	r.service = service

	return r
}

// Name A comma-separated list of token names to evict from the service account token
// caches.
// Use a wildcard (`*`) to evict all tokens that belong to a service account.
// It does not support other wildcard patterns.
// API Name: name
func (r *ClearCachedServiceTokens) _name(name string) *ClearCachedServiceTokens {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ClearCachedServiceTokens) ErrorTrace(errortrace bool) *ClearCachedServiceTokens {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ClearCachedServiceTokens) FilterPath(filterpaths ...string) *ClearCachedServiceTokens {
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
func (r *ClearCachedServiceTokens) Human(human bool) *ClearCachedServiceTokens {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ClearCachedServiceTokens) Pretty(pretty bool) *ClearCachedServiceTokens {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
