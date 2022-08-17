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

// ShardStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L33-L39
type ShardStatistics struct {
	Failed     uint           `json:"failed"`
	Failures   []ShardFailure `json:"failures,omitempty"`
	Skipped    *uint          `json:"skipped,omitempty"`
	Successful uint           `json:"successful"`
	Total      uint           `json:"total"`
}

// ShardStatisticsBuilder holds ShardStatistics struct and provides a builder API.
type ShardStatisticsBuilder struct {
	v *ShardStatistics
}

// NewShardStatistics provides a builder for the ShardStatistics struct.
func NewShardStatisticsBuilder() *ShardStatisticsBuilder {
	r := ShardStatisticsBuilder{
		&ShardStatistics{},
	}

	return &r
}

// Build finalize the chain and returns the ShardStatistics struct
func (rb *ShardStatisticsBuilder) Build() ShardStatistics {
	return *rb.v
}

func (rb *ShardStatisticsBuilder) Failed(failed uint) *ShardStatisticsBuilder {
	rb.v.Failed = failed
	return rb
}

func (rb *ShardStatisticsBuilder) Failures(failures []ShardFailureBuilder) *ShardStatisticsBuilder {
	tmp := make([]ShardFailure, len(failures))
	for _, value := range failures {
		tmp = append(tmp, value.Build())
	}
	rb.v.Failures = tmp
	return rb
}

func (rb *ShardStatisticsBuilder) Skipped(skipped uint) *ShardStatisticsBuilder {
	rb.v.Skipped = &skipped
	return rb
}

func (rb *ShardStatisticsBuilder) Successful(successful uint) *ShardStatisticsBuilder {
	rb.v.Successful = successful
	return rb
}

func (rb *ShardStatisticsBuilder) Total(total uint) *ShardStatisticsBuilder {
	rb.v.Total = total
	return rb
}
