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

// Updates the data lifecycle of the selected data streams.
package putdatalifecycle

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

type PutDataLifecycle struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int

	name string
}

// NewPutDataLifecycle type alias for index.
type NewPutDataLifecycle func(name string) *PutDataLifecycle

// NewPutDataLifecycleFunc returns a new instance of PutDataLifecycle with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutDataLifecycleFunc(tp elastictransport.Interface) NewPutDataLifecycle {
	return func(name string) *PutDataLifecycle {
		n := New(tp)

		n.Name(name)

		return n
	}
}

// Updates the data lifecycle of the selected data streams.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/dlm-put-lifecycle.html
func New(tp elastictransport.Interface) *PutDataLifecycle {
	r := &PutDataLifecycle{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutDataLifecycle) Raw(raw io.Reader) *PutDataLifecycle {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutDataLifecycle) Request(req *Request) *PutDataLifecycle {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutDataLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutDataLifecycle: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_data_stream")
		path.WriteString("/")

		path.WriteString(r.name)
		path.WriteString("/")
		path.WriteString("_lifecycle")

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r PutDataLifecycle) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutDataLifecycle query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putdatalifecycle.Response
func (r PutDataLifecycle) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutDataLifecycle headers map.
func (r *PutDataLifecycle) Header(key, value string) *PutDataLifecycle {
	r.headers.Set(key, value)

	return r
}

// Name A comma-separated list of data streams whose lifecycle will be updated; use
// `*` to set the lifecycle to all data streams
// API Name: name
func (r *PutDataLifecycle) Name(v string) *PutDataLifecycle {
	r.paramSet |= nameMask
	r.name = v

	return r
}

// ExpandWildcards Whether wildcard expressions should get expanded to open or closed indices
// (default: open)
// API name: expand_wildcards
func (r *PutDataLifecycle) ExpandWildcards(v string) *PutDataLifecycle {
	r.values.Set("expand_wildcards", v)

	return r
}

// MasterTimeout Specify timeout for connection to master
// API name: master_timeout
func (r *PutDataLifecycle) MasterTimeout(v string) *PutDataLifecycle {
	r.values.Set("master_timeout", v)

	return r
}

// Timeout Explicit timestamp for the document
// API name: timeout
func (r *PutDataLifecycle) Timeout(v string) *PutDataLifecycle {
	r.values.Set("timeout", v)

	return r
}
