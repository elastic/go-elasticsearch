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

// PercentilesAggregateBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L135-L137
type PercentilesAggregateBase struct {
	Meta   *Metadata   `json:"meta,omitempty"`
	Values Percentiles `json:"values"`
}

// PercentilesAggregateBaseBuilder holds PercentilesAggregateBase struct and provides a builder API.
type PercentilesAggregateBaseBuilder struct {
	v *PercentilesAggregateBase
}

// NewPercentilesAggregateBase provides a builder for the PercentilesAggregateBase struct.
func NewPercentilesAggregateBaseBuilder() *PercentilesAggregateBaseBuilder {
	r := PercentilesAggregateBaseBuilder{
		&PercentilesAggregateBase{},
	}

	return &r
}

// Build finalize the chain and returns the PercentilesAggregateBase struct
func (rb *PercentilesAggregateBaseBuilder) Build() PercentilesAggregateBase {
	return *rb.v
}

func (rb *PercentilesAggregateBaseBuilder) Meta(meta *MetadataBuilder) *PercentilesAggregateBaseBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *PercentilesAggregateBaseBuilder) Values(values *PercentilesBuilder) *PercentilesAggregateBaseBuilder {
	v := values.Build()
	rb.v.Values = v
	return rb
}
