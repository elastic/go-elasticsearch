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

// Check user profile privileges.
//
// Determine whether the users associated with the specified user profile IDs
// have all the requested privileges.
//
// NOTE: The user profile feature is designed only for use by Kibana and
// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
// Individual users and external applications should not call this API directly.
// Elastic reserves the right to change or remove this feature in future
// releases without prior notice.
package hasprivilegesuserprofile

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

type HasPrivilegesUserProfile struct {
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

// NewHasPrivilegesUserProfile type alias for index.
type NewHasPrivilegesUserProfile func() *HasPrivilegesUserProfile

// NewHasPrivilegesUserProfileFunc returns a new instance of HasPrivilegesUserProfile with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewHasPrivilegesUserProfileFunc(tp elastictransport.Interface) NewHasPrivilegesUserProfile {
	return func() *HasPrivilegesUserProfile {
		n := New(tp)

		return n
	}
}

// Check user profile privileges.
//
// Determine whether the users associated with the specified user profile IDs
// have all the requested privileges.
//
// NOTE: The user profile feature is designed only for use by Kibana and
// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
// Individual users and external applications should not call this API directly.
// Elastic reserves the right to change or remove this feature in future
// releases without prior notice.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-has-privileges-user-profile.html
func New(tp elastictransport.Interface) *HasPrivilegesUserProfile {
	r := &HasPrivilegesUserProfile{
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
func (r *HasPrivilegesUserProfile) Raw(raw io.Reader) *HasPrivilegesUserProfile {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *HasPrivilegesUserProfile) Request(req *Request) *HasPrivilegesUserProfile {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *HasPrivilegesUserProfile) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for HasPrivilegesUserProfile: %w", err)
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
		path.WriteString("profile")
		path.WriteString("/")
		path.WriteString("_has_privileges")

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
func (r HasPrivilegesUserProfile) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.has_privileges_user_profile")
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
		instrument.BeforeRequest(req, "security.has_privileges_user_profile")
		if reader := instrument.RecordRequestBody(ctx, "security.has_privileges_user_profile", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.has_privileges_user_profile")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the HasPrivilegesUserProfile query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a hasprivilegesuserprofile.Response
func (r HasPrivilegesUserProfile) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.has_privileges_user_profile")
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

// Header set a key, value pair in the HasPrivilegesUserProfile headers map.
func (r *HasPrivilegesUserProfile) Header(key, value string) *HasPrivilegesUserProfile {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *HasPrivilegesUserProfile) ErrorTrace(errortrace bool) *HasPrivilegesUserProfile {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *HasPrivilegesUserProfile) FilterPath(filterpaths ...string) *HasPrivilegesUserProfile {
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
func (r *HasPrivilegesUserProfile) Human(human bool) *HasPrivilegesUserProfile {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *HasPrivilegesUserProfile) Pretty(pretty bool) *HasPrivilegesUserProfile {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Privileges An object containing all the privileges to be checked.
// API name: privileges
func (r *HasPrivilegesUserProfile) Privileges(privileges *types.PrivilegesCheck) *HasPrivilegesUserProfile {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Privileges = *privileges

	return r
}

// Uids A list of profile IDs. The privileges are checked for associated users of the
// profiles.
// API name: uids
func (r *HasPrivilegesUserProfile) Uids(uids ...string) *HasPrivilegesUserProfile {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Uids = uids

	return r
}
