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

// Perform decision inference on the service.
//
// Answer a set of named questions about an input, such as whether a customer is
// complaining, which team should handle a request, or how angry a customer is.
// The questions are defined in the task settings of the inference endpoint, and
// can be replaced for a single request through `task_settings`.
//
// > info > Only the `elastic` service currently supports the `decision` task.
package decision

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
)

const (
	inferenceidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Decision struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	inferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewDecision type alias for index.
type NewDecision func(inferenceid string) *Decision

// NewDecisionFunc returns a new instance of Decision with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDecisionFunc(tp elastictransport.Interface) NewDecision {
	return func(inferenceid string) *Decision {
		n := New(tp)

		n._inferenceid(inferenceid)

		return n
	}
}

// Perform decision inference on the service.
//
// Answer a set of named questions about an input, such as whether a customer is
// complaining, which team should handle a request, or how angry a customer is.
// The questions are defined in the task settings of the inference endpoint, and
// can be replaced for a single request through `task_settings`.
//
// > info > Only the `elastic` service currently supports the `decision` task.
//
// [Elasticsearch] https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-inference-inference
//
// [Serverless] https://www.elastic.co/docs/api/doc/elasticsearch-serverless/operation/operation-inference-inference
func New(tp elastictransport.Interface) *Decision {
	r := &Decision{
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
func (r *Decision) Raw(raw io.Reader) *Decision {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Decision) Request(req *Request) *Decision {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Decision) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Decision: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == inferenceidMask:
		path.WriteString("/")
		path.WriteString("_inference")
		path.WriteString("/")
		path.WriteString("decision")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "inferenceid", r.inferenceid)
		}
		path.WriteString(r.inferenceid)

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
func (r Decision) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx = instrument.Start(providedCtx, "inference.decision")
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
		instrument.BeforeRequest(req, "inference.decision")
		if reader := instrument.RecordRequestBody(ctx, "inference.decision", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.decision")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Decision query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a decision.Response
func (r Decision) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.decision")
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

// Header set a key, value pair in the Decision headers map.
func (r *Decision) Header(key, value string) *Decision {
	r.headers.Set(key, value)

	return r
}

// InferenceId The unique identifier for the inference endpoint.
// API Name: inferenceid
func (r *Decision) _inferenceid(inferenceid string) *Decision {
	r.paramSet |= inferenceidMask
	r.inferenceid = inferenceid

	return r
}

// Timeout The amount of time to wait for the inference request to complete.
// API name: timeout
func (r *Decision) Timeout(duration string) *Decision {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Decision) ErrorTrace(errortrace bool) *Decision {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Decision) FilterPath(filterpaths ...string) *Decision {
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
func (r *Decision) Human(human bool) *Decision {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use this
// option for debugging only.
// API name: pretty
func (r *Decision) Pretty(pretty bool) *Decision {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// The input to answer the questions about. It can be specified as a string, or
// as an object. The object form additionally allows specifying images.
//
// string example:
//
//	"input": "I was charged twice for my order and nobody answers my emails."
//
// object example:
//
//	"input": {
//	  "type": "image",
//	  "format": "base64",
//	  "value": "data:image/jpeg;base64,..."
//	}
//
// API name: input
func (r *Decision) Input(decisioninput types.DecisionInputVariant) *Decision {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Input = *decisioninput.DecisionInputCaster()

	return r
}

// Task settings for the individual inference request. These settings override
// the task settings specified when creating the inference endpoint.
// API name: task_settings
func (r *Decision) TaskSettings(tasksettings types.DecisionTaskSettingsVariant) *Decision {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TaskSettings = tasksettings.DecisionTaskSettingsCaster()

	return r
}
