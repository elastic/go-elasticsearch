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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

// Get the recovery point objective series.
//
// Get, for each bucket in the half-open range `[start, end)`, the maximum age
// that the newest recoverable state reached during that bucket. This API is
// intended for internal operator use. The range is divided into at most 1,000
// buckets. A bucket has no value if no retained recovery point covers its
// start.
package getrposeries

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

type GetRpoSeries struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewGetRpoSeries type alias for index.
type NewGetRpoSeries func() *GetRpoSeries

// NewGetRpoSeriesFunc returns a new instance of GetRpoSeries with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetRpoSeriesFunc(tp elastictransport.Interface) NewGetRpoSeries {
	return func() *GetRpoSeries {
		n := New(tp)

		return n
	}
}

// Get the recovery point objective series.
//
// Get, for each bucket in the half-open range `[start, end)`, the maximum age
// that the newest recoverable state reached during that bucket. This API is
// intended for internal operator use. The range is divided into at most 1,000
// buckets. A bucket has no value if no retained recovery point covers its
// start.
//
// [Elasticsearch]
//
// [Serverless] https://www.elastic.co/docs/api/doc/elasticsearch-serverless/operation/operation-data-recovery-get-rpo-series
func New(tp elastictransport.Interface) *GetRpoSeries {
	r := &GetRpoSeries{
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
func (r *GetRpoSeries) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_data_recovery")
		path.WriteString("/")
		path.WriteString("rpo")

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
func (r GetRpoSeries) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx = instrument.Start(providedCtx, "data_recovery.get_rpo_series")
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
		instrument.BeforeRequest(req, "data_recovery.get_rpo_series")
		if reader := instrument.RecordRequestBody(ctx, "data_recovery.get_rpo_series", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "data_recovery.get_rpo_series")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetRpoSeries query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getrposeries.Response
func (r GetRpoSeries) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "data_recovery.get_rpo_series")
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
func (r GetRpoSeries) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "data_recovery.get_rpo_series")
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
		err := fmt.Errorf("an error happened during the GetRpoSeries query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the GetRpoSeries headers map.
func (r *GetRpoSeries) Header(key, value string) *GetRpoSeries {
	r.headers.Set(key, value)

	return r
}

// Start The inclusive start of the range, as a date or epoch milliseconds. When
// omitted, it is seven days before `end`, rounded down to a whole multiple of
// `bucket_duration`, so that repeated requests with default parameters return
// the same bucket boundaries.
// API name: start
func (r *GetRpoSeries) Start(datetime string) *GetRpoSeries {
	r.values.Set("start", datetime)

	return r
}

// End The exclusive end of the range, as a date or epoch milliseconds.
// API name: end
func (r *GetRpoSeries) End(datetime string) *GetRpoSeries {
	r.values.Set("end", datetime)

	return r
}

// BucketDuration The duration of each bucket. It must be greater than zero, and the range must
// divide into at most 1,000 buckets.
// API name: bucket_duration
func (r *GetRpoSeries) BucketDuration(duration string) *GetRpoSeries {
	r.values.Set("bucket_duration", duration)

	return r
}

// MasterTimeout The period to wait for the master node. If the master node is not available
// before the timeout expires, the request fails and returns an error. To
// indicate that the request should never timeout, set it to `-1`.
// API name: master_timeout
func (r *GetRpoSeries) MasterTimeout(duration string) *GetRpoSeries {
	r.values.Set("master_timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetRpoSeries) ErrorTrace(errortrace bool) *GetRpoSeries {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetRpoSeries) FilterPath(filterpaths ...string) *GetRpoSeries {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and `"exists_time_in_millis":
// 3600000` for computers. When disabled the human readable values will be
// omitted. This makes sense for responses being consumed only by machines.
// API name: human
func (r *GetRpoSeries) Human(human bool) *GetRpoSeries {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use this
// option for debugging only.
// API name: pretty
func (r *GetRpoSeries) Pretty(pretty bool) *GetRpoSeries {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
