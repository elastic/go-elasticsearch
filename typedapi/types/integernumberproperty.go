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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeseriesmetrictype"
)

// IntegerNumberProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/core.ts#L134-L137
type IntegerNumberProperty struct {
	Coerce           *bool                                      `json:"coerce,omitempty"`
	CopyTo           *Fields                                    `json:"copy_to,omitempty"`
	DocValues        *bool                                      `json:"doc_values,omitempty"`
	Dynamic          *dynamicmapping.DynamicMapping             `json:"dynamic,omitempty"`
	Fields           map[PropertyName]Property                  `json:"fields,omitempty"`
	IgnoreAbove      *int                                       `json:"ignore_above,omitempty"`
	IgnoreMalformed  *bool                                      `json:"ignore_malformed,omitempty"`
	Index            *bool                                      `json:"index,omitempty"`
	LocalMetadata    *Metadata                                  `json:"local_metadata,omitempty"`
	Meta             map[string]string                          `json:"meta,omitempty"`
	NullValue        *int                                       `json:"null_value,omitempty"`
	OnScriptError    *onscripterror.OnScriptError               `json:"on_script_error,omitempty"`
	Properties       map[PropertyName]Property                  `json:"properties,omitempty"`
	Script           *Script                                    `json:"script,omitempty"`
	Similarity       *string                                    `json:"similarity,omitempty"`
	Store            *bool                                      `json:"store,omitempty"`
	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type,omitempty"`
}

// IntegerNumberPropertyBuilder holds IntegerNumberProperty struct and provides a builder API.
type IntegerNumberPropertyBuilder struct {
	v *IntegerNumberProperty
}

// NewIntegerNumberProperty provides a builder for the IntegerNumberProperty struct.
func NewIntegerNumberPropertyBuilder() *IntegerNumberPropertyBuilder {
	r := IntegerNumberPropertyBuilder{
		&IntegerNumberProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "integer"

	return &r
}

// Build finalize the chain and returns the IntegerNumberProperty struct
func (rb *IntegerNumberPropertyBuilder) Build() IntegerNumberProperty {
	return *rb.v
}

func (rb *IntegerNumberPropertyBuilder) Coerce(coerce bool) *IntegerNumberPropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

func (rb *IntegerNumberPropertyBuilder) CopyTo(copyto *FieldsBuilder) *IntegerNumberPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *IntegerNumberPropertyBuilder) DocValues(docvalues bool) *IntegerNumberPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *IntegerNumberPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *IntegerNumberPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *IntegerNumberPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *IntegerNumberPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *IntegerNumberPropertyBuilder) IgnoreAbove(ignoreabove int) *IntegerNumberPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *IntegerNumberPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *IntegerNumberPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

func (rb *IntegerNumberPropertyBuilder) Index(index bool) *IntegerNumberPropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *IntegerNumberPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *IntegerNumberPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *IntegerNumberPropertyBuilder) Meta(value map[string]string) *IntegerNumberPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *IntegerNumberPropertyBuilder) NullValue(nullvalue int) *IntegerNumberPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

func (rb *IntegerNumberPropertyBuilder) OnScriptError(onscripterror onscripterror.OnScriptError) *IntegerNumberPropertyBuilder {
	rb.v.OnScriptError = &onscripterror
	return rb
}

func (rb *IntegerNumberPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *IntegerNumberPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *IntegerNumberPropertyBuilder) Script(script *ScriptBuilder) *IntegerNumberPropertyBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *IntegerNumberPropertyBuilder) Similarity(similarity string) *IntegerNumberPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *IntegerNumberPropertyBuilder) Store(store bool) *IntegerNumberPropertyBuilder {
	rb.v.Store = &store
	return rb
}

func (rb *IntegerNumberPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *IntegerNumberPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}
