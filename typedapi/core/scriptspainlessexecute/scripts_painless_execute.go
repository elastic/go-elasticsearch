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
// https://github.com/elastic/elasticsearch-specification/tree/f1932ce6b46a53a8342db522b1a7883bcc9e0996

// Run a script.
//
// Runs a script and returns a result.
// Use this API to build and test scripts, such as when defining a script for a
// runtime field.
// This API requires very few dependencies and is especially useful if you don't
// have permissions to write documents on a cluster.
//
// The API uses several _contexts_, which control how scripts are run, what
// variables are available at runtime, and what the return type is.
//
// Each context requires a script, but additional parameters depend on the
// context you're using for that script.
package scriptspainlessexecute

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/painlesscontext"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ScriptsPainlessExecute struct {
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

// NewScriptsPainlessExecute type alias for index.
type NewScriptsPainlessExecute func() *ScriptsPainlessExecute

// NewScriptsPainlessExecuteFunc returns a new instance of ScriptsPainlessExecute with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewScriptsPainlessExecuteFunc(tp elastictransport.Interface) NewScriptsPainlessExecute {
	return func() *ScriptsPainlessExecute {
		n := New(tp)

		return n
	}
}

// Run a script.
//
// Runs a script and returns a result.
// Use this API to build and test scripts, such as when defining a script for a
// runtime field.
// This API requires very few dependencies and is especially useful if you don't
// have permissions to write documents on a cluster.
//
// The API uses several _contexts_, which control how scripts are run, what
// variables are available at runtime, and what the return type is.
//
// Each context requires a script, but additional parameters depend on the
// context you're using for that script.
//
// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
func New(tp elastictransport.Interface) *ScriptsPainlessExecute {
	r := &ScriptsPainlessExecute{
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
func (r *ScriptsPainlessExecute) Raw(raw io.Reader) *ScriptsPainlessExecute {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *ScriptsPainlessExecute) Request(req *Request) *ScriptsPainlessExecute {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ScriptsPainlessExecute) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for ScriptsPainlessExecute: %w", err)
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
		path.WriteString("_scripts")
		path.WriteString("/")
		path.WriteString("painless")
		path.WriteString("/")
		path.WriteString("_execute")

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
func (r ScriptsPainlessExecute) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "scripts_painless_execute")
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
		instrument.BeforeRequest(req, "scripts_painless_execute")
		if reader := instrument.RecordRequestBody(ctx, "scripts_painless_execute", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "scripts_painless_execute")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ScriptsPainlessExecute query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a scriptspainlessexecute.Response
func (r ScriptsPainlessExecute) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "scripts_painless_execute")
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

// Header set a key, value pair in the ScriptsPainlessExecute headers map.
func (r *ScriptsPainlessExecute) Header(key, value string) *ScriptsPainlessExecute {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ScriptsPainlessExecute) ErrorTrace(errortrace bool) *ScriptsPainlessExecute {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ScriptsPainlessExecute) FilterPath(filterpaths ...string) *ScriptsPainlessExecute {
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
func (r *ScriptsPainlessExecute) Human(human bool) *ScriptsPainlessExecute {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ScriptsPainlessExecute) Pretty(pretty bool) *ScriptsPainlessExecute {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// The context that the script should run in.
// NOTE: Result ordering in the field contexts is not guaranteed.
// API name: context
func (r *ScriptsPainlessExecute) Context(context painlesscontext.PainlessContext) *ScriptsPainlessExecute {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Context = &context
	return r
}

// Additional parameters for the `context`.
// NOTE: This parameter is required for all contexts except `painless_test`,
// which is the default if no value is provided for `context`.
// API name: context_setup
func (r *ScriptsPainlessExecute) ContextSetup(contextsetup types.PainlessContextSetupVariant) *ScriptsPainlessExecute {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ContextSetup = contextsetup.PainlessContextSetupCaster()

	return r
}

// The Painless script to run.
// API name: script
func (r *ScriptsPainlessExecute) Script(script types.ScriptVariant) *ScriptsPainlessExecute {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Script = script.ScriptCaster()

	return r
}
