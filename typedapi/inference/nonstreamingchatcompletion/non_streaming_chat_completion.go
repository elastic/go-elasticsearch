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

// Perform non-streaming chat completion inference on the service.
//
// The chat completion inference API enables rich responses for chat completion
// tasks. It only works with the `chat_completion` task type.
//
// NOTE: The `chat_completion` task type supports both streaming and
// non-streaming. The Chat completion inference API provides more comprehensive
// customization options through more fields and function calling support. To
// determine whether a given inference service supports this task type, please
// see the page for that service.
//
// These services support non-streaming chat completion inference:
//
//   - AI21
//   - Azure OpenAI
//   - Deepseek
//   - Elastic
//   - FireworksAI
//   - Groq
//   - Huggingface
//   - IBMWatsonX
//   - Llama
//   - Mistral
//   - NVIDIA
//   - OpenAI
//   - OpenShiftAI
package nonstreamingchatcompletion

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

type NonStreamingChatCompletion struct {
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

// NewNonStreamingChatCompletion type alias for index.
type NewNonStreamingChatCompletion func(inferenceid string) *NonStreamingChatCompletion

// NewNonStreamingChatCompletionFunc returns a new instance of NonStreamingChatCompletion with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewNonStreamingChatCompletionFunc(tp elastictransport.Interface) NewNonStreamingChatCompletion {
	return func(inferenceid string) *NonStreamingChatCompletion {
		n := New(tp)

		n._inferenceid(inferenceid)

		return n
	}
}

// Perform non-streaming chat completion inference on the service.
//
// The chat completion inference API enables rich responses for chat completion
// tasks. It only works with the `chat_completion` task type.
//
// NOTE: The `chat_completion` task type supports both streaming and
// non-streaming. The Chat completion inference API provides more comprehensive
// customization options through more fields and function calling support. To
// determine whether a given inference service supports this task type, please
// see the page for that service.
//
// These services support non-streaming chat completion inference:
//
//   - AI21
//   - Azure OpenAI
//   - Deepseek
//   - Elastic
//   - FireworksAI
//   - Groq
//   - Huggingface
//   - IBMWatsonX
//   - Llama
//   - Mistral
//   - NVIDIA
//   - OpenAI
//   - OpenShiftAI
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-inference-non-streaming-chat-completion
func New(tp elastictransport.Interface) *NonStreamingChatCompletion {
	r := &NonStreamingChatCompletion{
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
func (r *NonStreamingChatCompletion) Raw(raw io.Reader) *NonStreamingChatCompletion {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *NonStreamingChatCompletion) Request(req *Request) *NonStreamingChatCompletion {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *NonStreamingChatCompletion) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for NonStreamingChatCompletion: %w", err)
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
		path.WriteString("chat_completion")
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
func (r NonStreamingChatCompletion) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx = instrument.Start(providedCtx, "inference.non_streaming_chat_completion")
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
		instrument.BeforeRequest(req, "inference.non_streaming_chat_completion")
		if reader := instrument.RecordRequestBody(ctx, "inference.non_streaming_chat_completion", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.non_streaming_chat_completion")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the NonStreamingChatCompletion query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a nonstreamingchatcompletion.Response
func (r NonStreamingChatCompletion) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.non_streaming_chat_completion")
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

// Header set a key, value pair in the NonStreamingChatCompletion headers map.
func (r *NonStreamingChatCompletion) Header(key, value string) *NonStreamingChatCompletion {
	r.headers.Set(key, value)

	return r
}

// InferenceId The inference Id
// API Name: inferenceid
func (r *NonStreamingChatCompletion) _inferenceid(inferenceid string) *NonStreamingChatCompletion {
	r.paramSet |= inferenceidMask
	r.inferenceid = inferenceid

	return r
}

// Timeout Specifies the amount of time to wait for the inference request to complete.
// API name: timeout
func (r *NonStreamingChatCompletion) Timeout(duration string) *NonStreamingChatCompletion {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *NonStreamingChatCompletion) ErrorTrace(errortrace bool) *NonStreamingChatCompletion {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *NonStreamingChatCompletion) FilterPath(filterpaths ...string) *NonStreamingChatCompletion {
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
func (r *NonStreamingChatCompletion) Human(human bool) *NonStreamingChatCompletion {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use this
// option for debugging only.
// API name: pretty
func (r *NonStreamingChatCompletion) Pretty(pretty bool) *NonStreamingChatCompletion {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// The prompt caching configuration for the completion request. Example:
//
//	{
//	  "cache_control": {
//	    "type": "ephemeral",
//	    "ttl": "1h"
//	  }
//	}
//
// Unknown fields in the object and invalid `ttl` values return an error.
// Currently supported only for `elastic` provider.
// API name: cache_control
func (r *NonStreamingChatCompletion) CacheControl(cachecontrol types.CacheControlVariant) *NonStreamingChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.CacheControl = cachecontrol.CacheControlCaster()

	return r
}

// The upper bound limit for the number of tokens that can be generated for a
// completion request.
// API name: max_completion_tokens
func (r *NonStreamingChatCompletion) MaxCompletionTokens(maxcompletiontokens int64) *NonStreamingChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.MaxCompletionTokens = &maxcompletiontokens

	return r
}

// A list of objects representing the conversation. Requests should generally
// only add new messages from the user (role `user`). The other message roles
// (`assistant`, `system`, or `tool`) should generally only be copied from the
// response to a previous completion request, such that the messages array is
// built up throughout a conversation.
// API name: messages
func (r *NonStreamingChatCompletion) Messages(messages ...types.MessageVariant) *NonStreamingChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range messages {

		r.req.Messages = append(r.req.Messages, *v.MessageCaster())

	}
	return r
}

func (r *NonStreamingChatCompletion) MessagesValues(messagesvalues []types.Message) *NonStreamingChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Messages = messagesvalues
	return r
}

// The ID of the model to use. By default, the model ID is set to the value
// included when creating the inference endpoint.
// API name: model
func (r *NonStreamingChatCompletion) Model(model string) *NonStreamingChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Model = &model

	return r
}

// The reasoning configuration for the completion request. This controls the
// model's reasoning process in one of two ways:
//
//   - By specifying the model’s reasoning effort level with the `effort`
//     field.
//   - By enabling reasoning with default settings by setting `enabled` field to
//     `true`.
//
// It also includes optional settings to control:
//
//   - The level of detail in the summary returned in the response with the
//     `summary` field.
//   - Whether reasoning details are included in the response at all with the
//     `exclude` field.
//
// Example (effort):
//
//	{
//	   "reasoning": {
//	       "effort": "high",
//	       "summary": "concise",
//	       "exclude": false
//	   }
//	}
//
// Example (enabled):
//
//	{
//	   "reasoning": {
//	       "enabled": true,
//	       "summary": "concise",
//	       "exclude": false
//	   }
//	}
//
// Currently supported only for `elastic` provider.
// API name: reasoning
func (r *NonStreamingChatCompletion) Reasoning(reasoning types.ReasoningVariant) *NonStreamingChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Reasoning = reasoning.ReasoningCaster()

	return r
}

