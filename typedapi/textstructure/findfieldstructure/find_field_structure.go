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

// Find the structure of a text field.
// Find the structure of a text field in an Elasticsearch index.
//
// This API provides a starting point for extracting further information from
// log messages already ingested into Elasticsearch.
// For example, if you have ingested data into a very simple index that has just
// `@timestamp` and message fields, you can use this API to see what common
// structure exists in the message field.
//
// The response from the API contains:
//
// * Sample messages.
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
// If the structure finder produces unexpected results, specify the `explain`
// query parameter and an explanation will appear in the response.
// It helps determine why the returned structure was chosen.
package findfieldstructure

import (
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

type FindFieldStructure struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewFindFieldStructure type alias for index.
type NewFindFieldStructure func() *FindFieldStructure

// NewFindFieldStructureFunc returns a new instance of FindFieldStructure with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewFindFieldStructureFunc(tp elastictransport.Interface) NewFindFieldStructure {
	return func() *FindFieldStructure {
		n := New(tp)

		return n
	}
}

// Find the structure of a text field.
// Find the structure of a text field in an Elasticsearch index.
//
// This API provides a starting point for extracting further information from
// log messages already ingested into Elasticsearch.
// For example, if you have ingested data into a very simple index that has just
// `@timestamp` and message fields, you can use this API to see what common
// structure exists in the message field.
//
// The response from the API contains:
//
// * Sample messages.
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
// If the structure finder produces unexpected results, specify the `explain`
// query parameter and an explanation will appear in the response.
// It helps determine why the returned structure was chosen.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-field-structure.html
func New(tp elastictransport.Interface) *FindFieldStructure {
	r := &FindFieldStructure{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *FindFieldStructure) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_text_structure")
		path.WriteString("/")
		path.WriteString("find_field_structure")

		method = http.MethodGet
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r FindFieldStructure) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "text_structure.find_field_structure")
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
		instrument.BeforeRequest(req, "text_structure.find_field_structure")
		if reader := instrument.RecordRequestBody(ctx, "text_structure.find_field_structure", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "text_structure.find_field_structure")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the FindFieldStructure query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a findfieldstructure.Response
func (r FindFieldStructure) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "text_structure.find_field_structure")
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r FindFieldStructure) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "text_structure.find_field_structure")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the FindFieldStructure query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the FindFieldStructure headers map.
func (r *FindFieldStructure) Header(key, value string) *FindFieldStructure {
	r.headers.Set(key, value)

	return r
}

// ColumnNames If `format` is set to `delimited`, you can specify the column names in a
// comma-separated list.
// If this parameter is not specified, the structure finder uses the column
// names from the header row of the text.
// If the text does not have a header row, columns are named "column1",
// "column2", "column3", for example.
// API name: column_names
func (r *FindFieldStructure) ColumnNames(columnnames string) *FindFieldStructure {
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
func (r *FindFieldStructure) Delimiter(delimiter string) *FindFieldStructure {
	r.values.Set("delimiter", delimiter)

	return r
}

// DocumentsToSample The number of documents to include in the structural analysis.
// The minimum value is 2.
// API name: documents_to_sample
func (r *FindFieldStructure) DocumentsToSample(documentstosample string) *FindFieldStructure {
	r.values.Set("documents_to_sample", documentstosample)

	return r
}

// EcsCompatibility The mode of compatibility with ECS compliant Grok patterns.
// Use this parameter to specify whether to use ECS Grok patterns instead of
// legacy ones when the structure finder creates a Grok pattern.
// This setting primarily has an impact when a whole message Grok pattern such
// as `%{CATALINALOG}` matches the input.
// If the structure finder identifies a common structure but has no idea of the
// meaning then generic field names such as `path`, `ipaddress`, `field1`, and
// `field2` are used in the `grok_pattern` output.
// The intention in that situation is that a user who knows the meanings will
// rename the fields before using them.
// API name: ecs_compatibility
func (r *FindFieldStructure) EcsCompatibility(ecscompatibility ecscompatibilitytype.EcsCompatibilityType) *FindFieldStructure {
	r.values.Set("ecs_compatibility", ecscompatibility.String())

	return r
}

// Explain If `true`, the response includes a field named `explanation`, which is an
// array of strings that indicate how the structure finder produced its result.
// API name: explain
func (r *FindFieldStructure) Explain(explain bool) *FindFieldStructure {
	r.values.Set("explain", strconv.FormatBool(explain))

	return r
}

// Field The field that should be analyzed.
// API name: field
func (r *FindFieldStructure) Field(field string) *FindFieldStructure {
	r.values.Set("field", field)

	return r
}

// Format The high level structure of the text.
// By default, the API chooses the format.
// In this default scenario, all rows must have the same number of fields for a
// delimited format to be detected.
// If the format is set to delimited and the delimiter is not set, however, the
// API tolerates up to 5% of rows that have a different number of columns than
// the first row.
// API name: format
func (r *FindFieldStructure) Format(format formattype.FormatType) *FindFieldStructure {
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
func (r *FindFieldStructure) GrokPattern(grokpattern string) *FindFieldStructure {
	r.values.Set("grok_pattern", grokpattern)

	return r
}

// Index The name of the index that contains the analyzed field.
// API name: index
func (r *FindFieldStructure) Index(indexname string) *FindFieldStructure {
	r.values.Set("index", indexname)

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
func (r *FindFieldStructure) Quote(quote string) *FindFieldStructure {
	r.values.Set("quote", quote)

	return r
}

// ShouldTrimFields If the format is `delimited`, you can specify whether values between
// delimiters should have whitespace trimmed from them.
// If this parameter is not specified and the delimiter is pipe (`|`), the
// default value is true.
// Otherwise, the default value is `false`.
// API name: should_trim_fields
func (r *FindFieldStructure) ShouldTrimFields(shouldtrimfields bool) *FindFieldStructure {
	r.values.Set("should_trim_fields", strconv.FormatBool(shouldtrimfields))

	return r
}

// Timeout The maximum amount of time that the structure analysis can take.
// If the analysis is still running when the timeout expires, it will be
// stopped.
// API name: timeout
func (r *FindFieldStructure) Timeout(duration string) *FindFieldStructure {
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
func (r *FindFieldStructure) TimestampField(field string) *FindFieldStructure {
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
func (r *FindFieldStructure) TimestampFormat(timestampformat string) *FindFieldStructure {
	r.values.Set("timestamp_format", timestampformat)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *FindFieldStructure) ErrorTrace(errortrace bool) *FindFieldStructure {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *FindFieldStructure) FilterPath(filterpaths ...string) *FindFieldStructure {
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
func (r *FindFieldStructure) Human(human bool) *FindFieldStructure {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *FindFieldStructure) Pretty(pretty bool) *FindFieldStructure {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
