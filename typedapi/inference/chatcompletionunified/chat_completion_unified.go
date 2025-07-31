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

// Perform chat completion inference
//
// The chat completion inference API enables real-time responses for chat
// completion tasks by delivering answers incrementally, reducing response times
// during computation.
// It only works with the `chat_completion` task type for `openai` and `elastic`
// inference services.
//
// NOTE: The `chat_completion` task type is only available within the _stream
// API and only supports streaming.
// The Chat completion inference API and the Stream inference API differ in
// their response structure and capabilities.
// The Chat completion inference API provides more comprehensive customization
// options through more fields and function calling support.
// If you use the `openai`, `hugging_face` or the `elastic` service, use the
// Chat completion inference API.
package chatcompletionunified

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

type ChatCompletionUnified struct {
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

// NewChatCompletionUnified type alias for index.
type NewChatCompletionUnified func(inferenceid string) *ChatCompletionUnified

// NewChatCompletionUnifiedFunc returns a new instance of ChatCompletionUnified with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewChatCompletionUnifiedFunc(tp elastictransport.Interface) NewChatCompletionUnified {
	return func(inferenceid string) *ChatCompletionUnified {
		n := New(tp)

		n._inferenceid(inferenceid)

		return n
	}
}

// Perform chat completion inference
//
// The chat completion inference API enables real-time responses for chat
// completion tasks by delivering answers incrementally, reducing response times
// during computation.
// It only works with the `chat_completion` task type for `openai` and `elastic`
// inference services.
//
// NOTE: The `chat_completion` task type is only available within the _stream
// API and only supports streaming.
// The Chat completion inference API and the Stream inference API differ in
// their response structure and capabilities.
// The Chat completion inference API provides more comprehensive customization
// options through more fields and function calling support.
// If you use the `openai`, `hugging_face` or the `elastic` service, use the
// Chat completion inference API.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/chat-completion-inference-api.html
func New(tp elastictransport.Interface) *ChatCompletionUnified {
	r := &ChatCompletionUnified{
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
func (r *ChatCompletionUnified) Raw(raw io.Reader) *ChatCompletionUnified {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *ChatCompletionUnified) Request(req *Request) *ChatCompletionUnified {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ChatCompletionUnified) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for ChatCompletionUnified: %w", err)
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
func (r ChatCompletionUnified) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.chat_completion_unified")
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
		instrument.BeforeRequest(req, "inference.chat_completion_unified")
		if reader := instrument.RecordRequestBody(ctx, "inference.chat_completion_unified", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.chat_completion_unified")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ChatCompletionUnified query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a chatcompletionunified.Response
func (r ChatCompletionUnified) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.chat_completion_unified")
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

// Header set a key, value pair in the ChatCompletionUnified headers map.
func (r *ChatCompletionUnified) Header(key, value string) *ChatCompletionUnified {
	r.headers.Set(key, value)

	return r
}

// InferenceId The inference Id
// API Name: inferenceid
func (r *ChatCompletionUnified) _inferenceid(inferenceid string) *ChatCompletionUnified {
	r.paramSet |= inferenceidMask
	r.inferenceid = inferenceid

	return r
}

// Timeout Specifies the amount of time to wait for the inference request to complete.
// API name: timeout
func (r *ChatCompletionUnified) Timeout(duration string) *ChatCompletionUnified {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ChatCompletionUnified) ErrorTrace(errortrace bool) *ChatCompletionUnified {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ChatCompletionUnified) FilterPath(filterpaths ...string) *ChatCompletionUnified {
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
func (r *ChatCompletionUnified) Human(human bool) *ChatCompletionUnified {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ChatCompletionUnified) Pretty(pretty bool) *ChatCompletionUnified {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// MaxCompletionTokens The upper bound limit for the number of tokens that can be generated for a
// completion request.
// API name: max_completion_tokens
func (r *ChatCompletionUnified) MaxCompletionTokens(maxcompletiontokens int64) *ChatCompletionUnified {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.MaxCompletionTokens = &maxcompletiontokens

	return r
}

// Messages A list of objects representing the conversation.
// Requests should generally only add new messages from the user (role `user`).
// The other message roles (`assistant`, `system`, or `tool`) should generally
// only be copied from the response to a previous completion request, such that
// the messages array is built up throughout a conversation.
// API name: messages
func (r *ChatCompletionUnified) Messages(messages ...types.Message) *ChatCompletionUnified {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Messages = messages

	return r
}

// Model The ID of the model to use.
// API name: model
func (r *ChatCompletionUnified) Model(model string) *ChatCompletionUnified {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Model = &model

	return r
}

// Stop A sequence of strings to control when the model should stop generating
// additional tokens.
// API name: stop
func (r *ChatCompletionUnified) Stop(stops ...string) *ChatCompletionUnified {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Stop = stops

	return r
}

// Temperature The sampling temperature to use.
// API name: temperature
func (r *ChatCompletionUnified) Temperature(temperature float32) *ChatCompletionUnified {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Temperature = &temperature

	return r
}

// ToolChoice Controls which tool is called by the model.
// String representation: One of `auto`, `none`, or `requrired`. `auto` allows
// the model to choose between calling tools and generating a message. `none`
// causes the model to not call any tools. `required` forces the model to call
// one or more tools.
// Example (object representation):
// ```
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
// ```
// API name: tool_choice
func (r *ChatCompletionUnified) ToolChoice(completiontooltype types.CompletionToolType) *ChatCompletionUnified {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ToolChoice = completiontooltype

	return r
}

// Tools A list of tools that the model can call.
// Example:
// ```
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
// ```
// API name: tools
func (r *ChatCompletionUnified) Tools(tools ...types.CompletionTool) *ChatCompletionUnified {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Tools = tools

	return r
}

// TopP Nucleus sampling, an alternative to sampling with temperature.
// API name: top_p
func (r *ChatCompletionUnified) TopP(topp float32) *ChatCompletionUnified {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TopP = &topp

	return r
}
