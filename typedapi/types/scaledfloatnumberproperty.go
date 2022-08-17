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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeseriesmetrictype"
)

// ScaledFloatNumberProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/core.ts#L159-L164
type ScaledFloatNumberProperty struct {
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
	Properties       map[PropertyName]Property                  `json:"properties,omitempty"`
	ScalingFactor    *float64                                   `json:"scaling_factor,omitempty"`
	Similarity       *string                                    `json:"similarity,omitempty"`
	Store            *bool                                      `json:"store,omitempty"`
	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type,omitempty"`
}

// ScaledFloatNumberPropertyBuilder holds ScaledFloatNumberProperty struct and provides a builder API.
type ScaledFloatNumberPropertyBuilder struct {
	v *ScaledFloatNumberProperty
}

// NewScaledFloatNumberProperty provides a builder for the ScaledFloatNumberProperty struct.
func NewScaledFloatNumberPropertyBuilder() *ScaledFloatNumberPropertyBuilder {
	r := ScaledFloatNumberPropertyBuilder{
		&ScaledFloatNumberProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "scaled_float"

	return &r
}

// Build finalize the chain and returns the ScaledFloatNumberProperty struct
func (rb *ScaledFloatNumberPropertyBuilder) Build() ScaledFloatNumberProperty {
	return *rb.v
}

func (rb *ScaledFloatNumberPropertyBuilder) Coerce(coerce bool) *ScaledFloatNumberPropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) CopyTo(copyto *FieldsBuilder) *ScaledFloatNumberPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) DocValues(docvalues bool) *ScaledFloatNumberPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *ScaledFloatNumberPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *ScaledFloatNumberPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) IgnoreAbove(ignoreabove int) *ScaledFloatNumberPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *ScaledFloatNumberPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) Index(index bool) *ScaledFloatNumberPropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *ScaledFloatNumberPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) Meta(value map[string]string) *ScaledFloatNumberPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) NullValue(nullvalue float64) *ScaledFloatNumberPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *ScaledFloatNumberPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) ScalingFactor(scalingfactor float64) *ScaledFloatNumberPropertyBuilder {
	rb.v.ScalingFactor = &scalingfactor
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) Similarity(similarity string) *ScaledFloatNumberPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) Store(store bool) *ScaledFloatNumberPropertyBuilder {
	rb.v.Store = &store
	return rb
}

func (rb *ScaledFloatNumberPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *ScaledFloatNumberPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}
