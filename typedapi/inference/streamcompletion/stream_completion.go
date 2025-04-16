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
// https://github.com/elastic/elasticsearch-specification/tree/f6a370d0fba975752c644fc730f7c45610e28f36

// Perform streaming inference.
// Get real-time responses for completion tasks by delivering answers
// incrementally, reducing response times during computation.
// This API works only with the completion task type.
//
// IMPORTANT: The inference APIs enable you to use certain services, such as
// built-in machine learning models (ELSER, E5), models uploaded through Eland,
// Cohere, OpenAI, Azure, Google AI Studio, Google Vertex AI, Anthropic,
// Watsonx.ai, or Hugging Face. For built-in models and models uploaded through
// Eland, the inference APIs offer an alternative way to use and manage trained
// models. However, if you do not plan to use the inference APIs to use these
// models or if you want to use non-NLP models, use the machine learning trained
// model APIs.
//
// This API requires the `monitor_inference` cluster privilege (the built-in
// `inference_admin` and `inference_user` roles grant this privilege). You must
// use a client that supports streaming.
package streamcompletion

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

const (
	inferenceidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type StreamCompletion struct {
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

// NewStreamCompletion type alias for index.
type NewStreamCompletion func(inferenceid string) *StreamCompletion

// NewStreamCompletionFunc returns a new instance of StreamCompletion with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStreamCompletionFunc(tp elastictransport.Interface) NewStreamCompletion {
	return func(inferenceid string) *StreamCompletion {
		n := New(tp)

		n._inferenceid(inferenceid)

		return n
	}
}

// Perform streaming inference.
// Get real-time responses for completion tasks by delivering answers
// incrementally, reducing response times during computation.
// This API works only with the completion task type.
//
// IMPORTANT: The inference APIs enable you to use certain services, such as
// built-in machine learning models (ELSER, E5), models uploaded through Eland,
// Cohere, OpenAI, Azure, Google AI Studio, Google Vertex AI, Anthropic,
// Watsonx.ai, or Hugging Face. For built-in models and models uploaded through
// Eland, the inference APIs offer an alternative way to use and manage trained
// models. However, if you do not plan to use the inference APIs to use these
// models or if you want to use non-NLP models, use the machine learning trained
// model APIs.
//
// This API requires the `monitor_inference` cluster privilege (the built-in
// `inference_admin` and `inference_user` roles grant this privilege). You must
// use a client that supports streaming.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/stream-inference-api.html
func New(tp elastictransport.Interface) *StreamCompletion {
	r := &StreamCompletion{
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
func (r *StreamCompletion) Raw(raw io.Reader) *StreamCompletion {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *StreamCompletion) Request(req *Request) *StreamCompletion {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *StreamCompletion) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for StreamCompletion: %w", err)
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
		path.WriteString("completion")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "inferenceid", r.inferenceid)
		}
		path.WriteString(r.inferenceid)
		path.WriteString("/")
		path.WriteString("_stream")

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
		req.Header.Set("Accept", "text/event-stream")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r StreamCompletion) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.stream_completion")
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
		instrument.BeforeRequest(req, "inference.stream_completion")
		if reader := instrument.RecordRequestBody(ctx, "inference.stream_completion", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.stream_completion")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the StreamCompletion query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a streamcompletion.Response
func (r StreamCompletion) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.stream_completion")
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
		response, err = io.ReadAll(res.Body)
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

// Header set a key, value pair in the StreamCompletion headers map.
func (r *StreamCompletion) Header(key, value string) *StreamCompletion {
	r.headers.Set(key, value)

	return r
}

// InferenceId The unique identifier for the inference endpoint.
// API Name: inferenceid
func (r *StreamCompletion) _inferenceid(inferenceid string) *StreamCompletion {
	r.paramSet |= inferenceidMask
	r.inferenceid = inferenceid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *StreamCompletion) ErrorTrace(errortrace bool) *StreamCompletion {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *StreamCompletion) FilterPath(filterpaths ...string) *StreamCompletion {
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
func (r *StreamCompletion) Human(human bool) *StreamCompletion {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *StreamCompletion) Pretty(pretty bool) *StreamCompletion {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// The text on which you want to perform the inference task.
// It can be a single string or an array.
//
// NOTE: Inference endpoints for the completion task type currently only support
// a single string as input.
// API name: input
func (r *StreamCompletion) Input(inputs ...string) *StreamCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Input = make([]string, len(inputs))
	r.req.Input = inputs

	return r
}

// Optional task settings
// API name: task_settings
func (r *StreamCompletion) TaskSettings(tasksettings json.RawMessage) *StreamCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TaskSettings = tasksettings

	return r
}
