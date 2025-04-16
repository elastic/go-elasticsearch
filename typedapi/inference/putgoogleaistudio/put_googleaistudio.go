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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

// Create an Google AI Studio inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `googleaistudio` service.
//
// When you create an inference endpoint, the associated machine learning model
// is automatically deployed if it is not already running.
// After creating the endpoint, wait for the model deployment to complete before
// using it.
// To verify the deployment status, use the get trained model statistics API.
// Look for `"state": "fully_allocated"` in the response and ensure that the
// `"allocation_count"` matches the `"target_allocation_count"`.
// Avoid creating multiple endpoints for the same model unless required, as each
// endpoint consumes significant resources.
package putgoogleaistudio

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/googleaiservicetype"
)

const (
	tasktypeMask = iota + 1

	googleaistudioinferenceidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutGoogleaistudio struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype                  string
	googleaistudioinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutGoogleaistudio type alias for index.
type NewPutGoogleaistudio func(tasktype, googleaistudioinferenceid string) *PutGoogleaistudio

// NewPutGoogleaistudioFunc returns a new instance of PutGoogleaistudio with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutGoogleaistudioFunc(tp elastictransport.Interface) NewPutGoogleaistudio {
	return func(tasktype, googleaistudioinferenceid string) *PutGoogleaistudio {
		n := New(tp)

		n._tasktype(tasktype)

		n._googleaistudioinferenceid(googleaistudioinferenceid)

		return n
	}
}

// Create an Google AI Studio inference endpoint.
//
// Create an inference endpoint to perform an inference task with the
// `googleaistudio` service.
//
// When you create an inference endpoint, the associated machine learning model
// is automatically deployed if it is not already running.
// After creating the endpoint, wait for the model deployment to complete before
// using it.
// To verify the deployment status, use the get trained model statistics API.
// Look for `"state": "fully_allocated"` in the response and ensure that the
// `"allocation_count"` matches the `"target_allocation_count"`.
// Avoid creating multiple endpoints for the same model unless required, as each
// endpoint consumes significant resources.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-inference-put-googleaistudio
func New(tp elastictransport.Interface) *PutGoogleaistudio {
	r := &PutGoogleaistudio{
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
func (r *PutGoogleaistudio) Raw(raw io.Reader) *PutGoogleaistudio {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutGoogleaistudio) Request(req *Request) *PutGoogleaistudio {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutGoogleaistudio) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutGoogleaistudio: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == tasktypeMask|googleaistudioinferenceidMask:
		path.WriteString("/")
		path.WriteString("_inference")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "tasktype", r.tasktype)
		}
		path.WriteString(r.tasktype)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "googleaistudioinferenceid", r.googleaistudioinferenceid)
		}
		path.WriteString(r.googleaistudioinferenceid)

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
func (r PutGoogleaistudio) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "inference.put_googleaistudio")
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
		instrument.BeforeRequest(req, "inference.put_googleaistudio")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_googleaistudio", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_googleaistudio")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutGoogleaistudio query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putgoogleaistudio.Response
func (r PutGoogleaistudio) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_googleaistudio")
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

// Header set a key, value pair in the PutGoogleaistudio headers map.
func (r *PutGoogleaistudio) Header(key, value string) *PutGoogleaistudio {
	r.headers.Set(key, value)

	return r
}

// TaskType The type of the inference task that the model will perform.
// API Name: tasktype
func (r *PutGoogleaistudio) _tasktype(tasktype string) *PutGoogleaistudio {
	r.paramSet |= tasktypeMask
	r.tasktype = tasktype

	return r
}

// GoogleaistudioInferenceId The unique identifier of the inference endpoint.
// API Name: googleaistudioinferenceid
func (r *PutGoogleaistudio) _googleaistudioinferenceid(googleaistudioinferenceid string) *PutGoogleaistudio {
	r.paramSet |= googleaistudioinferenceidMask
	r.googleaistudioinferenceid = googleaistudioinferenceid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutGoogleaistudio) ErrorTrace(errortrace bool) *PutGoogleaistudio {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutGoogleaistudio) FilterPath(filterpaths ...string) *PutGoogleaistudio {
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
func (r *PutGoogleaistudio) Human(human bool) *PutGoogleaistudio {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutGoogleaistudio) Pretty(pretty bool) *PutGoogleaistudio {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// The chunking configuration object.
// API name: chunking_settings
func (r *PutGoogleaistudio) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutGoogleaistudio {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ChunkingSettings = chunkingsettings.InferenceChunkingSettingsCaster()

	return r
}

// The type of service supported for the specified task type. In this case,
// `googleaistudio`.
// API name: service
func (r *PutGoogleaistudio) Service(service googleaiservicetype.GoogleAiServiceType) *PutGoogleaistudio {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Service = service
	return r
}

// Settings used to install the inference model. These settings are specific to
// the `googleaistudio` service.
// API name: service_settings
func (r *PutGoogleaistudio) ServiceSettings(servicesettings types.GoogleAiStudioServiceSettingsVariant) *PutGoogleaistudio {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ServiceSettings = *servicesettings.GoogleAiStudioServiceSettingsCaster()

	return r
}
