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

// Start a trained model deployment.
// It allocates the model to every machine learning node.
package starttrainedmodeldeployment

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/deploymentallocationstate"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/trainingpriority"
)

const (
	modelidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type StartTrainedModelDeployment struct {
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

// NewStartTrainedModelDeployment type alias for index.
type NewStartTrainedModelDeployment func(modelid string) *StartTrainedModelDeployment

// NewStartTrainedModelDeploymentFunc returns a new instance of StartTrainedModelDeployment with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStartTrainedModelDeploymentFunc(tp elastictransport.Interface) NewStartTrainedModelDeployment {
	return func(modelid string) *StartTrainedModelDeployment {
		n := New(tp)

		n._modelid(modelid)

		return n
	}
}

// Start a trained model deployment.
// It allocates the model to every machine learning node.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-ml-start-trained-model-deployment
func New(tp elastictransport.Interface) *StartTrainedModelDeployment {
	r := &StartTrainedModelDeployment{
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
func (r *StartTrainedModelDeployment) Raw(raw io.Reader) *StartTrainedModelDeployment {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *StartTrainedModelDeployment) Request(req *Request) *StartTrainedModelDeployment {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *StartTrainedModelDeployment) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for StartTrainedModelDeployment: %w", err)
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
		path.WriteString("_start")

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
func (r StartTrainedModelDeployment) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.start_trained_model_deployment")
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
		instrument.BeforeRequest(req, "ml.start_trained_model_deployment")
		if reader := instrument.RecordRequestBody(ctx, "ml.start_trained_model_deployment", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.start_trained_model_deployment")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the StartTrainedModelDeployment query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a starttrainedmodeldeployment.Response
func (r StartTrainedModelDeployment) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.start_trained_model_deployment")
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

// Header set a key, value pair in the StartTrainedModelDeployment headers map.
func (r *StartTrainedModelDeployment) Header(key, value string) *StartTrainedModelDeployment {
	r.headers.Set(key, value)

	return r
}

// ModelId The unique identifier of the trained model. Currently, only PyTorch models
// are supported.
// API Name: modelid
func (r *StartTrainedModelDeployment) _modelid(modelid string) *StartTrainedModelDeployment {
	r.paramSet |= modelidMask
	r.modelid = modelid

	return r
}

// CacheSize The inference cache size (in memory outside the JVM heap) per node for the
// model.
// The default value is the same size as the `model_size_bytes`. To disable the
// cache,
// `0b` can be provided.
// API name: cache_size
func (r *StartTrainedModelDeployment) CacheSize(bytesize string) *StartTrainedModelDeployment {
	r.values.Set("cache_size", bytesize)

	return r
}

// DeploymentId A unique identifier for the deployment of the model.
// API name: deployment_id
func (r *StartTrainedModelDeployment) DeploymentId(deploymentid string) *StartTrainedModelDeployment {
	r.values.Set("deployment_id", deploymentid)

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
func (r *StartTrainedModelDeployment) NumberOfAllocations(numberofallocations int) *StartTrainedModelDeployment {
	r.values.Set("number_of_allocations", strconv.Itoa(numberofallocations))

	return r
}

// Priority The deployment priority.
// API name: priority
func (r *StartTrainedModelDeployment) Priority(priority trainingpriority.TrainingPriority) *StartTrainedModelDeployment {
	r.values.Set("priority", priority.String())

	return r
}

// QueueCapacity Specifies the number of inference requests that are allowed in the queue.
// After the number of requests exceeds
// this value, new requests are rejected with a 429 error.
// API name: queue_capacity
func (r *StartTrainedModelDeployment) QueueCapacity(queuecapacity int) *StartTrainedModelDeployment {
	r.values.Set("queue_capacity", strconv.Itoa(queuecapacity))

	return r
}

// ThreadsPerAllocation Sets the number of threads used by each model allocation during inference.
// This generally increases
// the inference speed. The inference process is a compute-bound process; any
// number
// greater than the number of available hardware threads on the machine does not
// increase the
// inference speed. If this setting is greater than the number of hardware
// threads
// it will automatically be changed to a value less than the number of hardware
// threads.
// API name: threads_per_allocation
func (r *StartTrainedModelDeployment) ThreadsPerAllocation(threadsperallocation int) *StartTrainedModelDeployment {
	r.values.Set("threads_per_allocation", strconv.Itoa(threadsperallocation))

	return r
}

// Timeout Specifies the amount of time to wait for the model to deploy.
// API name: timeout
func (r *StartTrainedModelDeployment) Timeout(duration string) *StartTrainedModelDeployment {
	r.values.Set("timeout", duration)

	return r
}

// WaitFor Specifies the allocation status to wait for before returning.
// API name: wait_for
func (r *StartTrainedModelDeployment) WaitFor(waitfor deploymentallocationstate.DeploymentAllocationState) *StartTrainedModelDeployment {
	r.values.Set("wait_for", waitfor.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *StartTrainedModelDeployment) ErrorTrace(errortrace bool) *StartTrainedModelDeployment {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *StartTrainedModelDeployment) FilterPath(filterpaths ...string) *StartTrainedModelDeployment {
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
func (r *StartTrainedModelDeployment) Human(human bool) *StartTrainedModelDeployment {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *StartTrainedModelDeployment) Pretty(pretty bool) *StartTrainedModelDeployment {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Adaptive allocations configuration. When enabled, the number of allocations
// is set based on the current load.
// If adaptive_allocations is enabled, do not set the number of allocations
// manually.
// API name: adaptive_allocations
func (r *StartTrainedModelDeployment) AdaptiveAllocations(adaptiveallocations types.AdaptiveAllocationsSettingsVariant) *StartTrainedModelDeployment {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AdaptiveAllocations = adaptiveallocations.AdaptiveAllocationsSettingsCaster()

	return r
}
