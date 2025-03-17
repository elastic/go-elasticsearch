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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

// Update an inference endpoint.
//
// Modify `task_settings`, secrets (within `service_settings`), or
// `num_allocations` for an inference endpoint, depending on the specific
// endpoint service and `task_type`.
//
// IMPORTANT: The inference APIs enable you to use certain services, such as
// built-in machine learning models (ELSER, E5), models uploaded through Eland,
// Cohere, OpenAI, Azure, Google AI Studio, Google Vertex AI, Anthropic,
// Watsonx.ai, or Hugging Face.
// For built-in models and models uploaded through Eland, the inference APIs
// offer an alternative way to use and manage trained models.
// However, if you do not plan to use the inference APIs to use these models or
// if you want to use non-NLP models, use the machine learning trained model
// APIs.
package update

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

	tasktypeMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Update struct {
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
	tasktype    string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewUpdate type alias for index.
type NewUpdate func(inferenceid string) *Update

// NewUpdateFunc returns a new instance of Update with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateFunc(tp elastictransport.Interface) NewUpdate {
	return func(inferenceid string) *Update {
		n := New(tp)

		n._inferenceid(inferenceid)

		return n
	}
}

// Update an inference endpoint.
//
// Modify `task_settings`, secrets (within `service_settings`), or
// `num_allocations` for an inference endpoint, depending on the specific
// endpoint service and `task_type`.
//
// IMPORTANT: The inference APIs enable you to use certain services, such as
// built-in machine learning models (ELSER, E5), models uploaded through Eland,
// Cohere, OpenAI, Azure, Google AI Studio, Google Vertex AI, Anthropic,
// Watsonx.ai, or Hugging Face.
// For built-in models and models uploaded through Eland, the inference APIs
// offer an alternative way to use and manage trained models.
// However, if you do not plan to use the inference APIs to use these models or
// if you want to use non-NLP models, use the machine learning trained model
// APIs.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-inference-update
func New(tp elastictransport.Interface) *Update {
	r := &Update{
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
func (r *Update) Raw(raw io.Reader) *Update {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Update) Request(req *Request) *Update {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Update) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Update: %w", err)
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

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "inferenceid", r.inferenceid)
		}
		path.WriteString(r.inferenceid)
		path.WriteString("/")
		path.WriteString("_update")

		method = http.MethodPut
	case r.paramSet == tasktypeMask|inferenceidMask:
		path.WriteString("/")
		path.WriteString("_inference")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "tasktype", r.tasktype)
		}
		path.WriteString(r.tasktype)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "inferenceid", r.inferenceid)
		}
		path.WriteString(r.inferenceid)
		path.WriteString("/")
		path.WriteString("_update")

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
func (r Update) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.update")
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
		instrument.BeforeRequest(req, "inference.update")
		if reader := instrument.RecordRequestBody(ctx, "inference.update", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.update")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Update query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a update.Response
func (r Update) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.update")
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

// Header set a key, value pair in the Update headers map.
func (r *Update) Header(key, value string) *Update {
	r.headers.Set(key, value)

	return r
}

// InferenceId The unique identifier of the inference endpoint.
// API Name: inferenceid
func (r *Update) _inferenceid(inferenceid string) *Update {
	r.paramSet |= inferenceidMask
	r.inferenceid = inferenceid

	return r
}

// TaskType The type of inference task that the model performs.
// API Name: tasktype
func (r *Update) TaskType(tasktype string) *Update {
	r.paramSet |= tasktypeMask
	r.tasktype = tasktype

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Update) ErrorTrace(errortrace bool) *Update {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Update) FilterPath(filterpaths ...string) *Update {
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
func (r *Update) Human(human bool) *Update {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Update) Pretty(pretty bool) *Update {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Chunking configuration object
// API name: chunking_settings
func (r *Update) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *Update {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ChunkingSettings = chunkingsettings.InferenceChunkingSettingsCaster()

	return r
}

// The service type
// API name: service
func (r *Update) Service(service string) *Update {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Service = service

	return r
}

// Settings specific to the service
// API name: service_settings
func (r *Update) ServiceSettings(servicesettings json.RawMessage) *Update {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ServiceSettings = servicesettings

	return r
}

// Task settings specific to the service and task type
// API name: task_settings
func (r *Update) TaskSettings(tasksettings json.RawMessage) *Update {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TaskSettings = tasksettings

	return r
}
