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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

// Clear the cache of searchable snapshots.
package clearcache

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

type ClearCache struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	index string
}

// NewClearCache type alias for index.
type NewClearCache func() *ClearCache

// NewClearCacheFunc returns a new instance of ClearCache with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewClearCacheFunc(tp elastictransport.Interface) NewClearCache {
	return func() *ClearCache {
		n := New(tp)

		return n
	}
}

// Clear the cache of searchable snapshots.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/searchable-snapshots-apis.html
func New(tp elastictransport.Interface) *ClearCache {
	r := &ClearCache{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ClearCache) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_searchable_snapshots")
		path.WriteString("/")
		path.WriteString("cache")
		path.WriteString("/")
		path.WriteString("clear")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_searchable_snapshots")
		path.WriteString("/")
		path.WriteString("cache")
		path.WriteString("/")
		path.WriteString("clear")

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
func (r ClearCache) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the ClearCache query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a clearcache.Response
func (r ClearCache) Do(ctx context.Context) (Response, error) {

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
func (r ClearCache) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the ClearCache headers map.
func (r *ClearCache) Header(key, value string) *ClearCache {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names
// API Name: index
func (r *ClearCache) Index(v string) *ClearCache {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *ClearCache) ExpandWildcards(v string) *ClearCache {
	r.values.Set("expand_wildcards", v)

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *ClearCache) AllowNoIndices(b bool) *ClearCache {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *ClearCache) IgnoreUnavailable(b bool) *ClearCache {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// API name: pretty
func (r *ClearCache) Pretty(b bool) *ClearCache {
	r.values.Set("pretty", strconv.FormatBool(b))

	return r
}

// API name: human
func (r *ClearCache) Human(b bool) *ClearCache {
	r.values.Set("human", strconv.FormatBool(b))

	return r
}
