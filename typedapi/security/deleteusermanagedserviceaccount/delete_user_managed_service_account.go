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

// Delete user-managed service accounts.
//
// Delete a service account from a namespace of your own.
//
// Deleting an account that still has service tokens is rejected unless `force`
// is `true`. A forced delete leaves the tokens behind: they cannot authenticate
// while no account of that name exists, and recreating the account is rejected
// until they are deleted.
//
// NOTE: The `elastic` namespace is reserved for the built-in service accounts
// that ship with Elasticsearch. A name that no user-managed service account
// could have is rejected rather than reported as not found. The
// `manage_service_account` privilege does not authorize this API.
package deleteusermanagedserviceaccount

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

const (
	namespaceMask = iota + 1

	serviceMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteUserManagedServiceAccount struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	namespace string
	service   string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewDeleteUserManagedServiceAccount type alias for index.
type NewDeleteUserManagedServiceAccount func(namespace, service string) *DeleteUserManagedServiceAccount

// NewDeleteUserManagedServiceAccountFunc returns a new instance of DeleteUserManagedServiceAccount with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteUserManagedServiceAccountFunc(tp elastictransport.Interface) NewDeleteUserManagedServiceAccount {
	return func(namespace, service string) *DeleteUserManagedServiceAccount {
		n := New(tp)

		n._namespace(namespace)

		n._service(service)

		return n
	}
}

// Delete user-managed service accounts.
//
// Delete a service account from a namespace of your own.
//
// Deleting an account that still has service tokens is rejected unless `force`
// is `true`. A forced delete leaves the tokens behind: they cannot authenticate
// while no account of that name exists, and recreating the account is rejected
// until they are deleted.
//
// NOTE: The `elastic` namespace is reserved for the built-in service accounts
// that ship with Elasticsearch. A name that no user-managed service account
// could have is rejected rather than reported as not found. The
// `manage_service_account` privilege does not authorize this API.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-security-delete-user-managed-service-account
func New(tp elastictransport.Interface) *DeleteUserManagedServiceAccount {
	r := &DeleteUserManagedServiceAccount{
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
func (r *DeleteUserManagedServiceAccount) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r DeleteUserManagedServiceAccount) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx = instrument.Start(providedCtx, "security.delete_user_managed_service_account")
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
		instrument.BeforeRequest(req, "security.delete_user_managed_service_account")
		if reader := instrument.RecordRequestBody(ctx, "security.delete_user_managed_service_account", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.delete_user_managed_service_account")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the DeleteUserManagedServiceAccount query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a deleteusermanagedserviceaccount.Response
func (r DeleteUserManagedServiceAccount) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.delete_user_managed_service_account")
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
func (r DeleteUserManagedServiceAccount) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.delete_user_managed_service_account")
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
		err := fmt.Errorf("an error happened during the DeleteUserManagedServiceAccount query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the DeleteUserManagedServiceAccount headers map.
func (r *DeleteUserManagedServiceAccount) Header(key, value string) *DeleteUserManagedServiceAccount {
	r.headers.Set(key, value)

	return r
}

// Namespace The namespace, which is a top-level grouping of service accounts. It must
// start with a letter or digit and can contain only letters, digits, hyphens,
// and underscores, up to a maximum of 128 characters. It cannot be `elastic`,
// which is reserved for built-in service accounts.
// API Name: namespace
func (r *DeleteUserManagedServiceAccount) _namespace(namespace string) *DeleteUserManagedServiceAccount {
	r.paramSet |= namespaceMask
	r.namespace = namespace

	return r
}

// Service The service name. It must start with a letter or digit and can contain only
// letters, digits, hyphens, and underscores, up to a maximum of 128 characters.
// API Name: service
func (r *DeleteUserManagedServiceAccount) _service(service string) *DeleteUserManagedServiceAccount {
	r.paramSet |= serviceMask
	r.service = service

	return r
}

// Refresh If `wait_for` (the default) then wait for a refresh to make this operation
// visible to search, if `true` then refresh the affected shards to make this
// operation visible to search, if `false` then do nothing with refreshes.
// API name: refresh
func (r *DeleteUserManagedServiceAccount) Refresh(refresh refresh.Refresh) *DeleteUserManagedServiceAccount {
	r.values.Set("refresh", refresh.String())

	return r
}

// Force If `false` (the default), deleting a service account that still has service
// tokens is rejected. If `true`, the account is deleted and its tokens are left
// in place.
// API name: force
func (r *DeleteUserManagedServiceAccount) Force(force bool) *DeleteUserManagedServiceAccount {
	r.values.Set("force", strconv.FormatBool(force))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *DeleteUserManagedServiceAccount) ErrorTrace(errortrace bool) *DeleteUserManagedServiceAccount {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *DeleteUserManagedServiceAccount) FilterPath(filterpaths ...string) *DeleteUserManagedServiceAccount {
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
func (r *DeleteUserManagedServiceAccount) Human(human bool) *DeleteUserManagedServiceAccount {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use this
// option for debugging only.
// API name: pretty
func (r *DeleteUserManagedServiceAccount) Pretty(pretty bool) *DeleteUserManagedServiceAccount {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
