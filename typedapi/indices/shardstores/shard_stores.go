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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Provides store information for shard copies of indices.
package shardstores

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

type ShardStores struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	index string
}

// NewShardStores type alias for index.
type NewShardStores func() *ShardStores

// NewShardStoresFunc returns a new instance of ShardStores with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewShardStoresFunc(tp elastictransport.Interface) NewShardStores {
	return func() *ShardStores {
		n := New(tp)

		return n
	}
}

// Provides store information for shard copies of indices.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-shards-stores.html
func New(tp elastictransport.Interface) *ShardStores {
	r := &ShardStores{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ShardStores) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_shard_stores")

		method = http.MethodGet
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_shard_stores")

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
func (r ShardStores) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the ShardStores query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a shardstores.Response
func (r ShardStores) Do(ctx context.Context) (*Response, error) {

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
func (r ShardStores) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the ShardStores headers map.
func (r *ShardStores) Header(key, value string) *ShardStores {
	r.headers.Set(key, value)

	return r
}

// Index List of data streams, indices, and aliases used to limit the request.
// API Name: index
func (r *ShardStores) Index(v string) *ShardStores {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// AllowNoIndices If false, the request returns an error if any wildcard expression, index
// alias, or _all
// value targets only missing or closed indices. This behavior applies even if
// the request
// targets other open indices.
// API name: allow_no_indices
func (r *ShardStores) AllowNoIndices(b bool) *ShardStores {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match. If the request can target
// data streams,
// this argument determines whether wildcard expressions match hidden data
// streams.
// API name: expand_wildcards
func (r *ShardStores) ExpandWildcards(v string) *ShardStores {
	r.values.Set("expand_wildcards", v)

	return r
}

// IgnoreUnavailable If true, missing or closed indices are not included in the response.
// API name: ignore_unavailable
func (r *ShardStores) IgnoreUnavailable(b bool) *ShardStores {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// Status List of shard health statuses used to limit the request.
// API name: status
func (r *ShardStores) Status(v string) *ShardStores {
	r.values.Set("status", v)

	return r
}
