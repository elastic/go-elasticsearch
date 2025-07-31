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

// Update a trained model deployment.
package updatetrainedmodeldeployment

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
	modelidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateTrainedModelDeployment struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	modelid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewUpdateTrainedModelDeployment type alias for index.
type NewUpdateTrainedModelDeployment func(modelid string) *UpdateTrainedModelDeployment

// NewUpdateTrainedModelDeploymentFunc returns a new instance of UpdateTrainedModelDeployment with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateTrainedModelDeploymentFunc(tp elastictransport.Interface) NewUpdateTrainedModelDeployment {
	return func(modelid string) *UpdateTrainedModelDeployment {
		n := New(tp)

		n._modelid(modelid)

		return n
	}
}

// Update a trained model deployment.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-ml-update-trained-model-deployment
func New(tp elastictransport.Interface) *UpdateTrainedModelDeployment {
	r := &UpdateTrainedModelDeployment{
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
func (r *UpdateTrainedModelDeployment) Raw(raw io.Reader) *UpdateTrainedModelDeployment {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *UpdateTrainedModelDeployment) Request(req *Request) *UpdateTrainedModelDeployment {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *UpdateTrainedModelDeployment) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for UpdateTrainedModelDeployment: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == modelidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "modelid", r.modelid)
		}
		path.WriteString(r.modelid)
		path.WriteString("/")
		path.WriteString("deployment")
		path.WriteString("/")
		path.WriteString("_update")

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
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r UpdateTrainedModelDeployment) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.update_trained_model_deployment")
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
		instrument.BeforeRequest(req, "ml.update_trained_model_deployment")
		if reader := instrument.RecordRequestBody(ctx, "ml.update_trained_model_deployment", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.update_trained_model_deployment")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the UpdateTrainedModelDeployment query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a updatetrainedmodeldeployment.Response
func (r UpdateTrainedModelDeployment) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.update_trained_model_deployment")
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

// Header set a key, value pair in the UpdateTrainedModelDeployment headers map.
func (r *UpdateTrainedModelDeployment) Header(key, value string) *UpdateTrainedModelDeployment {
	r.headers.Set(key, value)

	return r
}

// ModelId The unique identifier of the trained model. Currently, only PyTorch models
// are supported.
// API Name: modelid
func (r *UpdateTrainedModelDeployment) _modelid(modelid string) *UpdateTrainedModelDeployment {
	r.paramSet |= modelidMask
	r.modelid = modelid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *UpdateTrainedModelDeployment) ErrorTrace(errortrace bool) *UpdateTrainedModelDeployment {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *UpdateTrainedModelDeployment) FilterPath(filterpaths ...string) *UpdateTrainedModelDeployment {
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
func (r *UpdateTrainedModelDeployment) Human(human bool) *UpdateTrainedModelDeployment {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *UpdateTrainedModelDeployment) Pretty(pretty bool) *UpdateTrainedModelDeployment {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// AdaptiveAllocations Adaptive allocations configuration. When enabled, the number of allocations
// is set based on the current load.
// If adaptive_allocations is enabled, do not set the number of allocations
// manually.
// API name: adaptive_allocations
func (r *UpdateTrainedModelDeployment) AdaptiveAllocations(adaptiveallocations *types.AdaptiveAllocationsSettings) *UpdateTrainedModelDeployment {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AdaptiveAllocations = adaptiveallocations

	return r
}

// NumberOfAllocations The number of model allocations on each node where the model is deployed.
// All allocations on a node share the same copy of the model in memory but use
// a separate set of threads to evaluate the model.
// Increasing this value generally increases the throughput.
// If this setting is greater than the number of hardware threads
// it will automatically be changed to a value less than the number of hardware
// threads.
// If adaptive_allocations is enabled, do not set this value, because it’s
// automatically set.
// API name: number_of_allocations
func (r *UpdateTrainedModelDeployment) NumberOfAllocations(numberofallocations int) *UpdateTrainedModelDeployment {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.NumberOfAllocations = &numberofallocations

	return r
}
