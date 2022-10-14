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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeseriesmetrictype"
)

// UnsignedLongNumberProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/mapping/core.ts#L166-L169
type UnsignedLongNumberProperty struct {
	Boost           *float64                       `json:"boost,omitempty"`
	Coerce          *bool                          `json:"coerce,omitempty"`
	CopyTo          *Fields                        `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	Index           *bool                          `json:"index,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	// Meta Metadata about the field.
	Meta          map[string]string            `json:"meta,omitempty"`
	NullValue     *uint64                      `json:"null_value,omitempty"`
	OnScriptError *onscripterror.OnScriptError `json:"on_script_error,omitempty"`
	Properties    map[PropertyName]Property    `json:"properties,omitempty"`
	Script        *Script                      `json:"script,omitempty"`
	Similarity    *string                      `json:"similarity,omitempty"`
	Store         *bool                        `json:"store,omitempty"`
	// TimeSeriesDimension For internal use by Elastic only. Marks the field as a time series dimension.
	// Defaults to false.
	TimeSeriesDimension *bool `json:"time_series_dimension,omitempty"`
	// TimeSeriesMetric For internal use by Elastic only. Marks the field as a time series dimension.
	// Defaults to false.
	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type,omitempty"`
}

// UnsignedLongNumberPropertyBuilder holds UnsignedLongNumberProperty struct and provides a builder API.
type UnsignedLongNumberPropertyBuilder struct {
	v *UnsignedLongNumberProperty
}

// NewUnsignedLongNumberProperty provides a builder for the UnsignedLongNumberProperty struct.
func NewUnsignedLongNumberPropertyBuilder() *UnsignedLongNumberPropertyBuilder {
	r := UnsignedLongNumberPropertyBuilder{
		&UnsignedLongNumberProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "unsigned_long"

	return &r
}

// Build finalize the chain and returns the UnsignedLongNumberProperty struct
func (rb *UnsignedLongNumberPropertyBuilder) Build() UnsignedLongNumberProperty {
	return *rb.v
}

func (rb *UnsignedLongNumberPropertyBuilder) Boost(boost float64) *UnsignedLongNumberPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) Coerce(coerce bool) *UnsignedLongNumberPropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) CopyTo(copyto *FieldsBuilder) *UnsignedLongNumberPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) DocValues(docvalues bool) *UnsignedLongNumberPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *UnsignedLongNumberPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *UnsignedLongNumberPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) IgnoreAbove(ignoreabove int) *UnsignedLongNumberPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *UnsignedLongNumberPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) Index(index bool) *UnsignedLongNumberPropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *UnsignedLongNumberPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

// Meta Metadata about the field.

func (rb *UnsignedLongNumberPropertyBuilder) Meta(value map[string]string) *UnsignedLongNumberPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) NullValue(nullvalue uint64) *UnsignedLongNumberPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) OnScriptError(onscripterror onscripterror.OnScriptError) *UnsignedLongNumberPropertyBuilder {
	rb.v.OnScriptError = &onscripterror
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *UnsignedLongNumberPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) Script(script *ScriptBuilder) *UnsignedLongNumberPropertyBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) Similarity(similarity string) *UnsignedLongNumberPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *UnsignedLongNumberPropertyBuilder) Store(store bool) *UnsignedLongNumberPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// TimeSeriesDimension For internal use by Elastic only. Marks the field as a time series dimension.
// Defaults to false.

func (rb *UnsignedLongNumberPropertyBuilder) TimeSeriesDimension(timeseriesdimension bool) *UnsignedLongNumberPropertyBuilder {
	rb.v.TimeSeriesDimension = &timeseriesdimension
	return rb
}

// TimeSeriesMetric For internal use by Elastic only. Marks the field as a time series dimension.
// Defaults to false.

func (rb *UnsignedLongNumberPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *UnsignedLongNumberPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}
