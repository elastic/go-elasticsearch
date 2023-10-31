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

// Returns information about why a specific matches (or doesn't match) a query.
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	id    string
	index string
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

// Returns information about why a specific matches (or doesn't match) a query.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/search-explain.html
func New(tp elastictransport.Interface) *Explain {
	r := &Explain{
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Explain: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_explain")
		path.WriteString("/")

		path.WriteString(r.id)

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
func (r Explain) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Explain query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a explain.Response
func (r Explain) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Explain headers map.
func (r *Explain) Header(key, value string) *Explain {
	r.headers.Set(key, value)

	return r
}

// Id Defines the document ID.
// API Name: id
func (r *Explain) _id(id string) *Explain {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index Index names used to limit the request.
// Only a single index name can be provided to this parameter.
// API Name: index
func (r *Explain) _index(index string) *Explain {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Analyzer Analyzer to use for the query string.
// This parameter can only be used when the `q` query string parameter is
// specified.
// API name: analyzer
func (r *Explain) Analyzer(analyzer string) *Explain {
	r.values.Set("analyzer", analyzer)

	return r
}

// AnalyzeWildcard If `true`, wildcard and prefix queries are analyzed.
// API name: analyze_wildcard
func (r *Explain) AnalyzeWildcard(analyzewildcard bool) *Explain {
	r.values.Set("analyze_wildcard", strconv.FormatBool(analyzewildcard))

	return r
}

// DefaultOperator The default operator for query string query: `AND` or `OR`.
// API name: default_operator
func (r *Explain) DefaultOperator(defaultoperator operator.Operator) *Explain {
	r.values.Set("default_operator", defaultoperator.String())

	return r
}

// Df Field to use as default where no field prefix is given in the query string.
// API name: df
func (r *Explain) Df(df string) *Explain {
	r.values.Set("df", df)

	return r
}

// Lenient If `true`, format-based query failures (such as providing text to a numeric
// field) in the query string will be ignored.
// API name: lenient
func (r *Explain) Lenient(lenient bool) *Explain {
	r.values.Set("lenient", strconv.FormatBool(lenient))

	return r
}

// Preference Specifies the node or shard the operation should be performed on.
// Random by default.
// API name: preference
func (r *Explain) Preference(preference string) *Explain {
	r.values.Set("preference", preference)

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *Explain) Routing(routing string) *Explain {
	r.values.Set("routing", routing)

	return r
}

// Source_ True or false to return the `_source` field or not, or a list of fields to
// return.
// API name: _source
func (r *Explain) Source_(sourceconfigparam string) *Explain {
	r.values.Set("_source", sourceconfigparam)

	return r
}

// SourceExcludes_ A comma-separated list of source fields to exclude from the response.
// API name: _source_excludes
func (r *Explain) SourceExcludes_(fields ...string) *Explain {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A comma-separated list of source fields to include in the response.
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

// Q Query in the Lucene query string syntax.
// API name: q
func (r *Explain) Q(q string) *Explain {
	r.values.Set("q", q)

	return r
}

// Query Defines the search definition using the Query DSL.
// API name: query
func (r *Explain) Query(query *types.Query) *Explain {

	r.req.Query = query

	return r
}
