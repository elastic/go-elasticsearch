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

// RareTermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L300-L308
type RareTermsAggregation struct {
	Exclude     *TermsExclude `json:"exclude,omitempty"`
	Field       *Field        `json:"field,omitempty"`
	Include     *TermsInclude `json:"include,omitempty"`
	MaxDocCount *int64        `json:"max_doc_count,omitempty"`
	Meta        *Metadata     `json:"meta,omitempty"`
	Missing     *Missing      `json:"missing,omitempty"`
	Name        *string       `json:"name,omitempty"`
	Precision   *float64      `json:"precision,omitempty"`
	ValueType   *string       `json:"value_type,omitempty"`
}

// RareTermsAggregationBuilder holds RareTermsAggregation struct and provides a builder API.
type RareTermsAggregationBuilder struct {
	v *RareTermsAggregation
}

// NewRareTermsAggregation provides a builder for the RareTermsAggregation struct.
func NewRareTermsAggregationBuilder() *RareTermsAggregationBuilder {
	r := RareTermsAggregationBuilder{
		&RareTermsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the RareTermsAggregation struct
func (rb *RareTermsAggregationBuilder) Build() RareTermsAggregation {
	return *rb.v
}

func (rb *RareTermsAggregationBuilder) Exclude(exclude *TermsExcludeBuilder) *RareTermsAggregationBuilder {
	v := exclude.Build()
	rb.v.Exclude = &v
	return rb
}

func (rb *RareTermsAggregationBuilder) Field(field Field) *RareTermsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *RareTermsAggregationBuilder) Include(include *TermsIncludeBuilder) *RareTermsAggregationBuilder {
	v := include.Build()
	rb.v.Include = &v
	return rb
}

func (rb *RareTermsAggregationBuilder) MaxDocCount(maxdoccount int64) *RareTermsAggregationBuilder {
	rb.v.MaxDocCount = &maxdoccount
	return rb
}

func (rb *RareTermsAggregationBuilder) Meta(meta *MetadataBuilder) *RareTermsAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *RareTermsAggregationBuilder) Missing(missing *MissingBuilder) *RareTermsAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *RareTermsAggregationBuilder) Name(name string) *RareTermsAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *RareTermsAggregationBuilder) Precision(precision float64) *RareTermsAggregationBuilder {
	rb.v.Precision = &precision
	return rb
}

func (rb *RareTermsAggregationBuilder) ValueType(valuetype string) *RareTermsAggregationBuilder {
	rb.v.ValueType = &valuetype
	return rb
}
