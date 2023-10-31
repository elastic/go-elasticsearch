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

// Creates or updates a pipeline.
package putpipeline

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
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutPipeline struct {
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

// NewPutPipeline type alias for index.
type NewPutPipeline func(id string) *PutPipeline

// NewPutPipelineFunc returns a new instance of PutPipeline with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutPipelineFunc(tp elastictransport.Interface) NewPutPipeline {
	return func(id string) *PutPipeline {
		n := New(tp)

		n._id(id)

		return n
	}
}

// Creates or updates a pipeline.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/ingest.html
func New(tp elastictransport.Interface) *PutPipeline {
	r := &PutPipeline{
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
func (r *PutPipeline) Raw(raw io.Reader) *PutPipeline {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutPipeline) Request(req *Request) *PutPipeline {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutPipeline) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutPipeline: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_ingest")
		path.WriteString("/")
		path.WriteString("pipeline")
		path.WriteString("/")

		path.WriteString(r.id)

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
func (r PutPipeline) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutPipeline query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putpipeline.Response
func (r PutPipeline) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutPipeline headers map.
func (r *PutPipeline) Header(key, value string) *PutPipeline {
	r.headers.Set(key, value)

	return r
}

// Id ID of the ingest pipeline to create or update.
// API Name: id
func (r *PutPipeline) _id(id string) *PutPipeline {
	r.paramSet |= idMask
	r.id = id

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *PutPipeline) MasterTimeout(duration string) *PutPipeline {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and returns an error.
// API name: timeout
func (r *PutPipeline) Timeout(duration string) *PutPipeline {
	r.values.Set("timeout", duration)

	return r
}

// IfVersion Required version for optimistic concurrency control for pipeline updates
// API name: if_version
func (r *PutPipeline) IfVersion(versionnumber string) *PutPipeline {
	r.values.Set("if_version", versionnumber)

	return r
}

// Description Description of the ingest pipeline.
// API name: description
func (r *PutPipeline) Description(description string) *PutPipeline {

	r.req.Description = &description

	return r
}

// Meta_ Optional metadata about the ingest pipeline. May have any contents. This map
// is not automatically generated by Elasticsearch.
// API name: _meta
func (r *PutPipeline) Meta_(metadata types.Metadata) *PutPipeline {
	r.req.Meta_ = metadata

	return r
}

// OnFailure Processors to run immediately after a processor failure. Each processor
// supports a processor-level `on_failure` value. If a processor without an
// `on_failure` value fails, Elasticsearch uses this pipeline-level parameter as
// a fallback. The processors in this parameter run sequentially in the order
// specified. Elasticsearch will not attempt to run the pipeline's remaining
// processors.
// API name: on_failure
func (r *PutPipeline) OnFailure(onfailures ...types.ProcessorContainer) *PutPipeline {
	r.req.OnFailure = onfailures

	return r
}

// Processors Processors used to perform transformations on documents before indexing.
// Processors run sequentially in the order specified.
// API name: processors
func (r *PutPipeline) Processors(processors ...types.ProcessorContainer) *PutPipeline {
	r.req.Processors = processors

	return r
}

// Version Version number used by external systems to track ingest pipelines. This
// parameter is intended for external systems only. Elasticsearch does not use
// or validate pipeline version numbers.
// API name: version
func (r *PutPipeline) Version(versionnumber int64) *PutPipeline {
	r.req.Version = &versionnumber

	return r
}
