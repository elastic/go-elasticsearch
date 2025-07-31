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

// Find the structure of a text file.
// The text file must contain data that is suitable to be ingested into
// Elasticsearch.
//
// This API provides a starting point for ingesting data into Elasticsearch in a
// format that is suitable for subsequent use with other Elastic Stack
// functionality.
// Unlike other Elasticsearch endpoints, the data that is posted to this
// endpoint does not need to be UTF-8 encoded and in JSON format.
// It must, however, be text; binary text formats are not currently supported.
// The size is limited to the Elasticsearch HTTP receive buffer size, which
// defaults to 100 Mb.
//
// The response from the API contains:
//
// * A couple of messages from the beginning of the text.
// * Statistics that reveal the most common values for all fields detected
// within the text and basic numeric statistics for numeric fields.
// * Information about the structure of the text, which is useful when you write
// ingest configurations to index it or similarly formatted text.
// * Appropriate mappings for an Elasticsearch index, which you could use to
// ingest the text.
//
// All this information can be calculated by the structure finder with no
// guidance.
// However, you can optionally override some of the decisions about the text
// structure by specifying one or more query parameters.
package findstructure

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

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type FindStructure struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewFindStructure type alias for index.
type NewFindStructure func() *FindStructure

// NewFindStructureFunc returns a new instance of FindStructure with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewFindStructureFunc(tp elastictransport.Interface) NewFindStructure {
	return func() *FindStructure {
		n := New(tp)

		return n
	}
}

// Find the structure of a text file.
// The text file must contain data that is suitable to be ingested into
// Elasticsearch.
//
// This API provides a starting point for ingesting data into Elasticsearch in a
// format that is suitable for subsequent use with other Elastic Stack
// functionality.
// Unlike other Elasticsearch endpoints, the data that is posted to this
// endpoint does not need to be UTF-8 encoded and in JSON format.
// It must, however, be text; binary text formats are not currently supported.
// The size is limited to the Elasticsearch HTTP receive buffer size, which
// defaults to 100 Mb.
//
// The response from the API contains:
//
// * A couple of messages from the beginning of the text.
// * Statistics that reveal the most common values for all fields detected
// within the text and basic numeric statistics for numeric fields.
// * Information about the structure of the text, which is useful when you write
// ingest configurations to index it or similarly formatted text.
// * Appropriate mappings for an Elasticsearch index, which you could use to
// ingest the text.
//
// All this information can be calculated by the structure finder with no
// guidance.
// However, you can optionally override some of the decisions about the text
// structure by specifying one or more query parameters.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-structure.html
func New(tp elastictransport.Interface) *FindStructure {
	r := &FindStructure{
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
func (r *FindStructure) Raw(raw io.Reader) *FindStructure {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *FindStructure) Request(req *Request) *FindStructure {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *FindStructure) HttpRequest(ctx context.Context) (*http.Request, error) {
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

		for _, elem := range *r.req {
			data, err := json.Marshal(elem)
			if err != nil {
				return nil, err
			}
			r.buf.Write(data)
			r.buf.Write([]byte("\n"))
		}

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for FindStructure: %w", err)
		}

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_text_structure")
		path.WriteString("/")
		path.WriteString("find_structure")

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
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+x-ndjson;compatible-with=8")
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
func (r FindStructure) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "text_structure.find_structure")
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
		instrument.BeforeRequest(req, "text_structure.find_structure")
		if reader := instrument.RecordRequestBody(ctx, "text_structure.find_structure", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "text_structure.find_structure")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the FindStructure query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a findstructure.Response
func (r FindStructure) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "text_structure.find_structure")
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

// Header set a key, value pair in the FindStructure headers map.
func (r *FindStructure) Header(key, value string) *FindStructure {
	r.headers.Set(key, value)

	return r
}

// Charset The text's character set.
// It must be a character set that is supported by the JVM that Elasticsearch
// uses.
// For example, `UTF-8`, `UTF-16LE`, `windows-1252`, or `EUC-JP`.
// If this parameter is not specified, the structure finder chooses an
// appropriate character set.
// API name: charset
func (r *FindStructure) Charset(charset string) *FindStructure {
	r.values.Set("charset", charset)

	return r
}

