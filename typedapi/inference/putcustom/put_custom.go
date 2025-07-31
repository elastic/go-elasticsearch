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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Create a custom inference endpoint.
//
// The custom service gives more control over how to interact with external
// inference services that aren't explicitly supported through dedicated
// integrations.
// The custom service gives you the ability to define the headers, url, query
// parameters, request body, and secrets.
// The custom service supports the template replacement functionality, which
// enables you to define a template that can be replaced with the value
// associated with that key.
// Templates are portions of a string that start with `${` and end with `}`.
// The parameters `secret_parameters` and `task_settings` are checked for keys
// for template replacement. Template replacement is supported in the `request`,
// `headers`, `url`, and `query_parameters`.
// If the definition (key) is not found for a template, an error message is
// returned.
// In case of an endpoint definition like the following:
// ```
// PUT _inference/text_embedding/test-text-embedding
//
//	{
//	  "service": "custom",
//	  "service_settings": {
//	     "secret_parameters": {
//	          "api_key": "<some api key>"
//	     },
//	     "url": "...endpoints.huggingface.cloud/v1/embeddings",
//	     "headers": {
//	         "Authorization": "Bearer ${api_key}",
//	         "Content-Type": "application/json"
//	     },
//	     "request": "{\"input\": ${input}}",
//	     "response": {
//	         "json_parser": {
//	             "text_embeddings":"$.data[*].embedding[*]"
//	         }
//	     }
//	  }
//	}
//
// ```
// To replace `${api_key}` the `secret_parameters` and `task_settings` are
// checked for a key named `api_key`.
//
// > info
// > Templates should not be surrounded by quotes.
//
// Pre-defined templates:
// * `${input}` refers to the array of input strings that comes from the `input`
// field of the subsequent inference requests.
// * `${input_type}` refers to the input type translation values.
// * `${query}` refers to the query field used specifically for reranking tasks.
// * `${top_n}` refers to the `top_n` field available when performing rerank
// requests.
// * `${return_documents}` refers to the `return_documents` field available when
// performing rerank requests.
package putcustom

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/customservicetype"
)

const (
	tasktypeMask = iota + 1

	custominferenceidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutCustom struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype          string
	custominferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutCustom type alias for index.
type NewPutCustom func(tasktype, custominferenceid string) *PutCustom

// NewPutCustomFunc returns a new instance of PutCustom with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutCustomFunc(tp elastictransport.Interface) NewPutCustom {
	return func(tasktype, custominferenceid string) *PutCustom {
		n := New(tp)

		n._tasktype(tasktype)

		n._custominferenceid(custominferenceid)

		return n
	}
}

// Create a custom inference endpoint.
//
// The custom service gives more control over how to interact with external
// inference services that aren't explicitly supported through dedicated
// integrations.
// The custom service gives you the ability to define the headers, url, query
// parameters, request body, and secrets.
// The custom service supports the template replacement functionality, which
// enables you to define a template that can be replaced with the value
// associated with that key.
// Templates are portions of a string that start with `${` and end with `}`.
// The parameters `secret_parameters` and `task_settings` are checked for keys
// for template replacement. Template replacement is supported in the `request`,
// `headers`, `url`, and `query_parameters`.
// If the definition (key) is not found for a template, an error message is
// returned.
// In case of an endpoint definition like the following:
// ```
// PUT _inference/text_embedding/test-text-embedding
//
//	{
//	  "service": "custom",
//	  "service_settings": {
//	     "secret_parameters": {
//	          "api_key": "<some api key>"
//	     },
//	     "url": "...endpoints.huggingface.cloud/v1/embeddings",
//	     "headers": {
//	         "Authorization": "Bearer ${api_key}",
//	         "Content-Type": "application/json"
//	     },
//	     "request": "{\"input\": ${input}}",
//	     "response": {
//	         "json_parser": {
//	             "text_embeddings":"$.data[*].embedding[*]"
//	         }
//	     }
//	  }
//	}
//
// ```
// To replace `${api_key}` the `secret_parameters` and `task_settings` are
// checked for a key named `api_key`.
//
// > info
// > Templates should not be surrounded by quotes.
//
// Pre-defined templates:
// * `${input}` refers to the array of input strings that comes from the `input`
// field of the subsequent inference requests.
// * `${input_type}` refers to the input type translation values.
// * `${query}` refers to the query field used specifically for reranking tasks.
// * `${top_n}` refers to the `top_n` field available when performing rerank
// requests.
// * `${return_documents}` refers to the `return_documents` field available when
// performing rerank requests.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-inference-put-custom
func New(tp elastictransport.Interface) *PutCustom {
	r := &PutCustom{
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
func (r *PutCustom) Raw(raw io.Reader) *PutCustom {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutCustom) Request(req *Request) *PutCustom {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutCustom) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutCustom: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == tasktypeMask|custominferenceidMask:
		path.WriteString("/")
		path.WriteString("_inference")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "tasktype", r.tasktype)
		}
		path.WriteString(r.tasktype)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "custominferenceid", r.custominferenceid)
		}
		path.WriteString(r.custominferenceid)

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
func (r PutCustom) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.put_custom")
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
		instrument.BeforeRequest(req, "inference.put_custom")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_custom", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_custom")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutCustom query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putcustom.Response
func (r PutCustom) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_custom")
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

// Header set a key, value pair in the PutCustom headers map.
func (r *PutCustom) Header(key, value string) *PutCustom {
	r.headers.Set(key, value)

	return r
}

// TaskType The type of the inference task that the model will perform.
// API Name: tasktype
func (r *PutCustom) _tasktype(tasktype string) *PutCustom {
	r.paramSet |= tasktypeMask
	r.tasktype = tasktype

	return r
}

// CustomInferenceId The unique identifier of the inference endpoint.
// API Name: custominferenceid
func (r *PutCustom) _custominferenceid(custominferenceid string) *PutCustom {
	r.paramSet |= custominferenceidMask
	r.custominferenceid = custominferenceid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutCustom) ErrorTrace(errortrace bool) *PutCustom {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutCustom) FilterPath(filterpaths ...string) *PutCustom {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"exists_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *PutCustom) Human(human bool) *PutCustom {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutCustom) Pretty(pretty bool) *PutCustom {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// The chunking configuration object.
// API name: chunking_settings
func (r *PutCustom) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutCustom {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ChunkingSettings = chunkingsettings.InferenceChunkingSettingsCaster()

	return r
}

// The type of service supported for the specified task type. In this case,
// `custom`.
// API name: service
func (r *PutCustom) Service(service customservicetype.CustomServiceType) *PutCustom {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Service = service
	return r
}

// Settings used to install the inference model.
// These settings are specific to the `custom` service.
// API name: service_settings
func (r *PutCustom) ServiceSettings(servicesettings types.CustomServiceSettingsVariant) *PutCustom {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ServiceSettings = *servicesettings.CustomServiceSettingsCaster()

	return r
}

// Settings to configure the inference task.
// These settings are specific to the task type you specified.
// API name: task_settings
func (r *PutCustom) TaskSettings(tasksettings types.CustomTaskSettingsVariant) *PutCustom {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TaskSettings = tasksettings.CustomTaskSettingsCaster()

	return r
}
