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

// Get API key information.
//
// Retrieves information for one or more API keys.
// NOTE: If you have only the `manage_own_api_key` privilege, this API returns
// only the API keys that you own.
// If you have `read_security`, `manage_api_key` or greater privileges
// (including `manage_security`), this API returns all API keys regardless of
// ownership.
package getapikey

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

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetApiKey struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewGetApiKey type alias for index.
type NewGetApiKey func() *GetApiKey

// NewGetApiKeyFunc returns a new instance of GetApiKey with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetApiKeyFunc(tp elastictransport.Interface) NewGetApiKey {
	return func() *GetApiKey {
		n := New(tp)

		return n
	}
}

// Get API key information.
//
// Retrieves information for one or more API keys.
// NOTE: If you have only the `manage_own_api_key` privilege, this API returns
// only the API keys that you own.
// If you have `read_security`, `manage_api_key` or greater privileges
// (including `manage_security`), this API returns all API keys regardless of
// ownership.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-api-key.html
func New(tp elastictransport.Interface) *GetApiKey {
	r := &GetApiKey{
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
func (r *GetApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("api_key")

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
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r GetApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.get_api_key")
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
		instrument.BeforeRequest(req, "security.get_api_key")
		if reader := instrument.RecordRequestBody(ctx, "security.get_api_key", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.get_api_key")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetApiKey query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getapikey.Response
func (r GetApiKey) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.get_api_key")
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
func (r GetApiKey) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.get_api_key")
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
		err := fmt.Errorf("an error happened during the GetApiKey query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the GetApiKey headers map.
func (r *GetApiKey) Header(key, value string) *GetApiKey {
	r.headers.Set(key, value)

	return r
}

// Id An API key id.
// This parameter cannot be used with any of `name`, `realm_name` or `username`.
// API name: id
func (r *GetApiKey) Id(id string) *GetApiKey {
	r.values.Set("id", id)

	return r
}

// Name An API key name.
// This parameter cannot be used with any of `id`, `realm_name` or `username`.
// It supports prefix search with wildcard.
// API name: name
func (r *GetApiKey) Name(name string) *GetApiKey {
	r.values.Set("name", name)

	return r
}

// Owner A boolean flag that can be used to query API keys owned by the currently
// authenticated user.
// The `realm_name` or `username` parameters cannot be specified when this
// parameter is set to `true` as they are assumed to be the currently
// authenticated ones.
// API name: owner
func (r *GetApiKey) Owner(owner bool) *GetApiKey {
	r.values.Set("owner", strconv.FormatBool(owner))

	return r
}

// RealmName The name of an authentication realm.
// This parameter cannot be used with either `id` or `name` or when `owner` flag
// is set to `true`.
// API name: realm_name
func (r *GetApiKey) RealmName(name string) *GetApiKey {
	r.values.Set("realm_name", name)

	return r
}

// Username The username of a user.
// This parameter cannot be used with either `id` or `name` or when `owner` flag
// is set to `true`.
// API name: username
func (r *GetApiKey) Username(username string) *GetApiKey {
	r.values.Set("username", username)

	return r
}

// WithLimitedBy Return the snapshot of the owner user's role descriptors
// associated with the API key. An API key's actual
// permission is the intersection of its assigned role
// descriptors and the owner user's role descriptors.
// API name: with_limited_by
func (r *GetApiKey) WithLimitedBy(withlimitedby bool) *GetApiKey {
	r.values.Set("with_limited_by", strconv.FormatBool(withlimitedby))

	return r
}

// ActiveOnly A boolean flag that can be used to query API keys that are currently active.
// An API key is considered active if it is neither invalidated, nor expired at
// query time. You can specify this together with other parameters such as
// `owner` or `name`. If `active_only` is false, the response will include both
// active and inactive (expired or invalidated) keys.
// API name: active_only
func (r *GetApiKey) ActiveOnly(activeonly bool) *GetApiKey {
	r.values.Set("active_only", strconv.FormatBool(activeonly))

	return r
}

// WithProfileUid Determines whether to also retrieve the profile uid, for the API key owner
// principal, if it exists.
// API name: with_profile_uid
func (r *GetApiKey) WithProfileUid(withprofileuid bool) *GetApiKey {
	r.values.Set("with_profile_uid", strconv.FormatBool(withprofileuid))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetApiKey) ErrorTrace(errortrace bool) *GetApiKey {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetApiKey) FilterPath(filterpaths ...string) *GetApiKey {
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
func (r *GetApiKey) Human(human bool) *GetApiKey {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GetApiKey) Pretty(pretty bool) *GetApiKey {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
