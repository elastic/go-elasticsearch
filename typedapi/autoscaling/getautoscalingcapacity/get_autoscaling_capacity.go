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

// Get the autoscaling capacity.
//
// NOTE: This feature is designed for indirect use by Elasticsearch Service,
// Elastic Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
// supported.
//
// This API gets the current autoscaling capacity based on the configured
// autoscaling policy.
// It will return information to size the cluster appropriately to the current
// workload.
//
// The `required_capacity` is calculated as the maximum of the
// `required_capacity` result of all individual deciders that are enabled for
// the policy.
//
// The operator should verify that the `current_nodes` match the operator’s
// knowledge of the cluster to avoid making autoscaling decisions based on stale
// or incomplete information.
//
// The response contains decider-specific information you can use to diagnose
// how and why autoscaling determined a certain capacity was required.
// This information is provided for diagnosis only.
// Do not use this information to make autoscaling decisions.
package getautoscalingcapacity

import (
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
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetAutoscalingCapacity struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewGetAutoscalingCapacity type alias for index.
type NewGetAutoscalingCapacity func() *GetAutoscalingCapacity

// NewGetAutoscalingCapacityFunc returns a new instance of GetAutoscalingCapacity with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetAutoscalingCapacityFunc(tp elastictransport.Interface) NewGetAutoscalingCapacity {
	return func() *GetAutoscalingCapacity {
		n := New(tp)

		return n
	}
}

// Get the autoscaling capacity.
//
// NOTE: This feature is designed for indirect use by Elasticsearch Service,
// Elastic Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
// supported.
//
// This API gets the current autoscaling capacity based on the configured
// autoscaling policy.
// It will return information to size the cluster appropriately to the current
// workload.
//
// The `required_capacity` is calculated as the maximum of the
// `required_capacity` result of all individual deciders that are enabled for
// the policy.
//
// The operator should verify that the `current_nodes` match the operator’s
// knowledge of the cluster to avoid making autoscaling decisions based on stale
// or incomplete information.
//
// The response contains decider-specific information you can use to diagnose
// how and why autoscaling determined a certain capacity was required.
// This information is provided for diagnosis only.
// Do not use this information to make autoscaling decisions.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-autoscaling-get-autoscaling-capacity
func New(tp elastictransport.Interface) *GetAutoscalingCapacity {
	r := &GetAutoscalingCapacity{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetAutoscalingCapacity) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_autoscaling")
		path.WriteString("/")
		path.WriteString("capacity")

		method = http.MethodGet
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r GetAutoscalingCapacity) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "autoscaling.get_autoscaling_capacity")
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
		instrument.BeforeRequest(req, "autoscaling.get_autoscaling_capacity")
		if reader := instrument.RecordRequestBody(ctx, "autoscaling.get_autoscaling_capacity", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "autoscaling.get_autoscaling_capacity")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetAutoscalingCapacity query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getautoscalingcapacity.Response
func (r GetAutoscalingCapacity) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "autoscaling.get_autoscaling_capacity")
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetAutoscalingCapacity) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "autoscaling.get_autoscaling_capacity")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the GetAutoscalingCapacity query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the GetAutoscalingCapacity headers map.
func (r *GetAutoscalingCapacity) Header(key, value string) *GetAutoscalingCapacity {
	r.headers.Set(key, value)

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *GetAutoscalingCapacity) MasterTimeout(duration string) *GetAutoscalingCapacity {
	r.values.Set("master_timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetAutoscalingCapacity) ErrorTrace(errortrace bool) *GetAutoscalingCapacity {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetAutoscalingCapacity) FilterPath(filterpaths ...string) *GetAutoscalingCapacity {
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
func (r *GetAutoscalingCapacity) Human(human bool) *GetAutoscalingCapacity {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GetAutoscalingCapacity) Pretty(pretty bool) *GetAutoscalingCapacity {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
