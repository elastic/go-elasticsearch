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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// TypeMapping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/mapping/TypeMapping.ts#L34-L55
type TypeMapping struct {
	AllField             *AllField                      `json:"all_field,omitempty"`
	DataStreamTimestamp_ *DataStreamTimestamp           `json:"_data_stream_timestamp,omitempty"`
	DateDetection        *bool                          `json:"date_detection,omitempty"`
	Dynamic              *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	DynamicDateFormats   []string                       `json:"dynamic_date_formats,omitempty"`
	DynamicTemplates     []map[string]DynamicTemplate   `json:"dynamic_templates,omitempty"`
	Enabled              *bool                          `json:"enabled,omitempty"`
	FieldNames_          *FieldNamesField               `json:"_field_names,omitempty"`
	IndexField           *IndexField                    `json:"index_field,omitempty"`
	Meta_                map[string]interface{}         `json:"_meta,omitempty"`
	NumericDetection     *bool                          `json:"numeric_detection,omitempty"`
	Properties           map[string]Property            `json:"properties,omitempty"`
	Routing_             *RoutingField                  `json:"_routing,omitempty"`
	Runtime              map[string]RuntimeField        `json:"runtime,omitempty"`
	Size_                *SizeField                     `json:"_size,omitempty"`
	Source_              *SourceField                   `json:"_source,omitempty"`
}

// NewTypeMapping returns a TypeMapping.
func NewTypeMapping() *TypeMapping {
	r := &TypeMapping{
		Properties: make(map[string]Property, 0),
		Runtime:    make(map[string]RuntimeField, 0),
	}

	return r
}
