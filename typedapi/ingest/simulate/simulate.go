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

// Allows to simulate a pipeline with example documents.
package simulate

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
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Simulate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	id string
}

// NewSimulate type alias for index.
type NewSimulate func() *Simulate

// NewSimulateFunc returns a new instance of Simulate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSimulateFunc(tp elastictransport.Interface) NewSimulate {
	return func() *Simulate {
		n := New(tp)

		return n
	}
}

// Allows to simulate a pipeline with example documents.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/simulate-pipeline-api.html
func New(tp elastictransport.Interface) *Simulate {
	r := &Simulate{
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
func (r *Simulate) Raw(raw io.Reader) *Simulate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Simulate) Request(req *Request) *Simulate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Simulate) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Simulate: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ingest")
		path.WriteString("/")
		path.WriteString("pipeline")
		path.WriteString("/")
		path.WriteString("_simulate")

		method = http.MethodPost
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_ingest")
		path.WriteString("/")
		path.WriteString("pipeline")
		path.WriteString("/")

		path.WriteString(r.id)
		path.WriteString("/")
		path.WriteString("_simulate")

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
func (r Simulate) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Simulate query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a simulate.Response
func (r Simulate) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Simulate headers map.
func (r *Simulate) Header(key, value string) *Simulate {
	r.headers.Set(key, value)

	return r
}

// Id Pipeline to test.
// If you don’t specify a `pipeline` in the request body, this parameter is
// required.
// API Name: id
func (r *Simulate) Id(id string) *Simulate {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Verbose If `true`, the response includes output data for each processor in the
// executed pipeline.
// API name: verbose
func (r *Simulate) Verbose(verbose bool) *Simulate {
	r.values.Set("verbose", strconv.FormatBool(verbose))

	return r
}

// Docs Sample documents to test in the pipeline.
// API name: docs
func (r *Simulate) Docs(docs ...types.Document) *Simulate {
	r.req.Docs = docs

	return r
}

// Pipeline Pipeline to test.
// If you don’t specify the `pipeline` request path parameter, this parameter is
// required.
// If you specify both this and the request path parameter, the API only uses
// the request path parameter.
// API name: pipeline
func (r *Simulate) Pipeline(pipeline *types.IngestPipeline) *Simulate {

	r.req.Pipeline = pipeline

	return r
}
