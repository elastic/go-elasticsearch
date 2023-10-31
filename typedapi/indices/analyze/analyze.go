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

// Performs the analysis process on a text and return the tokens breakdown of
// the text.
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	index string
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

// Performs the analysis process on a text and return the tokens breakdown of
// the text.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/indices-analyze.html
func New(tp elastictransport.Interface) *Analyze {
	r := &Analyze{
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Analyze: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_analyze")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

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
func (r Analyze) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Analyze query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a analyze.Response
func (r Analyze) Do(ctx context.Context) (*Response, error) {

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

// Analyzer The name of the analyzer that should be applied to the provided `text`.
// This could be a built-in analyzer, or an analyzer that’s been configured in
// the index.
// API name: analyzer
func (r *Analyze) Analyzer(analyzer string) *Analyze {

	r.req.Analyzer = &analyzer

	return r
}

// Attributes Array of token attributes used to filter the output of the `explain`
// parameter.
// API name: attributes
func (r *Analyze) Attributes(attributes ...string) *Analyze {
	r.req.Attributes = attributes

	return r
}

// CharFilter Array of character filters used to preprocess characters before the
// tokenizer.
// API name: char_filter
func (r *Analyze) CharFilter(charfilters ...types.CharFilter) *Analyze {
	r.req.CharFilter = charfilters

	return r
}

// Explain If `true`, the response includes token attributes and additional details.
// API name: explain
func (r *Analyze) Explain(explain bool) *Analyze {
	r.req.Explain = &explain

	return r
}

// Field Field used to derive the analyzer.
// To use this parameter, you must specify an index.
// If specified, the `analyzer` parameter overrides this value.
// API name: field
func (r *Analyze) Field(field string) *Analyze {
	r.req.Field = &field

	return r
}

// Filter Array of token filters used to apply after the tokenizer.
// API name: filter
func (r *Analyze) Filter(filters ...types.TokenFilter) *Analyze {
	r.req.Filter = filters

	return r
}

// Normalizer Normalizer to use to convert text into a single token.
// API name: normalizer
func (r *Analyze) Normalizer(normalizer string) *Analyze {

	r.req.Normalizer = &normalizer

	return r
}

// Text Text to analyze.
// If an array of strings is provided, it is analyzed as a multi-value field.
// API name: text
func (r *Analyze) Text(texttoanalyzes ...string) *Analyze {
	r.req.Text = texttoanalyzes

	return r
}

// Tokenizer Tokenizer to use to convert text into tokens.
// API name: tokenizer
func (r *Analyze) Tokenizer(tokenizer types.Tokenizer) *Analyze {
	r.req.Tokenizer = tokenizer

	return r
}
