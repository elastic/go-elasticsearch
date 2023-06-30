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

	req *Request
	raw io.Reader

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

		n.Id(id)

		n.Index(index)

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

// Id The document ID
// API Name: id
func (r *Explain) Id(v string) *Explain {
	r.paramSet |= idMask
	r.id = v

	return r
}

// Index The name of the index
// API Name: index
func (r *Explain) Index(v string) *Explain {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// Analyzer The analyzer for the query string query
// API name: analyzer
func (r *Explain) Analyzer(v string) *Explain {
	r.values.Set("analyzer", v)

	return r
}

// AnalyzeWildcard Specify whether wildcards and prefix queries in the query string query should
// be analyzed (default: false)
// API name: analyze_wildcard
func (r *Explain) AnalyzeWildcard(b bool) *Explain {
	r.values.Set("analyze_wildcard", strconv.FormatBool(b))

	return r
}

// DefaultOperator The default operator for query string query (AND or OR)
// API name: default_operator
func (r *Explain) DefaultOperator(enum operator.Operator) *Explain {
	r.values.Set("default_operator", enum.String())

	return r
}

// Df The default field for query string query (default: _all)
// API name: df
func (r *Explain) Df(v string) *Explain {
	r.values.Set("df", v)

	return r
}

// Lenient Specify whether format-based query failures (such as providing text to a
// numeric field) should be ignored
// API name: lenient
func (r *Explain) Lenient(b bool) *Explain {
	r.values.Set("lenient", strconv.FormatBool(b))

	return r
}

// Preference Specify the node or shard the operation should be performed on (default:
// random)
// API name: preference
func (r *Explain) Preference(v string) *Explain {
	r.values.Set("preference", v)

	return r
}

// Routing Specific routing value
// API name: routing
func (r *Explain) Routing(v string) *Explain {
	r.values.Set("routing", v)

	return r
}

// Source_ True or false to return the _source field or not, or a list of fields to
// return
// API name: _source
func (r *Explain) Source_(v string) *Explain {
	r.values.Set("_source", v)

	return r
}

// SourceExcludes_ A list of fields to exclude from the returned _source field
// API name: _source_excludes
func (r *Explain) SourceExcludes_(v string) *Explain {
	r.values.Set("_source_excludes", v)

	return r
}

// SourceIncludes_ A list of fields to extract and return from the _source field
// API name: _source_includes
func (r *Explain) SourceIncludes_(v string) *Explain {
	r.values.Set("_source_includes", v)

	return r
}

// StoredFields A comma-separated list of stored fields to return in the response
// API name: stored_fields
func (r *Explain) StoredFields(v string) *Explain {
	r.values.Set("stored_fields", v)

	return r
}

// Q Query in the Lucene query string syntax
// API name: q
func (r *Explain) Q(v string) *Explain {
	r.values.Set("q", v)

	return r
}
