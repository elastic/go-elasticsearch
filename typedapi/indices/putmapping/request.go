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


package putmapping

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// Request holds the request body struct for the package putmapping
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/put_mapping/IndicesPutMappingRequest.ts#L42-L116
type Request struct {

	// DateDetection Controls whether dynamic date detection is enabled.
	DateDetection *bool `json:"date_detection,omitempty"`

	// Dynamic Controls whether new fields are added dynamically.
	Dynamic *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`

	// DynamicDateFormats If date detection is enabled then new string fields are checked
	// against 'dynamic_date_formats' and if the value matches then
	// a new date field is added instead of string.
	DynamicDateFormats []string `json:"dynamic_date_formats,omitempty"`

	// DynamicTemplates Specify dynamic templates for the mapping.
	DynamicTemplates []map[string]types.DynamicTemplate `json:"dynamic_templates,omitempty"`

	// FieldNames_ Control whether field names are enabled for the index.
	FieldNames_ *types.FieldNamesField `json:"_field_names,omitempty"`

	// Meta_ A mapping type can have custom meta data associated with it. These are
	// not used at all by Elasticsearch, but can be used to store
	// application-specific metadata.
	Meta_ *types.Metadata `json:"_meta,omitempty"`

	// NumericDetection Automatically map strings into numeric data types for all fields.
	NumericDetection *bool `json:"numeric_detection,omitempty"`

	// Properties Mapping for a field. For new fields, this mapping can include:
	//
	// - Field name
	// - Field data type
	// - Mapping parameters
	Properties map[types.PropertyName]types.Property `json:"properties,omitempty"`

	// Routing_ Enable making a routing value required on indexed documents.
	Routing_ *types.RoutingField `json:"_routing,omitempty"`

	// Runtime Mapping of runtime fields for the index.
	Runtime *types.RuntimeFields `json:"runtime,omitempty"`

	// Source_ Control whether the _source field is enabled on the index.
	Source_ *types.SourceField `json:"_source,omitempty"`
}

// RequestBuilder is the builder API for the putmapping.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			Properties: make(map[types.PropertyName]types.Property, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putmapping request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) DateDetection(datedetection bool) *RequestBuilder {
	rb.v.DateDetection = &datedetection
	return rb
}

func (rb *RequestBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *RequestBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *RequestBuilder) DynamicDateFormats(dynamic_date_formats ...string) *RequestBuilder {
	rb.v.DynamicDateFormats = dynamic_date_formats
	return rb
}

func (rb *RequestBuilder) DynamicTemplates(arg []map[string]types.DynamicTemplate) *RequestBuilder {
	rb.v.DynamicTemplates = arg
	return rb
}

func (rb *RequestBuilder) FieldNames_(fieldnames_ *types.FieldNamesFieldBuilder) *RequestBuilder {
	v := fieldnames_.Build()
	rb.v.FieldNames_ = &v
	return rb
}

func (rb *RequestBuilder) Meta_(meta_ *types.MetadataBuilder) *RequestBuilder {
	v := meta_.Build()
	rb.v.Meta_ = &v
	return rb
}

func (rb *RequestBuilder) NumericDetection(numericdetection bool) *RequestBuilder {
	rb.v.NumericDetection = &numericdetection
	return rb
}

func (rb *RequestBuilder) Properties(values map[types.PropertyName]*types.PropertyBuilder) *RequestBuilder {
	tmp := make(map[types.PropertyName]types.Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *RequestBuilder) Routing_(routing_ *types.RoutingFieldBuilder) *RequestBuilder {
	v := routing_.Build()
	rb.v.Routing_ = &v
	return rb
}

func (rb *RequestBuilder) Runtime(runtime *types.RuntimeFieldsBuilder) *RequestBuilder {
	v := runtime.Build()
	rb.v.Runtime = &v
	return rb
}

func (rb *RequestBuilder) Source_(source_ *types.SourceFieldBuilder) *RequestBuilder {
	v := source_.Build()
	rb.v.Source_ = &v
	return rb
}
