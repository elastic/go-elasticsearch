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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Returns information about indices: number of primaries and replicas, document
// counts, disk size, ...
package indices

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeunit"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Indices struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	index string
}

// NewIndices type alias for index.
type NewIndices func() *Indices

// NewIndicesFunc returns a new instance of Indices with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewIndicesFunc(tp elastictransport.Interface) NewIndices {
	return func() *Indices {
		n := New(tp)

		return n
	}
}

// Returns information about indices: number of primaries and replicas, document
// counts, disk size, ...
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/cat-indices.html
func New(tp elastictransport.Interface) *Indices {
	r := &Indices{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Indices) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("indices")

		method = http.MethodGet
	case r.paramSet == indexMask:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("indices")
		path.WriteString("/")

		path.WriteString(r.index)

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
func (r Indices) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Indices query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a indices.Response
func (r Indices) Do(ctx context.Context) (Response, error) {

	response := NewResponse()

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
func (r Indices) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the Indices headers map.
func (r *Indices) Header(key, value string) *Indices {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and aliases used to limit the
// request.
// Supports wildcards (`*`). To target all data streams and indices, omit this
// parameter or use `*` or `_all`.
// API Name: index
func (r *Indices) Index(index string) *Indices {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Bytes The unit used to display byte values.
// API name: bytes
func (r *Indices) Bytes(bytes bytes.Bytes) *Indices {
	r.values.Set("bytes", bytes.String())

	return r
}

// ExpandWildcards The type of index that wildcard patterns can match.
// API name: expand_wildcards
func (r *Indices) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Indices {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// Health The health status used to limit returned indices. By default, the response
// includes indices of any health status.
// API name: health
func (r *Indices) Health(health healthstatus.HealthStatus) *Indices {
	r.values.Set("health", health.String())

	return r
}

// IncludeUnloadedSegments If true, the response includes information from segments that are not loaded
// into memory.
// API name: include_unloaded_segments
func (r *Indices) IncludeUnloadedSegments(includeunloadedsegments bool) *Indices {
	r.values.Set("include_unloaded_segments", strconv.FormatBool(includeunloadedsegments))

	return r
}

// Pri If true, the response only includes information from primary shards.
// API name: pri
func (r *Indices) Pri(pri bool) *Indices {
	r.values.Set("pri", strconv.FormatBool(pri))

	return r
}

// Time The unit used to display time values.
// API name: time
func (r *Indices) Time(time timeunit.TimeUnit) *Indices {
	r.values.Set("time", time.String())

	return r
}
