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

// QueryCacheStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L150-L159
type QueryCacheStats struct {
	CacheCount        int       `json:"cache_count"`
	CacheSize         int       `json:"cache_size"`
	Evictions         int       `json:"evictions"`
	HitCount          int       `json:"hit_count"`
	MemorySize        *ByteSize `json:"memory_size,omitempty"`
	MemorySizeInBytes int       `json:"memory_size_in_bytes"`
	MissCount         int       `json:"miss_count"`
	TotalCount        int       `json:"total_count"`
}

// QueryCacheStatsBuilder holds QueryCacheStats struct and provides a builder API.
type QueryCacheStatsBuilder struct {
	v *QueryCacheStats
}

// NewQueryCacheStats provides a builder for the QueryCacheStats struct.
func NewQueryCacheStatsBuilder() *QueryCacheStatsBuilder {
	r := QueryCacheStatsBuilder{
		&QueryCacheStats{},
	}

	return &r
}

// Build finalize the chain and returns the QueryCacheStats struct
func (rb *QueryCacheStatsBuilder) Build() QueryCacheStats {
	return *rb.v
}

func (rb *QueryCacheStatsBuilder) CacheCount(cachecount int) *QueryCacheStatsBuilder {
	rb.v.CacheCount = cachecount
	return rb
}

func (rb *QueryCacheStatsBuilder) CacheSize(cachesize int) *QueryCacheStatsBuilder {
	rb.v.CacheSize = cachesize
	return rb
}

func (rb *QueryCacheStatsBuilder) Evictions(evictions int) *QueryCacheStatsBuilder {
	rb.v.Evictions = evictions
	return rb
}

func (rb *QueryCacheStatsBuilder) HitCount(hitcount int) *QueryCacheStatsBuilder {
	rb.v.HitCount = hitcount
	return rb
}

func (rb *QueryCacheStatsBuilder) MemorySize(memorysize *ByteSizeBuilder) *QueryCacheStatsBuilder {
	v := memorysize.Build()
	rb.v.MemorySize = &v
	return rb
}

func (rb *QueryCacheStatsBuilder) MemorySizeInBytes(memorysizeinbytes int) *QueryCacheStatsBuilder {
	rb.v.MemorySizeInBytes = memorysizeinbytes
	return rb
}

func (rb *QueryCacheStatsBuilder) MissCount(misscount int) *QueryCacheStatsBuilder {
	rb.v.MissCount = misscount
	return rb
}

func (rb *QueryCacheStatsBuilder) TotalCount(totalcount int) *QueryCacheStatsBuilder {
	rb.v.TotalCount = totalcount
	return rb
}
