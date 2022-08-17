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

// ShardProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/profile.ts#L130-L135
type ShardProfile struct {
	Aggregations []AggregationProfile `json:"aggregations"`
	Fetch        *FetchProfile        `json:"fetch,omitempty"`
	Id           string               `json:"id"`
	Searches     []SearchProfile      `json:"searches"`
}

// ShardProfileBuilder holds ShardProfile struct and provides a builder API.
type ShardProfileBuilder struct {
	v *ShardProfile
}

// NewShardProfile provides a builder for the ShardProfile struct.
func NewShardProfileBuilder() *ShardProfileBuilder {
	r := ShardProfileBuilder{
		&ShardProfile{},
	}

	return &r
}

// Build finalize the chain and returns the ShardProfile struct
func (rb *ShardProfileBuilder) Build() ShardProfile {
	return *rb.v
}

func (rb *ShardProfileBuilder) Aggregations(aggregations []AggregationProfileBuilder) *ShardProfileBuilder {
	tmp := make([]AggregationProfile, len(aggregations))
	for _, value := range aggregations {
		tmp = append(tmp, value.Build())
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *ShardProfileBuilder) Fetch(fetch *FetchProfileBuilder) *ShardProfileBuilder {
	v := fetch.Build()
	rb.v.Fetch = &v
	return rb
}

func (rb *ShardProfileBuilder) Id(id string) *ShardProfileBuilder {
	rb.v.Id = id
	return rb
}

func (rb *ShardProfileBuilder) Searches(searches []SearchProfileBuilder) *ShardProfileBuilder {
	tmp := make([]SearchProfile, len(searches))
	for _, value := range searches {
		tmp = append(tmp, value.Build())
	}
	rb.v.Searches = tmp
	return rb
}
