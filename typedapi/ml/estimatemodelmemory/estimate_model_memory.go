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

// Estimate job model memory usage.
//
// Make an estimation of the memory usage for an anomaly detection job model.
// The estimate is based on analysis configuration details for the job and
// cardinality
// estimates for the fields it references.
package estimatemodelmemory

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
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type EstimateModelMemory struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewEstimateModelMemory type alias for index.
type NewEstimateModelMemory func() *EstimateModelMemory

// NewEstimateModelMemoryFunc returns a new instance of EstimateModelMemory with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewEstimateModelMemoryFunc(tp elastictransport.Interface) NewEstimateModelMemory {
	return func() *EstimateModelMemory {
		n := New(tp)

		return n
	}
}

// Estimate job model memory usage.
//
// Make an estimation of the memory usage for an anomaly detection job model.
// The estimate is based on analysis configuration details for the job and
// cardinality
// estimates for the fields it references.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-ml-estimate-model-memory
func New(tp elastictransport.Interface) *EstimateModelMemory {
	r := &EstimateModelMemory{
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
func (r *EstimateModelMemory) Raw(raw io.Reader) *EstimateModelMemory {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *EstimateModelMemory) Request(req *Request) *EstimateModelMemory {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *EstimateModelMemory) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for EstimateModelMemory: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")
		path.WriteString("_estimate_model_memory")

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
func (r EstimateModelMemory) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.estimate_model_memory")
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
		instrument.BeforeRequest(req, "ml.estimate_model_memory")
		if reader := instrument.RecordRequestBody(ctx, "ml.estimate_model_memory", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.estimate_model_memory")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the EstimateModelMemory query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a estimatemodelmemory.Response
func (r EstimateModelMemory) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.estimate_model_memory")
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

// Header set a key, value pair in the EstimateModelMemory headers map.
func (r *EstimateModelMemory) Header(key, value string) *EstimateModelMemory {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *EstimateModelMemory) ErrorTrace(errortrace bool) *EstimateModelMemory {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *EstimateModelMemory) FilterPath(filterpaths ...string) *EstimateModelMemory {
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
func (r *EstimateModelMemory) Human(human bool) *EstimateModelMemory {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *EstimateModelMemory) Pretty(pretty bool) *EstimateModelMemory {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// For a list of the properties that you can specify in the
// `analysis_config` component of the body of this API.
// API name: analysis_config
func (r *EstimateModelMemory) AnalysisConfig(analysisconfig types.AnalysisConfigVariant) *EstimateModelMemory {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AnalysisConfig = analysisconfig.AnalysisConfigCaster()

	return r
}

// Estimates of the highest cardinality in a single bucket that is observed
// for influencer fields over the time period that the job analyzes data.
// To produce a good answer, values must be provided for all influencer
// fields. Providing values for fields that are not listed as `influencers`
// has no effect on the estimation.
// API name: max_bucket_cardinality
func (r *EstimateModelMemory) MaxBucketCardinality(maxbucketcardinality map[string]int64) *EstimateModelMemory {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxBucketCardinality = maxbucketcardinality
	return r
}

func (r *EstimateModelMemory) AddMaxBucketCardinality(key string, value int64) *EstimateModelMemory {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	var tmp map[string]int64
	if r.req.MaxBucketCardinality == nil {
		r.req.MaxBucketCardinality = make(map[string]int64)
	} else {
		tmp = r.req.MaxBucketCardinality
	}

	tmp[key] = value

	r.req.MaxBucketCardinality = tmp
	return r
}

// Estimates of the cardinality that is observed for fields over the whole
// time period that the job analyzes data. To produce a good answer, values
// must be provided for fields referenced in the `by_field_name`,
// `over_field_name` and `partition_field_name` of any detectors. Providing
// values for other fields has no effect on the estimation. It can be
// omitted from the request if no detectors have a `by_field_name`,
// `over_field_name` or `partition_field_name`.
// API name: overall_cardinality
func (r *EstimateModelMemory) OverallCardinality(overallcardinality map[string]int64) *EstimateModelMemory {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.OverallCardinality = overallcardinality
	return r
}

func (r *EstimateModelMemory) AddOverallCardinality(key string, value int64) *EstimateModelMemory {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	var tmp map[string]int64
	if r.req.OverallCardinality == nil {
		r.req.OverallCardinality = make(map[string]int64)
	} else {
		tmp = r.req.OverallCardinality
	}

	tmp[key] = value

	r.req.OverallCardinality = tmp
	return r
}
