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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Returns results matching a query expressed in Event Query Language (EQL)
package search

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Search struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	index string
}

// NewSearch type alias for index.
type NewSearch func(index string) *Search

// NewSearchFunc returns a new instance of Search with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSearchFunc(tp elastictransport.Interface) NewSearch {
	return func(index string) *Search {
		n := New(tp)

		n.Index(index)

		return n
	}
}

// Returns results matching a query expressed in Event Query Language (EQL)
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
func New(tp elastictransport.Interface) *Search {
	r := &Search{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Search) Raw(raw json.RawMessage) *Search {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Search) Request(req *Request) *Search {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Search) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Search: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_eql")
		path.WriteString("/")
		path.WriteString("search")

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

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r Search) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Search query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the Search headers map.
func (r *Search) Header(key, value string) *Search {
	r.headers.Set(key, value)

	return r
}

// Index The name of the index to scope the operation
// API Name: index
func (r *Search) Index(v string) *Search {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// API name: allow_no_indices
func (r *Search) AllowNoIndices(b bool) *Search {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// API name: expand_wildcards
func (r *Search) ExpandWildcards(value string) *Search {
	r.values.Set("expand_wildcards", value)

	return r
}

// IgnoreUnavailable If true, missing or closed indices are not included in the response.
// API name: ignore_unavailable
func (r *Search) IgnoreUnavailable(b bool) *Search {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// KeepAlive Period for which the search and its results are stored on the cluster.
// API name: keep_alive
func (r *Search) KeepAlive(value string) *Search {
	r.values.Set("keep_alive", value)

	return r
}

// KeepOnCompletion If true, the search and its results are stored on the cluster.
// API name: keep_on_completion
func (r *Search) KeepOnCompletion(b bool) *Search {
	r.values.Set("keep_on_completion", strconv.FormatBool(b))

	return r
}

// WaitForCompletionTimeout Timeout duration to wait for the request to finish. Defaults to no timeout,
// meaning the request waits for complete search results.
// API name: wait_for_completion_timeout
func (r *Search) WaitForCompletionTimeout(value string) *Search {
	r.values.Set("wait_for_completion_timeout", value)

	return r
}
