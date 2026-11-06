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

// List recoveries.
//
// List the recoveries that were started with the start recovery API, including
// those that have finished. This API is intended for internal operator use.
// Recoveries are ordered by start time, and the most recent are returned first
// by default. To retrieve further pages, pass the start time and identifier of
// the last recovery of a page as `search_after_started_at` and
// `search_after_recovery_id`.
package listrecoveries

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/recoveryrunstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ListRecoveries struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewListRecoveries type alias for index.
type NewListRecoveries func() *ListRecoveries

// NewListRecoveriesFunc returns a new instance of ListRecoveries with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewListRecoveriesFunc(tp elastictransport.Interface) NewListRecoveries {
	return func() *ListRecoveries {
		n := New(tp)

		return n
	}
}

// List recoveries.
//
// List the recoveries that were started with the start recovery API, including
// those that have finished. This API is intended for internal operator use.
// Recoveries are ordered by start time, and the most recent are returned first
// by default. To retrieve further pages, pass the start time and identifier of
// the last recovery of a page as `search_after_started_at` and
// `search_after_recovery_id`.
func New(tp elastictransport.Interface) *ListRecoveries {
	r := &ListRecoveries{
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
func (r *ListRecoveries) HttpRequest(ctx context.Context) (*http.Request, error) {
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
		path.WriteString("recoveries")

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
func (r ListRecoveries) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx = instrument.Start(providedCtx, "data_recovery.list_recoveries")
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
		instrument.BeforeRequest(req, "data_recovery.list_recoveries")
		if reader := instrument.RecordRequestBody(ctx, "data_recovery.list_recoveries", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "data_recovery.list_recoveries")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ListRecoveries query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a listrecoveries.Response
func (r ListRecoveries) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "data_recovery.list_recoveries")
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
func (r ListRecoveries) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "data_recovery.list_recoveries")
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
		err := fmt.Errorf("an error happened during the ListRecoveries query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the ListRecoveries headers map.
func (r *ListRecoveries) Header(key, value string) *ListRecoveries {
	r.headers.Set(key, value)

	return r
}

// Size The maximum number of recoveries to return. The value must be between 1 and
// 1000.
// API name: size
func (r *ListRecoveries) Size(size int) *ListRecoveries {
	r.values.Set("size", strconv.Itoa(size))

	return r
}

// StartedAtGte Return only recoveries that started at or after this time, as an ISO-8601
// timestamp.
// API name: started_at_gte
func (r *ListRecoveries) StartedAtGte(datetime string) *ListRecoveries {
	r.values.Set("started_at_gte", datetime)

	return r
}

// StartedAtLt Return only recoveries that started before this time, as an ISO-8601
// timestamp. It must not be earlier than `started_at_gte`.
// API name: started_at_lt
func (r *ListRecoveries) StartedAtLt(datetime string) *ListRecoveries {
	r.values.Set("started_at_lt", datetime)

	return r
}

// Order The order of the recoveries by start time.
// API name: order
func (r *ListRecoveries) Order(order sortorder.SortOrder) *ListRecoveries {
	r.values.Set("order", order.String())

	return r
}

// SearchAfterStartedAt Return only recoveries after this start time in the sort order, to retrieve
// the next page. It must be provided together with `search_after_recovery_id`.
// API name: search_after_started_at
func (r *ListRecoveries) SearchAfterStartedAt(datetime string) *ListRecoveries {
	r.values.Set("search_after_started_at", datetime)

	return r
}

// SearchAfterRecoveryId Return only recoveries after this identifier in the sort order for recoveries
// with the same start time, to retrieve the next page. It must be provided
// together with `search_after_started_at`.
// API name: search_after_recovery_id
func (r *ListRecoveries) SearchAfterRecoveryId(id string) *ListRecoveries {
	r.values.Set("search_after_recovery_id", id)

	return r
}

// Status A comma-separated list of statuses. Return only recoveries with one of these
// statuses. By default, recoveries with any status are returned.
// API name: status
func (r *ListRecoveries) Status(statuses ...recoveryrunstatus.RecoveryRunStatus) *ListRecoveries {
	tmp := []string{}
	for _, item := range statuses {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("status", strings.Join(tmp, ","))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ListRecoveries) ErrorTrace(errortrace bool) *ListRecoveries {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ListRecoveries) FilterPath(filterpaths ...string) *ListRecoveries {
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
func (r *ListRecoveries) Human(human bool) *ListRecoveries {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use this
// option for debugging only.
// API name: pretty
func (r *ListRecoveries) Pretty(pretty bool) *ListRecoveries {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