// An identifier that groups related completion requests into a session.
// Currently supported only for `elastic` provider.
// API name: session_id
func (r *NonStreamingChatCompletion) SessionId(sessionid string) *NonStreamingChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.SessionId = &sessionid

	return r
}

// A sequence of strings to control when the model should stop generating
// additional tokens.
// API name: stop
func (r *NonStreamingChatCompletion) Stop(stops ...string) *NonStreamingChatCompletion {
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
func (r *NonStreamingChatCompletion) Temperature(temperature float32) *NonStreamingChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Temperature = &temperature

	return r
}

// Controls which tool is called by the model. String representation: One of
// `auto`, `none`, or `requrired`. `auto` allows the model to choose between
// calling tools and generating a message. `none` causes the model to not call
// any tools. `required` forces the model to call one or more tools. Example
// (object representation):
//
//	{
//	  "tool_choice": {
//	      "type": "function",
//	      "function": {
//	          "name": "get_current_weather"
//	      }
//	  }
//	}
//
// API name: tool_choice
func (r *NonStreamingChatCompletion) ToolChoice(completiontooltype types.CompletionToolTypeVariant) *NonStreamingChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ToolChoice = *completiontooltype.CompletionToolTypeCaster()

	return r
}

// A list of tools that the model can call. Example:
//
//	{
//	  "tools": [
//	      {
//	          "type": "function",
//	          "function": {
//	              "name": "get_price_of_item",
//	              "description": "Get the current price of an item",
//	              "parameters": {
//	                  "type": "object",
//	                  "properties": {
//	                      "item": {
//	                          "id": "12345"
//	                      },
//	                      "unit": {
//	                          "type": "currency"
//	                      }
//	                  }
//	              }
//	          }
//	      }
//	  ]
//	}
//
// API name: tools
func (r *NonStreamingChatCompletion) Tools(tools ...types.CompletionToolVariant) *NonStreamingChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range tools {

		r.req.Tools = append(r.req.Tools, *v.CompletionToolCaster())

	}
	return r
}

func (r *NonStreamingChatCompletion) ToolsValues(toolsvalues []types.CompletionTool) *NonStreamingChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Tools = toolsvalues
	return r
}

// Nucleus sampling, an alternative to sampling with temperature.
// API name: top_p
func (r *NonStreamingChatCompletion) TopP(topp float32) *NonStreamingChatCompletion {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TopP = &topp

	return r
}
