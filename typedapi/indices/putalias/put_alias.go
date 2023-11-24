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

// Creates or updates an alias.
package putalias

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

	nameMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAlias struct {
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
	name  string
}

// NewPutAlias type alias for index.
type NewPutAlias func(index, name string) *PutAlias

// NewPutAliasFunc returns a new instance of PutAlias with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutAliasFunc(tp elastictransport.Interface) NewPutAlias {
	return func(index, name string) *PutAlias {
		n := New(tp)

		n._index(index)

		n._name(name)

		return n
	}
}

// Creates or updates an alias.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
func New(tp elastictransport.Interface) *PutAlias {
	r := &PutAlias{
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
func (r *PutAlias) Raw(raw io.Reader) *PutAlias {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutAlias) Request(req *Request) *PutAlias {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutAlias) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutAlias: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|nameMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_alias")
		path.WriteString("/")

		path.WriteString(r.name)

		method = http.MethodPut
	case r.paramSet == indexMask|nameMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_aliases")
		path.WriteString("/")

		path.WriteString(r.name)

		method = http.MethodPut
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
func (r PutAlias) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutAlias query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putalias.Response
func (r PutAlias) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutAlias headers map.
func (r *PutAlias) Header(key, value string) *PutAlias {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams or indices to add.
// Supports wildcards (`*`).
// Wildcard patterns that match both data streams and indices return an error.
// API Name: index
func (r *PutAlias) _index(index string) *PutAlias {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Name Alias to update.
// If the alias doesn’t exist, the request creates it.
// Index alias names support date math.
// API Name: name
func (r *PutAlias) _name(name string) *PutAlias {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *PutAlias) MasterTimeout(duration string) *PutAlias {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *PutAlias) Timeout(duration string) *PutAlias {
	r.values.Set("timeout", duration)

	return r
}

// Filter Query used to limit documents the alias can access.
// API name: filter
func (r *PutAlias) Filter(filter *types.Query) *PutAlias {

	r.req.Filter = filter

	return r
}

// IndexRouting Value used to route indexing operations to a specific shard.
// If specified, this overwrites the `routing` value for indexing operations.
// Data stream aliases don’t support this parameter.
// API name: index_routing
func (r *PutAlias) IndexRouting(routing string) *PutAlias {
	r.req.IndexRouting = &routing

	return r
}

// IsWriteIndex If `true`, sets the write index or data stream for the alias.
// If an alias points to multiple indices or data streams and `is_write_index`
// isn’t set, the alias rejects write requests.
// If an index alias points to one index and `is_write_index` isn’t set, the
// index automatically acts as the write index.
// Data stream aliases don’t automatically set a write data stream, even if the
// alias points to one data stream.
// API name: is_write_index
func (r *PutAlias) IsWriteIndex(iswriteindex bool) *PutAlias {
	r.req.IsWriteIndex = &iswriteindex

	return r
}

// Routing Value used to route indexing and search operations to a specific shard.
// Data stream aliases don’t support this parameter.
// API name: routing
func (r *PutAlias) Routing(routing string) *PutAlias {
	r.req.Routing = &routing

	return r
}

// SearchRouting Value used to route search operations to a specific shard.
// If specified, this overwrites the `routing` value for search operations.
// Data stream aliases don’t support this parameter.
// API name: search_routing
func (r *PutAlias) SearchRouting(routing string) *PutAlias {
	r.req.SearchRouting = &routing

	return r
}
