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

// Get data frame analytics job stats.
package getdataframeanalyticsstats

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetDataFrameAnalyticsStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewGetDataFrameAnalyticsStats type alias for index.
type NewGetDataFrameAnalyticsStats func() *GetDataFrameAnalyticsStats

// NewGetDataFrameAnalyticsStatsFunc returns a new instance of GetDataFrameAnalyticsStats with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetDataFrameAnalyticsStatsFunc(tp elastictransport.Interface) NewGetDataFrameAnalyticsStats {
	return func() *GetDataFrameAnalyticsStats {
		n := New(tp)

		return n
	}
}

// Get data frame analytics job stats.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-dfanalytics-stats.html
func New(tp elastictransport.Interface) *GetDataFrameAnalyticsStats {
	r := &GetDataFrameAnalyticsStats{
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
func (r *GetDataFrameAnalyticsStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("data_frame")
		path.WriteString("/")
		path.WriteString("analytics")
		path.WriteString("/")
		path.WriteString("_stats")

		method = http.MethodGet
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("data_frame")
		path.WriteString("/")
		path.WriteString("analytics")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)
		path.WriteString("/")
		path.WriteString("_stats")

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
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r GetDataFrameAnalyticsStats) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.get_data_frame_analytics_stats")
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
		instrument.BeforeRequest(req, "ml.get_data_frame_analytics_stats")
		if reader := instrument.RecordRequestBody(ctx, "ml.get_data_frame_analytics_stats", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.get_data_frame_analytics_stats")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetDataFrameAnalyticsStats query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getdataframeanalyticsstats.Response
func (r GetDataFrameAnalyticsStats) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_data_frame_analytics_stats")
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
func (r GetDataFrameAnalyticsStats) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_data_frame_analytics_stats")
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
		err := fmt.Errorf("an error happened during the GetDataFrameAnalyticsStats query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the GetDataFrameAnalyticsStats headers map.
func (r *GetDataFrameAnalyticsStats) Header(key, value string) *GetDataFrameAnalyticsStats {
	r.headers.Set(key, value)

	return r
}

// Id Identifier for the data frame analytics job. If you do not specify this
// option, the API returns information for the first hundred data frame
// analytics jobs.
// API Name: id
func (r *GetDataFrameAnalyticsStats) Id(id string) *GetDataFrameAnalyticsStats {
	r.paramSet |= idMask
	r.id = id

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
// 1. Contains wildcard expressions and there are no data frame analytics
// jobs that match.
// 2. Contains the `_all` string or no identifiers and there are no matches.
// 3. Contains wildcard expressions and there are only partial matches.
//
// The default value returns an empty data_frame_analytics array when there
// are no matches and the subset of results when there are partial matches.
// If this parameter is `false`, the request returns a 404 status code when
// there are no matches or only partial matches.
// API name: allow_no_match
func (r *GetDataFrameAnalyticsStats) AllowNoMatch(allownomatch bool) *GetDataFrameAnalyticsStats {
	r.values.Set("allow_no_match", strconv.FormatBool(allownomatch))

	return r
}

// From Skips the specified number of data frame analytics jobs.
// API name: from
func (r *GetDataFrameAnalyticsStats) From(from int) *GetDataFrameAnalyticsStats {
	r.values.Set("from", strconv.Itoa(from))

	return r
}

// Size Specifies the maximum number of data frame analytics jobs to obtain.
// API name: size
func (r *GetDataFrameAnalyticsStats) Size(size int) *GetDataFrameAnalyticsStats {
	r.values.Set("size", strconv.Itoa(size))

	return r
}

// Verbose Defines whether the stats response should be verbose.
// API name: verbose
func (r *GetDataFrameAnalyticsStats) Verbose(verbose bool) *GetDataFrameAnalyticsStats {
	r.values.Set("verbose", strconv.FormatBool(verbose))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetDataFrameAnalyticsStats) ErrorTrace(errortrace bool) *GetDataFrameAnalyticsStats {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetDataFrameAnalyticsStats) FilterPath(filterpaths ...string) *GetDataFrameAnalyticsStats {
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
func (r *GetDataFrameAnalyticsStats) Human(human bool) *GetDataFrameAnalyticsStats {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GetDataFrameAnalyticsStats) Pretty(pretty bool) *GetDataFrameAnalyticsStats {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
