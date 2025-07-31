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

// Create an AlibabaCloud AI Search inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `alibabacloud-ai-search` service.
package putalibabacloud

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/alibabacloudservicetype"
)

const (
	tasktypeMask = iota + 1

	alibabacloudinferenceidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAlibabacloud struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype                string
	alibabacloudinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutAlibabacloud type alias for index.
type NewPutAlibabacloud func(tasktype, alibabacloudinferenceid string) *PutAlibabacloud

// NewPutAlibabacloudFunc returns a new instance of PutAlibabacloud with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutAlibabacloudFunc(tp elastictransport.Interface) NewPutAlibabacloud {
	return func(tasktype, alibabacloudinferenceid string) *PutAlibabacloud {
		n := New(tp)

		n._tasktype(tasktype)

		n._alibabacloudinferenceid(alibabacloudinferenceid)

		return n
	}
}

// Create an AlibabaCloud AI Search inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `alibabacloud-ai-search` service.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-alibabacloud-ai-search.html
func New(tp elastictransport.Interface) *PutAlibabacloud {
	r := &PutAlibabacloud{
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
func (r *PutAlibabacloud) Raw(raw io.Reader) *PutAlibabacloud {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutAlibabacloud) Request(req *Request) *PutAlibabacloud {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutAlibabacloud) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutAlibabacloud: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == tasktypeMask|alibabacloudinferenceidMask:
		path.WriteString("/")
		path.WriteString("_inference")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "tasktype", r.tasktype)
		}
		path.WriteString(r.tasktype)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "alibabacloudinferenceid", r.alibabacloudinferenceid)
		}
		path.WriteString(r.alibabacloudinferenceid)

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
func (r PutAlibabacloud) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.put_alibabacloud")
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
		instrument.BeforeRequest(req, "inference.put_alibabacloud")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_alibabacloud", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_alibabacloud")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutAlibabacloud query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putalibabacloud.Response
func (r PutAlibabacloud) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_alibabacloud")
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

// Header set a key, value pair in the PutAlibabacloud headers map.
func (r *PutAlibabacloud) Header(key, value string) *PutAlibabacloud {
	r.headers.Set(key, value)

	return r
}

// TaskType The type of the inference task that the model will perform.
// API Name: tasktype
func (r *PutAlibabacloud) _tasktype(tasktype string) *PutAlibabacloud {
	r.paramSet |= tasktypeMask
	r.tasktype = tasktype

	return r
}

// AlibabacloudInferenceId The unique identifier of the inference endpoint.
// API Name: alibabacloudinferenceid
func (r *PutAlibabacloud) _alibabacloudinferenceid(alibabacloudinferenceid string) *PutAlibabacloud {
	r.paramSet |= alibabacloudinferenceidMask
	r.alibabacloudinferenceid = alibabacloudinferenceid

	return r
}

// Timeout Specifies the amount of time to wait for the inference endpoint to be
// created.
// API name: timeout
func (r *PutAlibabacloud) Timeout(duration string) *PutAlibabacloud {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutAlibabacloud) ErrorTrace(errortrace bool) *PutAlibabacloud {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutAlibabacloud) FilterPath(filterpaths ...string) *PutAlibabacloud {
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
func (r *PutAlibabacloud) Human(human bool) *PutAlibabacloud {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutAlibabacloud) Pretty(pretty bool) *PutAlibabacloud {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// ChunkingSettings The chunking configuration object.
// API name: chunking_settings
func (r *PutAlibabacloud) ChunkingSettings(chunkingsettings *types.InferenceChunkingSettings) *PutAlibabacloud {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ChunkingSettings = chunkingsettings

	return r
}

// Service The type of service supported for the specified task type. In this case,
// `alibabacloud-ai-search`.
// API name: service
func (r *PutAlibabacloud) Service(service alibabacloudservicetype.AlibabaCloudServiceType) *PutAlibabacloud {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Service = service

	return r
}

// ServiceSettings Settings used to install the inference model. These settings are specific to
// the `alibabacloud-ai-search` service.
// API name: service_settings
func (r *PutAlibabacloud) ServiceSettings(servicesettings *types.AlibabaCloudServiceSettings) *PutAlibabacloud {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ServiceSettings = *servicesettings

	return r
}

// TaskSettings Settings to configure the inference task.
// These settings are specific to the task type you specified.
// API name: task_settings
func (r *PutAlibabacloud) TaskSettings(tasksettings *types.AlibabaCloudTaskSettings) *PutAlibabacloud {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TaskSettings = tasksettings

	return r
}
