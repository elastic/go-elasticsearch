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

// IndexingPressureMemorySummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L302-L311
type IndexingPressureMemorySummary struct {
	AllInBytes                            int64  `json:"all_in_bytes"`
	CombinedCoordinatingAndPrimaryInBytes int64  `json:"combined_coordinating_and_primary_in_bytes"`
	CoordinatingInBytes                   int64  `json:"coordinating_in_bytes"`
	CoordinatingRejections                *int64 `json:"coordinating_rejections,omitempty"`
	PrimaryInBytes                        int64  `json:"primary_in_bytes"`
	PrimaryRejections                     *int64 `json:"primary_rejections,omitempty"`
	ReplicaInBytes                        int64  `json:"replica_in_bytes"`
	ReplicaRejections                     *int64 `json:"replica_rejections,omitempty"`
}

// IndexingPressureMemorySummaryBuilder holds IndexingPressureMemorySummary struct and provides a builder API.
type IndexingPressureMemorySummaryBuilder struct {
	v *IndexingPressureMemorySummary
}

// NewIndexingPressureMemorySummary provides a builder for the IndexingPressureMemorySummary struct.
func NewIndexingPressureMemorySummaryBuilder() *IndexingPressureMemorySummaryBuilder {
	r := IndexingPressureMemorySummaryBuilder{
		&IndexingPressureMemorySummary{},
	}

	return &r
}

// Build finalize the chain and returns the IndexingPressureMemorySummary struct
func (rb *IndexingPressureMemorySummaryBuilder) Build() IndexingPressureMemorySummary {
	return *rb.v
}

func (rb *IndexingPressureMemorySummaryBuilder) AllInBytes(allinbytes int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.AllInBytes = allinbytes
	return rb
}

func (rb *IndexingPressureMemorySummaryBuilder) CombinedCoordinatingAndPrimaryInBytes(combinedcoordinatingandprimaryinbytes int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.CombinedCoordinatingAndPrimaryInBytes = combinedcoordinatingandprimaryinbytes
	return rb
}

func (rb *IndexingPressureMemorySummaryBuilder) CoordinatingInBytes(coordinatinginbytes int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.CoordinatingInBytes = coordinatinginbytes
	return rb
}

func (rb *IndexingPressureMemorySummaryBuilder) CoordinatingRejections(coordinatingrejections int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.CoordinatingRejections = &coordinatingrejections
	return rb
}

func (rb *IndexingPressureMemorySummaryBuilder) PrimaryInBytes(primaryinbytes int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.PrimaryInBytes = primaryinbytes
	return rb
}

func (rb *IndexingPressureMemorySummaryBuilder) PrimaryRejections(primaryrejections int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.PrimaryRejections = &primaryrejections
	return rb
}

func (rb *IndexingPressureMemorySummaryBuilder) ReplicaInBytes(replicainbytes int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.ReplicaInBytes = replicainbytes
	return rb
}

func (rb *IndexingPressureMemorySummaryBuilder) ReplicaRejections(replicarejections int64) *IndexingPressureMemorySummaryBuilder {
	rb.v.ReplicaRejections = &replicarejections
	return rb
}