// ColumnNames If you have set format to `delimited`, you can specify the column names in a
// comma-separated list.
// If this parameter is not specified, the structure finder uses the column
// names from the header row of the text.
// If the text does not have a header role, columns are named "column1",
// "column2", "column3", for example.
// API name: column_names
func (r *FindStructure) ColumnNames(columnnames string) *FindStructure {
	r.values.Set("column_names", columnnames)

	return r
}

// Delimiter If you have set `format` to `delimited`, you can specify the character used
// to delimit the values in each row.
// Only a single character is supported; the delimiter cannot have multiple
// characters.
// By default, the API considers the following possibilities: comma, tab,
// semi-colon, and pipe (`|`).
// In this default scenario, all rows must have the same number of fields for
// the delimited format to be detected.
// If you specify a delimiter, up to 10% of the rows can have a different number
// of columns than the first row.
// API name: delimiter
func (r *FindStructure) Delimiter(delimiter string) *FindStructure {
	r.values.Set("delimiter", delimiter)

	return r
}

// EcsCompatibility The mode of compatibility with ECS compliant Grok patterns.
// Use this parameter to specify whether to use ECS Grok patterns instead of
// legacy ones when the structure finder creates a Grok pattern.
// Valid values are `disabled` and `v1`.
// This setting primarily has an impact when a whole message Grok pattern such
// as `%{CATALINALOG}` matches the input.
// If the structure finder identifies a common structure but has no idea of
// meaning then generic field names such as `path`, `ipaddress`, `field1`, and
// `field2` are used in the `grok_pattern` output, with the intention that a
// user who knows the meanings rename these fields before using it.
// API name: ecs_compatibility
func (r *FindStructure) EcsCompatibility(ecscompatibility string) *FindStructure {
	r.values.Set("ecs_compatibility", ecscompatibility)

	return r
}

// Explain If this parameter is set to `true`, the response includes a field named
// explanation, which is an array of strings that indicate how the structure
// finder produced its result.
// If the structure finder produces unexpected results for some text, use this
// query parameter to help you determine why the returned structure was chosen.
// API name: explain
func (r *FindStructure) Explain(explain bool) *FindStructure {
	r.values.Set("explain", strconv.FormatBool(explain))

	return r
}

// Format The high level structure of the text.
// Valid values are `ndjson`, `xml`, `delimited`, and `semi_structured_text`.
// By default, the API chooses the format.
// In this default scenario, all rows must have the same number of fields for a
// delimited format to be detected.
// If the format is set to `delimited` and the delimiter is not set, however,
// the API tolerates up to 5% of rows that have a different number of columns
// than the first row.
// API name: format
func (r *FindStructure) Format(format string) *FindStructure {
	r.values.Set("format", format)

	return r
}

// GrokPattern If you have set `format` to `semi_structured_text`, you can specify a Grok
// pattern that is used to extract fields from every message in the text.
// The name of the timestamp field in the Grok pattern must match what is
// specified in the `timestamp_field` parameter.
// If that parameter is not specified, the name of the timestamp field in the
// Grok pattern must match "timestamp".
// If `grok_pattern` is not specified, the structure finder creates a Grok
// pattern.
// API name: grok_pattern
func (r *FindStructure) GrokPattern(grokpattern string) *FindStructure {
	r.values.Set("grok_pattern", grokpattern)

	return r
}

// HasHeaderRow If you have set `format` to `delimited`, you can use this parameter to
// indicate whether the column names are in the first row of the text.
// If this parameter is not specified, the structure finder guesses based on the
// similarity of the first row of the text to other rows.
// API name: has_header_row
func (r *FindStructure) HasHeaderRow(hasheaderrow bool) *FindStructure {
	r.values.Set("has_header_row", strconv.FormatBool(hasheaderrow))

	return r
}

// LineMergeSizeLimit The maximum number of characters in a message when lines are merged to form
// messages while analyzing semi-structured text.
// If you have extremely long messages you may need to increase this, but be
// aware that this may lead to very long processing times if the way to group
// lines into messages is misdetected.
// API name: line_merge_size_limit
func (r *FindStructure) LineMergeSizeLimit(linemergesizelimit string) *FindStructure {
	r.values.Set("line_merge_size_limit", linemergesizelimit)

	return r
}

