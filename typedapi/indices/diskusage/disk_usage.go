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

// Analyzes the disk usage of each field of an index or data stream
package diskusage

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

	buf *gobytes.Buffer

	paramSet int

	index string
}

// NewDiskUsage type alias for index.
type NewDiskUsage func(index string) *DiskUsage

// NewDiskUsageFunc returns a new instance of DiskUsage with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDiskUsageFunc(tp elastictransport.Interface) NewDiskUsage {
	return func(index string) *DiskUsage {
		n := New(tp)

		n.Index(index)

		return n
	}
}

// Analyzes the disk usage of each field of an index or data stream
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-disk-usage.html
func New(tp elastictransport.Interface) *DiskUsage {
	r := &DiskUsage{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
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
func (r DiskUsage) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the DiskUsage query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a diskusage.Response
func (r DiskUsage) Do(ctx context.Context) (Response, error) {

	response := new(Response)

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return nil, err
		}

		return *response, nil

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
func (r DiskUsage) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the DiskUsage headers map.
func (r *DiskUsage) Header(key, value string) *DiskUsage {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and aliases used to limit the
// request. It’s recommended to execute this API with a single index (or the
// latest backing index of a data stream) as the API consumes resources
// significantly.
// API Name: index
func (r *DiskUsage) Index(v string) *DiskUsage {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// AllowNoIndices If false, the request returns an error if any wildcard expression, index
// alias, or _all value targets only missing or closed indices. This behavior
// applies even if the request targets other open indices. For example, a
// request targeting foo*,bar* returns an error if an index starts with foo but
// no index starts with bar.
// API name: allow_no_indices
func (r *DiskUsage) AllowNoIndices(b bool) *DiskUsage {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match. If the request can target
// data streams, this argument determines whether wildcard expressions match
// hidden data streams. Supports comma-separated values, such as open,hidden.
// API name: expand_wildcards
func (r *DiskUsage) ExpandWildcards(v string) *DiskUsage {
	r.values.Set("expand_wildcards", v)

	return r
}

// Flush If true, the API performs a flush before analysis. If false, the response may
// not include uncommitted data.
// API name: flush
func (r *DiskUsage) Flush(b bool) *DiskUsage {
	r.values.Set("flush", strconv.FormatBool(b))

	return r
}

// IgnoreUnavailable If true, missing or closed indices are not included in the response.
// API name: ignore_unavailable
func (r *DiskUsage) IgnoreUnavailable(b bool) *DiskUsage {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// RunExpensiveTasks Analyzing field disk usage is resource-intensive. To use the API, this
// parameter must be set to true.
// API name: run_expensive_tasks
func (r *DiskUsage) RunExpensiveTasks(b bool) *DiskUsage {
	r.values.Set("run_expensive_tasks", strconv.FormatBool(b))

	return r
}
