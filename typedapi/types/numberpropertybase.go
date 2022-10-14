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

// NumberPropertyBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/mapping/core.ts#L106-L124
type NumberPropertyBase struct {
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
}

// NumberPropertyBaseBuilder holds NumberPropertyBase struct and provides a builder API.
type NumberPropertyBaseBuilder struct {
	v *NumberPropertyBase
}

// NewNumberPropertyBase provides a builder for the NumberPropertyBase struct.
func NewNumberPropertyBaseBuilder() *NumberPropertyBaseBuilder {
	r := NumberPropertyBaseBuilder{
		&NumberPropertyBase{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NumberPropertyBase struct
func (rb *NumberPropertyBaseBuilder) Build() NumberPropertyBase {
	return *rb.v
}

func (rb *NumberPropertyBaseBuilder) Boost(boost float64) *NumberPropertyBaseBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *NumberPropertyBaseBuilder) Coerce(coerce bool) *NumberPropertyBaseBuilder {
	rb.v.Coerce = &coerce
	return rb
}

func (rb *NumberPropertyBaseBuilder) CopyTo(copyto *FieldsBuilder) *NumberPropertyBaseBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *NumberPropertyBaseBuilder) DocValues(docvalues bool) *NumberPropertyBaseBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *NumberPropertyBaseBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *NumberPropertyBaseBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *NumberPropertyBaseBuilder) Fields(values map[PropertyName]*PropertyBuilder) *NumberPropertyBaseBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *NumberPropertyBaseBuilder) IgnoreAbove(ignoreabove int) *NumberPropertyBaseBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *NumberPropertyBaseBuilder) IgnoreMalformed(ignoremalformed bool) *NumberPropertyBaseBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

func (rb *NumberPropertyBaseBuilder) Index(index bool) *NumberPropertyBaseBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *NumberPropertyBaseBuilder) LocalMetadata(localmetadata *MetadataBuilder) *NumberPropertyBaseBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

// Meta Metadata about the field.

func (rb *NumberPropertyBaseBuilder) Meta(value map[string]string) *NumberPropertyBaseBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *NumberPropertyBaseBuilder) OnScriptError(onscripterror onscripterror.OnScriptError) *NumberPropertyBaseBuilder {
	rb.v.OnScriptError = &onscripterror
	return rb
}

func (rb *NumberPropertyBaseBuilder) Properties(values map[PropertyName]*PropertyBuilder) *NumberPropertyBaseBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *NumberPropertyBaseBuilder) Script(script *ScriptBuilder) *NumberPropertyBaseBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *NumberPropertyBaseBuilder) Similarity(similarity string) *NumberPropertyBaseBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *NumberPropertyBaseBuilder) Store(store bool) *NumberPropertyBaseBuilder {
	rb.v.Store = &store
	return rb
}

// TimeSeriesDimension For internal use by Elastic only. Marks the field as a time series dimension.
// Defaults to false.

func (rb *NumberPropertyBaseBuilder) TimeSeriesDimension(timeseriesdimension bool) *NumberPropertyBaseBuilder {
	rb.v.TimeSeriesDimension = &timeseriesdimension
	return rb
}

// TimeSeriesMetric For internal use by Elastic only. Marks the field as a time series dimension.
// Defaults to false.

func (rb *NumberPropertyBaseBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *NumberPropertyBaseBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}
