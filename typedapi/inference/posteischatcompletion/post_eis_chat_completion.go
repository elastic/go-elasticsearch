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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

// Perform a chat completion task through the Elastic Inference Service (EIS).
//
// Perform a chat completion inference task with the `elastic` service.
package posteischatcompletion

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
	eisinferenceidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PostEisChatCompletion struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	eisinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPostEisChatCompletion type alias for index.
type NewPostEisChatCompletion func(eisinferenceid string) *PostEisChatCompletion

// NewPostEisChatCompletionFunc returns a new instance of PostEisChatCompletion with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPostEisChatCompletionFunc(tp elastictransport.Interface) NewPostEisChatCompletion {
	return func(eisinferenceid string) *PostEisChatCompletion {
		n := New(tp)

		n._eisinferenceid(eisinferenceid)

		return n
	}
}

// Perform a chat completion task through the Elastic Inference Service (EIS).
//
// Perform a chat completion inference task with the `elastic` service.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-inference-post-eis-chat-completion
func New(tp elastictransport.Interface) *PostEisChatCompletion {
	r := &PostEisChatCompletion{
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
func (r *PostEisChatCompletion) Raw(raw io.Reader) *PostEisChatCompletion {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PostEisChatCompletion) Request(req *Request) *PostEisChatCompletion {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PostEisChatCompletion) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PostEisChatCompletion: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == eisinferenceidMask:
		path.WriteString("/")
		path.WriteString("_inference")
		path.WriteString("/")
		path.WriteString("chat_completion")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "eisinferenceid", r.eisinferenceid)
		}
		path.WriteString(r.eisinferenceid)
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
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r PostEisChatCompletion) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.post_eis_chat_completion")
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
		instrument.BeforeRequest(req, "inference.post_eis_chat_completion")
		if reader := instrument.RecordRequestBody(ctx, "inference.post_eis_chat_completion", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.post_eis_chat_completion")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PostEisChatCompletion query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a posteischatcompletion.Response
func (r PostEisChatCompletion) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.post_eis_chat_completion")
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

// Header set a key, value pair in the PostEisChatCompletion headers map.
func (r *PostEisChatCompletion) Header(key, value string) *PostEisChatCompletion {
	r.headers.Set(key, value)

	return r
}

// EisInferenceId The unique identifier of the inference endpoint.
// API Name: eisinferenceid
func (r *PostEisChatCompletion) _eisinferenceid(eisinferenceid string) *PostEisChatCompletion {
	r.paramSet |= eisinferenceidMask
	r.eisinferenceid = eisinferenceid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PostEisChatCompletion) ErrorTrace(errortrace bool) *PostEisChatCompletion {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PostEisChatCompletion) FilterPath(filterpaths ...string) *PostEisChatCompletion {
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
func (r *PostEisChatCompletion) Human(human bool) *PostEisChatCompletion {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PostEisChatCompletion) Pretty(pretty bool) *PostEisChatCompletion {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// The upper bound limit for the number of tokens that can be generated for a
// completion request.
// API name: max_completion_tokens
func (r *PostEisChatCompletion) MaxCompletionTokens(maxcompletiontokens int64) *PostEisChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.MaxCompletionTokens = &maxcompletiontokens

	return r
}

// A list of objects representing the conversation.
// API name: messages
func (r *PostEisChatCompletion) Messages(messages ...types.MessageVariant) *PostEisChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range messages {

		r.req.Messages = append(r.req.Messages, *v.MessageCaster())

	}
	return r
}

// The ID of the model to use.
// API name: model
func (r *PostEisChatCompletion) Model(model string) *PostEisChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Model = &model

	return r
}

// A sequence of strings to control when the model should stop generating
// additional tokens.
// API name: stop
func (r *PostEisChatCompletion) Stop(stops ...string) *PostEisChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range stops {

		r.req.Stop = append(r.req.Stop, v)

	}
	return r
}

// The sampling temperature to use.
// API name: temperature
func (r *PostEisChatCompletion) Temperature(temperature float32) *PostEisChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Temperature = &temperature

	return r
}

// Controls which tool is called by the model.
// API name: tool_choice
func (r *PostEisChatCompletion) ToolChoice(completiontooltype types.CompletionToolTypeVariant) *PostEisChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ToolChoice = *completiontooltype.CompletionToolTypeCaster()

	return r
}

// A list of tools that the model can call.
// API name: tools
func (r *PostEisChatCompletion) Tools(tools ...types.CompletionToolVariant) *PostEisChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range tools {

		r.req.Tools = append(r.req.Tools, *v.CompletionToolCaster())

	}
	return r
}

// Nucleus sampling, an alternative to sampling with temperature.
// API name: top_p
func (r *PostEisChatCompletion) TopP(topp float32) *PostEisChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TopP = &topp

	return r
}
