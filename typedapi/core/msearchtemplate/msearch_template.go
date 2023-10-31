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

// Allows to execute several search template operations in one request.
package msearchtemplate

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/searchtype"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MsearchTemplate struct {
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

// NewMsearchTemplate type alias for index.
type NewMsearchTemplate func() *MsearchTemplate

// NewMsearchTemplateFunc returns a new instance of MsearchTemplate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMsearchTemplateFunc(tp elastictransport.Interface) NewMsearchTemplate {
	return func() *MsearchTemplate {
		n := New(tp)

		return n
	}
}

// Allows to execute several search template operations in one request.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
func New(tp elastictransport.Interface) *MsearchTemplate {
	r := &MsearchTemplate{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *MsearchTemplate) Raw(raw io.Reader) *MsearchTemplate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *MsearchTemplate) Request(req *Request) *MsearchTemplate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *MsearchTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
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

		for _, elem := range *r.req {
			data, err := json.Marshal(elem)
			if err != nil {
				return nil, err
			}
			r.buf.Write(data)
			r.buf.Write([]byte("\n"))
		}

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for MsearchTemplate: %w", err)
		}

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_msearch")
		path.WriteString("/")
		path.WriteString("template")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_msearch")
		path.WriteString("/")
		path.WriteString("template")

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
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+x-ndjson;compatible-with=8")
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
func (r MsearchTemplate) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the MsearchTemplate query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a msearchtemplate.Response
func (r MsearchTemplate) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	r.TypedKeys(true)

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

// Header set a key, value pair in the MsearchTemplate headers map.
func (r *MsearchTemplate) Header(key, value string) *MsearchTemplate {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and aliases to search.
// Supports wildcards (`*`).
// To search all data streams and indices, omit this parameter or use `*`.
// API Name: index
func (r *MsearchTemplate) Index(index string) *MsearchTemplate {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// CcsMinimizeRoundtrips If `true`, network round-trips are minimized for cross-cluster search
// requests.
// API name: ccs_minimize_roundtrips
func (r *MsearchTemplate) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *MsearchTemplate {
	r.values.Set("ccs_minimize_roundtrips", strconv.FormatBool(ccsminimizeroundtrips))

	return r
}

// MaxConcurrentSearches Maximum number of concurrent searches the API can run.
// API name: max_concurrent_searches
func (r *MsearchTemplate) MaxConcurrentSearches(maxconcurrentsearches string) *MsearchTemplate {
	r.values.Set("max_concurrent_searches", maxconcurrentsearches)

	return r
}

// SearchType The type of the search operation.
// Available options: `query_then_fetch`, `dfs_query_then_fetch`.
// API name: search_type
func (r *MsearchTemplate) SearchType(searchtype searchtype.SearchType) *MsearchTemplate {
	r.values.Set("search_type", searchtype.String())

	return r
}

// RestTotalHitsAsInt If `true`, the response returns `hits.total` as an integer.
// If `false`, it returns `hits.total` as an object.
// API name: rest_total_hits_as_int
func (r *MsearchTemplate) RestTotalHitsAsInt(resttotalhitsasint bool) *MsearchTemplate {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(resttotalhitsasint))

	return r
}

// TypedKeys If `true`, the response prefixes aggregation and suggester names with their
// respective types.
// API name: typed_keys
func (r *MsearchTemplate) TypedKeys(typedkeys bool) *MsearchTemplate {
	r.values.Set("typed_keys", strconv.FormatBool(typedkeys))

	return r
}
