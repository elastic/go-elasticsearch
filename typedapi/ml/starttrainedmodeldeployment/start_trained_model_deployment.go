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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Start a trained model deployment.
package starttrainedmodeldeployment

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/deploymentallocationstate"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/trainingpriority"
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

	buf *gobytes.Buffer

	paramSet int

	modelid string
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
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/start-trained-model-deployment.html
func New(tp elastictransport.Interface) *StartTrainedModelDeployment {
	r := &StartTrainedModelDeployment{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *StartTrainedModelDeployment) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == modelidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")

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
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
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
func (r StartTrainedModelDeployment) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the StartTrainedModelDeployment query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a starttrainedmodeldeployment.Response
func (r StartTrainedModelDeployment) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r StartTrainedModelDeployment) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
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

// NumberOfAllocations The number of model allocations on each node where the model is deployed.
// All allocations on a node share the same copy of the model in memory but use
// a separate set of threads to evaluate the model.
// Increasing this value generally increases the throughput.
// If this setting is greater than the number of hardware threads
// it will automatically be changed to a value less than the number of hardware
// threads.
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
