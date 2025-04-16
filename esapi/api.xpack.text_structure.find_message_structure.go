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
//
// Code generated from specification version 8.19.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newTextStructureFindMessageStructureFunc(t Transport) TextStructureFindMessageStructure {
	return func(body io.Reader, o ...func(*TextStructureFindMessageStructureRequest)) (*Response, error) {
		var r = TextStructureFindMessageStructureRequest{Body: body}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.Instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// TextStructureFindMessageStructure - Finds the structure of a list of messages. The messages must contain data that is suitable to be ingested into Elasticsearch.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/find-message-structure.html.
type TextStructureFindMessageStructure func(body io.Reader, o ...func(*TextStructureFindMessageStructureRequest)) (*Response, error)

// TextStructureFindMessageStructureRequest configures the Text Structure Find Message Structure API request.
type TextStructureFindMessageStructureRequest struct {
	Body io.Reader

	ColumnNames      []string
	Delimiter        string
	EcsCompatibility string
	Explain          *bool
	Format           string
	GrokPattern      string
	Quote            string
	ShouldTrimFields *bool
	Timeout          time.Duration
	TimestampField   string
	TimestampFormat  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r TextStructureFindMessageStructureRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "text_structure.find_message_structure")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + len("/_text_structure/find_message_structure"))
	path.WriteString("http://")
	path.WriteString("/_text_structure/find_message_structure")

	params = make(map[string]string)

	if len(r.ColumnNames) > 0 {
		params["column_names"] = strings.Join(r.ColumnNames, ",")
	}

	if r.Delimiter != "" {
		params["delimiter"] = r.Delimiter
	}

	if r.EcsCompatibility != "" {
		params["ecs_compatibility"] = r.EcsCompatibility
	}

	if r.Explain != nil {
		params["explain"] = strconv.FormatBool(*r.Explain)
	}

	if r.Format != "" {
		params["format"] = r.Format
	}

	if r.GrokPattern != "" {
		params["grok_pattern"] = r.GrokPattern
	}

	if r.Quote != "" {
		params["quote"] = r.Quote
	}

	if r.ShouldTrimFields != nil {
		params["should_trim_fields"] = strconv.FormatBool(*r.ShouldTrimFields)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.TimestampField != "" {
		params["timestamp_field"] = r.TimestampField
	}

	if r.TimestampFormat != "" {
		params["timestamp_format"] = r.TimestampFormat
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "text_structure.find_message_structure")
		if reader := instrument.RecordRequestBody(ctx, "text_structure.find_message_structure", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "text_structure.find_message_structure")
	}
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
func (f TextStructureFindMessageStructure) WithContext(v context.Context) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.ctx = v
	}
}

// WithColumnNames - optional parameter containing a comma separated list of the column names for a delimited file.
func (f TextStructureFindMessageStructure) WithColumnNames(v ...string) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.ColumnNames = v
	}
}

// WithDelimiter - optional parameter to specify the delimiter character for a delimited file - must be a single character.
func (f TextStructureFindMessageStructure) WithDelimiter(v string) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.Delimiter = v
	}
}

// WithEcsCompatibility - optional parameter to specify the compatibility mode with ecs grok patterns - may be either 'v1' or 'disabled'.
func (f TextStructureFindMessageStructure) WithEcsCompatibility(v string) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.EcsCompatibility = v
	}
}

// WithExplain - whether to include a commentary on how the structure was derived.
func (f TextStructureFindMessageStructure) WithExplain(v bool) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.Explain = &v
	}
}

// WithFormat - optional parameter to specify the high level file format.
func (f TextStructureFindMessageStructure) WithFormat(v string) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.Format = v
	}
}

// WithGrokPattern - optional parameter to specify the grok pattern that should be used to extract fields from messages in a semi-structured text file.
func (f TextStructureFindMessageStructure) WithGrokPattern(v string) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.GrokPattern = v
	}
}

// WithQuote - optional parameter to specify the quote character for a delimited file - must be a single character.
func (f TextStructureFindMessageStructure) WithQuote(v string) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.Quote = v
	}
}

// WithShouldTrimFields - optional parameter to specify whether the values between delimiters in a delimited file should have whitespace trimmed from them.
func (f TextStructureFindMessageStructure) WithShouldTrimFields(v bool) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.ShouldTrimFields = &v
	}
}

// WithTimeout - timeout after which the analysis will be aborted.
func (f TextStructureFindMessageStructure) WithTimeout(v time.Duration) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.Timeout = v
	}
}

// WithTimestampField - optional parameter to specify the timestamp field in the file.
func (f TextStructureFindMessageStructure) WithTimestampField(v string) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.TimestampField = v
	}
}

// WithTimestampFormat - optional parameter to specify the timestamp format in the file - may be either a joda or java time format.
func (f TextStructureFindMessageStructure) WithTimestampFormat(v string) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.TimestampFormat = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f TextStructureFindMessageStructure) WithPretty() func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f TextStructureFindMessageStructure) WithHuman() func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f TextStructureFindMessageStructure) WithErrorTrace() func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f TextStructureFindMessageStructure) WithFilterPath(v ...string) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f TextStructureFindMessageStructure) WithHeader(h map[string]string) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f TextStructureFindMessageStructure) WithOpaqueID(s string) func(*TextStructureFindMessageStructureRequest) {
	return func(r *TextStructureFindMessageStructureRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
