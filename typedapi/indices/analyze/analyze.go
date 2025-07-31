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

// Get tokens from text analysis.
// The analyze API performs analysis on a text string and returns the resulting
// tokens.
//
// Generating excessive amount of tokens may cause a node to run out of memory.
// The `index.analyze.max_token_count` setting enables you to limit the number
// of tokens that can be produced.
// If more than this limit of tokens gets generated, an error occurs.
// The `_analyze` endpoint without a specified index will always use `10000` as
// its limit.
package analyze

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

type Analyze struct {
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

// NewAnalyze type alias for index.
type NewAnalyze func() *Analyze

// NewAnalyzeFunc returns a new instance of Analyze with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewAnalyzeFunc(tp elastictransport.Interface) NewAnalyze {
	return func() *Analyze {
		n := New(tp)

		return n
	}
}

// Get tokens from text analysis.
// The analyze API performs analysis on a text string and returns the resulting
// tokens.
//
// Generating excessive amount of tokens may cause a node to run out of memory.
// The `index.analyze.max_token_count` setting enables you to limit the number
// of tokens that can be produced.
// If more than this limit of tokens gets generated, an error occurs.
// The `_analyze` endpoint without a specified index will always use `10000` as
// its limit.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-analyze
func New(tp elastictransport.Interface) *Analyze {
	r := &Analyze{
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
func (r *Analyze) Raw(raw io.Reader) *Analyze {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Analyze) Request(req *Request) *Analyze {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Analyze) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Analyze: %w", err)
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
		path.WriteString("_analyze")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_analyze")

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
func (r Analyze) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.analyze")
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
		instrument.BeforeRequest(req, "indices.analyze")
		if reader := instrument.RecordRequestBody(ctx, "indices.analyze", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.analyze")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Analyze query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a analyze.Response
func (r Analyze) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.analyze")
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

// Header set a key, value pair in the Analyze headers map.
func (r *Analyze) Header(key, value string) *Analyze {
	r.headers.Set(key, value)

	return r
}

// Index Index used to derive the analyzer.
// If specified, the `analyzer` or field parameter overrides this value.
// If no index is specified or the index does not have a default analyzer, the
// analyze API uses the standard analyzer.
// API Name: index
func (r *Analyze) Index(index string) *Analyze {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Analyze) ErrorTrace(errortrace bool) *Analyze {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Analyze) FilterPath(filterpaths ...string) *Analyze {
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
func (r *Analyze) Human(human bool) *Analyze {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Analyze) Pretty(pretty bool) *Analyze {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Analyzer The name of the analyzer that should be applied to the provided `text`.
// This could be a built-in analyzer, or an analyzer that’s been configured in
// the index.
// API name: analyzer
func (r *Analyze) Analyzer(analyzer string) *Analyze {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Analyzer = &analyzer

	return r
}

// Attributes Array of token attributes used to filter the output of the `explain`
// parameter.
// API name: attributes
func (r *Analyze) Attributes(attributes ...string) *Analyze {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Attributes = attributes

	return r
}

// CharFilter Array of character filters used to preprocess characters before the
// tokenizer.
// API name: char_filter
func (r *Analyze) CharFilter(charfilters ...types.CharFilter) *Analyze {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.CharFilter = charfilters

	return r
}

// Explain If `true`, the response includes token attributes and additional details.
// API name: explain
func (r *Analyze) Explain(explain bool) *Analyze {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Explain = &explain

	return r
}

// Field Field used to derive the analyzer.
// To use this parameter, you must specify an index.
// If specified, the `analyzer` parameter overrides this value.
// API name: field
func (r *Analyze) Field(field string) *Analyze {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Field = &field

	return r
}

// Filter Array of token filters used to apply after the tokenizer.
// API name: filter
func (r *Analyze) Filter(filters ...types.TokenFilter) *Analyze {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Filter = filters

	return r
}

// Normalizer Normalizer to use to convert text into a single token.
// API name: normalizer
func (r *Analyze) Normalizer(normalizer string) *Analyze {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Normalizer = &normalizer

	return r
}

// Text Text to analyze.
// If an array of strings is provided, it is analyzed as a multi-value field.
// API name: text
func (r *Analyze) Text(texttoanalyzes ...string) *Analyze {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Text = texttoanalyzes

	return r
}

// Tokenizer Tokenizer to use to convert text into tokens.
// API name: tokenizer
func (r *Analyze) Tokenizer(tokenizer types.Tokenizer) *Analyze {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Tokenizer = tokenizer

	return r
}
