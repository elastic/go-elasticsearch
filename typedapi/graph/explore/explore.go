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

// Explore extracted and summarized information about the documents and terms in
// an index.
package explore

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

type Explore struct {
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

// NewExplore type alias for index.
type NewExplore func(index string) *Explore

// NewExploreFunc returns a new instance of Explore with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExploreFunc(tp elastictransport.Interface) NewExplore {
	return func(index string) *Explore {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Explore extracted and summarized information about the documents and terms in
// an index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/graph-explore-api.html
func New(tp elastictransport.Interface) *Explore {
	r := &Explore{
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
func (r *Explore) Raw(raw io.Reader) *Explore {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Explore) Request(req *Request) *Explore {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Explore) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Explore: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_graph")
		path.WriteString("/")
		path.WriteString("explore")

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
func (r Explore) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Explore query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a explore.Response
func (r Explore) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Explore headers map.
func (r *Explore) Header(key, value string) *Explore {
	r.headers.Set(key, value)

	return r
}

// Index Name of the index.
// API Name: index
func (r *Explore) _index(index string) *Explore {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *Explore) Routing(routing string) *Explore {
	r.values.Set("routing", routing)

	return r
}

// Timeout Specifies the period of time to wait for a response from each shard.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// Defaults to no timeout.
// API name: timeout
func (r *Explore) Timeout(duration string) *Explore {
	r.values.Set("timeout", duration)

	return r
}

// Connections Specifies or more fields from which you want to extract terms that are
// associated with the specified vertices.
// API name: connections
func (r *Explore) Connections(connections *types.Hop) *Explore {

	r.req.Connections = connections

	return r
}

// Controls Direct the Graph API how to build the graph.
// API name: controls
func (r *Explore) Controls(controls *types.ExploreControls) *Explore {

	r.req.Controls = controls

	return r
}

// Query A seed query that identifies the documents of interest. Can be any valid
// Elasticsearch query.
// API name: query
func (r *Explore) Query(query *types.Query) *Explore {

	r.req.Query = query

	return r
}

// Vertices Specifies one or more fields that contain the terms you want to include in
// the graph as vertices.
// API name: vertices
func (r *Explore) Vertices(vertices ...types.VertexDefinition) *Explore {
	r.req.Vertices = vertices

	return r
}
