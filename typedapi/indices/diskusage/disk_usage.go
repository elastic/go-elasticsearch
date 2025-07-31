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

// Analyze the index disk usage.
// Analyze the disk usage of each field of an index or data stream.
// This API might not support indices created in previous Elasticsearch
// versions.
// The result of a small index can be inaccurate as some parts of an index might
// not be analyzed by the API.
//
// NOTE: The total size of fields of the analyzed shards of the index in the
// response is usually smaller than the index `store_size` value because some
// small metadata files are ignored and some parts of data files might not be
// scanned by the API.
// Since stored fields are stored together in a compressed format, the sizes of
// stored fields are also estimates and can be inaccurate.
// The stored size of the `_id` field is likely underestimated while the
// `_source` field is overestimated.
package diskusage

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DiskUsage struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewDiskUsage type alias for index.
type NewDiskUsage func(index string) *DiskUsage

// NewDiskUsageFunc returns a new instance of DiskUsage with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDiskUsageFunc(tp elastictransport.Interface) NewDiskUsage {
	return func(index string) *DiskUsage {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Analyze the index disk usage.
// Analyze the disk usage of each field of an index or data stream.
// This API might not support indices created in previous Elasticsearch
// versions.
// The result of a small index can be inaccurate as some parts of an index might
// not be analyzed by the API.
//
// NOTE: The total size of fields of the analyzed shards of the index in the
// response is usually smaller than the index `store_size` value because some
// small metadata files are ignored and some parts of data files might not be
// scanned by the API.
// Since stored fields are stored together in a compressed format, the sizes of
// stored fields are also estimates and can be inaccurate.
// The stored size of the `_id` field is likely underestimated while the
// `_source` field is overestimated.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-disk-usage.html
func New(tp elastictransport.Interface) *DiskUsage {
	r := &DiskUsage{
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
func (r *DiskUsage) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_disk_usage")

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r DiskUsage) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.disk_usage")
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
		instrument.BeforeRequest(req, "indices.disk_usage")
		if reader := instrument.RecordRequestBody(ctx, "indices.disk_usage", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.disk_usage")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the DiskUsage query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a diskusage.Response
func (r DiskUsage) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.disk_usage")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := new(Response)

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return *response, nil
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
func (r DiskUsage) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.disk_usage")
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
		err := fmt.Errorf("an error happened during the DiskUsage query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the DiskUsage headers map.
func (r *DiskUsage) Header(key, value string) *DiskUsage {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and aliases used to limit the
// request.
// It’s recommended to execute this API with a single index (or the latest
// backing index of a data stream) as the API consumes resources significantly.
// API Name: index
func (r *DiskUsage) _index(index string) *DiskUsage {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If false, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices.
// This behavior applies even if the request targets other open indices.
// For example, a request targeting `foo*,bar*` returns an error if an index
// starts with `foo` but no index starts with `bar`.
// API name: allow_no_indices
func (r *DiskUsage) AllowNoIndices(allownoindices bool) *DiskUsage {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *DiskUsage) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *DiskUsage {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// Flush If `true`, the API performs a flush before analysis.
// If `false`, the response may not include uncommitted data.
// API name: flush
func (r *DiskUsage) Flush(flush bool) *DiskUsage {
	r.values.Set("flush", strconv.FormatBool(flush))

	return r
}

// IgnoreUnavailable If `true`, missing or closed indices are not included in the response.
// API name: ignore_unavailable
func (r *DiskUsage) IgnoreUnavailable(ignoreunavailable bool) *DiskUsage {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// RunExpensiveTasks Analyzing field disk usage is resource-intensive.
// To use the API, this parameter must be set to `true`.
// API name: run_expensive_tasks
func (r *DiskUsage) RunExpensiveTasks(runexpensivetasks bool) *DiskUsage {
	r.values.Set("run_expensive_tasks", strconv.FormatBool(runexpensivetasks))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *DiskUsage) ErrorTrace(errortrace bool) *DiskUsage {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *DiskUsage) FilterPath(filterpaths ...string) *DiskUsage {
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
func (r *DiskUsage) Human(human bool) *DiskUsage {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *DiskUsage) Pretty(pretty bool) *DiskUsage {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
