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

// CompositeAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/aggregations/bucket.ts#L79-L84
type CompositeAggregation struct {
	After   *CompositeAggregateKey                  `json:"after,omitempty"`
	Meta    *Metadata                               `json:"meta,omitempty"`
	Name    *string                                 `json:"name,omitempty"`
	Size    *int                                    `json:"size,omitempty"`
	Sources []map[string]CompositeAggregationSource `json:"sources,omitempty"`
}

// CompositeAggregationBuilder holds CompositeAggregation struct and provides a builder API.
type CompositeAggregationBuilder struct {
	v *CompositeAggregation
}

// NewCompositeAggregation provides a builder for the CompositeAggregation struct.
func NewCompositeAggregationBuilder() *CompositeAggregationBuilder {
	r := CompositeAggregationBuilder{
		&CompositeAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the CompositeAggregation struct
func (rb *CompositeAggregationBuilder) Build() CompositeAggregation {
	return *rb.v
}

func (rb *CompositeAggregationBuilder) After(after *CompositeAggregateKeyBuilder) *CompositeAggregationBuilder {
	v := after.Build()
	rb.v.After = &v
	return rb
}

func (rb *CompositeAggregationBuilder) Meta(meta *MetadataBuilder) *CompositeAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *CompositeAggregationBuilder) Name(name string) *CompositeAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *CompositeAggregationBuilder) Size(size int) *CompositeAggregationBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *CompositeAggregationBuilder) Sources(value ...map[string]CompositeAggregationSource) *CompositeAggregationBuilder {
	rb.v.Sources = value
	return rb
}
