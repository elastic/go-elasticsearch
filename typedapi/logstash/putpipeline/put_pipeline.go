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

// Adds and updates Logstash Pipelines used for Central Management
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

// Adds and updates Logstash Pipelines used for Central Management
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-put-pipeline.html
func New(tp elastictransport.Interface) *PutPipeline {
	r := &PutPipeline{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
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
		path.WriteString("_logstash")
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

// Header set a key, value pair in the PutPipeline headers map.
func (r *PutPipeline) Header(key, value string) *PutPipeline {
	r.headers.Set(key, value)

	return r
}

// Id Identifier for the pipeline.
// API Name: id
func (r *PutPipeline) _id(id string) *PutPipeline {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Description Description of the pipeline.
// This description is not used by Elasticsearch or Logstash.
// API name: description
func (r *PutPipeline) Description(description string) *PutPipeline {

	r.req.Description = description

	return r
}

// LastModified Date the pipeline was last updated.
// Must be in the `yyyy-MM-dd'T'HH:mm:ss.SSSZZ` strict_date_time format.
// API name: last_modified
func (r *PutPipeline) LastModified(datetime types.DateTime) *PutPipeline {
	r.req.LastModified = datetime

	return r
}

// Pipeline Configuration for the pipeline.
// API name: pipeline
func (r *PutPipeline) Pipeline(pipeline string) *PutPipeline {

	r.req.Pipeline = pipeline

	return r
}

// PipelineMetadata Optional metadata about the pipeline.
// May have any contents.
// This metadata is not generated or used by Elasticsearch or Logstash.
// API name: pipeline_metadata
func (r *PutPipeline) PipelineMetadata(pipelinemetadata *types.PipelineMetadata) *PutPipeline {

	r.req.PipelineMetadata = *pipelinemetadata

	return r
}

// PipelineSettings Settings for the pipeline.
// Supports only flat keys in dot notation.
// API name: pipeline_settings
func (r *PutPipeline) PipelineSettings(pipelinesettings *types.PipelineSettings) *PutPipeline {

	r.req.PipelineSettings = *pipelinesettings

	return r
}

// Username User who last updated the pipeline.
// API name: username
func (r *PutPipeline) Username(username string) *PutPipeline {

	r.req.Username = username

	return r
}
