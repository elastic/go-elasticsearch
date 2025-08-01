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

// Get index statistics.
// For data streams, the API retrieves statistics for the stream's backing
// indices.
//
// By default, the returned statistics are index-level with `primaries` and
// `total` aggregations.
// `primaries` are the values for only the primary shards.
// `total` are the accumulated values for both primary and replica shards.
//
// To get shard-level statistics, set the `level` parameter to `shards`.
//
// NOTE: When moving to another node, the shard-level statistics for a shard are
// cleared.
// Although the shard is no longer part of the node, that node retains any
// node-level statistics to which the shard contributed.
package stats

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/level"
)

const (
	metricMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Stats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	metric string
	index  string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewStats type alias for index.
type NewStats func() *Stats

// NewStatsFunc returns a new instance of Stats with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStatsFunc(tp elastictransport.Interface) NewStats {
	return func() *Stats {
		n := New(tp)

		return n
	}
}

// Get index statistics.
// For data streams, the API retrieves statistics for the stream's backing
// indices.
//
// By default, the returned statistics are index-level with `primaries` and
// `total` aggregations.
// `primaries` are the values for only the primary shards.
// `total` are the accumulated values for both primary and replica shards.
//
// To get shard-level statistics, set the `level` parameter to `shards`.
//
// NOTE: When moving to another node, the shard-level statistics for a shard are
// cleared.
// Although the shard is no longer part of the node, that node retains any
// node-level statistics to which the shard contributed.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-indices-stats
func New(tp elastictransport.Interface) *Stats {
	r := &Stats{
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
func (r *Stats) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_stats")

		method = http.MethodGet
	case r.paramSet == metricMask:
		path.WriteString("/")
		path.WriteString("_stats")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "metric", r.metric)
		}
		path.WriteString(r.metric)

		method = http.MethodGet
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_stats")

		method = http.MethodGet
	case r.paramSet == indexMask|metricMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_stats")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "metric", r.metric)
		}
		path.WriteString(r.metric)

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
func (r Stats) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.stats")
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
		instrument.BeforeRequest(req, "indices.stats")
		if reader := instrument.RecordRequestBody(ctx, "indices.stats", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.stats")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Stats query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a stats.Response
func (r Stats) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.stats")
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
func (r Stats) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.stats")
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
		err := fmt.Errorf("an error happened during the Stats query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Stats headers map.
func (r *Stats) Header(key, value string) *Stats {
	r.headers.Set(key, value)

	return r
}

// Metric Limit the information returned the specific metrics.
// API Name: metric
func (r *Stats) Metric(metric string) *Stats {
	r.paramSet |= metricMask
	r.metric = metric

	return r
}

// Index A comma-separated list of index names; use `_all` or empty string to perform
// the operation on all indices
// API Name: index
func (r *Stats) Index(index string) *Stats {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// CompletionFields Comma-separated list or wildcard expressions of fields to include in
// fielddata and suggest statistics.
// API name: completion_fields
func (r *Stats) CompletionFields(fields ...string) *Stats {
	r.values.Set("completion_fields", strings.Join(fields, ","))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match. If the request can target
// data streams, this argument
// determines whether wildcard expressions match hidden data streams. Supports
// comma-separated values,
// such as `open,hidden`.
// API name: expand_wildcards
func (r *Stats) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Stats {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// FielddataFields Comma-separated list or wildcard expressions of fields to include in
// fielddata statistics.
// API name: fielddata_fields
func (r *Stats) FielddataFields(fields ...string) *Stats {
	r.values.Set("fielddata_fields", strings.Join(fields, ","))

	return r
}

// Fields Comma-separated list or wildcard expressions of fields to include in the
// statistics.
// API name: fields
func (r *Stats) Fields(fields ...string) *Stats {
	r.values.Set("fields", strings.Join(fields, ","))

	return r
}

// ForbidClosedIndices If true, statistics are not collected from closed indices.
// API name: forbid_closed_indices
func (r *Stats) ForbidClosedIndices(forbidclosedindices bool) *Stats {
	r.values.Set("forbid_closed_indices", strconv.FormatBool(forbidclosedindices))

	return r
}

// Groups Comma-separated list of search groups to include in the search statistics.
// API name: groups
func (r *Stats) Groups(groups ...string) *Stats {
	tmp := []string{}
	for _, item := range groups {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("groups", strings.Join(tmp, ","))

	return r
}

// IncludeSegmentFileSizes If true, the call reports the aggregated disk usage of each one of the Lucene
// index files (only applies if segment stats are requested).
// API name: include_segment_file_sizes
func (r *Stats) IncludeSegmentFileSizes(includesegmentfilesizes bool) *Stats {
	r.values.Set("include_segment_file_sizes", strconv.FormatBool(includesegmentfilesizes))

	return r
}

// IncludeUnloadedSegments If true, the response includes information from segments that are not loaded
// into memory.
// API name: include_unloaded_segments
func (r *Stats) IncludeUnloadedSegments(includeunloadedsegments bool) *Stats {
	r.values.Set("include_unloaded_segments", strconv.FormatBool(includeunloadedsegments))

	return r
}

// Level Indicates whether statistics are aggregated at the cluster, index, or shard
// level.
// API name: level
func (r *Stats) Level(level level.Level) *Stats {
	r.values.Set("level", level.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Stats) ErrorTrace(errortrace bool) *Stats {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Stats) FilterPath(filterpaths ...string) *Stats {
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
func (r *Stats) Human(human bool) *Stats {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Stats) Pretty(pretty bool) *Stats {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
