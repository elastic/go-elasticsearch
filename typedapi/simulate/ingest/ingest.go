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

// Simulate data ingestion.
// Run ingest pipelines against a set of provided documents, optionally with
// substitute pipeline definitions, to simulate ingesting data into an index.
//
// This API is meant to be used for troubleshooting or pipeline development, as
// it does not actually index any data into Elasticsearch.
//
// The API runs the default and final pipeline for that index against a set of
// documents provided in the body of the request.
// If a pipeline contains a reroute processor, it follows that reroute processor
// to the new index, running that index's pipelines as well the same way that a
// non-simulated ingest would.
// No data is indexed into Elasticsearch.
// Instead, the transformed document is returned, along with the list of
// pipelines that have been run and the name of the index where the document
// would have been indexed if this were not a simulation.
// The transformed document is validated against the mappings that would apply
// to this index, and any validation error is reported in the result.
//
// This API differs from the simulate pipeline API in that you specify a single
// pipeline for that API, and it runs only that one pipeline.
// The simulate pipeline API is more useful for developing a single pipeline,
// while the simulate ingest API is more useful for troubleshooting the
// interaction of the various pipelines that get applied when ingesting into an
// index.
//
// By default, the pipeline definitions that are currently in the system are
// used.
// However, you can supply substitute pipeline definitions in the body of the
// request.
// These will be used in place of the pipeline definitions that are already in
// the system. This can be used to replace existing pipeline definitions or to
// create new ones. The pipeline substitutions are used only within this
// request.
package ingest

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

type Ingest struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewIngest type alias for index.
type NewIngest func() *Ingest

// NewIngestFunc returns a new instance of Ingest with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewIngestFunc(tp elastictransport.Interface) NewIngest {
	return func() *Ingest {
		n := New(tp)

		return n
	}
}

// Simulate data ingestion.
// Run ingest pipelines against a set of provided documents, optionally with
// substitute pipeline definitions, to simulate ingesting data into an index.
//
// This API is meant to be used for troubleshooting or pipeline development, as
// it does not actually index any data into Elasticsearch.
//
// The API runs the default and final pipeline for that index against a set of
// documents provided in the body of the request.
// If a pipeline contains a reroute processor, it follows that reroute processor
// to the new index, running that index's pipelines as well the same way that a
// non-simulated ingest would.
// No data is indexed into Elasticsearch.
// Instead, the transformed document is returned, along with the list of
// pipelines that have been run and the name of the index where the document
// would have been indexed if this were not a simulation.
// The transformed document is validated against the mappings that would apply
// to this index, and any validation error is reported in the result.
//
// This API differs from the simulate pipeline API in that you specify a single
// pipeline for that API, and it runs only that one pipeline.
// The simulate pipeline API is more useful for developing a single pipeline,
// while the simulate ingest API is more useful for troubleshooting the
// interaction of the various pipelines that get applied when ingesting into an
// index.
//
// By default, the pipeline definitions that are currently in the system are
// used.
// However, you can supply substitute pipeline definitions in the body of the
// request.
// These will be used in place of the pipeline definitions that are already in
// the system. This can be used to replace existing pipeline definitions or to
// create new ones. The pipeline substitutions are used only within this
// request.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/simulate-ingest-api.html
func New(tp elastictransport.Interface) *Ingest {
	r := &Ingest{
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
func (r *Ingest) Raw(raw io.Reader) *Ingest {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Ingest) Request(req *Request) *Ingest {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Ingest) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Ingest: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ingest")
		path.WriteString("/")
		path.WriteString("_simulate")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")
		path.WriteString("_ingest")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
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
func (r Ingest) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "simulate.ingest")
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
		instrument.BeforeRequest(req, "simulate.ingest")
		if reader := instrument.RecordRequestBody(ctx, "simulate.ingest", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "simulate.ingest")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Ingest query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a ingest.Response
func (r Ingest) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "simulate.ingest")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// Header set a key, value pair in the Ingest headers map.
func (r *Ingest) Header(key, value string) *Ingest {
	r.headers.Set(key, value)

	return r
}

// Index The index to simulate ingesting into.
// This value can be overridden by specifying an index on each document.
// If you specify this parameter in the request path, it is used for any
// documents that do not explicitly specify an index argument.
// API Name: index
func (r *Ingest) Index(index string) *Ingest {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Pipeline The pipeline to use as the default pipeline.
// This value can be used to override the default pipeline of the index.
// API name: pipeline
func (r *Ingest) Pipeline(pipelinename string) *Ingest {
	r.values.Set("pipeline", pipelinename)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Ingest) ErrorTrace(errortrace bool) *Ingest {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Ingest) FilterPath(filterpaths ...string) *Ingest {
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
func (r *Ingest) Human(human bool) *Ingest {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Ingest) Pretty(pretty bool) *Ingest {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// ComponentTemplateSubstitutions A map of component template names to substitute component template definition
// objects.
// API name: component_template_substitutions
func (r *Ingest) ComponentTemplateSubstitutions(componenttemplatesubstitutions map[string]types.ComponentTemplateNode) *Ingest {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ComponentTemplateSubstitutions = componenttemplatesubstitutions

	return r
}

// Docs Sample documents to test in the pipeline.
// API name: docs
func (r *Ingest) Docs(docs ...types.Document) *Ingest {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Docs = docs

	return r
}

// IndexTemplateSubstitutions A map of index template names to substitute index template definition
// objects.
// API name: index_template_substitutions
func (r *Ingest) IndexTemplateSubstitutions(indextemplatesubstitutions map[string]types.IndexTemplate) *Ingest {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexTemplateSubstitutions = indextemplatesubstitutions

	return r
}

// API name: mapping_addition
func (r *Ingest) MappingAddition(mappingaddition *types.TypeMapping) *Ingest {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.MappingAddition = mappingaddition

	return r
}

// PipelineSubstitutions Pipelines to test.
// If you don’t specify the `pipeline` request path parameter, this parameter is
// required.
// If you specify both this and the request path parameter, the API only uses
// the request path parameter.
// API name: pipeline_substitutions
func (r *Ingest) PipelineSubstitutions(pipelinesubstitutions map[string]types.IngestPipeline) *Ingest {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.PipelineSubstitutions = pipelinesubstitutions

	return r
}
