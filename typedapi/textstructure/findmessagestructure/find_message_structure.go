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

// Find the structure of text messages.
// Find the structure of a list of text messages.
// The messages must contain data that is suitable to be ingested into
// Elasticsearch.
//
// This API provides a starting point for ingesting data into Elasticsearch in a
// format that is suitable for subsequent use with other Elastic Stack
// functionality.
// Use this API rather than the find text structure API if your input text has
// already been split up into separate messages by some other process.
//
// The response from the API contains:
//
// * Sample messages.
// * Statistics that reveal the most common values for all fields detected
// within the text and basic numeric statistics for numeric fields.
// * Information about the structure of the text, which is useful when you write
// ingest configurations to index it or similarly formatted text.
// Appropriate mappings for an Elasticsearch index, which you could use to
// ingest the text.
//
// All this information can be calculated by the structure finder with no
// guidance.
// However, you can optionally override some of the decisions about the text
// structure by specifying one or more query parameters.
//
// If the structure finder produces unexpected results, specify the `explain`
// query parameter and an explanation will appear in the response.
// It helps determine why the returned structure was chosen.
package findmessagestructure

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/ecscompatibilitytype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/formattype"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type FindMessageStructure struct {
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

// NewFindMessageStructure type alias for index.
type NewFindMessageStructure func() *FindMessageStructure

// NewFindMessageStructureFunc returns a new instance of FindMessageStructure with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewFindMessageStructureFunc(tp elastictransport.Interface) NewFindMessageStructure {
	return func() *FindMessageStructure {
		n := New(tp)

		return n
	}
}

// Find the structure of text messages.
// Find the structure of a list of text messages.
// The messages must contain data that is suitable to be ingested into
// Elasticsearch.
//
// This API provides a starting point for ingesting data into Elasticsearch in a
// format that is suitable for subsequent use with other Elastic Stack
// functionality.
// Use this API rather than the find text structure API if your input text has
// already been split up into separate messages by some other process.
//
// The response from the API contains:
//
// * Sample messages.
// * Statistics that reveal the most common values for all fields detected
// within the text and basic numeric statistics for numeric fields.
// * Information about the structure of the text, which is useful when you write
// ingest configurations to index it or similarly formatted text.
// Appropriate mappings for an Elasticsearch index, which you could use to
// ingest the text.
//
// All this information can be calculated by the structure finder with no
// guidance.
// However, you can optionally override some of the decisions about the text
// structure by specifying one or more query parameters.
//
// If the structure finder produces unexpected results, specify the `explain`
// query parameter and an explanation will appear in the response.
// It helps determine why the returned structure was chosen.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-message-structure.html
func New(tp elastictransport.Interface) *FindMessageStructure {
	r := &FindMessageStructure{
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
func (r *FindMessageStructure) Raw(raw io.Reader) *FindMessageStructure {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *FindMessageStructure) Request(req *Request) *FindMessageStructure {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *FindMessageStructure) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for FindMessageStructure: %w", err)
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
		path.WriteString("_text_structure")
		path.WriteString("/")
		path.WriteString("find_message_structure")

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
func (r FindMessageStructure) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "text_structure.find_message_structure")
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
		instrument.BeforeRequest(req, "text_structure.find_message_structure")
		if reader := instrument.RecordRequestBody(ctx, "text_structure.find_message_structure", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "text_structure.find_message_structure")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the FindMessageStructure query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a findmessagestructure.Response
func (r FindMessageStructure) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "text_structure.find_message_structure")
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

// Header set a key, value pair in the FindMessageStructure headers map.
func (r *FindMessageStructure) Header(key, value string) *FindMessageStructure {
	r.headers.Set(key, value)

	return r
}

// ColumnNames If the format is `delimited`, you can specify the column names in a
// comma-separated list.
// If this parameter is not specified, the structure finder uses the column
// names from the header row of the text.
// If the text does not have a header role, columns are named "column1",
// "column2", "column3", for example.
// API name: column_names
func (r *FindMessageStructure) ColumnNames(columnnames string) *FindMessageStructure {
	r.values.Set("column_names", columnnames)

	return r
}

// Delimiter If you the format is `delimited`, you can specify the character used to
// delimit the values in each row.
// Only a single character is supported; the delimiter cannot have multiple
// characters.
// By default, the API considers the following possibilities: comma, tab,
// semi-colon, and pipe (`|`).
// In this default scenario, all rows must have the same number of fields for
// the delimited format to be detected.
// If you specify a delimiter, up to 10% of the rows can have a different number
// of columns than the first row.
// API name: delimiter
func (r *FindMessageStructure) Delimiter(delimiter string) *FindMessageStructure {
	r.values.Set("delimiter", delimiter)

	return r
}

