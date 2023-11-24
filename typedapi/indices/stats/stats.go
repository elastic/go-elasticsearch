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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Provides statistics on operations happening in an index.
package stats

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/level"
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

	buf *gobytes.Buffer

	paramSet int

	metric string
	index  string
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

// Provides statistics on operations happening in an index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-stats.html
func New(tp elastictransport.Interface) *Stats {
	r := &Stats{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
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

		path.WriteString(r.metric)

		method = http.MethodGet
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_stats")

		method = http.MethodGet
	case r.paramSet == indexMask|metricMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_stats")
		path.WriteString("/")

		path.WriteString(r.metric)

		method = http.MethodGet
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Stats) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Stats query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a stats.Response
func (r Stats) Do(ctx context.Context) (*Response, error) {

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
func (r Stats) IsSuccess(ctx context.Context) (bool, error) {
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
