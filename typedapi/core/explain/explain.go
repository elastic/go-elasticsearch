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

// Explain a document match result.
// Get information about why a specific document matches, or doesn't match, a
// query.
// It computes a score explanation for a query and a specific document.
package explain

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Explain struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id    string
	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewExplain type alias for index.
type NewExplain func(index, id string) *Explain

// NewExplainFunc returns a new instance of Explain with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExplainFunc(tp elastictransport.Interface) NewExplain {
	return func(index, id string) *Explain {
		n := New(tp)

		n._id(id)

		n._index(index)

		return n
	}
}

// Explain a document match result.
// Get information about why a specific document matches, or doesn't match, a
// query.
// It computes a score explanation for a query and a specific document.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
func New(tp elastictransport.Interface) *Explain {
	r := &Explain{
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
func (r *Explain) Raw(raw io.Reader) *Explain {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Explain) Request(req *Request) *Explain {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Explain) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Explain: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_explain")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

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
func (r Explain) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "explain")
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
		instrument.BeforeRequest(req, "explain")
		if reader := instrument.RecordRequestBody(ctx, "explain", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "explain")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Explain query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a explain.Response
func (r Explain) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "explain")
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

// Header set a key, value pair in the Explain headers map.
func (r *Explain) Header(key, value string) *Explain {
	r.headers.Set(key, value)

	return r
}

// Id The document identifier.
// API Name: id
func (r *Explain) _id(id string) *Explain {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index Index names that are used to limit the request.
// Only a single index name can be provided to this parameter.
// API Name: index
func (r *Explain) _index(index string) *Explain {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Analyzer The analyzer to use for the query string.
// This parameter can be used only when the `q` query string parameter is
// specified.
// API name: analyzer
func (r *Explain) Analyzer(analyzer string) *Explain {
	r.values.Set("analyzer", analyzer)

	return r
}

// AnalyzeWildcard If `true`, wildcard and prefix queries are analyzed.
// This parameter can be used only when the `q` query string parameter is
// specified.
// API name: analyze_wildcard
func (r *Explain) AnalyzeWildcard(analyzewildcard bool) *Explain {
	r.values.Set("analyze_wildcard", strconv.FormatBool(analyzewildcard))

	return r
}

// DefaultOperator The default operator for query string query: `AND` or `OR`.
// This parameter can be used only when the `q` query string parameter is
// specified.
// API name: default_operator
func (r *Explain) DefaultOperator(defaultoperator operator.Operator) *Explain {
	r.values.Set("default_operator", defaultoperator.String())

	return r
}

// Df The field to use as default where no field prefix is given in the query
// string.
// This parameter can be used only when the `q` query string parameter is
// specified.
// API name: df
func (r *Explain) Df(df string) *Explain {
	r.values.Set("df", df)

	return r
}

// Lenient If `true`, format-based query failures (such as providing text to a numeric
// field) in the query string will be ignored.
// This parameter can be used only when the `q` query string parameter is
// specified.
// API name: lenient
func (r *Explain) Lenient(lenient bool) *Explain {
	r.values.Set("lenient", strconv.FormatBool(lenient))

	return r
}

// Preference The node or shard the operation should be performed on.
// It is random by default.
// API name: preference
func (r *Explain) Preference(preference string) *Explain {
	r.values.Set("preference", preference)

	return r
}

// Routing A custom value used to route operations to a specific shard.
// API name: routing
func (r *Explain) Routing(routing string) *Explain {
	r.values.Set("routing", routing)

	return r
}

// Source_ `True` or `false` to return the `_source` field or not or a list of fields to
// return.
// API name: _source
func (r *Explain) Source_(sourceconfigparam string) *Explain {
	r.values.Set("_source", sourceconfigparam)

	return r
}

// SourceExcludes_ A comma-separated list of source fields to exclude from the response.
// You can also use this parameter to exclude fields from the subset specified
// in `_source_includes` query parameter.
// If the `_source` parameter is `false`, this parameter is ignored.
// API name: _source_excludes
func (r *Explain) SourceExcludes_(fields ...string) *Explain {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A comma-separated list of source fields to include in the response.
// If this parameter is specified, only these source fields are returned.
// You can exclude fields from this subset using the `_source_excludes` query
// parameter.
// If the `_source` parameter is `false`, this parameter is ignored.
// API name: _source_includes
func (r *Explain) SourceIncludes_(fields ...string) *Explain {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// StoredFields A comma-separated list of stored fields to return in the response.
// API name: stored_fields
func (r *Explain) StoredFields(fields ...string) *Explain {
	r.values.Set("stored_fields", strings.Join(fields, ","))

	return r
}

// Q The query in the Lucene query string syntax.
// API name: q
func (r *Explain) Q(q string) *Explain {
	r.values.Set("q", q)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Explain) ErrorTrace(errortrace bool) *Explain {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Explain) FilterPath(filterpaths ...string) *Explain {
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
func (r *Explain) Human(human bool) *Explain {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Explain) Pretty(pretty bool) *Explain {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Query Defines the search definition using the Query DSL.
// API name: query
func (r *Explain) Query(query *types.Query) *Explain {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = query

	return r
}
