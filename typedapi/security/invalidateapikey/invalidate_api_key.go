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

// Invalidate API keys.
//
// This API invalidates API keys created by the create API key or grant API key
// APIs.
// Invalidated API keys fail authentication, but they can still be viewed using
// the get API key information and query API key information APIs, for at least
// the configured retention period, until they are automatically deleted.
//
// To use this API, you must have at least the `manage_security`,
// `manage_api_key`, or `manage_own_api_key` cluster privileges.
// The `manage_security` privilege allows deleting any API key, including both
// REST and cross cluster API keys.
// The `manage_api_key` privilege allows deleting any REST API key, but not
// cross cluster API keys.
// The `manage_own_api_key` only allows deleting REST API keys that are owned by
// the user.
// In addition, with the `manage_own_api_key` privilege, an invalidation request
// must be issued in one of the three formats:
//
// - Set the parameter `owner=true`.
// - Or, set both `username` and `realm_name` to match the user's identity.
// - Or, if the request is issued by an API key, that is to say an API key
// invalidates itself, specify its ID in the `ids` field.
package invalidateapikey

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

type InvalidateApiKey struct {
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

// NewInvalidateApiKey type alias for index.
type NewInvalidateApiKey func() *InvalidateApiKey

// NewInvalidateApiKeyFunc returns a new instance of InvalidateApiKey with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewInvalidateApiKeyFunc(tp elastictransport.Interface) NewInvalidateApiKey {
	return func() *InvalidateApiKey {
		n := New(tp)

		return n
	}
}

// Invalidate API keys.
//
// This API invalidates API keys created by the create API key or grant API key
// APIs.
// Invalidated API keys fail authentication, but they can still be viewed using
// the get API key information and query API key information APIs, for at least
// the configured retention period, until they are automatically deleted.
//
// To use this API, you must have at least the `manage_security`,
// `manage_api_key`, or `manage_own_api_key` cluster privileges.
// The `manage_security` privilege allows deleting any API key, including both
// REST and cross cluster API keys.
// The `manage_api_key` privilege allows deleting any REST API key, but not
// cross cluster API keys.
// The `manage_own_api_key` only allows deleting REST API keys that are owned by
// the user.
// In addition, with the `manage_own_api_key` privilege, an invalidation request
// must be issued in one of the three formats:
//
// - Set the parameter `owner=true`.
// - Or, set both `username` and `realm_name` to match the user's identity.
// - Or, if the request is issued by an API key, that is to say an API key
// invalidates itself, specify its ID in the `ids` field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-invalidate-api-key.html
func New(tp elastictransport.Interface) *InvalidateApiKey {
	r := &InvalidateApiKey{
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
func (r *InvalidateApiKey) Raw(raw io.Reader) *InvalidateApiKey {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *InvalidateApiKey) Request(req *Request) *InvalidateApiKey {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *InvalidateApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for InvalidateApiKey: %w", err)
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
		path.WriteString("api_key")

		method = http.MethodDelete
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
func (r InvalidateApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.invalidate_api_key")
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
		instrument.BeforeRequest(req, "security.invalidate_api_key")
		if reader := instrument.RecordRequestBody(ctx, "security.invalidate_api_key", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.invalidate_api_key")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the InvalidateApiKey query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a invalidateapikey.Response
func (r InvalidateApiKey) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.invalidate_api_key")
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

// Header set a key, value pair in the InvalidateApiKey headers map.
func (r *InvalidateApiKey) Header(key, value string) *InvalidateApiKey {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *InvalidateApiKey) ErrorTrace(errortrace bool) *InvalidateApiKey {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *InvalidateApiKey) FilterPath(filterpaths ...string) *InvalidateApiKey {
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
func (r *InvalidateApiKey) Human(human bool) *InvalidateApiKey {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *InvalidateApiKey) Pretty(pretty bool) *InvalidateApiKey {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// API name: id
func (r *InvalidateApiKey) Id(id string) *InvalidateApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Id = &id

	return r
}

// Ids A list of API key ids.
// This parameter cannot be used with any of `name`, `realm_name`, or
// `username`.
// API name: ids
func (r *InvalidateApiKey) Ids(ids ...string) *InvalidateApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Ids = ids

	return r
}

// Name An API key name.
// This parameter cannot be used with any of `ids`, `realm_name` or `username`.
// API name: name
func (r *InvalidateApiKey) Name(name string) *InvalidateApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Name = &name

	return r
}

// Owner Query API keys owned by the currently authenticated user.
// The `realm_name` or `username` parameters cannot be specified when this
// parameter is set to `true` as they are assumed to be the currently
// authenticated ones.
//
// NOTE: At least one of `ids`, `name`, `username`, and `realm_name` must be
// specified if `owner` is `false`.
// API name: owner
func (r *InvalidateApiKey) Owner(owner bool) *InvalidateApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Owner = &owner

	return r
}

// RealmName The name of an authentication realm.
// This parameter cannot be used with either `ids` or `name`, or when `owner`
// flag is set to `true`.
// API name: realm_name
func (r *InvalidateApiKey) RealmName(realmname string) *InvalidateApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RealmName = &realmname

	return r
}

// Username The username of a user.
// This parameter cannot be used with either `ids` or `name` or when `owner`
// flag is set to `true`.
// API name: username
func (r *InvalidateApiKey) Username(username string) *InvalidateApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Username = &username

	return r
}
