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

// Performs a kNN search.
package knnsearch

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type KnnSearch struct {
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

// NewKnnSearch type alias for index.
type NewKnnSearch func(index string) *KnnSearch

// NewKnnSearchFunc returns a new instance of KnnSearch with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewKnnSearchFunc(tp elastictransport.Interface) NewKnnSearch {
	return func(index string) *KnnSearch {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Performs a kNN search.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/search-search.html
func New(tp elastictransport.Interface) *KnnSearch {
	r := &KnnSearch{
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
func (r *KnnSearch) Raw(raw io.Reader) *KnnSearch {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *KnnSearch) Request(req *Request) *KnnSearch {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *KnnSearch) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for KnnSearch: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_knn_search")

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
func (r KnnSearch) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the KnnSearch query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a knnsearch.Response
func (r KnnSearch) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the KnnSearch headers map.
func (r *KnnSearch) Header(key, value string) *KnnSearch {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names to search;
// use `_all` or to perform the operation on all indices
// API Name: index
func (r *KnnSearch) _index(index string) *KnnSearch {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Routing A comma-separated list of specific routing values
// API name: routing
func (r *KnnSearch) Routing(routing string) *KnnSearch {
	r.values.Set("routing", routing)

	return r
}

// DocvalueFields The request returns doc values for field names matching these patterns
// in the hits.fields property of the response. Accepts wildcard (*) patterns.
// API name: docvalue_fields
func (r *KnnSearch) DocvalueFields(docvaluefields ...types.FieldAndFormat) *KnnSearch {
	r.req.DocvalueFields = docvaluefields

	return r
}

// Fields The request returns values for field names matching these patterns
// in the hits.fields property of the response. Accepts wildcard (*) patterns.
// API name: fields
func (r *KnnSearch) Fields(fields ...string) *KnnSearch {
	r.req.Fields = fields

	return r
}

// Filter Query to filter the documents that can match. The kNN search will return the
// top
// `k` documents that also match this filter. The value can be a single query or
// a
// list of queries. If `filter` isn't provided, all documents are allowed to
// match.
// API name: filter
func (r *KnnSearch) Filter(filters ...types.Query) *KnnSearch {
	r.req.Filter = filters

	return r
}

// Knn kNN query to execute
// API name: knn
func (r *KnnSearch) Knn(knn *types.CoreKnnQuery) *KnnSearch {

	r.req.Knn = *knn

	return r
}

// Source_ Indicates which source fields are returned for matching documents. These
// fields are returned in the hits._source property of the search response.
// API name: _source
func (r *KnnSearch) Source_(sourceconfig types.SourceConfig) *KnnSearch {
	r.req.Source_ = sourceconfig

	return r
}

// StoredFields List of stored fields to return as part of a hit. If no fields are specified,
// no stored fields are included in the response. If this field is specified,
// the _source
// parameter defaults to false. You can pass _source: true to return both source
// fields
// and stored fields in the search response.
// API name: stored_fields
func (r *KnnSearch) StoredFields(fields ...string) *KnnSearch {
	r.req.StoredFields = fields

	return r
}