// EcsCompatibility The mode of compatibility with ECS compliant Grok patterns.
// Use this parameter to specify whether to use ECS Grok patterns instead of
// legacy ones when the structure finder creates a Grok pattern.
// This setting primarily has an impact when a whole message Grok pattern such
// as `%{CATALINALOG}` matches the input.
// If the structure finder identifies a common structure but has no idea of
// meaning then generic field names such as `path`, `ipaddress`, `field1`, and
// `field2` are used in the `grok_pattern` output, with the intention that a
// user who knows the meanings rename these fields before using it.
// API name: ecs_compatibility
func (r *FindMessageStructure) EcsCompatibility(ecscompatibility ecscompatibilitytype.EcsCompatibilityType) *FindMessageStructure {
	r.values.Set("ecs_compatibility", ecscompatibility.String())

	return r
}

// Explain If this parameter is set to true, the response includes a field named
// `explanation`, which is an array of strings that indicate how the structure
// finder produced its result.
// API name: explain
func (r *FindMessageStructure) Explain(explain bool) *FindMessageStructure {
	r.values.Set("explain", strconv.FormatBool(explain))

	return r
}

// Format The high level structure of the text.
// By default, the API chooses the format.
// In this default scenario, all rows must have the same number of fields for a
// delimited format to be detected.
// If the format is `delimited` and the delimiter is not set, however, the API
// tolerates up to 5% of rows that have a different number of columns than the
// first row.
// API name: format
func (r *FindMessageStructure) Format(format formattype.FormatType) *FindMessageStructure {
	r.values.Set("format", format.String())

	return r
}

// GrokPattern If the format is `semi_structured_text`, you can specify a Grok pattern that
// is used to extract fields from every message in the text.
// The name of the timestamp field in the Grok pattern must match what is
// specified in the `timestamp_field` parameter.
// If that parameter is not specified, the name of the timestamp field in the
// Grok pattern must match "timestamp".
// If `grok_pattern` is not specified, the structure finder creates a Grok
// pattern.
// API name: grok_pattern
func (r *FindMessageStructure) GrokPattern(grokpattern string) *FindMessageStructure {
	r.values.Set("grok_pattern", grokpattern)

	return r
}

// Quote If the format is `delimited`, you can specify the character used to quote the
// values in each row if they contain newlines or the delimiter character.
// Only a single character is supported.
// If this parameter is not specified, the default value is a double quote
// (`"`).
// If your delimited text format does not use quoting, a workaround is to set
// this argument to a character that does not appear anywhere in the sample.
// API name: quote
func (r *FindMessageStructure) Quote(quote string) *FindMessageStructure {
	r.values.Set("quote", quote)

	return r
}

// ShouldTrimFields If the format is `delimited`, you can specify whether values between
// delimiters should have whitespace trimmed from them.
// If this parameter is not specified and the delimiter is pipe (`|`), the
// default value is true.
// Otherwise, the default value is `false`.
// API name: should_trim_fields
func (r *FindMessageStructure) ShouldTrimFields(shouldtrimfields bool) *FindMessageStructure {
	r.values.Set("should_trim_fields", strconv.FormatBool(shouldtrimfields))

	return r
}

// Timeout The maximum amount of time that the structure analysis can take.
// If the analysis is still running when the timeout expires, it will be
// stopped.
// API name: timeout
func (r *FindMessageStructure) Timeout(duration string) *FindMessageStructure {
	r.values.Set("timeout", duration)

	return r
}

// TimestampField The name of the field that contains the primary timestamp of each record in
// the text.
// In particular, if the text was ingested into an index, this is the field that
// would be used to populate the `@timestamp` field.
//
// If the format is `semi_structured_text`, this field must match the name of
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
func (r *FindMessageStructure) TimestampField(field string) *FindMessageStructure {
	r.values.Set("timestamp_field", field)

	return r
}

// TimestampFormat The Java time format of the timestamp field in the text.
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
// supported providing they occur after `ss` and are separated from the `ss` by
// a period (`.`), comma (`,`), or colon (`:`).
// Spacing and punctuation is also permitted with the exception a question mark
// (`?`), newline, and carriage return, together with literal text enclosed in
// single quotes.
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
// If the special value `null` is specified, the structure finder will not look
// for a primary timestamp in the text.
// When the format is semi-structured text, this will result in the structure
// finder treating the text as single-line messages.
// API name: timestamp_format
func (r *FindMessageStructure) TimestampFormat(timestampformat string) *FindMessageStructure {
	r.values.Set("timestamp_format", timestampformat)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *FindMessageStructure) ErrorTrace(errortrace bool) *FindMessageStructure {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *FindMessageStructure) FilterPath(filterpaths ...string) *FindMessageStructure {
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
func (r *FindMessageStructure) Human(human bool) *FindMessageStructure {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *FindMessageStructure) Pretty(pretty bool) *FindMessageStructure {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Messages The list of messages you want to analyze.
// API name: messages
func (r *FindMessageStructure) Messages(messages ...string) *FindMessageStructure {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Messages = messages

	return r
}
