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

// Create an Elasticsearch inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `elasticsearch` service.
//
// > info
// > Your Elasticsearch deployment contains preconfigured ELSER and E5 inference
// endpoints, you only need to create the enpoints using the API if you want to
// customize the settings.
//
// If you use the ELSER or the E5 model through the `elasticsearch` service, the
// API request will automatically download and deploy the model if it isn't
// downloaded yet.
//
// > info
// > You might see a 502 bad gateway error in the response when using the Kibana
// Console. This error usually just reflects a timeout, while the model
// downloads in the background. You can check the download progress in the
// Machine Learning UI. If using the Python client, you can set the timeout
// parameter to a higher value.
//
// After creating the endpoint, wait for the model deployment to complete before
// using it.
// To verify the deployment status, use the get trained model statistics API.
// Look for `"state": "fully_allocated"` in the response and ensure that the
// `"allocation_count"` matches the `"target_allocation_count"`.
// Avoid creating multiple endpoints for the same model unless required, as each
// endpoint consumes significant resources.
package putelasticsearch

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/elasticsearchservicetype"
)

const (
	tasktypeMask = iota + 1

	elasticsearchinferenceidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutElasticsearch struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype                 string
	elasticsearchinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutElasticsearch type alias for index.
type NewPutElasticsearch func(tasktype, elasticsearchinferenceid string) *PutElasticsearch

// NewPutElasticsearchFunc returns a new instance of PutElasticsearch with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutElasticsearchFunc(tp elastictransport.Interface) NewPutElasticsearch {
	return func(tasktype, elasticsearchinferenceid string) *PutElasticsearch {
		n := New(tp)

		n._tasktype(tasktype)

		n._elasticsearchinferenceid(elasticsearchinferenceid)

		return n
	}
}

// Create an Elasticsearch inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `elasticsearch` service.
//
// > info
// > Your Elasticsearch deployment contains preconfigured ELSER and E5 inference
// endpoints, you only need to create the enpoints using the API if you want to
// customize the settings.
//
// If you use the ELSER or the E5 model through the `elasticsearch` service, the
// API request will automatically download and deploy the model if it isn't
// downloaded yet.
//
// > info
// > You might see a 502 bad gateway error in the response when using the Kibana
// Console. This error usually just reflects a timeout, while the model
// downloads in the background. You can check the download progress in the
// Machine Learning UI. If using the Python client, you can set the timeout
// parameter to a higher value.
//
// After creating the endpoint, wait for the model deployment to complete before
// using it.
// To verify the deployment status, use the get trained model statistics API.
// Look for `"state": "fully_allocated"` in the response and ensure that the
// `"allocation_count"` matches the `"target_allocation_count"`.
// Avoid creating multiple endpoints for the same model unless required, as each
// endpoint consumes significant resources.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-elasticsearch.html
func New(tp elastictransport.Interface) *PutElasticsearch {
	r := &PutElasticsearch{
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
func (r *PutElasticsearch) Raw(raw io.Reader) *PutElasticsearch {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutElasticsearch) Request(req *Request) *PutElasticsearch {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutElasticsearch) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutElasticsearch: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == tasktypeMask|elasticsearchinferenceidMask:
		path.WriteString("/")
		path.WriteString("_inference")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "tasktype", r.tasktype)
		}
		path.WriteString(r.tasktype)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "elasticsearchinferenceid", r.elasticsearchinferenceid)
		}
		path.WriteString(r.elasticsearchinferenceid)

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
func (r PutElasticsearch) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.put_elasticsearch")
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
		instrument.BeforeRequest(req, "inference.put_elasticsearch")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_elasticsearch", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_elasticsearch")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutElasticsearch query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putelasticsearch.Response
func (r PutElasticsearch) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_elasticsearch")
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

// Header set a key, value pair in the PutElasticsearch headers map.
func (r *PutElasticsearch) Header(key, value string) *PutElasticsearch {
	r.headers.Set(key, value)

	return r
}

// TaskType The type of the inference task that the model will perform.
// API Name: tasktype
func (r *PutElasticsearch) _tasktype(tasktype string) *PutElasticsearch {
	r.paramSet |= tasktypeMask
	r.tasktype = tasktype

	return r
}

// ElasticsearchInferenceId The unique identifier of the inference endpoint.
// The must not match the `model_id`.
// API Name: elasticsearchinferenceid
func (r *PutElasticsearch) _elasticsearchinferenceid(elasticsearchinferenceid string) *PutElasticsearch {
	r.paramSet |= elasticsearchinferenceidMask
	r.elasticsearchinferenceid = elasticsearchinferenceid

	return r
}

// Timeout Specifies the amount of time to wait for the inference endpoint to be
// created.
// API name: timeout
func (r *PutElasticsearch) Timeout(duration string) *PutElasticsearch {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutElasticsearch) ErrorTrace(errortrace bool) *PutElasticsearch {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutElasticsearch) FilterPath(filterpaths ...string) *PutElasticsearch {
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
func (r *PutElasticsearch) Human(human bool) *PutElasticsearch {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutElasticsearch) Pretty(pretty bool) *PutElasticsearch {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// ChunkingSettings The chunking configuration object.
// API name: chunking_settings
func (r *PutElasticsearch) ChunkingSettings(chunkingsettings *types.InferenceChunkingSettings) *PutElasticsearch {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ChunkingSettings = chunkingsettings

	return r
}

// Service The type of service supported for the specified task type. In this case,
// `elasticsearch`.
// API name: service
func (r *PutElasticsearch) Service(service elasticsearchservicetype.ElasticsearchServiceType) *PutElasticsearch {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Service = service

	return r
}

// ServiceSettings Settings used to install the inference model. These settings are specific to
// the `elasticsearch` service.
// API name: service_settings
func (r *PutElasticsearch) ServiceSettings(servicesettings *types.ElasticsearchServiceSettings) *PutElasticsearch {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ServiceSettings = *servicesettings

	return r
}

// TaskSettings Settings to configure the inference task.
// These settings are specific to the task type you specified.
// API name: task_settings
func (r *PutElasticsearch) TaskSettings(tasksettings *types.ElasticsearchTaskSettings) *PutElasticsearch {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TaskSettings = tasksettings

	return r
}
