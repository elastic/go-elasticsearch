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

// Get multiple term vectors.
//
// Get multiple term vectors with a single request.
// You can specify existing documents by index and ID or provide artificial
// documents in the body of the request.
// You can specify the index in the request body or request URI.
// The response contains a `docs` array with all the fetched termvectors.
// Each element has the structure provided by the termvectors API.
//
// **Artificial documents**
//
// You can also use `mtermvectors` to generate term vectors for artificial
// documents provided in the body of the request.
// The mapping used is determined by the specified `_index`.
package mtermvectors

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Mtermvectors struct {
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

// NewMtermvectors type alias for index.
type NewMtermvectors func() *Mtermvectors

// NewMtermvectorsFunc returns a new instance of Mtermvectors with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMtermvectorsFunc(tp elastictransport.Interface) NewMtermvectors {
	return func() *Mtermvectors {
		n := New(tp)

		return n
	}
}

// Get multiple term vectors.
//
// Get multiple term vectors with a single request.
// You can specify existing documents by index and ID or provide artificial
// documents in the body of the request.
// You can specify the index in the request body or request URI.
// The response contains a `docs` array with all the fetched termvectors.
// Each element has the structure provided by the termvectors API.
//
// **Artificial documents**
//
// You can also use `mtermvectors` to generate term vectors for artificial
// documents provided in the body of the request.
// The mapping used is determined by the specified `_index`.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
func New(tp elastictransport.Interface) *Mtermvectors {
	r := &Mtermvectors{
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
func (r *Mtermvectors) Raw(raw io.Reader) *Mtermvectors {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Mtermvectors) Request(req *Request) *Mtermvectors {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Mtermvectors) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Mtermvectors: %w", err)
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
		path.WriteString("_mtermvectors")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_mtermvectors")

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
func (r Mtermvectors) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "mtermvectors")
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
		instrument.BeforeRequest(req, "mtermvectors")
		if reader := instrument.RecordRequestBody(ctx, "mtermvectors", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "mtermvectors")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Mtermvectors query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a mtermvectors.Response
func (r Mtermvectors) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "mtermvectors")
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

// Header set a key, value pair in the Mtermvectors headers map.
func (r *Mtermvectors) Header(key, value string) *Mtermvectors {
	r.headers.Set(key, value)

	return r
}

// Index The name of the index that contains the documents.
// API Name: index
func (r *Mtermvectors) Index(index string) *Mtermvectors {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Fields A comma-separated list or wildcard expressions of fields to include in the
// statistics.
// It is used as the default list unless a specific field list is provided in
// the `completion_fields` or `fielddata_fields` parameters.
// API name: fields
func (r *Mtermvectors) Fields(fields ...string) *Mtermvectors {
	r.values.Set("fields", strings.Join(fields, ","))

	return r
}

// FieldStatistics If `true`, the response includes the document count, sum of document
// frequencies, and sum of total term frequencies.
// API name: field_statistics
func (r *Mtermvectors) FieldStatistics(fieldstatistics bool) *Mtermvectors {
	r.values.Set("field_statistics", strconv.FormatBool(fieldstatistics))

	return r
}

// Offsets If `true`, the response includes term offsets.
// API name: offsets
func (r *Mtermvectors) Offsets(offsets bool) *Mtermvectors {
	r.values.Set("offsets", strconv.FormatBool(offsets))

	return r
}

// Payloads If `true`, the response includes term payloads.
// API name: payloads
func (r *Mtermvectors) Payloads(payloads bool) *Mtermvectors {
	r.values.Set("payloads", strconv.FormatBool(payloads))

	return r
}

// Positions If `true`, the response includes term positions.
// API name: positions
func (r *Mtermvectors) Positions(positions bool) *Mtermvectors {
	r.values.Set("positions", strconv.FormatBool(positions))

	return r
}

// Preference The node or shard the operation should be performed on.
// It is random by default.
// API name: preference
func (r *Mtermvectors) Preference(preference string) *Mtermvectors {
	r.values.Set("preference", preference)

	return r
}

// Realtime If true, the request is real-time as opposed to near-real-time.
// API name: realtime
func (r *Mtermvectors) Realtime(realtime bool) *Mtermvectors {
	r.values.Set("realtime", strconv.FormatBool(realtime))

	return r
}

// Routing A custom value used to route operations to a specific shard.
// API name: routing
func (r *Mtermvectors) Routing(routing string) *Mtermvectors {
	r.values.Set("routing", routing)

	return r
}

// TermStatistics If true, the response includes term frequency and document frequency.
// API name: term_statistics
func (r *Mtermvectors) TermStatistics(termstatistics bool) *Mtermvectors {
	r.values.Set("term_statistics", strconv.FormatBool(termstatistics))

	return r
}

// Version If `true`, returns the document version as part of a hit.
// API name: version
func (r *Mtermvectors) Version(versionnumber string) *Mtermvectors {
	r.values.Set("version", versionnumber)

	return r
}

// VersionType The version type.
// API name: version_type
func (r *Mtermvectors) VersionType(versiontype versiontype.VersionType) *Mtermvectors {
	r.values.Set("version_type", versiontype.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Mtermvectors) ErrorTrace(errortrace bool) *Mtermvectors {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Mtermvectors) FilterPath(filterpaths ...string) *Mtermvectors {
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
func (r *Mtermvectors) Human(human bool) *Mtermvectors {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Mtermvectors) Pretty(pretty bool) *Mtermvectors {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Docs An array of existing or artificial documents.
// API name: docs
func (r *Mtermvectors) Docs(docs ...types.MTermVectorsOperation) *Mtermvectors {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Docs = docs

	return r
}

// Ids A simplified syntax to specify documents by their ID if they're in the same
// index.
// API name: ids
func (r *Mtermvectors) Ids(ids ...string) *Mtermvectors {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Ids = ids

	return r
}
