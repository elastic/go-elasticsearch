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

// Create an API key.
//
// Create an API key for access without requiring basic authentication.
//
// IMPORTANT: If the credential that is used to authenticate this request is an
// API key, the derived API key cannot have any privileges.
// If you specify privileges, the API returns an error.
//
// A successful request returns a JSON structure that contains the API key, its
// unique id, and its name.
// If applicable, it also returns expiration information for the API key in
// milliseconds.
//
// NOTE: By default, API keys never expire. You can specify expiration
// information when you create the API keys.
//
// The API keys are created by the Elasticsearch API key service, which is
// automatically enabled.
// To configure or turn off the API key service, refer to API key service
// setting documentation.
package createapikey

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CreateApiKey struct {
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

// NewCreateApiKey type alias for index.
type NewCreateApiKey func() *CreateApiKey

// NewCreateApiKeyFunc returns a new instance of CreateApiKey with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCreateApiKeyFunc(tp elastictransport.Interface) NewCreateApiKey {
	return func() *CreateApiKey {
		n := New(tp)

		return n
	}
}

// Create an API key.
//
// Create an API key for access without requiring basic authentication.
//
// IMPORTANT: If the credential that is used to authenticate this request is an
// API key, the derived API key cannot have any privileges.
// If you specify privileges, the API returns an error.
//
// A successful request returns a JSON structure that contains the API key, its
// unique id, and its name.
// If applicable, it also returns expiration information for the API key in
// milliseconds.
//
// NOTE: By default, API keys never expire. You can specify expiration
// information when you create the API keys.
//
// The API keys are created by the Elasticsearch API key service, which is
// automatically enabled.
// To configure or turn off the API key service, refer to API key service
// setting documentation.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html
func New(tp elastictransport.Interface) *CreateApiKey {
	r := &CreateApiKey{
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
func (r *CreateApiKey) Raw(raw io.Reader) *CreateApiKey {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *CreateApiKey) Request(req *Request) *CreateApiKey {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *CreateApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for CreateApiKey: %w", err)
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
func (r CreateApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.create_api_key")
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
		instrument.BeforeRequest(req, "security.create_api_key")
		if reader := instrument.RecordRequestBody(ctx, "security.create_api_key", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.create_api_key")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the CreateApiKey query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a createapikey.Response
func (r CreateApiKey) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.create_api_key")
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

// Header set a key, value pair in the CreateApiKey headers map.
func (r *CreateApiKey) Header(key, value string) *CreateApiKey {
	r.headers.Set(key, value)

	return r
}

// Refresh If `true` (the default) then refresh the affected shards to make this
// operation visible to search, if `wait_for` then wait for a refresh to make
// this operation visible to search, if `false` then do nothing with refreshes.
// API name: refresh
func (r *CreateApiKey) Refresh(refresh refresh.Refresh) *CreateApiKey {
	r.values.Set("refresh", refresh.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *CreateApiKey) ErrorTrace(errortrace bool) *CreateApiKey {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *CreateApiKey) FilterPath(filterpaths ...string) *CreateApiKey {
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
func (r *CreateApiKey) Human(human bool) *CreateApiKey {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *CreateApiKey) Pretty(pretty bool) *CreateApiKey {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Expiration The expiration time for the API key.
// By default, API keys never expire.
// API name: expiration
func (r *CreateApiKey) Expiration(duration types.Duration) *CreateApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Expiration = duration

	return r
}

// Metadata Arbitrary metadata that you want to associate with the API key. It supports
// nested data structure. Within the metadata object, keys beginning with `_`
// are reserved for system usage.
// API name: metadata
func (r *CreateApiKey) Metadata(metadata types.Metadata) *CreateApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Metadata = metadata

	return r
}

// Name A name for the API key.
// API name: name
func (r *CreateApiKey) Name(name string) *CreateApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Name = &name

	return r
}

// RoleDescriptors An array of role descriptors for this API key.
// When it is not specified or it is an empty array, the API key will have a
// point in time snapshot of permissions of the authenticated user.
// If you supply role descriptors, the resultant permissions are an intersection
// of API keys permissions and the authenticated user's permissions thereby
// limiting the access scope for API keys.
// The structure of role descriptor is the same as the request for the create
// role API.
// For more details, refer to the create or update roles API.
//
// NOTE: Due to the way in which this permission intersection is calculated, it
// is not possible to create an API key that is a child of another API key,
// unless the derived key is created without any privileges.
// In this case, you must explicitly specify a role descriptor with no
// privileges.
// The derived API key can be used for authentication; it will not have
// authority to call Elasticsearch APIs.
// API name: role_descriptors
func (r *CreateApiKey) RoleDescriptors(roledescriptors map[string]types.RoleDescriptor) *CreateApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RoleDescriptors = roledescriptors

	return r
}
