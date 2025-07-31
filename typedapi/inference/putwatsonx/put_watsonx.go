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

// Create a Watsonx inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `watsonxai` service.
// You need an IBM Cloud Databases for Elasticsearch deployment to use the
// `watsonxai` inference service.
// You can provision one through the IBM catalog, the Cloud Databases CLI
// plug-in, the Cloud Databases API, or Terraform.
package putwatsonx

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/watsonxservicetype"
)

const (
	tasktypeMask = iota + 1

	watsonxinferenceidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutWatsonx struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype           string
	watsonxinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutWatsonx type alias for index.
type NewPutWatsonx func(tasktype, watsonxinferenceid string) *PutWatsonx

// NewPutWatsonxFunc returns a new instance of PutWatsonx with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutWatsonxFunc(tp elastictransport.Interface) NewPutWatsonx {
	return func(tasktype, watsonxinferenceid string) *PutWatsonx {
		n := New(tp)

		n._tasktype(tasktype)

		n._watsonxinferenceid(watsonxinferenceid)

		return n
	}
}

// Create a Watsonx inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `watsonxai` service.
// You need an IBM Cloud Databases for Elasticsearch deployment to use the
// `watsonxai` inference service.
// You can provision one through the IBM catalog, the Cloud Databases CLI
// plug-in, the Cloud Databases API, or Terraform.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-watsonx-ai.html
func New(tp elastictransport.Interface) *PutWatsonx {
	r := &PutWatsonx{
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
func (r *PutWatsonx) Raw(raw io.Reader) *PutWatsonx {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutWatsonx) Request(req *Request) *PutWatsonx {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutWatsonx) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutWatsonx: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == tasktypeMask|watsonxinferenceidMask:
		path.WriteString("/")
		path.WriteString("_inference")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "tasktype", r.tasktype)
		}
		path.WriteString(r.tasktype)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "watsonxinferenceid", r.watsonxinferenceid)
		}
		path.WriteString(r.watsonxinferenceid)

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
func (r PutWatsonx) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.put_watsonx")
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
		instrument.BeforeRequest(req, "inference.put_watsonx")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_watsonx", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_watsonx")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutWatsonx query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putwatsonx.Response
func (r PutWatsonx) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_watsonx")
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

// Header set a key, value pair in the PutWatsonx headers map.
func (r *PutWatsonx) Header(key, value string) *PutWatsonx {
	r.headers.Set(key, value)

	return r
}

// TaskType The task type.
// The only valid task type for the model to perform is `text_embedding`.
// API Name: tasktype
func (r *PutWatsonx) _tasktype(tasktype string) *PutWatsonx {
	r.paramSet |= tasktypeMask
	r.tasktype = tasktype

	return r
}

// WatsonxInferenceId The unique identifier of the inference endpoint.
// API Name: watsonxinferenceid
func (r *PutWatsonx) _watsonxinferenceid(watsonxinferenceid string) *PutWatsonx {
	r.paramSet |= watsonxinferenceidMask
	r.watsonxinferenceid = watsonxinferenceid

	return r
}

// Timeout Specifies the amount of time to wait for the inference endpoint to be
// created.
// API name: timeout
func (r *PutWatsonx) Timeout(duration string) *PutWatsonx {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutWatsonx) ErrorTrace(errortrace bool) *PutWatsonx {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutWatsonx) FilterPath(filterpaths ...string) *PutWatsonx {
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
func (r *PutWatsonx) Human(human bool) *PutWatsonx {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutWatsonx) Pretty(pretty bool) *PutWatsonx {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Service The type of service supported for the specified task type. In this case,
// `watsonxai`.
// API name: service
func (r *PutWatsonx) Service(service watsonxservicetype.WatsonxServiceType) *PutWatsonx {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Service = service

	return r
}

// ServiceSettings Settings used to install the inference model. These settings are specific to
// the `watsonxai` service.
// API name: service_settings
func (r *PutWatsonx) ServiceSettings(servicesettings *types.WatsonxServiceSettings) *PutWatsonx {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ServiceSettings = *servicesettings

	return r
}
