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

// MemStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/get_memory_stats/types.ts#L65-L88
type MemStats struct {
	// AdjustedTotal If the amount of physical memory has been overridden using the
	// es.total_memory_bytes system property
	// then this reports the overridden value. Otherwise it reports the same value
	// as total.
	AdjustedTotal *ByteSize `json:"adjusted_total,omitempty"`
	// AdjustedTotalInBytes If the amount of physical memory has been overridden using the
	// `es.total_memory_bytes` system property
	// then this reports the overridden value in bytes. Otherwise it reports the
	// same value as `total_in_bytes`.
	AdjustedTotalInBytes int `json:"adjusted_total_in_bytes"`
	// Ml Contains statistics about machine learning use of native memory on the node.
	Ml MemMlStats `json:"ml"`
	// Total Total amount of physical memory.
	Total *ByteSize `json:"total,omitempty"`
	// TotalInBytes Total amount of physical memory in bytes.
	TotalInBytes int `json:"total_in_bytes"`
}

// MemStatsBuilder holds MemStats struct and provides a builder API.
type MemStatsBuilder struct {
	v *MemStats
}

// NewMemStats provides a builder for the MemStats struct.
func NewMemStatsBuilder() *MemStatsBuilder {
	r := MemStatsBuilder{
		&MemStats{},
	}

	return &r
}

// Build finalize the chain and returns the MemStats struct
func (rb *MemStatsBuilder) Build() MemStats {
	return *rb.v
}

// AdjustedTotal If the amount of physical memory has been overridden using the
// es.total_memory_bytes system property
// then this reports the overridden value. Otherwise it reports the same value
// as total.

func (rb *MemStatsBuilder) AdjustedTotal(adjustedtotal *ByteSizeBuilder) *MemStatsBuilder {
	v := adjustedtotal.Build()
	rb.v.AdjustedTotal = &v
	return rb
}

// AdjustedTotalInBytes If the amount of physical memory has been overridden using the
// `es.total_memory_bytes` system property
// then this reports the overridden value in bytes. Otherwise it reports the
// same value as `total_in_bytes`.

func (rb *MemStatsBuilder) AdjustedTotalInBytes(adjustedtotalinbytes int) *MemStatsBuilder {
	rb.v.AdjustedTotalInBytes = adjustedtotalinbytes
	return rb
}

// Ml Contains statistics about machine learning use of native memory on the node.

func (rb *MemStatsBuilder) Ml(ml *MemMlStatsBuilder) *MemStatsBuilder {
	v := ml.Build()
	rb.v.Ml = v
	return rb
}

// Total Total amount of physical memory.

func (rb *MemStatsBuilder) Total(total *ByteSizeBuilder) *MemStatsBuilder {
	v := total.Build()
	rb.v.Total = &v
	return rb
}

// TotalInBytes Total amount of physical memory in bytes.

func (rb *MemStatsBuilder) TotalInBytes(totalinbytes int) *MemStatsBuilder {
	rb.v.TotalInBytes = totalinbytes
	return rb
}
