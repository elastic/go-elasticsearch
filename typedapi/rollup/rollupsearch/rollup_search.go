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

// Enables searching rolled-up data using the standard query DSL.
package rollupsearch

import (
	gobytes "bytes"
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
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type RollupSearch struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	index string
}

// NewRollupSearch type alias for index.
type NewRollupSearch func(index string) *RollupSearch

// NewRollupSearchFunc returns a new instance of RollupSearch with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRollupSearchFunc(tp elastictransport.Interface) NewRollupSearch {
	return func(index string) *RollupSearch {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Enables searching rolled-up data using the standard query DSL.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/rollup-search.html
func New(tp elastictransport.Interface) *RollupSearch {
	r := &RollupSearch{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *RollupSearch) Raw(raw io.Reader) *RollupSearch {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *RollupSearch) Request(req *Request) *RollupSearch {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *RollupSearch) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for RollupSearch: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_rollup_search")

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

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r RollupSearch) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the RollupSearch query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a rollupsearch.Response
func (r RollupSearch) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	r.TypedKeys(true)

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

// Header set a key, value pair in the RollupSearch headers map.
func (r *RollupSearch) Header(key, value string) *RollupSearch {
	r.headers.Set(key, value)

	return r
}

// Index Enables searching rolled-up data using the standard Query DSL.
// API Name: index
func (r *RollupSearch) _index(index string) *RollupSearch {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// RestTotalHitsAsInt Indicates whether hits.total should be rendered as an integer or an object in
// the rest search response
// API name: rest_total_hits_as_int
func (r *RollupSearch) RestTotalHitsAsInt(resttotalhitsasint bool) *RollupSearch {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(resttotalhitsasint))

	return r
}

// TypedKeys Specify whether aggregation and suggester names should be prefixed by their
// respective types in the response
// API name: typed_keys
func (r *RollupSearch) TypedKeys(typedkeys bool) *RollupSearch {
	r.values.Set("typed_keys", strconv.FormatBool(typedkeys))

	return r
}

// Aggregations Specifies aggregations.
// API name: aggregations
func (r *RollupSearch) Aggregations(aggregations map[string]types.Aggregations) *RollupSearch {

	r.req.Aggregations = aggregations

	return r
}

// Query Specifies a DSL query.
// API name: query
func (r *RollupSearch) Query(query *types.Query) *RollupSearch {

	r.req.Query = query

	return r
}

// Size Must be zero if set, as rollups work on pre-aggregated data.
// API name: size
func (r *RollupSearch) Size(size int) *RollupSearch {
	r.req.Size = &size

	return r
}
