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

// Bulk update API keys.
// Update the attributes for multiple API keys.
//
// IMPORTANT: It is not possible to use an API key as the authentication
// credential for this API. To update API keys, the owner user's credentials are
// required.
//
// This API is similar to the update API key API but enables you to apply the
// same update to multiple API keys in one API call. This operation can greatly
// improve performance over making individual updates.
//
// It is not possible to update expired or invalidated API keys.
//
// This API supports updates to API key access scope, metadata and expiration.
// The access scope of each API key is derived from the `role_descriptors` you
// specify in the request and a snapshot of the owner user's permissions at the
// time of the request.
// The snapshot of the owner's permissions is updated automatically on every
// call.
//
// IMPORTANT: If you don't specify `role_descriptors` in the request, a call to
// this API might still change an API key's access scope. This change can occur
// if the owner user's permissions have changed since the API key was created or
// last modified.
//
// A successful request returns a JSON structure that contains the IDs of all
// updated API keys, the IDs of API keys that already had the requested changes
// and did not require an update, and error details for any failed update.
package bulkupdateapikeys

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

type BulkUpdateApiKeys struct {
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

// NewBulkUpdateApiKeys type alias for index.
type NewBulkUpdateApiKeys func() *BulkUpdateApiKeys

// NewBulkUpdateApiKeysFunc returns a new instance of BulkUpdateApiKeys with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewBulkUpdateApiKeysFunc(tp elastictransport.Interface) NewBulkUpdateApiKeys {
	return func() *BulkUpdateApiKeys {
		n := New(tp)

		return n
	}
}

// Bulk update API keys.
// Update the attributes for multiple API keys.
//
// IMPORTANT: It is not possible to use an API key as the authentication
// credential for this API. To update API keys, the owner user's credentials are
// required.
//
// This API is similar to the update API key API but enables you to apply the
// same update to multiple API keys in one API call. This operation can greatly
// improve performance over making individual updates.
//
// It is not possible to update expired or invalidated API keys.
//
// This API supports updates to API key access scope, metadata and expiration.
// The access scope of each API key is derived from the `role_descriptors` you
// specify in the request and a snapshot of the owner user's permissions at the
// time of the request.
// The snapshot of the owner's permissions is updated automatically on every
// call.
//
// IMPORTANT: If you don't specify `role_descriptors` in the request, a call to
// this API might still change an API key's access scope. This change can occur
// if the owner user's permissions have changed since the API key was created or
// last modified.
//
// A successful request returns a JSON structure that contains the IDs of all
// updated API keys, the IDs of API keys that already had the requested changes
// and did not require an update, and error details for any failed update.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-bulk-update-api-keys.html
func New(tp elastictransport.Interface) *BulkUpdateApiKeys {
	r := &BulkUpdateApiKeys{
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
func (r *BulkUpdateApiKeys) Raw(raw io.Reader) *BulkUpdateApiKeys {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *BulkUpdateApiKeys) Request(req *Request) *BulkUpdateApiKeys {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *BulkUpdateApiKeys) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for BulkUpdateApiKeys: %w", err)
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
		path.WriteString("/")
		path.WriteString("_bulk_update")

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
func (r BulkUpdateApiKeys) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.bulk_update_api_keys")
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
		instrument.BeforeRequest(req, "security.bulk_update_api_keys")
		if reader := instrument.RecordRequestBody(ctx, "security.bulk_update_api_keys", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.bulk_update_api_keys")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the BulkUpdateApiKeys query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a bulkupdateapikeys.Response
func (r BulkUpdateApiKeys) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.bulk_update_api_keys")
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

// Header set a key, value pair in the BulkUpdateApiKeys headers map.
func (r *BulkUpdateApiKeys) Header(key, value string) *BulkUpdateApiKeys {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *BulkUpdateApiKeys) ErrorTrace(errortrace bool) *BulkUpdateApiKeys {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *BulkUpdateApiKeys) FilterPath(filterpaths ...string) *BulkUpdateApiKeys {
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
func (r *BulkUpdateApiKeys) Human(human bool) *BulkUpdateApiKeys {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *BulkUpdateApiKeys) Pretty(pretty bool) *BulkUpdateApiKeys {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Expiration Expiration time for the API keys.
// By default, API keys never expire.
// This property can be omitted to leave the value unchanged.
// API name: expiration
func (r *BulkUpdateApiKeys) Expiration(duration types.Duration) *BulkUpdateApiKeys {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Expiration = duration

	return r
}

// Ids The API key identifiers.
// API name: ids
func (r *BulkUpdateApiKeys) Ids(ids ...string) *BulkUpdateApiKeys {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Ids = ids

	return r
}

// Metadata Arbitrary nested metadata to associate with the API keys.
// Within the `metadata` object, top-level keys beginning with an underscore
// (`_`) are reserved for system usage.
// Any information specified with this parameter fully replaces metadata
// previously associated with the API key.
// API name: metadata
func (r *BulkUpdateApiKeys) Metadata(metadata types.Metadata) *BulkUpdateApiKeys {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Metadata = metadata

	return r
}

// RoleDescriptors The role descriptors to assign to the API keys.
// An API key's effective permissions are an intersection of its assigned
// privileges and the point-in-time snapshot of permissions of the owner user.
// You can assign new privileges by specifying them in this parameter.
// To remove assigned privileges, supply the `role_descriptors` parameter as an
// empty object `{}`.
// If an API key has no assigned privileges, it inherits the owner user's full
// permissions.
// The snapshot of the owner's permissions is always updated, whether you supply
// the `role_descriptors` parameter.
// The structure of a role descriptor is the same as the request for the create
// API keys API.
// API name: role_descriptors
func (r *BulkUpdateApiKeys) RoleDescriptors(roledescriptors map[string]types.RoleDescriptor) *BulkUpdateApiKeys {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RoleDescriptors = roledescriptors

	return r
}
