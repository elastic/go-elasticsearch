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
// https://github.com/elastic/elasticsearch-specification/tree/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c

// Allows a user to validate a potentially expensive query without executing it.
package validatequery

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ValidateQuery struct {
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

// NewValidateQuery type alias for index.
type NewValidateQuery func() *ValidateQuery

// NewValidateQueryFunc returns a new instance of ValidateQuery with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewValidateQueryFunc(tp elastictransport.Interface) NewValidateQuery {
	return func() *ValidateQuery {
		n := New(tp)

		return n
	}
}

// Allows a user to validate a potentially expensive query without executing it.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/search-validate.html
func New(tp elastictransport.Interface) *ValidateQuery {
	r := &ValidateQuery{
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
func (r *ValidateQuery) Raw(raw io.Reader) *ValidateQuery {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *ValidateQuery) Request(req *Request) *ValidateQuery {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ValidateQuery) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for ValidateQuery: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_validate")
		path.WriteString("/")
		path.WriteString("query")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_validate")
		path.WriteString("/")
		path.WriteString("query")

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
func (r ValidateQuery) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the ValidateQuery query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a validatequery.Response
func (r ValidateQuery) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the ValidateQuery headers map.
func (r *ValidateQuery) Header(key, value string) *ValidateQuery {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names to restrict the operation; use `_all`
// or empty string to perform the operation on all indices
// API Name: index
func (r *ValidateQuery) Index(index string) *ValidateQuery {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *ValidateQuery) AllowNoIndices(allownoindices bool) *ValidateQuery {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// AllShards Execute validation on all shards instead of one random shard per index
// API name: all_shards
func (r *ValidateQuery) AllShards(allshards bool) *ValidateQuery {
	r.values.Set("all_shards", strconv.FormatBool(allshards))

	return r
}

// Analyzer The analyzer to use for the query string
// API name: analyzer
func (r *ValidateQuery) Analyzer(analyzer string) *ValidateQuery {
	r.values.Set("analyzer", analyzer)

	return r
}

// AnalyzeWildcard Specify whether wildcard and prefix queries should be analyzed (default:
// false)
// API name: analyze_wildcard
func (r *ValidateQuery) AnalyzeWildcard(analyzewildcard bool) *ValidateQuery {
	r.values.Set("analyze_wildcard", strconv.FormatBool(analyzewildcard))

	return r
}

// DefaultOperator The default operator for query string query (AND or OR)
// API name: default_operator
func (r *ValidateQuery) DefaultOperator(defaultoperator operator.Operator) *ValidateQuery {
	r.values.Set("default_operator", defaultoperator.String())

	return r
}

// Df The field to use as default where no field prefix is given in the query
// string
// API name: df
func (r *ValidateQuery) Df(df string) *ValidateQuery {
	r.values.Set("df", df)

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *ValidateQuery) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ValidateQuery {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// Explain Return detailed information about the error
// API name: explain
func (r *ValidateQuery) Explain(explain bool) *ValidateQuery {
	r.values.Set("explain", strconv.FormatBool(explain))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *ValidateQuery) IgnoreUnavailable(ignoreunavailable bool) *ValidateQuery {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Lenient Specify whether format-based query failures (such as providing text to a
// numeric field) should be ignored
// API name: lenient
func (r *ValidateQuery) Lenient(lenient bool) *ValidateQuery {
	r.values.Set("lenient", strconv.FormatBool(lenient))

	return r
}

// Rewrite Provide a more detailed explanation showing the actual Lucene query that will
// be executed.
// API name: rewrite
func (r *ValidateQuery) Rewrite(rewrite bool) *ValidateQuery {
	r.values.Set("rewrite", strconv.FormatBool(rewrite))

	return r
}

// Q Query in the Lucene query string syntax
// API name: q
func (r *ValidateQuery) Q(q string) *ValidateQuery {
	r.values.Set("q", q)

	return r
}

// API name: query
func (r *ValidateQuery) Query(query *types.Query) *ValidateQuery {

	r.req.Query = query

	return r
}
