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

// Returns number of documents matching a query.
package count

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

type Count struct {
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

// NewCount type alias for index.
type NewCount func() *Count

// NewCountFunc returns a new instance of Count with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCountFunc(tp elastictransport.Interface) NewCount {
	return func() *Count {
		n := New(tp)

		return n
	}
}

// Returns number of documents matching a query.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
func New(tp elastictransport.Interface) *Count {
	r := &Count{
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
func (r *Count) Raw(raw io.Reader) *Count {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Count) Request(req *Request) *Count {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Count) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Count: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_count")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_count")

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
func (r Count) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Count query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a count.Response
func (r Count) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Count headers map.
func (r *Count) Header(key, value string) *Count {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and aliases to search.
// Supports wildcards (`*`).
// To search all data streams and indices, omit this parameter or use `*` or
// `_all`.
// API Name: index
func (r *Count) Index(index string) *Count {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices.
// This behavior applies even if the request targets other open indices.
// API name: allow_no_indices
func (r *Count) AllowNoIndices(allownoindices bool) *Count {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// Analyzer Analyzer to use for the query string.
// This parameter can only be used when the `q` query string parameter is
// specified.
// API name: analyzer
func (r *Count) Analyzer(analyzer string) *Count {
	r.values.Set("analyzer", analyzer)

	return r
}

// AnalyzeWildcard If `true`, wildcard and prefix queries are analyzed.
// This parameter can only be used when the `q` query string parameter is
// specified.
// API name: analyze_wildcard
func (r *Count) AnalyzeWildcard(analyzewildcard bool) *Count {
	r.values.Set("analyze_wildcard", strconv.FormatBool(analyzewildcard))

	return r
}

// DefaultOperator The default operator for query string query: `AND` or `OR`.
// This parameter can only be used when the `q` query string parameter is
// specified.
// API name: default_operator
func (r *Count) DefaultOperator(defaultoperator operator.Operator) *Count {
	r.values.Set("default_operator", defaultoperator.String())

	return r
}

// Df Field to use as default where no field prefix is given in the query string.
// This parameter can only be used when the `q` query string parameter is
// specified.
// API name: df
func (r *Count) Df(df string) *Count {
	r.values.Set("df", df)

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *Count) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Count {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreThrottled If `true`, concrete, expanded or aliased indices are ignored when frozen.
// API name: ignore_throttled
func (r *Count) IgnoreThrottled(ignorethrottled bool) *Count {
	r.values.Set("ignore_throttled", strconv.FormatBool(ignorethrottled))

	return r
}

// IgnoreUnavailable If `false`, the request returns an error if it targets a missing or closed
// index.
// API name: ignore_unavailable
func (r *Count) IgnoreUnavailable(ignoreunavailable bool) *Count {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Lenient If `true`, format-based query failures (such as providing text to a numeric
// field) in the query string will be ignored.
// API name: lenient
func (r *Count) Lenient(lenient bool) *Count {
	r.values.Set("lenient", strconv.FormatBool(lenient))

	return r
}

// MinScore Sets the minimum `_score` value that documents must have to be included in
// the result.
// API name: min_score
func (r *Count) MinScore(minscore string) *Count {
	r.values.Set("min_score", minscore)

	return r
}

// Preference Specifies the node or shard the operation should be performed on.
// Random by default.
// API name: preference
func (r *Count) Preference(preference string) *Count {
	r.values.Set("preference", preference)

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *Count) Routing(routing string) *Count {
	r.values.Set("routing", routing)

	return r
}

// TerminateAfter Maximum number of documents to collect for each shard.
// If a query reaches this limit, Elasticsearch terminates the query early.
// Elasticsearch collects documents before sorting.
// API name: terminate_after
func (r *Count) TerminateAfter(terminateafter string) *Count {
	r.values.Set("terminate_after", terminateafter)

	return r
}

// Q Query in the Lucene query string syntax.
// API name: q
func (r *Count) Q(q string) *Count {
	r.values.Set("q", q)

	return r
}

// Query Defines the search definition using the Query DSL.
// API name: query
func (r *Count) Query(query *types.Query) *Count {

	r.req.Query = query

	return r
}
