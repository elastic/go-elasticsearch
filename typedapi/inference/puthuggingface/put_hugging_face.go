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

// Create a Hugging Face inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `hugging_face` service.
// Supported tasks include: `text_embedding`, `completion`, and
// `chat_completion`.
//
// To configure the endpoint, first visit the Hugging Face Inference Endpoints
// page and create a new endpoint.
// Select a model that supports the task you intend to use.
//
// For Elastic's `text_embedding` task:
// The selected model must support the `Sentence Embeddings` task. On the new
// endpoint creation page, select the `Sentence Embeddings` task under the
// `Advanced Configuration` section.
// After the endpoint has initialized, copy the generated endpoint URL.
// Recommended models for `text_embedding` task:
//
// * `all-MiniLM-L6-v2`
// * `all-MiniLM-L12-v2`
// * `all-mpnet-base-v2`
// * `e5-base-v2`
// * `e5-small-v2`
// * `multilingual-e5-base`
// * `multilingual-e5-small`
//
// For Elastic's `chat_completion` and `completion` tasks:
// The selected model must support the `Text Generation` task and expose OpenAI
// API. HuggingFace supports both serverless and dedicated endpoints for `Text
// Generation`. When creating dedicated endpoint select the `Text Generation`
// task.
// After the endpoint is initialized (for dedicated) or ready (for serverless),
// ensure it supports the OpenAI API and includes `/v1/chat/completions` part in
// URL. Then, copy the full endpoint URL for use.
// Recommended models for `chat_completion` and `completion` tasks:
//
// * `Mistral-7B-Instruct-v0.2`
// * `QwQ-32B`
// * `Phi-3-mini-128k-instruct`
//
// For Elastic's `rerank` task:
// The selected model must support the `sentence-ranking` task and expose OpenAI
// API.
// HuggingFace supports only dedicated (not serverless) endpoints for `Rerank`
// so far.
// After the endpoint is initialized, copy the full endpoint URL for use.
// Tested models for `rerank` task:
//
// * `bge-reranker-base`
// * `jina-reranker-v1-turbo-en-GGUF`
package puthuggingface

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/huggingfaceservicetype"
)

const (
	tasktypeMask = iota + 1

	huggingfaceinferenceidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutHuggingFace struct {
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
	huggingfaceinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutHuggingFace type alias for index.
type NewPutHuggingFace func(tasktype, huggingfaceinferenceid string) *PutHuggingFace

// NewPutHuggingFaceFunc returns a new instance of PutHuggingFace with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutHuggingFaceFunc(tp elastictransport.Interface) NewPutHuggingFace {
	return func(tasktype, huggingfaceinferenceid string) *PutHuggingFace {
		n := New(tp)

		n._tasktype(tasktype)

		n._huggingfaceinferenceid(huggingfaceinferenceid)

		return n
	}
}

// Create a Hugging Face inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `hugging_face` service.
// Supported tasks include: `text_embedding`, `completion`, and
// `chat_completion`.
//
// To configure the endpoint, first visit the Hugging Face Inference Endpoints
// page and create a new endpoint.
// Select a model that supports the task you intend to use.
//
// For Elastic's `text_embedding` task:
// The selected model must support the `Sentence Embeddings` task. On the new
// endpoint creation page, select the `Sentence Embeddings` task under the
// `Advanced Configuration` section.
// After the endpoint has initialized, copy the generated endpoint URL.
// Recommended models for `text_embedding` task:
//
// * `all-MiniLM-L6-v2`
// * `all-MiniLM-L12-v2`
// * `all-mpnet-base-v2`
// * `e5-base-v2`
// * `e5-small-v2`
// * `multilingual-e5-base`
// * `multilingual-e5-small`
//
// For Elastic's `chat_completion` and `completion` tasks:
// The selected model must support the `Text Generation` task and expose OpenAI
// API. HuggingFace supports both serverless and dedicated endpoints for `Text
// Generation`. When creating dedicated endpoint select the `Text Generation`
// task.
// After the endpoint is initialized (for dedicated) or ready (for serverless),
// ensure it supports the OpenAI API and includes `/v1/chat/completions` part in
// URL. Then, copy the full endpoint URL for use.
// Recommended models for `chat_completion` and `completion` tasks:
//
// * `Mistral-7B-Instruct-v0.2`
// * `QwQ-32B`
// * `Phi-3-mini-128k-instruct`
//
// For Elastic's `rerank` task:
// The selected model must support the `sentence-ranking` task and expose OpenAI
// API.
// HuggingFace supports only dedicated (not serverless) endpoints for `Rerank`
// so far.
// After the endpoint is initialized, copy the full endpoint URL for use.
// Tested models for `rerank` task:
//
// * `bge-reranker-base`
// * `jina-reranker-v1-turbo-en-GGUF`
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-hugging-face.html
func New(tp elastictransport.Interface) *PutHuggingFace {
	r := &PutHuggingFace{
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
func (r *PutHuggingFace) Raw(raw io.Reader) *PutHuggingFace {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutHuggingFace) Request(req *Request) *PutHuggingFace {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutHuggingFace) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutHuggingFace: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == tasktypeMask|huggingfaceinferenceidMask:
		path.WriteString("/")
		path.WriteString("_inference")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "tasktype", r.tasktype)
		}
		path.WriteString(r.tasktype)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "huggingfaceinferenceid", r.huggingfaceinferenceid)
		}
		path.WriteString(r.huggingfaceinferenceid)

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
func (r PutHuggingFace) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.put_hugging_face")
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
		instrument.BeforeRequest(req, "inference.put_hugging_face")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_hugging_face", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_hugging_face")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutHuggingFace query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a puthuggingface.Response
