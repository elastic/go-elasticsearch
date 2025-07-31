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

// Create a service account token.
//
// Create a service accounts token for access without requiring basic
// authentication.
//
// NOTE: Service account tokens never expire.
// You must actively delete them if they are no longer needed.
package createservicetoken

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

const (
	namespaceMask = iota + 1

	serviceMask

	nameMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CreateServiceToken struct {
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

// NewCreateServiceToken type alias for index.
type NewCreateServiceToken func(namespace, service string) *CreateServiceToken

// NewCreateServiceTokenFunc returns a new instance of CreateServiceToken with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCreateServiceTokenFunc(tp elastictransport.Interface) NewCreateServiceToken {
	return func(namespace, service string) *CreateServiceToken {
		n := New(tp)

		n._namespace(namespace)

		n._service(service)

		return n
	}
}

// Create a service account token.
//
// Create a service accounts token for access without requiring basic
// authentication.
//
// NOTE: Service account tokens never expire.
// You must actively delete them if they are no longer needed.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-service-token.html
func New(tp elastictransport.Interface) *CreateServiceToken {
	r := &CreateServiceToken{
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
func (r *CreateServiceToken) HttpRequest(ctx context.Context) (*http.Request, error) {
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

		method = http.MethodPut
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
		path.WriteString("/")
		path.WriteString("credential")
		path.WriteString("/")
		path.WriteString("token")

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
func (r CreateServiceToken) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.create_service_token")
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
		instrument.BeforeRequest(req, "security.create_service_token")
		if reader := instrument.RecordRequestBody(ctx, "security.create_service_token", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.create_service_token")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the CreateServiceToken query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a createservicetoken.Response
func (r CreateServiceToken) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.create_service_token")
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
func (r CreateServiceToken) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.create_service_token")
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
		err := fmt.Errorf("an error happened during the CreateServiceToken query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the CreateServiceToken headers map.
func (r *CreateServiceToken) Header(key, value string) *CreateServiceToken {
	r.headers.Set(key, value)

	return r
}

// Namespace The name of the namespace, which is a top-level grouping of service accounts.
// API Name: namespace
func (r *CreateServiceToken) _namespace(namespace string) *CreateServiceToken {
	r.paramSet |= namespaceMask
	r.namespace = namespace

	return r
}

// Service The name of the service.
// API Name: service
func (r *CreateServiceToken) _service(service string) *CreateServiceToken {
	r.paramSet |= serviceMask
	r.service = service

	return r
}

// Name The name for the service account token.
// If omitted, a random name will be generated.
//
// Token names must be at least one and no more than 256 characters.
// They can contain alphanumeric characters (a-z, A-Z, 0-9), dashes (`-`), and
// underscores (`_`), but cannot begin with an underscore.
//
// NOTE: Token names must be unique in the context of the associated service
// account.
// They must also be globally unique with their fully qualified names, which are
// comprised of the service account principal and token name, such as
// `<namespace>/<service>/<token-name>`.
// API Name: name
func (r *CreateServiceToken) Name(name string) *CreateServiceToken {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// Refresh If `true` then refresh the affected shards to make this operation visible to
// search, if `wait_for` (the default) then wait for a refresh to make this
// operation visible to search, if `false` then do nothing with refreshes.
// API name: refresh
func (r *CreateServiceToken) Refresh(refresh refresh.Refresh) *CreateServiceToken {
	r.values.Set("refresh", refresh.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *CreateServiceToken) ErrorTrace(errortrace bool) *CreateServiceToken {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *CreateServiceToken) FilterPath(filterpaths ...string) *CreateServiceToken {
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
func (r *CreateServiceToken) Human(human bool) *CreateServiceToken {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *CreateServiceToken) Pretty(pretty bool) *CreateServiceToken {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
