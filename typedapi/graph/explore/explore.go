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
// https://github.com/elastic/elasticsearch-specification/tree/e0ea3dc890d394d682096cc862b3bd879d9422e9


// Explore extracted and summarized information about the documents and terms in
// an index.
package explore

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
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

	req *Request
	raw json.RawMessage

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

		n.Index(index)

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
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Explore) Raw(raw json.RawMessage) *Explore {
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

	if r.raw != nil {
		r.buf.Write(r.raw)
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
		path.WriteString(url.PathEscape(r.index))
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

	if r.buf.Len() > 0 {
		req.Header.Set("content-type", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	req.Header.Set("accept", "application/vnd.elasticsearch+json;compatible-with=8")

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r Explore) Do(ctx context.Context) (*http.Response, error) {
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

// Header set a key, value pair in the Explore headers map.
func (r *Explore) Header(key, value string) *Explore {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names to search; use `_all` or empty string
// to perform the operation on all indices
// API Name: index
func (r *Explore) Index(v string) *Explore {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// Routing Specific routing value
// API name: routing
func (r *Explore) Routing(value string) *Explore {
	r.values.Set("routing", value)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *Explore) Timeout(value string) *Explore {
	r.values.Set("timeout", value)

	return r
}