func (r PutHuggingFace) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_hugging_face")
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

// Header set a key, value pair in the PutHuggingFace headers map.
func (r *PutHuggingFace) Header(key, value string) *PutHuggingFace {
	r.headers.Set(key, value)

	return r
}

// TaskType The type of the inference task that the model will perform.
// API Name: tasktype
func (r *PutHuggingFace) _tasktype(tasktype string) *PutHuggingFace {
	r.paramSet |= tasktypeMask
	r.tasktype = tasktype

	return r
}

// HuggingfaceInferenceId The unique identifier of the inference endpoint.
// API Name: huggingfaceinferenceid
func (r *PutHuggingFace) _huggingfaceinferenceid(huggingfaceinferenceid string) *PutHuggingFace {
	r.paramSet |= huggingfaceinferenceidMask
	r.huggingfaceinferenceid = huggingfaceinferenceid

	return r
}

// Timeout Specifies the amount of time to wait for the inference endpoint to be
// created.
// API name: timeout
func (r *PutHuggingFace) Timeout(duration string) *PutHuggingFace {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutHuggingFace) ErrorTrace(errortrace bool) *PutHuggingFace {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutHuggingFace) FilterPath(filterpaths ...string) *PutHuggingFace {
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
func (r *PutHuggingFace) Human(human bool) *PutHuggingFace {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutHuggingFace) Pretty(pretty bool) *PutHuggingFace {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// ChunkingSettings The chunking configuration object.
// API name: chunking_settings
func (r *PutHuggingFace) ChunkingSettings(chunkingsettings *types.InferenceChunkingSettings) *PutHuggingFace {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ChunkingSettings = chunkingsettings

	return r
}

// Service The type of service supported for the specified task type. In this case,
// `hugging_face`.
// API name: service
func (r *PutHuggingFace) Service(service huggingfaceservicetype.HuggingFaceServiceType) *PutHuggingFace {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Service = service

	return r
}

// ServiceSettings Settings used to install the inference model. These settings are specific to
// the `hugging_face` service.
// API name: service_settings
func (r *PutHuggingFace) ServiceSettings(servicesettings *types.HuggingFaceServiceSettings) *PutHuggingFace {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ServiceSettings = *servicesettings

	return r
}

// TaskSettings Settings to configure the inference task.
// These settings are specific to the task type you specified.
// API name: task_settings
func (r *PutHuggingFace) TaskSettings(tasksettings *types.HuggingFaceTaskSettings) *PutHuggingFace {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TaskSettings = tasksettings

	return r
}
