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
// https://github.com/elastic/elasticsearch-specification/tree/f1932ce6b46a53a8342db522b1a7883bcc9e0996

package findfieldstructure

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/ecscompatibilitytype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/formattype"
)

// Response holds the response body struct for the package findfieldstructure
//
// https://github.com/elastic/elasticsearch-specification/blob/f1932ce6b46a53a8342db522b1a7883bcc9e0996/specification/text_structure/find_field_structure/FindFieldStructureResponse.ts#L31-L49
type Response struct {
	Charset               string                                     `json:"charset"`
	EcsCompatibility      *ecscompatibilitytype.EcsCompatibilityType `json:"ecs_compatibility,omitempty"`
	FieldStats            map[string]types.FieldStat                 `json:"field_stats"`
	Format                formattype.FormatType                      `json:"format"`
	GrokPattern           *string                                    `json:"grok_pattern,omitempty"`
	IngestPipeline        types.PipelineConfig                       `json:"ingest_pipeline"`
	JavaTimestampFormats  []string                                   `json:"java_timestamp_formats,omitempty"`
	JodaTimestampFormats  []string                                   `json:"joda_timestamp_formats,omitempty"`
	Mappings              types.TypeMapping                          `json:"mappings"`
	MultilineStartPattern *string                                    `json:"multiline_start_pattern,omitempty"`
	NeedClientTimezone    bool                                       `json:"need_client_timezone"`
	NumLinesAnalyzed      int                                        `json:"num_lines_analyzed"`
	NumMessagesAnalyzed   int                                        `json:"num_messages_analyzed"`
	SampleStart           string                                     `json:"sample_start"`
	TimestampField        *string                                    `json:"timestamp_field,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		FieldStats: make(map[string]types.FieldStat, 0),
	}
	return r
}
