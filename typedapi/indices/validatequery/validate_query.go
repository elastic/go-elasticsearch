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

// Validate a query.
// Validates a query without running it.
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

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
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

// Validate a query.
// Validates a query without running it.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-validate.html
func New(tp elastictransport.Interface) *ValidateQuery {
	r := &ValidateQuery{
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for ValidateQuery: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
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

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
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
func (r ValidateQuery) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.validate_query")
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
		instrument.BeforeRequest(req, "indices.validate_query")
		if reader := instrument.RecordRequestBody(ctx, "indices.validate_query", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.validate_query")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ValidateQuery query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a validatequery.Response
func (r ValidateQuery) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.validate_query")
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

// Header set a key, value pair in the ValidateQuery headers map.
func (r *ValidateQuery) Header(key, value string) *ValidateQuery {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and aliases to search.
// Supports wildcards (`*`).
// To search all data streams or indices, omit this parameter or use `*` or
// `_all`.
// API Name: index
func (r *ValidateQuery) Index(index string) *ValidateQuery {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices.
// This behavior applies even if the request targets other open indices.
// API name: allow_no_indices
func (r *ValidateQuery) AllowNoIndices(allownoindices bool) *ValidateQuery {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// AllShards If `true`, the validation is executed on all shards instead of one random
// shard per index.
// API name: all_shards
func (r *ValidateQuery) AllShards(allshards bool) *ValidateQuery {
	r.values.Set("all_shards", strconv.FormatBool(allshards))

	return r
}

// Analyzer Analyzer to use for the query string.
// This parameter can only be used when the `q` query string parameter is
// specified.
// API name: analyzer
func (r *ValidateQuery) Analyzer(analyzer string) *ValidateQuery {
	r.values.Set("analyzer", analyzer)

	return r
}

// AnalyzeWildcard If `true`, wildcard and prefix queries are analyzed.
// API name: analyze_wildcard
func (r *ValidateQuery) AnalyzeWildcard(analyzewildcard bool) *ValidateQuery {
	r.values.Set("analyze_wildcard", strconv.FormatBool(analyzewildcard))

	return r
}

// DefaultOperator The default operator for query string query: `AND` or `OR`.
// API name: default_operator
func (r *ValidateQuery) DefaultOperator(defaultoperator operator.Operator) *ValidateQuery {
	r.values.Set("default_operator", defaultoperator.String())

	return r
}

// Df Field to use as default where no field prefix is given in the query string.
// This parameter can only be used when the `q` query string parameter is
// specified.
// API name: df
func (r *ValidateQuery) Df(df string) *ValidateQuery {
	r.values.Set("df", df)

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *ValidateQuery) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ValidateQuery {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// Explain If `true`, the response returns detailed information if an error has
// occurred.
// API name: explain
func (r *ValidateQuery) Explain(explain bool) *ValidateQuery {
	r.values.Set("explain", strconv.FormatBool(explain))

	return r
}

// IgnoreUnavailable If `false`, the request returns an error if it targets a missing or closed
// index.
// API name: ignore_unavailable
func (r *ValidateQuery) IgnoreUnavailable(ignoreunavailable bool) *ValidateQuery {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Lenient If `true`, format-based query failures (such as providing text to a numeric
// field) in the query string will be ignored.
// API name: lenient
func (r *ValidateQuery) Lenient(lenient bool) *ValidateQuery {
	r.values.Set("lenient", strconv.FormatBool(lenient))

	return r
}

// Rewrite If `true`, returns a more detailed explanation showing the actual Lucene
// query that will be executed.
// API name: rewrite
func (r *ValidateQuery) Rewrite(rewrite bool) *ValidateQuery {
	r.values.Set("rewrite", strconv.FormatBool(rewrite))

	return r
}

// Q Query in the Lucene query string syntax.
// API name: q
func (r *ValidateQuery) Q(q string) *ValidateQuery {
	r.values.Set("q", q)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ValidateQuery) ErrorTrace(errortrace bool) *ValidateQuery {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ValidateQuery) FilterPath(filterpaths ...string) *ValidateQuery {
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
func (r *ValidateQuery) Human(human bool) *ValidateQuery {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ValidateQuery) Pretty(pretty bool) *ValidateQuery {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Query Query in the Lucene query string syntax.
// API name: query
func (r *ValidateQuery) Query(query *types.Query) *ValidateQuery {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = query

	return r
}
