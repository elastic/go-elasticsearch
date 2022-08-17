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

// ShardsStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/_types/SnapshotShardsStats.ts#L22-L29
type ShardsStats struct {
	Done         int64 `json:"done"`
	Failed       int64 `json:"failed"`
	Finalizing   int64 `json:"finalizing"`
	Initializing int64 `json:"initializing"`
	Started      int64 `json:"started"`
	Total        int64 `json:"total"`
}

// ShardsStatsBuilder holds ShardsStats struct and provides a builder API.
type ShardsStatsBuilder struct {
	v *ShardsStats
}

// NewShardsStats provides a builder for the ShardsStats struct.
func NewShardsStatsBuilder() *ShardsStatsBuilder {
	r := ShardsStatsBuilder{
		&ShardsStats{},
	}

	return &r
}

// Build finalize the chain and returns the ShardsStats struct
func (rb *ShardsStatsBuilder) Build() ShardsStats {
	return *rb.v
}

func (rb *ShardsStatsBuilder) Done(done int64) *ShardsStatsBuilder {
	rb.v.Done = done
	return rb
}

func (rb *ShardsStatsBuilder) Failed(failed int64) *ShardsStatsBuilder {
	rb.v.Failed = failed
	return rb
}

func (rb *ShardsStatsBuilder) Finalizing(finalizing int64) *ShardsStatsBuilder {
	rb.v.Finalizing = finalizing
	return rb
}

func (rb *ShardsStatsBuilder) Initializing(initializing int64) *ShardsStatsBuilder {
	rb.v.Initializing = initializing
	return rb
}

func (rb *ShardsStatsBuilder) Started(started int64) *ShardsStatsBuilder {
	rb.v.Started = started
	return rb
}

func (rb *ShardsStatsBuilder) Total(total int64) *ShardsStatsBuilder {
	rb.v.Total = total
	return rb
}
