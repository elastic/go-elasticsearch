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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Run a knn search.
//
// NOTE: The kNN search API has been replaced by the `knn` option in the search
// API.
//
// Perform a k-nearest neighbor (kNN) search on a dense_vector field and return
// the matching documents.
// Given a query vector, the API finds the k closest vectors and returns those
// documents as search hits.
//
// Elasticsearch uses the HNSW algorithm to support efficient kNN search.
// Like most kNN algorithms, HNSW is an approximate method that sacrifices
// result accuracy for improved search speed.
// This means the results returned are not always the true k closest neighbors.
//
// The kNN search API supports restricting the search using a filter.
// The search will return the top k documents that also match the filter query.
//
// A kNN search response has the exact same structure as a search API response.
// However, certain sections have a meaning specific to kNN search:
//
// * The document `_score` is determined by the similarity between the query and
// document vector.
// * The `hits.total` object contains the total number of nearest neighbor
// candidates considered, which is `num_candidates * num_shards`. The
// `hits.total.relation` will always be `eq`, indicating an exact value.
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

type KnnSearch struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
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

// Run a knn search.
//
// NOTE: The kNN search API has been replaced by the `knn` option in the search
// API.
//
// Perform a k-nearest neighbor (kNN) search on a dense_vector field and return
// the matching documents.
// Given a query vector, the API finds the k closest vectors and returns those
// documents as search hits.
//
// Elasticsearch uses the HNSW algorithm to support efficient kNN search.
// Like most kNN algorithms, HNSW is an approximate method that sacrifices
// result accuracy for improved search speed.
// This means the results returned are not always the true k closest neighbors.
//
// The kNN search API supports restricting the search using a filter.
// The search will return the top k documents that also match the filter query.
//
// A kNN search response has the exact same structure as a search API response.
// However, certain sections have a meaning specific to kNN search:
//
// * The document `_score` is determined by the similarity between the query and
// document vector.
// * The `hits.total` object contains the total number of nearest neighbor
// candidates considered, which is `num_candidates * num_shards`. The
// `hits.total.relation` will always be `eq`, indicating an exact value.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/knn-search-api.html
func New(tp elastictransport.Interface) *KnnSearch {
	r := &KnnSearch{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for KnnSearch: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
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
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
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
func (r KnnSearch) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "knn_search")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "knn_search")
		if reader := instrument.RecordRequestBody(ctx, "knn_search", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "knn_search")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the KnnSearch query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a knnsearch.Response
func (r KnnSearch) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "knn_search")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// Header set a key, value pair in the KnnSearch headers map.
func (r *KnnSearch) Header(key, value string) *KnnSearch {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names to search;
// use `_all` or to perform the operation on all indices.
// API Name: index
func (r *KnnSearch) _index(index string) *KnnSearch {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Routing A comma-separated list of specific routing values.
// API name: routing
func (r *KnnSearch) Routing(routing string) *KnnSearch {
	r.values.Set("routing", routing)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *KnnSearch) ErrorTrace(errortrace bool) *KnnSearch {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *KnnSearch) FilterPath(filterpaths ...string) *KnnSearch {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *KnnSearch) Human(human bool) *KnnSearch {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *KnnSearch) Pretty(pretty bool) *KnnSearch {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// DocvalueFields The request returns doc values for field names matching these patterns
// in the `hits.fields` property of the response.
// It accepts wildcard (`*`) patterns.
// API name: docvalue_fields
func (r *KnnSearch) DocvalueFields(docvaluefields ...types.FieldAndFormat) *KnnSearch {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.DocvalueFields = docvaluefields

	return r
}

// Fields The request returns values for field names matching these patterns
// in the `hits.fields` property of the response.
// It accepts wildcard (`*`) patterns.
// API name: fields
func (r *KnnSearch) Fields(fields ...string) *KnnSearch {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Fields = fields

	return r
}

// Filter A query to filter the documents that can match. The kNN search will return
// the top
// `k` documents that also match this filter. The value can be a single query or
// a
// list of queries. If `filter` isn't provided, all documents are allowed to
// match.
// API name: filter
func (r *KnnSearch) Filter(filters ...types.Query) *KnnSearch {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Filter = filters

	return r
}

// Knn The kNN query to run.
// API name: knn
func (r *KnnSearch) Knn(knn *types.CoreKnnQuery) *KnnSearch {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Knn = *knn

	return r
}

// Source_ Indicates which source fields are returned for matching documents. These
// fields are returned in the `hits._source` property of the search response.
// API name: _source
func (r *KnnSearch) Source_(sourceconfig types.SourceConfig) *KnnSearch {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Source_ = sourceconfig

	return r
}

// StoredFields A list of stored fields to return as part of a hit. If no fields are
// specified,
// no stored fields are included in the response. If this field is specified,
// the `_source`
// parameter defaults to `false`. You can pass `_source: true` to return both
// source fields
// and stored fields in the search response.
// API name: stored_fields
func (r *KnnSearch) StoredFields(fields ...string) *KnnSearch {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.StoredFields = fields

	return r
}
