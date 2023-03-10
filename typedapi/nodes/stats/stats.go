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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Returns statistical information about nodes in the cluster.
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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/level"
)

const (
	nodeidMask = iota + 1

	metricMask

	indexmetricMask
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

	nodeid      string
	metric      string
	indexmetric string
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

// Returns statistical information about nodes in the cluster.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/cluster-nodes-stats.html
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
		path.WriteString("_nodes")
		path.WriteString("/")
		path.WriteString("stats")

		method = http.MethodGet
	case r.paramSet == nodeidMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")

		path.WriteString(r.nodeid)
		path.WriteString("/")
		path.WriteString("stats")

		method = http.MethodGet
	case r.paramSet == metricMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")
		path.WriteString("stats")
		path.WriteString("/")

		path.WriteString(r.metric)

		method = http.MethodGet
	case r.paramSet == nodeidMask|metricMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")

		path.WriteString(r.nodeid)
		path.WriteString("/")
		path.WriteString("stats")
		path.WriteString("/")

		path.WriteString(r.metric)

		method = http.MethodGet
	case r.paramSet == metricMask|indexmetricMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")
		path.WriteString("stats")
		path.WriteString("/")

		path.WriteString(r.metric)
		path.WriteString("/")

		path.WriteString(r.indexmetric)

		method = http.MethodGet
	case r.paramSet == nodeidMask|metricMask|indexmetricMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")

		path.WriteString(r.nodeid)
		path.WriteString("/")
		path.WriteString("stats")
		path.WriteString("/")

		path.WriteString(r.metric)
		path.WriteString("/")

		path.WriteString(r.indexmetric)

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

// NodeId Comma-separated list of node IDs or names used to limit returned information.
// API Name: nodeid
func (r *Stats) NodeId(v string) *Stats {
	r.paramSet |= nodeidMask
	r.nodeid = v

	return r
}

// Metric Limit the information returned to the specified metrics
// API Name: metric
func (r *Stats) Metric(v string) *Stats {
	r.paramSet |= metricMask
	r.metric = v

	return r
}

// IndexMetric Limit the information returned for indices metric to the specific index
// metrics. It can be used only if indices (or all) metric is specified.
// API Name: indexmetric
func (r *Stats) IndexMetric(v string) *Stats {
	r.paramSet |= indexmetricMask
	r.indexmetric = v

	return r
}

// CompletionFields Comma-separated list or wildcard expressions of fields to include in
// fielddata and suggest statistics.
// API name: completion_fields
func (r *Stats) CompletionFields(v string) *Stats {
	r.values.Set("completion_fields", v)

	return r
}

// FielddataFields Comma-separated list or wildcard expressions of fields to include in
// fielddata statistics.
// API name: fielddata_fields
func (r *Stats) FielddataFields(v string) *Stats {
	r.values.Set("fielddata_fields", v)

	return r
}

// Fields Comma-separated list or wildcard expressions of fields to include in the
// statistics.
// API name: fields
func (r *Stats) Fields(v string) *Stats {
	r.values.Set("fields", v)

	return r
}

// Groups Comma-separated list of search groups to include in the search statistics.
// API name: groups
func (r *Stats) Groups(b bool) *Stats {
	r.values.Set("groups", strconv.FormatBool(b))

	return r
}

// IncludeSegmentFileSizes If true, the call reports the aggregated disk usage of each one of the Lucene
// index files (only applies if segment stats are requested).
// API name: include_segment_file_sizes
func (r *Stats) IncludeSegmentFileSizes(b bool) *Stats {
	r.values.Set("include_segment_file_sizes", strconv.FormatBool(b))

	return r
}

// Level Indicates whether statistics are aggregated at the cluster, index, or shard
// level.
// API name: level
func (r *Stats) Level(enum level.Level) *Stats {
	r.values.Set("level", enum.String())

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *Stats) MasterTimeout(v string) *Stats {
	r.values.Set("master_timeout", v)

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and returns an error.
// API name: timeout
func (r *Stats) Timeout(v string) *Stats {
	r.values.Set("timeout", v)

	return r
}

// Types A comma-separated list of document types for the indexing index metric.
// API name: types
func (r *Stats) Types(v string) *Stats {
	r.values.Set("types", v)

	return r
}

// IncludeUnloadedSegments If set to true segment stats will include stats for segments that are not
// currently loaded into memory
// API name: include_unloaded_segments
func (r *Stats) IncludeUnloadedSegments(b bool) *Stats {
	r.values.Set("include_unloaded_segments", strconv.FormatBool(b))

	return r
}
