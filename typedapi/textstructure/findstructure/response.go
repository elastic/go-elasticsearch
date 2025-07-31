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

package findstructure

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package findstructure
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/text_structure/find_structure/FindStructureResponse.ts#L27-L97
type Response struct {

	// Charset The character encoding used to parse the text.
	Charset string `json:"charset"`
	// ColumnNames If `format` is `delimited`, the `column_names` field lists the column names
	// in the order they appear in the sample.
	ColumnNames         []string `json:"column_names,omitempty"`
	Delimiter           *string  `json:"delimiter,omitempty"`
	ExcludeLinesPattern *string  `json:"exclude_lines_pattern,omitempty"`
	Explanation         []string `json:"explanation,omitempty"`
	// FieldStats The most common values of each field, plus basic numeric statistics for the
	// numeric `page_count` field.
	// This information may provide clues that the data needs to be cleaned or
	// transformed prior to use by other Elastic Stack functionality.
	FieldStats map[string]types.FieldStat `json:"field_stats"`
	// Format Valid values include `ndjson`, `xml`, `delimited`, and
	// `semi_structured_text`.
	Format      string  `json:"format"`
	GrokPattern *string `json:"grok_pattern,omitempty"`
	// HasByteOrderMarker For UTF character encodings, it indicates whether the text begins with a byte
	// order marker.
	HasByteOrderMarker bool                 `json:"has_byte_order_marker"`
	HasHeaderRow       *bool                `json:"has_header_row,omitempty"`
	IngestPipeline     types.PipelineConfig `json:"ingest_pipeline"`
	// JavaTimestampFormats The Java time formats recognized in the time fields.
	// Elasticsearch mappings and ingest pipelines use this format.
	JavaTimestampFormats []string `json:"java_timestamp_formats,omitempty"`
	// JodaTimestampFormats Information that is used to tell Logstash how to parse timestamps.
	JodaTimestampFormats []string `json:"joda_timestamp_formats,omitempty"`
	// Mappings Some suitable mappings for an index into which the data could be ingested.
	Mappings              types.TypeMapping `json:"mappings"`
	MultilineStartPattern *string           `json:"multiline_start_pattern,omitempty"`
	// NeedClientTimezone If a timestamp format is detected that does not include a timezone,
	// `need_client_timezone` is `true`.
	// The server that parses the text must therefore be told the correct timezone
	// by the client.
	NeedClientTimezone bool `json:"need_client_timezone"`
	// NumLinesAnalyzed The number of lines of the text that were analyzed.
	NumLinesAnalyzed int `json:"num_lines_analyzed"`
	// NumMessagesAnalyzed The number of distinct messages the lines contained.
	// For NDJSON, this value is the same as `num_lines_analyzed`.
	// For other text formats, messages can span several lines.
	NumMessagesAnalyzed int     `json:"num_messages_analyzed"`
	Quote               *string `json:"quote,omitempty"`
	// SampleStart The first two messages in the text verbatim.
	// This may help diagnose parse errors or accidental uploads of the wrong text.
	SampleStart      string `json:"sample_start"`
	ShouldTrimFields *bool  `json:"should_trim_fields,omitempty"`
	// TimestampField The field considered most likely to be the primary timestamp of each
	// document.
	TimestampField *string `json:"timestamp_field,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		FieldStats: make(map[string]types.FieldStat, 0),
	}
	return r
}
