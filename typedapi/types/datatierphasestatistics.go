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

// DataTierPhaseStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L81-L92
type DataTierPhaseStatistics struct {
	DocCount                    int64 `json:"doc_count"`
	IndexCount                  int64 `json:"index_count"`
	NodeCount                   int64 `json:"node_count"`
	PrimaryShardCount           int64 `json:"primary_shard_count"`
	PrimaryShardSizeAvgBytes    int64 `json:"primary_shard_size_avg_bytes"`
	PrimaryShardSizeMadBytes    int64 `json:"primary_shard_size_mad_bytes"`
	PrimaryShardSizeMedianBytes int64 `json:"primary_shard_size_median_bytes"`
	PrimarySizeBytes            int64 `json:"primary_size_bytes"`
	TotalShardCount             int64 `json:"total_shard_count"`
	TotalSizeBytes              int64 `json:"total_size_bytes"`
}

// DataTierPhaseStatisticsBuilder holds DataTierPhaseStatistics struct and provides a builder API.
type DataTierPhaseStatisticsBuilder struct {
	v *DataTierPhaseStatistics
}

// NewDataTierPhaseStatistics provides a builder for the DataTierPhaseStatistics struct.
func NewDataTierPhaseStatisticsBuilder() *DataTierPhaseStatisticsBuilder {
	r := DataTierPhaseStatisticsBuilder{
		&DataTierPhaseStatistics{},
	}

	return &r
}

// Build finalize the chain and returns the DataTierPhaseStatistics struct
func (rb *DataTierPhaseStatisticsBuilder) Build() DataTierPhaseStatistics {
	return *rb.v
}

func (rb *DataTierPhaseStatisticsBuilder) DocCount(doccount int64) *DataTierPhaseStatisticsBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *DataTierPhaseStatisticsBuilder) IndexCount(indexcount int64) *DataTierPhaseStatisticsBuilder {
	rb.v.IndexCount = indexcount
	return rb
}

func (rb *DataTierPhaseStatisticsBuilder) NodeCount(nodecount int64) *DataTierPhaseStatisticsBuilder {
	rb.v.NodeCount = nodecount
	return rb
}

func (rb *DataTierPhaseStatisticsBuilder) PrimaryShardCount(primaryshardcount int64) *DataTierPhaseStatisticsBuilder {
	rb.v.PrimaryShardCount = primaryshardcount
	return rb
}

func (rb *DataTierPhaseStatisticsBuilder) PrimaryShardSizeAvgBytes(primaryshardsizeavgbytes int64) *DataTierPhaseStatisticsBuilder {
	rb.v.PrimaryShardSizeAvgBytes = primaryshardsizeavgbytes
	return rb
}

func (rb *DataTierPhaseStatisticsBuilder) PrimaryShardSizeMadBytes(primaryshardsizemadbytes int64) *DataTierPhaseStatisticsBuilder {
	rb.v.PrimaryShardSizeMadBytes = primaryshardsizemadbytes
	return rb
}

func (rb *DataTierPhaseStatisticsBuilder) PrimaryShardSizeMedianBytes(primaryshardsizemedianbytes int64) *DataTierPhaseStatisticsBuilder {
	rb.v.PrimaryShardSizeMedianBytes = primaryshardsizemedianbytes
	return rb
}

func (rb *DataTierPhaseStatisticsBuilder) PrimarySizeBytes(primarysizebytes int64) *DataTierPhaseStatisticsBuilder {
	rb.v.PrimarySizeBytes = primarysizebytes
	return rb
}

func (rb *DataTierPhaseStatisticsBuilder) TotalShardCount(totalshardcount int64) *DataTierPhaseStatisticsBuilder {
	rb.v.TotalShardCount = totalshardcount
	return rb
}

func (rb *DataTierPhaseStatisticsBuilder) TotalSizeBytes(totalsizebytes int64) *DataTierPhaseStatisticsBuilder {
	rb.v.TotalSizeBytes = totalsizebytes
	return rb
}
