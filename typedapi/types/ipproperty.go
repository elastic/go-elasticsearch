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
)

// IpProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/mapping/specialized.ts#L59-L72
type IpProperty struct {
	Boost           *float64                       `json:"boost,omitempty"`
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
	NullValue     *string                      `json:"null_value,omitempty"`
	OnScriptError *onscripterror.OnScriptError `json:"on_script_error,omitempty"`
	Properties    map[PropertyName]Property    `json:"properties,omitempty"`
	Script        *Script                      `json:"script,omitempty"`
	Similarity    *string                      `json:"similarity,omitempty"`
	Store         *bool                        `json:"store,omitempty"`
	// TimeSeriesDimension For internal use by Elastic only. Marks the field as a time series dimension.
	// Defaults to false.
	TimeSeriesDimension *bool  `json:"time_series_dimension,omitempty"`
	Type                string `json:"type,omitempty"`
}

// IpPropertyBuilder holds IpProperty struct and provides a builder API.
type IpPropertyBuilder struct {
	v *IpProperty
}

// NewIpProperty provides a builder for the IpProperty struct.
func NewIpPropertyBuilder() *IpPropertyBuilder {
	r := IpPropertyBuilder{
		&IpProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "ip"

	return &r
}

// Build finalize the chain and returns the IpProperty struct
func (rb *IpPropertyBuilder) Build() IpProperty {
	return *rb.v
}

func (rb *IpPropertyBuilder) Boost(boost float64) *IpPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *IpPropertyBuilder) CopyTo(copyto *FieldsBuilder) *IpPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *IpPropertyBuilder) DocValues(docvalues bool) *IpPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *IpPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *IpPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *IpPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *IpPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *IpPropertyBuilder) IgnoreAbove(ignoreabove int) *IpPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *IpPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *IpPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

func (rb *IpPropertyBuilder) Index(index bool) *IpPropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *IpPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *IpPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

// Meta Metadata about the field.

func (rb *IpPropertyBuilder) Meta(value map[string]string) *IpPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *IpPropertyBuilder) NullValue(nullvalue string) *IpPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

func (rb *IpPropertyBuilder) OnScriptError(onscripterror onscripterror.OnScriptError) *IpPropertyBuilder {
	rb.v.OnScriptError = &onscripterror
	return rb
}

func (rb *IpPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *IpPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *IpPropertyBuilder) Script(script *ScriptBuilder) *IpPropertyBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *IpPropertyBuilder) Similarity(similarity string) *IpPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *IpPropertyBuilder) Store(store bool) *IpPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// TimeSeriesDimension For internal use by Elastic only. Marks the field as a time series dimension.
// Defaults to false.

func (rb *IpPropertyBuilder) TimeSeriesDimension(timeseriesdimension bool) *IpPropertyBuilder {
	rb.v.TimeSeriesDimension = &timeseriesdimension
	return rb
}
