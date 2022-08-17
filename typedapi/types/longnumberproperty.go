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

// LongNumberProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/core.ts#L139-L142
type LongNumberProperty struct {
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
	NullValue        *int64                                     `json:"null_value,omitempty"`
	OnScriptError    *onscripterror.OnScriptError               `json:"on_script_error,omitempty"`
	Properties       map[PropertyName]Property                  `json:"properties,omitempty"`
	Script           *Script                                    `json:"script,omitempty"`
	Similarity       *string                                    `json:"similarity,omitempty"`
	Store            *bool                                      `json:"store,omitempty"`
	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type,omitempty"`
}

// LongNumberPropertyBuilder holds LongNumberProperty struct and provides a builder API.
type LongNumberPropertyBuilder struct {
	v *LongNumberProperty
}

// NewLongNumberProperty provides a builder for the LongNumberProperty struct.
func NewLongNumberPropertyBuilder() *LongNumberPropertyBuilder {
	r := LongNumberPropertyBuilder{
		&LongNumberProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "long"

	return &r
}

// Build finalize the chain and returns the LongNumberProperty struct
func (rb *LongNumberPropertyBuilder) Build() LongNumberProperty {
	return *rb.v
}

func (rb *LongNumberPropertyBuilder) Coerce(coerce bool) *LongNumberPropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

func (rb *LongNumberPropertyBuilder) CopyTo(copyto *FieldsBuilder) *LongNumberPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *LongNumberPropertyBuilder) DocValues(docvalues bool) *LongNumberPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *LongNumberPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *LongNumberPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *LongNumberPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *LongNumberPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *LongNumberPropertyBuilder) IgnoreAbove(ignoreabove int) *LongNumberPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *LongNumberPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *LongNumberPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

func (rb *LongNumberPropertyBuilder) Index(index bool) *LongNumberPropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *LongNumberPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *LongNumberPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *LongNumberPropertyBuilder) Meta(value map[string]string) *LongNumberPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *LongNumberPropertyBuilder) NullValue(nullvalue int64) *LongNumberPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

func (rb *LongNumberPropertyBuilder) OnScriptError(onscripterror onscripterror.OnScriptError) *LongNumberPropertyBuilder {
	rb.v.OnScriptError = &onscripterror
	return rb
}

func (rb *LongNumberPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *LongNumberPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *LongNumberPropertyBuilder) Script(script *ScriptBuilder) *LongNumberPropertyBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *LongNumberPropertyBuilder) Similarity(similarity string) *LongNumberPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *LongNumberPropertyBuilder) Store(store bool) *LongNumberPropertyBuilder {
	rb.v.Store = &store
	return rb
}

func (rb *LongNumberPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *LongNumberPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}
