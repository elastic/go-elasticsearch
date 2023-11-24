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

// The terms enum API  can be used to discover terms in the index that begin
// with the provided string. It is designed for low-latency look-ups used in
// auto-complete scenarios.
package termsenum

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

type TermsEnum struct {
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

// NewTermsEnum type alias for index.
type NewTermsEnum func(index string) *TermsEnum

// NewTermsEnumFunc returns a new instance of TermsEnum with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewTermsEnumFunc(tp elastictransport.Interface) NewTermsEnum {
	return func(index string) *TermsEnum {
		n := New(tp)

		n._index(index)

		return n
	}
}

// The terms enum API  can be used to discover terms in the index that begin
// with the provided string. It is designed for low-latency look-ups used in
// auto-complete scenarios.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-terms-enum.html
func New(tp elastictransport.Interface) *TermsEnum {
	r := &TermsEnum{
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
func (r *TermsEnum) Raw(raw io.Reader) *TermsEnum {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *TermsEnum) Request(req *Request) *TermsEnum {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *TermsEnum) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for TermsEnum: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_terms_enum")

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
func (r TermsEnum) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the TermsEnum query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a termsenum.Response
func (r TermsEnum) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the TermsEnum headers map.
func (r *TermsEnum) Header(key, value string) *TermsEnum {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and index aliases to search.
// Wildcard (*) expressions are supported.
// API Name: index
func (r *TermsEnum) _index(index string) *TermsEnum {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// CaseInsensitive When true the provided search string is matched against index terms without
// case sensitivity.
// API name: case_insensitive
func (r *TermsEnum) CaseInsensitive(caseinsensitive bool) *TermsEnum {
	r.req.CaseInsensitive = &caseinsensitive

	return r
}

// Field The string to match at the start of indexed terms. If not provided, all terms
// in the field are considered.
// API name: field
func (r *TermsEnum) Field(field string) *TermsEnum {
	r.req.Field = field

	return r
}

// IndexFilter Allows to filter an index shard if the provided query rewrites to match_none.
// API name: index_filter
func (r *TermsEnum) IndexFilter(indexfilter *types.Query) *TermsEnum {

	r.req.IndexFilter = indexfilter

	return r
}

// API name: search_after
func (r *TermsEnum) SearchAfter(searchafter string) *TermsEnum {

	r.req.SearchAfter = &searchafter

	return r
}

// Size How many matching terms to return.
// API name: size
func (r *TermsEnum) Size(size int) *TermsEnum {
	r.req.Size = &size

	return r
}

// String The string after which terms in the index should be returned. Allows for a
// form of pagination if the last result from one request is passed as the
// search_after parameter for a subsequent request.
// API name: string
func (r *TermsEnum) String(string string) *TermsEnum {

	r.req.String = &string

	return r
}

// Timeout The maximum length of time to spend collecting results. Defaults to "1s" (one
// second). If the timeout is exceeded the complete flag set to false in the
// response and the results may be partial or empty.
// API name: timeout
func (r *TermsEnum) Timeout(duration types.Duration) *TermsEnum {
	r.req.Timeout = duration

	return r
}
