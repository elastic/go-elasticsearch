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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

// Create user-managed service accounts.
//
// Create a service account in a namespace of your own, or replace one that
// already exists. A replacement is not a partial update: every write applies
// the defaults, so an account that was disabled and is then written again
// without `enabled` comes back enabled.
//
// Creating an account whose name still has leftover service tokens is rejected.
// Delete those tokens first.
//
// NOTE: The `elastic` namespace is reserved for the built-in service accounts
// that ship with Elasticsearch. The `manage_service_account` privilege does not
// authorize this API.
package putusermanagedserviceaccount

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

const (
	namespaceMask = iota + 1

	serviceMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutUserManagedServiceAccount struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	namespace string
	service   string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutUserManagedServiceAccount type alias for index.
type NewPutUserManagedServiceAccount func(namespace, service string) *PutUserManagedServiceAccount

// NewPutUserManagedServiceAccountFunc returns a new instance of PutUserManagedServiceAccount with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutUserManagedServiceAccountFunc(tp elastictransport.Interface) NewPutUserManagedServiceAccount {
	return func(namespace, service string) *PutUserManagedServiceAccount {
		n := New(tp)

		n._namespace(namespace)

		n._service(service)

		return n
	}
}

// Create user-managed service accounts.
//
// Create a service account in a namespace of your own, or replace one that
// already exists. A replacement is not a partial update: every write applies
// the defaults, so an account that was disabled and is then written again
// without `enabled` comes back enabled.
//
// Creating an account whose name still has leftover service tokens is rejected.
// Delete those tokens first.
//
// NOTE: The `elastic` namespace is reserved for the built-in service accounts
// that ship with Elasticsearch. The `manage_service_account` privilege does not
// authorize this API.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-security-put-user-managed-service-account
func New(tp elastictransport.Interface) *PutUserManagedServiceAccount {
	r := &PutUserManagedServiceAccount{
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
func (r *PutUserManagedServiceAccount) Raw(raw io.Reader) *PutUserManagedServiceAccount {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutUserManagedServiceAccount) Request(req *Request) *PutUserManagedServiceAccount {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutUserManagedServiceAccount) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutUserManagedServiceAccount: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == namespaceMask|serviceMask:
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
func (r PutUserManagedServiceAccount) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx = instrument.Start(providedCtx, "security.put_user_managed_service_account")
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
		instrument.BeforeRequest(req, "security.put_user_managed_service_account")
		if reader := instrument.RecordRequestBody(ctx, "security.put_user_managed_service_account", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.put_user_managed_service_account")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutUserManagedServiceAccount query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putusermanagedserviceaccount.Response
func (r PutUserManagedServiceAccount) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.put_user_managed_service_account")
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

// Header set a key, value pair in the PutUserManagedServiceAccount headers map.
func (r *PutUserManagedServiceAccount) Header(key, value string) *PutUserManagedServiceAccount {
	r.headers.Set(key, value)

	return r
}

// Namespace The namespace, which is a top-level grouping of service accounts. It must
// start with a letter or digit and can contain only letters, digits, hyphens,
// and underscores, up to a maximum of 128 characters. It cannot be `elastic`,
// which is reserved for built-in service accounts.
// API Name: namespace
func (r *PutUserManagedServiceAccount) _namespace(namespace string) *PutUserManagedServiceAccount {
	r.paramSet |= namespaceMask
	r.namespace = namespace

	return r
}

// Service The service name. It must start with a letter or digit and can contain only
// letters, digits, hyphens, and underscores, up to a maximum of 128 characters.
// API Name: service
func (r *PutUserManagedServiceAccount) _service(service string) *PutUserManagedServiceAccount {
	r.paramSet |= serviceMask
	r.service = service

	return r
}

// Refresh If `wait_for` (the default) then wait for a refresh to make this operation
// visible to search, if `true` then refresh the affected shards to make this
// operation visible to search, if `false` then do nothing with refreshes.
// API name: refresh
func (r *PutUserManagedServiceAccount) Refresh(refresh refresh.Refresh) *PutUserManagedServiceAccount {
	r.values.Set("refresh", refresh.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutUserManagedServiceAccount) ErrorTrace(errortrace bool) *PutUserManagedServiceAccount {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutUserManagedServiceAccount) FilterPath(filterpaths ...string) *PutUserManagedServiceAccount {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and `"exists_time_in_millis":
// 3600000` for computers. When disabled the human readable values will be
// omitted. This makes sense for responses being consumed only by machines.
// API name: human
func (r *PutUserManagedServiceAccount) Human(human bool) *PutUserManagedServiceAccount {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use this
// option for debugging only.
// API name: pretty
func (r *PutUserManagedServiceAccount) Pretty(pretty bool) *PutUserManagedServiceAccount {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// A free-text description of the account, as sent on the last PUT of the
// account. It has no meaning to Elasticsearch. Absent when the account has no
// description.
// API name: description
func (r *PutUserManagedServiceAccount) Description(description string) *PutUserManagedServiceAccount {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Description = &description

	return r
}

// Whether the account can authenticate. Tokens can still be created for a
// disabled account; they just cannot be used until the account is enabled.
// API name: enabled
func (r *PutUserManagedServiceAccount) Enabled(enabled bool) *PutUserManagedServiceAccount {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Enabled = &enabled

	return r
}

// The names of the roles to grant to the service account, up to a maximum of
// 1000. The roles are resolved when the account authenticates, so they do not
// have to exist yet.
// API name: roles
func (r *PutUserManagedServiceAccount) Roles(roles ...string) *PutUserManagedServiceAccount {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range roles {

		r.req.Roles = append(r.req.Roles, v)

	}
	return r
}
