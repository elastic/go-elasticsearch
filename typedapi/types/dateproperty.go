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

// DateProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/core.ts#L61-L71
type DateProperty struct {
	Boost           *float64                       `json:"boost,omitempty"`
	CopyTo          *Fields                        `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fielddata       *NumericFielddata              `json:"fielddata,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	Format          *string                        `json:"format,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	Index           *bool                          `json:"index,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	Locale          *string                        `json:"locale,omitempty"`
	Meta            map[string]string              `json:"meta,omitempty"`
	NullValue       *DateTime                      `json:"null_value,omitempty"`
	PrecisionStep   *int                           `json:"precision_step,omitempty"`
	Properties      map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity      *string                        `json:"similarity,omitempty"`
	Store           *bool                          `json:"store,omitempty"`
	Type            string                         `json:"type,omitempty"`
}

// DatePropertyBuilder holds DateProperty struct and provides a builder API.
type DatePropertyBuilder struct {
	v *DateProperty
}

// NewDateProperty provides a builder for the DateProperty struct.
func NewDatePropertyBuilder() *DatePropertyBuilder {
	r := DatePropertyBuilder{
		&DateProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "date"

	return &r
}

// Build finalize the chain and returns the DateProperty struct
func (rb *DatePropertyBuilder) Build() DateProperty {
	return *rb.v
}

func (rb *DatePropertyBuilder) Boost(boost float64) *DatePropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *DatePropertyBuilder) CopyTo(copyto *FieldsBuilder) *DatePropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *DatePropertyBuilder) DocValues(docvalues bool) *DatePropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *DatePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *DatePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *DatePropertyBuilder) Fielddata(fielddata *NumericFielddataBuilder) *DatePropertyBuilder {
	v := fielddata.Build()
	rb.v.Fielddata = &v
	return rb
}

func (rb *DatePropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *DatePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *DatePropertyBuilder) Format(format string) *DatePropertyBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *DatePropertyBuilder) IgnoreAbove(ignoreabove int) *DatePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *DatePropertyBuilder) IgnoreMalformed(ignoremalformed bool) *DatePropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

func (rb *DatePropertyBuilder) Index(index bool) *DatePropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *DatePropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *DatePropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *DatePropertyBuilder) Locale(locale string) *DatePropertyBuilder {
	rb.v.Locale = &locale
	return rb
}

func (rb *DatePropertyBuilder) Meta(value map[string]string) *DatePropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *DatePropertyBuilder) NullValue(nullvalue *DateTimeBuilder) *DatePropertyBuilder {
	v := nullvalue.Build()
	rb.v.NullValue = &v
	return rb
}

func (rb *DatePropertyBuilder) PrecisionStep(precisionstep int) *DatePropertyBuilder {
	rb.v.PrecisionStep = &precisionstep
	return rb
}

func (rb *DatePropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *DatePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *DatePropertyBuilder) Similarity(similarity string) *DatePropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *DatePropertyBuilder) Store(store bool) *DatePropertyBuilder {
	rb.v.Store = &store
	return rb
}
