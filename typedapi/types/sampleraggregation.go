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

// SamplerAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L314-L316
type SamplerAggregation struct {
	Meta      *Metadata `json:"meta,omitempty"`
	Name      *string   `json:"name,omitempty"`
	ShardSize *int      `json:"shard_size,omitempty"`
}

// SamplerAggregationBuilder holds SamplerAggregation struct and provides a builder API.
type SamplerAggregationBuilder struct {
	v *SamplerAggregation
}

// NewSamplerAggregation provides a builder for the SamplerAggregation struct.
func NewSamplerAggregationBuilder() *SamplerAggregationBuilder {
	r := SamplerAggregationBuilder{
		&SamplerAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the SamplerAggregation struct
func (rb *SamplerAggregationBuilder) Build() SamplerAggregation {
	return *rb.v
}

func (rb *SamplerAggregationBuilder) Meta(meta *MetadataBuilder) *SamplerAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *SamplerAggregationBuilder) Name(name string) *SamplerAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *SamplerAggregationBuilder) ShardSize(shardsize int) *SamplerAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}
