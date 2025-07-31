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

// Suggest a user profile.
//
// Get suggestions for user profiles that match specified search criteria.
//
// NOTE: The user profile feature is designed only for use by Kibana and
// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
// Individual users and external applications should not call this API directly.
// Elastic reserves the right to change or remove this feature in future
// releases without prior notice.
package suggestuserprofiles

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

type SuggestUserProfiles struct {
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

// NewSuggestUserProfiles type alias for index.
type NewSuggestUserProfiles func() *SuggestUserProfiles

// NewSuggestUserProfilesFunc returns a new instance of SuggestUserProfiles with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSuggestUserProfilesFunc(tp elastictransport.Interface) NewSuggestUserProfiles {
	return func() *SuggestUserProfiles {
		n := New(tp)

		return n
	}
}

// Suggest a user profile.
//
// Get suggestions for user profiles that match specified search criteria.
//
// NOTE: The user profile feature is designed only for use by Kibana and
// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
// Individual users and external applications should not call this API directly.
// Elastic reserves the right to change or remove this feature in future
// releases without prior notice.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-suggest-user-profile.html
func New(tp elastictransport.Interface) *SuggestUserProfiles {
	r := &SuggestUserProfiles{
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
func (r *SuggestUserProfiles) Raw(raw io.Reader) *SuggestUserProfiles {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *SuggestUserProfiles) Request(req *Request) *SuggestUserProfiles {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *SuggestUserProfiles) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for SuggestUserProfiles: %w", err)
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
		path.WriteString("_suggest")

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
func (r SuggestUserProfiles) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.suggest_user_profiles")
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
		instrument.BeforeRequest(req, "security.suggest_user_profiles")
		if reader := instrument.RecordRequestBody(ctx, "security.suggest_user_profiles", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.suggest_user_profiles")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the SuggestUserProfiles query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a suggestuserprofiles.Response
func (r SuggestUserProfiles) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.suggest_user_profiles")
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

// Header set a key, value pair in the SuggestUserProfiles headers map.
func (r *SuggestUserProfiles) Header(key, value string) *SuggestUserProfiles {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *SuggestUserProfiles) ErrorTrace(errortrace bool) *SuggestUserProfiles {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *SuggestUserProfiles) FilterPath(filterpaths ...string) *SuggestUserProfiles {
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
func (r *SuggestUserProfiles) Human(human bool) *SuggestUserProfiles {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *SuggestUserProfiles) Pretty(pretty bool) *SuggestUserProfiles {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Data A comma-separated list of filters for the `data` field of the profile
// document.
// To return all content use `data=*`.
// To return a subset of content, use `data=<key>` to retrieve content nested
// under the specified `<key>`.
// By default, the API returns no `data` content.
// It is an error to specify `data` as both the query parameter and the request
// body field.
// API name: data
func (r *SuggestUserProfiles) Data(data ...string) *SuggestUserProfiles {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Data = data

	return r
}

// Hint Extra search criteria to improve relevance of the suggestion result.
// Profiles matching the spcified hint are ranked higher in the response.
// Profiles not matching the hint aren't excluded from the response as long as
// the profile matches the `name` field query.
// API name: hint
func (r *SuggestUserProfiles) Hint(hint *types.Hint) *SuggestUserProfiles {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Hint = hint

	return r
}

// Name A query string used to match name-related fields in user profile documents.
// Name-related fields are the user's `username`, `full_name`, and `email`.
// API name: name
func (r *SuggestUserProfiles) Name(name string) *SuggestUserProfiles {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Name = &name

	return r
}

// Size The number of profiles to return.
// API name: size
func (r *SuggestUserProfiles) Size(size int64) *SuggestUserProfiles {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Size = &size

	return r
}