// LinesToSample The number of lines to include in the structural analysis, starting from the
// beginning of the text.
// The minimum is 2.
// If the value of this parameter is greater than the number of lines in the
// text, the analysis proceeds (as long as there are at least two lines in the
// text) for all of the lines.
//
// NOTE: The number of lines and the variation of the lines affects the speed of
// the analysis.
// For example, if you upload text where the first 1000 lines are all variations
// on the same message, the analysis will find more commonality than would be
// seen with a bigger sample.
// If possible, however, it is more efficient to upload sample text with more
// variety in the first 1000 lines than to request analysis of 100000 lines to
// achieve some variety.
// API name: lines_to_sample
func (r *FindStructure) LinesToSample(linestosample string) *FindStructure {
	r.values.Set("lines_to_sample", linestosample)

	return r
}

// Quote If you have set `format` to `delimited`, you can specify the character used
// to quote the values in each row if they contain newlines or the delimiter
// character.
// Only a single character is supported.
// If this parameter is not specified, the default value is a double quote
// (`"`).
// If your delimited text format does not use quoting, a workaround is to set
// this argument to a character that does not appear anywhere in the sample.
// API name: quote
func (r *FindStructure) Quote(quote string) *FindStructure {
	r.values.Set("quote", quote)

	return r
}

// ShouldTrimFields If you have set `format` to `delimited`, you can specify whether values
// between delimiters should have whitespace trimmed from them.
// If this parameter is not specified and the delimiter is pipe (`|`), the
// default value is `true`.
// Otherwise, the default value is `false`.
// API name: should_trim_fields
func (r *FindStructure) ShouldTrimFields(shouldtrimfields bool) *FindStructure {
	r.values.Set("should_trim_fields", strconv.FormatBool(shouldtrimfields))

	return r
}

// Timeout The maximum amount of time that the structure analysis can take.
// If the analysis is still running when the timeout expires then it will be
// stopped.
// API name: timeout
func (r *FindStructure) Timeout(duration string) *FindStructure {
	r.values.Set("timeout", duration)

	return r
}

// TimestampField The name of the field that contains the primary timestamp of each record in
// the text.
// In particular, if the text were ingested into an index, this is the field
// that would be used to populate the `@timestamp` field.
//
// If the `format` is `semi_structured_text`, this field must match the name of
// the appropriate extraction in the `grok_pattern`.
// Therefore, for semi-structured text, it is best not to specify this parameter
// unless `grok_pattern` is also specified.
//
// For structured text, if you specify this parameter, the field must exist
// within the text.
//
// If this parameter is not specified, the structure finder makes a decision
// about which field (if any) is the primary timestamp field.
// For structured text, it is not compulsory to have a timestamp in the text.
// API name: timestamp_field
func (r *FindStructure) TimestampField(field string) *FindStructure {
	r.values.Set("timestamp_field", field)

	return r
}

// TimestampFormat The Java time format of the timestamp field in the text.
//
// Only a subset of Java time format letter groups are supported:
//
// * `a`
// * `d`
// * `dd`
// * `EEE`
// * `EEEE`
// * `H`
// * `HH`
// * `h`
// * `M`
// * `MM`
// * `MMM`
// * `MMMM`
// * `mm`
// * `ss`
// * `XX`
// * `XXX`
// * `yy`
// * `yyyy`
// * `zzz`
//
// Additionally `S` letter groups (fractional seconds) of length one to nine are
// supported providing they occur after `ss` and separated from the `ss` by a
// `.`, `,` or `:`.
// Spacing and punctuation is also permitted with the exception of `?`, newline
// and carriage return, together with literal text enclosed in single quotes.
// For example, `MM/dd HH.mm.ss,SSSSSS 'in' yyyy` is a valid override format.
//
// One valuable use case for this parameter is when the format is
// semi-structured text, there are multiple timestamp formats in the text, and
// you know which format corresponds to the primary timestamp, but you do not
// want to specify the full `grok_pattern`.
// Another is when the timestamp format is one that the structure finder does
// not consider by default.
//
// If this parameter is not specified, the structure finder chooses the best
// format from a built-in set.
//
// If the special value `null` is specified the structure finder will not look
// for a primary timestamp in the text.
// When the format is semi-structured text this will result in the structure
// finder treating the text as single-line messages.
// API name: timestamp_format
func (r *FindStructure) TimestampFormat(timestampformat string) *FindStructure {
	r.values.Set("timestamp_format", timestampformat)

	return r
}
