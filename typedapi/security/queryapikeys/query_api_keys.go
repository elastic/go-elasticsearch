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

// Retrieves information for API keys using a subset of query DSL
package queryapikeys

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

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type QueryApiKeys struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int
}

// NewQueryApiKeys type alias for index.
type NewQueryApiKeys func() *QueryApiKeys

// NewQueryApiKeysFunc returns a new instance of QueryApiKeys with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewQueryApiKeysFunc(tp elastictransport.Interface) NewQueryApiKeys {
	return func() *QueryApiKeys {
		n := New(tp)

		return n
	}
}

// Retrieves information for API keys using a subset of query DSL
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-api-key.html
func New(tp elastictransport.Interface) *QueryApiKeys {
	r := &QueryApiKeys{
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
func (r *QueryApiKeys) Raw(raw io.Reader) *QueryApiKeys {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *QueryApiKeys) Request(req *Request) *QueryApiKeys {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *QueryApiKeys) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for QueryApiKeys: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("_query")
		path.WriteString("/")
		path.WriteString("api_key")

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
func (r QueryApiKeys) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the QueryApiKeys query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a queryapikeys.Response
func (r QueryApiKeys) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the QueryApiKeys headers map.
func (r *QueryApiKeys) Header(key, value string) *QueryApiKeys {
	r.headers.Set(key, value)

	return r
}

// WithLimitedBy Return the snapshot of the owner user's role descriptors associated with the
// API key.
// An API key's actual permission is the intersection of its assigned role
// descriptors and the owner user's role descriptors.
// API name: with_limited_by
func (r *QueryApiKeys) WithLimitedBy(withlimitedby bool) *QueryApiKeys {
	r.values.Set("with_limited_by", strconv.FormatBool(withlimitedby))

	return r
}

// From Starting document offset.
// By default, you cannot page through more than 10,000 hits using the from and
// size parameters.
// To page through more hits, use the `search_after` parameter.
// API name: from
func (r *QueryApiKeys) From(from int) *QueryApiKeys {
	r.req.From = &from

	return r
}

// Query A query to filter which API keys to return.
// The query supports a subset of query types, including `match_all`, `bool`,
// `term`, `terms`, `ids`, `prefix`, `wildcard`, and `range`.
// You can query all public information associated with an API key.
// API name: query
func (r *QueryApiKeys) Query(query *types.Query) *QueryApiKeys {

	r.req.Query = query

	return r
}

// SearchAfter Search after definition
// API name: search_after
func (r *QueryApiKeys) SearchAfter(sortresults ...types.FieldValue) *QueryApiKeys {
	r.req.SearchAfter = sortresults

	return r
}

// Size The number of hits to return.
// By default, you cannot page through more than 10,000 hits using the `from`
// and `size` parameters.
// To page through more hits, use the `search_after` parameter.
// API name: size
func (r *QueryApiKeys) Size(size int) *QueryApiKeys {
	r.req.Size = &size

	return r
}

// Sort Other than `id`, all public fields of an API key are eligible for sorting.
// In addition, sort can also be applied to the `_doc` field to sort by index
// order.
// API name: sort
func (r *QueryApiKeys) Sort(sorts ...types.SortCombinations) *QueryApiKeys {
	r.req.Sort = sorts

	return r
}
