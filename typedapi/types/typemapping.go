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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// TypeMapping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/TypeMapping.ts#L34-L55
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
	Meta_                *Metadata                      `json:"_meta,omitempty"`
	NumericDetection     *bool                          `json:"numeric_detection,omitempty"`
	Properties           map[PropertyName]Property      `json:"properties,omitempty"`
	Routing_             *RoutingField                  `json:"_routing,omitempty"`
	Runtime              map[string]RuntimeField        `json:"runtime,omitempty"`
	Size_                *SizeField                     `json:"_size,omitempty"`
	Source_              *SourceField                   `json:"_source,omitempty"`
}

// TypeMappingBuilder holds TypeMapping struct and provides a builder API.
type TypeMappingBuilder struct {
	v *TypeMapping
}

// NewTypeMapping provides a builder for the TypeMapping struct.
func NewTypeMappingBuilder() *TypeMappingBuilder {
	r := TypeMappingBuilder{
		&TypeMapping{
			Properties: make(map[PropertyName]Property, 0),
			Runtime:    make(map[string]RuntimeField, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TypeMapping struct
func (rb *TypeMappingBuilder) Build() TypeMapping {
	return *rb.v
}

func (rb *TypeMappingBuilder) AllField(allfield *AllFieldBuilder) *TypeMappingBuilder {
	v := allfield.Build()
	rb.v.AllField = &v
	return rb
}

func (rb *TypeMappingBuilder) DataStreamTimestamp_(datastreamtimestamp_ *DataStreamTimestampBuilder) *TypeMappingBuilder {
	v := datastreamtimestamp_.Build()
	rb.v.DataStreamTimestamp_ = &v
	return rb
}

func (rb *TypeMappingBuilder) DateDetection(datedetection bool) *TypeMappingBuilder {
	rb.v.DateDetection = &datedetection
	return rb
}

func (rb *TypeMappingBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *TypeMappingBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *TypeMappingBuilder) DynamicDateFormats(dynamic_date_formats ...string) *TypeMappingBuilder {
	rb.v.DynamicDateFormats = dynamic_date_formats
	return rb
}

func (rb *TypeMappingBuilder) DynamicTemplates(arg []map[string]DynamicTemplate) *TypeMappingBuilder {
	rb.v.DynamicTemplates = arg
	return rb
}

func (rb *TypeMappingBuilder) Enabled(enabled bool) *TypeMappingBuilder {
	rb.v.Enabled = &enabled
	return rb
}

func (rb *TypeMappingBuilder) FieldNames_(fieldnames_ *FieldNamesFieldBuilder) *TypeMappingBuilder {
	v := fieldnames_.Build()
	rb.v.FieldNames_ = &v
	return rb
}

func (rb *TypeMappingBuilder) IndexField(indexfield *IndexFieldBuilder) *TypeMappingBuilder {
	v := indexfield.Build()
	rb.v.IndexField = &v
	return rb
}

func (rb *TypeMappingBuilder) Meta_(meta_ *MetadataBuilder) *TypeMappingBuilder {
	v := meta_.Build()
	rb.v.Meta_ = &v
	return rb
}

func (rb *TypeMappingBuilder) NumericDetection(numericdetection bool) *TypeMappingBuilder {
	rb.v.NumericDetection = &numericdetection
	return rb
}

func (rb *TypeMappingBuilder) Properties(values map[PropertyName]*PropertyBuilder) *TypeMappingBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *TypeMappingBuilder) Routing_(routing_ *RoutingFieldBuilder) *TypeMappingBuilder {
	v := routing_.Build()
	rb.v.Routing_ = &v
	return rb
}

func (rb *TypeMappingBuilder) Runtime(values map[string]*RuntimeFieldBuilder) *TypeMappingBuilder {
	tmp := make(map[string]RuntimeField, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Runtime = tmp
	return rb
}

func (rb *TypeMappingBuilder) Size_(size_ *SizeFieldBuilder) *TypeMappingBuilder {
	v := size_.Build()
	rb.v.Size_ = &v
	return rb
}

func (rb *TypeMappingBuilder) Source_(source_ *SourceFieldBuilder) *TypeMappingBuilder {
	v := source_.Build()
	rb.v.Source_ = &v
	return rb
}
