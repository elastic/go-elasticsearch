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

// StatsBucketAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L246-L247
type StatsBucketAggregate struct {
	Avg         float64   `json:"avg,omitempty"`
	AvgAsString *string   `json:"avg_as_string,omitempty"`
	Count       int64     `json:"count"`
	Max         float64   `json:"max,omitempty"`
	MaxAsString *string   `json:"max_as_string,omitempty"`
	Meta        *Metadata `json:"meta,omitempty"`
	Min         float64   `json:"min,omitempty"`
	MinAsString *string   `json:"min_as_string,omitempty"`
	Sum         float64   `json:"sum"`
	SumAsString *string   `json:"sum_as_string,omitempty"`
}

// StatsBucketAggregateBuilder holds StatsBucketAggregate struct and provides a builder API.
type StatsBucketAggregateBuilder struct {
	v *StatsBucketAggregate
}

// NewStatsBucketAggregate provides a builder for the StatsBucketAggregate struct.
func NewStatsBucketAggregateBuilder() *StatsBucketAggregateBuilder {
	r := StatsBucketAggregateBuilder{
		&StatsBucketAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the StatsBucketAggregate struct
func (rb *StatsBucketAggregateBuilder) Build() StatsBucketAggregate {
	return *rb.v
}

func (rb *StatsBucketAggregateBuilder) Avg(avg float64) *StatsBucketAggregateBuilder {
	rb.v.Avg = avg
	return rb
}

func (rb *StatsBucketAggregateBuilder) AvgAsString(avgasstring string) *StatsBucketAggregateBuilder {
	rb.v.AvgAsString = &avgasstring
	return rb
}

func (rb *StatsBucketAggregateBuilder) Count(count int64) *StatsBucketAggregateBuilder {
	rb.v.Count = count
	return rb
}

func (rb *StatsBucketAggregateBuilder) Max(max float64) *StatsBucketAggregateBuilder {
	rb.v.Max = max
	return rb
}

func (rb *StatsBucketAggregateBuilder) MaxAsString(maxasstring string) *StatsBucketAggregateBuilder {
	rb.v.MaxAsString = &maxasstring
	return rb
}

func (rb *StatsBucketAggregateBuilder) Meta(meta *MetadataBuilder) *StatsBucketAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *StatsBucketAggregateBuilder) Min(min float64) *StatsBucketAggregateBuilder {
	rb.v.Min = min
	return rb
}

func (rb *StatsBucketAggregateBuilder) MinAsString(minasstring string) *StatsBucketAggregateBuilder {
	rb.v.MinAsString = &minasstring
	return rb
}

func (rb *StatsBucketAggregateBuilder) Sum(sum float64) *StatsBucketAggregateBuilder {
	rb.v.Sum = sum
	return rb
}

func (rb *StatsBucketAggregateBuilder) SumAsString(sumasstring string) *StatsBucketAggregateBuilder {
	rb.v.SumAsString = &sumasstring
	return rb
}
