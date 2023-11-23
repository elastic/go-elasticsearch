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

// Open a point in time that can be used in subsequent searches
package openpointintime

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
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type OpenPointInTime struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	index string
}

// NewOpenPointInTime type alias for index.
type NewOpenPointInTime func(index string) *OpenPointInTime

// NewOpenPointInTimeFunc returns a new instance of OpenPointInTime with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewOpenPointInTimeFunc(tp elastictransport.Interface) NewOpenPointInTime {
	return func(index string) *OpenPointInTime {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Open a point in time that can be used in subsequent searches
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
func New(tp elastictransport.Interface) *OpenPointInTime {
	r := &OpenPointInTime{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *OpenPointInTime) HttpRequest(ctx context.Context) (*http.Request, error) {
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
		path.WriteString("_pit")

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
func (r OpenPointInTime) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the OpenPointInTime query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a openpointintime.Response
func (r OpenPointInTime) Do(ctx context.Context) (*Response, error) {

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
func (r OpenPointInTime) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the OpenPointInTime headers map.
func (r *OpenPointInTime) Header(key, value string) *OpenPointInTime {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names to open point in time; use `_all` or
// empty string to perform the operation on all indices
// API Name: index
func (r *OpenPointInTime) _index(index string) *OpenPointInTime {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// KeepAlive Extends the time to live of the corresponding point in time.
// API name: keep_alive
func (r *OpenPointInTime) KeepAlive(duration string) *OpenPointInTime {
	r.values.Set("keep_alive", duration)

	return r
}

// IgnoreUnavailable If `false`, the request returns an error if it targets a missing or closed
// index.
// API name: ignore_unavailable
func (r *OpenPointInTime) IgnoreUnavailable(ignoreunavailable bool) *OpenPointInTime {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Preference Specifies the node or shard the operation should be performed on.
// Random by default.
// API name: preference
func (r *OpenPointInTime) Preference(preference string) *OpenPointInTime {
	r.values.Set("preference", preference)

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *OpenPointInTime) Routing(routing string) *OpenPointInTime {
	r.values.Set("routing", routing)

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`. Valid values are:
// `all`, `open`, `closed`, `hidden`, `none`.
// API name: expand_wildcards
func (r *OpenPointInTime) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *OpenPointInTime {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}
