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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Create or update a Logstash pipeline.
//
// Create a pipeline that is used for Logstash Central Management.
// If the specified pipeline exists, it is replaced.
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

type PutPipeline struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
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

// Create or update a Logstash pipeline.
//
// Create a pipeline that is used for Logstash Central Management.
// If the specified pipeline exists, it is replaced.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-put-pipeline.html
func New(tp elastictransport.Interface) *PutPipeline {
	r := &PutPipeline{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutPipeline: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_logstash")
		path.WriteString("/")
		path.WriteString("pipeline")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
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
func (r PutPipeline) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "logstash.put_pipeline")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "logstash.put_pipeline")
		if reader := instrument.RecordRequestBody(ctx, "logstash.put_pipeline", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "logstash.put_pipeline")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutPipeline query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Header set a key, value pair in the PutPipeline headers map.
func (r *PutPipeline) Header(key, value string) *PutPipeline {
	r.headers.Set(key, value)

	return r
}

// Id An identifier for the pipeline.
// API Name: id
func (r *PutPipeline) _id(id string) *PutPipeline {
	r.paramSet |= idMask
	r.id = id

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutPipeline) ErrorTrace(errortrace bool) *PutPipeline {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutPipeline) FilterPath(filterpaths ...string) *PutPipeline {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *PutPipeline) Human(human bool) *PutPipeline {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutPipeline) Pretty(pretty bool) *PutPipeline {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Description A description of the pipeline.
// This description is not used by Elasticsearch or Logstash.
// API name: description
func (r *PutPipeline) Description(description string) *PutPipeline {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Description = description

	return r
}

// LastModified The date the pipeline was last updated.
// It must be in the `yyyy-MM-dd'T'HH:mm:ss.SSSZZ` strict_date_time format.
// API name: last_modified
func (r *PutPipeline) LastModified(datetime types.DateTime) *PutPipeline {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LastModified = datetime

	return r
}

// Pipeline The configuration for the pipeline.
// API name: pipeline
func (r *PutPipeline) Pipeline(pipeline string) *PutPipeline {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Pipeline = pipeline

	return r
}

// PipelineMetadata Optional metadata about the pipeline, which can have any contents.
// This metadata is not generated or used by Elasticsearch or Logstash.
// API name: pipeline_metadata
func (r *PutPipeline) PipelineMetadata(pipelinemetadata *types.PipelineMetadata) *PutPipeline {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.PipelineMetadata = *pipelinemetadata

	return r
}

// PipelineSettings Settings for the pipeline.
// It supports only flat keys in dot notation.
// API name: pipeline_settings
func (r *PutPipeline) PipelineSettings(pipelinesettings *types.PipelineSettings) *PutPipeline {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.PipelineSettings = *pipelinesettings

	return r
}

// Username The user who last updated the pipeline.
// API name: username
func (r *PutPipeline) Username(username string) *PutPipeline {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Username = username

	return r
}
