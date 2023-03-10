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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Returns the information about the capabilities of fields among multiple
// indices.
package fieldcaps

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

type FieldCaps struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int

	index string
}

// NewFieldCaps type alias for index.
type NewFieldCaps func() *FieldCaps

// NewFieldCapsFunc returns a new instance of FieldCaps with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewFieldCapsFunc(tp elastictransport.Interface) NewFieldCaps {
	return func() *FieldCaps {
		n := New(tp)

		return n
	}
}

// Returns the information about the capabilities of fields among multiple
// indices.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/search-field-caps.html
func New(tp elastictransport.Interface) *FieldCaps {
	r := &FieldCaps{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *FieldCaps) Raw(raw io.Reader) *FieldCaps {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *FieldCaps) Request(req *Request) *FieldCaps {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *FieldCaps) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for FieldCaps: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_field_caps")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_field_caps")

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
func (r FieldCaps) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the FieldCaps query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a fieldcaps.Response
func (r FieldCaps) Do(ctx context.Context) (*Response, error) {

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

	return nil, errorResponse
}

// Header set a key, value pair in the FieldCaps headers map.
func (r *FieldCaps) Header(key, value string) *FieldCaps {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and aliases used to limit the
// request. Supports wildcards (*). To target all data streams and indices, omit
// this parameter or use * or _all.
// API Name: index
func (r *FieldCaps) Index(v string) *FieldCaps {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// AllowNoIndices If false, the request returns an error if any wildcard expression, index
// alias,
// or `_all` value targets only missing or closed indices. This behavior applies
// even if the request targets other open indices. For example, a request
// targeting `foo*,bar*` returns an error if an index starts with foo but no
// index starts with bar.
// API name: allow_no_indices
func (r *FieldCaps) AllowNoIndices(b bool) *FieldCaps {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match. If the request can target
// data streams, this argument determines whether wildcard expressions match
// hidden data streams. Supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *FieldCaps) ExpandWildcards(v string) *FieldCaps {
	r.values.Set("expand_wildcards", v)

	return r
}

// Fields Comma-separated list of fields to retrieve capabilities for. Wildcard (`*`)
// expressions are supported.
// API name: fields
func (r *FieldCaps) Fields(v string) *FieldCaps {
	r.values.Set("fields", v)

	return r
}

// IgnoreUnavailable If `true`, missing or closed indices are not included in the response.
// API name: ignore_unavailable
func (r *FieldCaps) IgnoreUnavailable(b bool) *FieldCaps {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// IncludeUnmapped If true, unmapped fields are included in the response.
// API name: include_unmapped
func (r *FieldCaps) IncludeUnmapped(b bool) *FieldCaps {
	r.values.Set("include_unmapped", strconv.FormatBool(b))

	return r
}

// Filters An optional set of filters: can include
// +metadata,-metadata,-nested,-multifield,-parent
// API name: filters
func (r *FieldCaps) Filters(v string) *FieldCaps {
	r.values.Set("filters", v)

	return r
}

// Types Only return results for fields that have one of the types in the list
// API name: types
func (r *FieldCaps) Types(v string) *FieldCaps {
	r.values.Set("types", v)

	return r
}
