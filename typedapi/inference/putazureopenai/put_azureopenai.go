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

// Create an Azure OpenAI inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `azureopenai` service.
//
// The list of chat completion models that you can choose from in your Azure
// OpenAI deployment include:
//
// * [GPT-4 and GPT-4 Turbo
// models](https://learn.microsoft.com/en-us/azure/ai-services/openai/concepts/models?tabs=global-standard%2Cstandard-chat-completions#gpt-4-and-gpt-4-turbo-models)
// *
// [GPT-3.5](https://learn.microsoft.com/en-us/azure/ai-services/openai/concepts/models?tabs=global-standard%2Cstandard-chat-completions#gpt-35)
//
// The list of embeddings models that you can choose from in your deployment can
// be found in the [Azure models
// documentation](https://learn.microsoft.com/en-us/azure/ai-services/openai/concepts/models?tabs=global-standard%2Cstandard-chat-completions#embeddings).
package putazureopenai

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/azureopenaiservicetype"
)

const (
	tasktypeMask = iota + 1

	azureopenaiinferenceidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAzureopenai struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype               string
	azureopenaiinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutAzureopenai type alias for index.
type NewPutAzureopenai func(tasktype, azureopenaiinferenceid string) *PutAzureopenai

// NewPutAzureopenaiFunc returns a new instance of PutAzureopenai with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutAzureopenaiFunc(tp elastictransport.Interface) NewPutAzureopenai {
	return func(tasktype, azureopenaiinferenceid string) *PutAzureopenai {
		n := New(tp)

		n._tasktype(tasktype)

		n._azureopenaiinferenceid(azureopenaiinferenceid)

		return n
	}
}

// Create an Azure OpenAI inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `azureopenai` service.
//
// The list of chat completion models that you can choose from in your Azure
// OpenAI deployment include:
//
// * [GPT-4 and GPT-4 Turbo
// models](https://learn.microsoft.com/en-us/azure/ai-services/openai/concepts/models?tabs=global-standard%2Cstandard-chat-completions#gpt-4-and-gpt-4-turbo-models)
// *
// [GPT-3.5](https://learn.microsoft.com/en-us/azure/ai-services/openai/concepts/models?tabs=global-standard%2Cstandard-chat-completions#gpt-35)
//
// The list of embeddings models that you can choose from in your deployment can
// be found in the [Azure models
// documentation](https://learn.microsoft.com/en-us/azure/ai-services/openai/concepts/models?tabs=global-standard%2Cstandard-chat-completions#embeddings).
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-azure-openai.html
func New(tp elastictransport.Interface) *PutAzureopenai {
	r := &PutAzureopenai{
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
func (r *PutAzureopenai) Raw(raw io.Reader) *PutAzureopenai {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutAzureopenai) Request(req *Request) *PutAzureopenai {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutAzureopenai) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutAzureopenai: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == tasktypeMask|azureopenaiinferenceidMask:
		path.WriteString("/")
		path.WriteString("_inference")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "tasktype", r.tasktype)
		}
		path.WriteString(r.tasktype)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "azureopenaiinferenceid", r.azureopenaiinferenceid)
		}
		path.WriteString(r.azureopenaiinferenceid)

		method = http.MethodPut
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
func (r PutAzureopenai) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.put_azureopenai")
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
		instrument.BeforeRequest(req, "inference.put_azureopenai")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_azureopenai", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_azureopenai")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutAzureopenai query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putazureopenai.Response
func (r PutAzureopenai) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_azureopenai")
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

// Header set a key, value pair in the PutAzureopenai headers map.
func (r *PutAzureopenai) Header(key, value string) *PutAzureopenai {
	r.headers.Set(key, value)

	return r
}

// TaskType The type of the inference task that the model will perform.
// NOTE: The `chat_completion` task type only supports streaming and only
// through the _stream API.
// API Name: tasktype
func (r *PutAzureopenai) _tasktype(tasktype string) *PutAzureopenai {
	r.paramSet |= tasktypeMask
	r.tasktype = tasktype

	return r
}

// AzureopenaiInferenceId The unique identifier of the inference endpoint.
// API Name: azureopenaiinferenceid
func (r *PutAzureopenai) _azureopenaiinferenceid(azureopenaiinferenceid string) *PutAzureopenai {
	r.paramSet |= azureopenaiinferenceidMask
	r.azureopenaiinferenceid = azureopenaiinferenceid

	return r
}

// Timeout Specifies the amount of time to wait for the inference endpoint to be
// created.
// API name: timeout
func (r *PutAzureopenai) Timeout(duration string) *PutAzureopenai {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutAzureopenai) ErrorTrace(errortrace bool) *PutAzureopenai {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutAzureopenai) FilterPath(filterpaths ...string) *PutAzureopenai {
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
func (r *PutAzureopenai) Human(human bool) *PutAzureopenai {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutAzureopenai) Pretty(pretty bool) *PutAzureopenai {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// ChunkingSettings The chunking configuration object.
// API name: chunking_settings
func (r *PutAzureopenai) ChunkingSettings(chunkingsettings *types.InferenceChunkingSettings) *PutAzureopenai {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ChunkingSettings = chunkingsettings

	return r
}

// Service The type of service supported for the specified task type. In this case,
// `azureopenai`.
// API name: service
func (r *PutAzureopenai) Service(service azureopenaiservicetype.AzureOpenAIServiceType) *PutAzureopenai {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Service = service

	return r
}

// ServiceSettings Settings used to install the inference model. These settings are specific to
// the `azureopenai` service.
// API name: service_settings
func (r *PutAzureopenai) ServiceSettings(servicesettings *types.AzureOpenAIServiceSettings) *PutAzureopenai {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ServiceSettings = *servicesettings

	return r
}

// TaskSettings Settings to configure the inference task.
// These settings are specific to the task type you specified.
// API name: task_settings
func (r *PutAzureopenai) TaskSettings(tasksettings *types.AzureOpenAITaskSettings) *PutAzureopenai {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TaskSettings = tasksettings

	return r
}
