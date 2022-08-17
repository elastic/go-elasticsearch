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

// DoubleNumberProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/core.ts#L129-L132
type DoubleNumberProperty struct {
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
	NullValue        *float64                                   `json:"null_value,omitempty"`
	OnScriptError    *onscripterror.OnScriptError               `json:"on_script_error,omitempty"`
	Properties       map[PropertyName]Property                  `json:"properties,omitempty"`
	Script           *Script                                    `json:"script,omitempty"`
	Similarity       *string                                    `json:"similarity,omitempty"`
	Store            *bool                                      `json:"store,omitempty"`
	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type,omitempty"`
}

// DoubleNumberPropertyBuilder holds DoubleNumberProperty struct and provides a builder API.
type DoubleNumberPropertyBuilder struct {
	v *DoubleNumberProperty
}

// NewDoubleNumberProperty provides a builder for the DoubleNumberProperty struct.
func NewDoubleNumberPropertyBuilder() *DoubleNumberPropertyBuilder {
	r := DoubleNumberPropertyBuilder{
		&DoubleNumberProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "double"

	return &r
}

// Build finalize the chain and returns the DoubleNumberProperty struct
func (rb *DoubleNumberPropertyBuilder) Build() DoubleNumberProperty {
	return *rb.v
}

func (rb *DoubleNumberPropertyBuilder) Coerce(coerce bool) *DoubleNumberPropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

func (rb *DoubleNumberPropertyBuilder) CopyTo(copyto *FieldsBuilder) *DoubleNumberPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *DoubleNumberPropertyBuilder) DocValues(docvalues bool) *DoubleNumberPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *DoubleNumberPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *DoubleNumberPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *DoubleNumberPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *DoubleNumberPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *DoubleNumberPropertyBuilder) IgnoreAbove(ignoreabove int) *DoubleNumberPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *DoubleNumberPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *DoubleNumberPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

func (rb *DoubleNumberPropertyBuilder) Index(index bool) *DoubleNumberPropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *DoubleNumberPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *DoubleNumberPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *DoubleNumberPropertyBuilder) Meta(value map[string]string) *DoubleNumberPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *DoubleNumberPropertyBuilder) NullValue(nullvalue float64) *DoubleNumberPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

func (rb *DoubleNumberPropertyBuilder) OnScriptError(onscripterror onscripterror.OnScriptError) *DoubleNumberPropertyBuilder {
	rb.v.OnScriptError = &onscripterror
	return rb
}

func (rb *DoubleNumberPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *DoubleNumberPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *DoubleNumberPropertyBuilder) Script(script *ScriptBuilder) *DoubleNumberPropertyBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *DoubleNumberPropertyBuilder) Similarity(similarity string) *DoubleNumberPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *DoubleNumberPropertyBuilder) Store(store bool) *DoubleNumberPropertyBuilder {
	rb.v.Store = &store
	return rb
}

func (rb *DoubleNumberPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *DoubleNumberPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}
