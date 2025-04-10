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
// Code generated from specification version 9.0.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newTextStructureFindFieldStructureFunc(t Transport) TextStructureFindFieldStructure {
	return func(index string, field string, o ...func(*TextStructureFindFieldStructureRequest)) (*Response, error) {
		var r = TextStructureFindFieldStructureRequest{Field: field, Index: index}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// TextStructureFindFieldStructure - Finds the structure of a text field in an index.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/find-field-structure.html.
type TextStructureFindFieldStructure func(index string, field string, o ...func(*TextStructureFindFieldStructureRequest)) (*Response, error)

// TextStructureFindFieldStructureRequest configures the Text Structure Find Field Structure API request.
type TextStructureFindFieldStructureRequest struct {
	ColumnNames       []string
	Delimiter         string
	DocumentsToSample *int
	EcsCompatibility  string
	Explain           *bool
	Field             string
	Format            string
	GrokPattern       string
	Index             string
	Quote             string
	ShouldTrimFields  *bool
	Timeout           time.Duration
	TimestampField    string
	TimestampFormat   string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r TextStructureFindFieldStructureRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "text_structure.find_field_structure")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + len("/_text_structure/find_field_structure"))
	path.WriteString("http://")
	path.WriteString("/_text_structure/find_field_structure")

	params = make(map[string]string)

	if len(r.ColumnNames) > 0 {
		params["column_names"] = strings.Join(r.ColumnNames, ",")
	}

	if r.Delimiter != "" {
		params["delimiter"] = r.Delimiter
	}

	if r.DocumentsToSample != nil {
		params["documents_to_sample"] = strconv.FormatInt(int64(*r.DocumentsToSample), 10)
	}

	if r.EcsCompatibility != "" {
		params["ecs_compatibility"] = r.EcsCompatibility
	}

	if r.Explain != nil {
		params["explain"] = strconv.FormatBool(*r.Explain)
	}

	if r.Field != "" {
		params["field"] = r.Field
	}

	if r.Format != "" {
		params["format"] = r.Format
	}

	if r.GrokPattern != "" {
		params["grok_pattern"] = r.GrokPattern
	}

	if r.Index != "" {
		params["index"] = r.Index
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

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		if instrument, ok := r.instrument.(Instrumentation); ok {
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

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "text_structure.find_field_structure")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "text_structure.find_field_structure")
	}
	if err != nil {
		if instrument, ok := r.instrument.(Instrumentation); ok {
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
func (f TextStructureFindFieldStructure) WithContext(v context.Context) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.ctx = v
	}
}

// WithColumnNames - optional parameter containing a comma separated list of the column names for a delimited file.
func (f TextStructureFindFieldStructure) WithColumnNames(v ...string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.ColumnNames = v
	}
}

// WithDelimiter - optional parameter to specify the delimiter character for a delimited file - must be a single character.
func (f TextStructureFindFieldStructure) WithDelimiter(v string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.Delimiter = v
	}
}

// WithDocumentsToSample - how many documents should be included in the analysis.
func (f TextStructureFindFieldStructure) WithDocumentsToSample(v int) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.DocumentsToSample = &v
	}
}

// WithEcsCompatibility - optional parameter to specify the compatibility mode with ecs grok patterns - may be either 'v1' or 'disabled'.
func (f TextStructureFindFieldStructure) WithEcsCompatibility(v string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.EcsCompatibility = v
	}
}

// WithExplain - whether to include a commentary on how the structure was derived.
func (f TextStructureFindFieldStructure) WithExplain(v bool) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.Explain = &v
	}
}

// WithField - the field that should be analyzed.
func (f TextStructureFindFieldStructure) WithField(v string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.Field = v
	}
}

// WithFormat - optional parameter to specify the high level file format.
func (f TextStructureFindFieldStructure) WithFormat(v string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.Format = v
	}
}

// WithGrokPattern - optional parameter to specify the grok pattern that should be used to extract fields from messages in a semi-structured text file.
func (f TextStructureFindFieldStructure) WithGrokPattern(v string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.GrokPattern = v
	}
}

// WithIndex - the index containing the analyzed field.
func (f TextStructureFindFieldStructure) WithIndex(v string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.Index = v
	}
}

// WithQuote - optional parameter to specify the quote character for a delimited file - must be a single character.
func (f TextStructureFindFieldStructure) WithQuote(v string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.Quote = v
	}
}

// WithShouldTrimFields - optional parameter to specify whether the values between delimiters in a delimited file should have whitespace trimmed from them.
func (f TextStructureFindFieldStructure) WithShouldTrimFields(v bool) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.ShouldTrimFields = &v
	}
}

// WithTimeout - timeout after which the analysis will be aborted.
func (f TextStructureFindFieldStructure) WithTimeout(v time.Duration) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.Timeout = v
	}
}

// WithTimestampField - optional parameter to specify the timestamp field in the file.
func (f TextStructureFindFieldStructure) WithTimestampField(v string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.TimestampField = v
	}
}

// WithTimestampFormat - optional parameter to specify the timestamp format in the file - may be either a joda or java time format.
func (f TextStructureFindFieldStructure) WithTimestampFormat(v string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.TimestampFormat = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f TextStructureFindFieldStructure) WithPretty() func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f TextStructureFindFieldStructure) WithHuman() func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f TextStructureFindFieldStructure) WithErrorTrace() func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f TextStructureFindFieldStructure) WithFilterPath(v ...string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f TextStructureFindFieldStructure) WithHeader(h map[string]string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f TextStructureFindFieldStructure) WithOpaqueID(s string) func(*TextStructureFindFieldStructureRequest) {
	return func(r *TextStructureFindFieldStructureRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
