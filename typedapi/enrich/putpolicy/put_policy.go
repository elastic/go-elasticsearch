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

// Creates a new enrich policy.
package putpolicy

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
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutPolicy struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	name string
}

// NewPutPolicy type alias for index.
type NewPutPolicy func(name string) *PutPolicy

// NewPutPolicyFunc returns a new instance of PutPolicy with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutPolicyFunc(tp elastictransport.Interface) NewPutPolicy {
	return func(name string) *PutPolicy {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Creates a new enrich policy.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/put-enrich-policy-api.html
func New(tp elastictransport.Interface) *PutPolicy {
	r := &PutPolicy{
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
func (r *PutPolicy) Raw(raw io.Reader) *PutPolicy {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutPolicy) Request(req *Request) *PutPolicy {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutPolicy) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutPolicy: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_enrich")
		path.WriteString("/")
		path.WriteString("policy")
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
func (r PutPolicy) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutPolicy query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putpolicy.Response
func (r PutPolicy) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutPolicy headers map.
func (r *PutPolicy) Header(key, value string) *PutPolicy {
	r.headers.Set(key, value)

	return r
}

// Name Name of the enrich policy to create or update.
// API Name: name
func (r *PutPolicy) _name(name string) *PutPolicy {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// GeoMatch Matches enrich data to incoming documents based on a `geo_shape` query.
// API name: geo_match
func (r *PutPolicy) GeoMatch(geomatch *types.EnrichPolicy) *PutPolicy {

	r.req.GeoMatch = geomatch

	return r
}

// Match Matches enrich data to incoming documents based on a `term` query.
// API name: match
func (r *PutPolicy) Match(match *types.EnrichPolicy) *PutPolicy {

	r.req.Match = match

	return r
}

// Range Matches a number, date, or IP address in incoming documents to a range in the
// enrich index based on a `term` query.
// API name: range
func (r *PutPolicy) Range(range_ *types.EnrichPolicy) *PutPolicy {

	r.req.Range = range_

	return r
}
