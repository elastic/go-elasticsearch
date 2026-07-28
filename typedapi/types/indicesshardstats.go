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
// https://github.com/elastic/elasticsearch-specification/tree/7560c979602e6941815872bdaec801200bc7ec4e

package types

import (
	"encoding/json"
)

// IndicesShardStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7560c979602e6941815872bdaec801200bc7ec4e/specification/indices/stats/types.ts#L197-L230
type IndicesShardStats struct {
	Bulk            *BulkStats                   `json:"bulk,omitempty"`
	Commit          *ShardCommit                 `json:"commit,omitempty"`
	Completion      *CompletionStats             `json:"completion,omitempty"`
	DenseVector     *DenseVectorStats            `json:"dense_vector,omitempty"`
	Docs            *DocStats                    `json:"docs,omitempty"`
	Fielddata       *FielddataStats              `json:"fielddata,omitempty"`
	Flush           *FlushStats                  `json:"flush,omitempty"`
	Get             *GetStats                    `json:"get,omitempty"`
	Indexing        *IndexingStats               `json:"indexing,omitempty"`
	Indices         map[string]IndicesShardStats `json:"indices,omitempty"`
	Mappings        *MappingStats                `json:"mappings,omitempty"`
	Merges          *MergesStats                 `json:"merges,omitempty"`
	QueryCache      *ShardQueryCache             `json:"query_cache,omitempty"`
	Recovery        *RecoveryStats               `json:"recovery,omitempty"`
	Refresh         *RefreshStats                `json:"refresh,omitempty"`
	RequestCache    *RequestCacheStats           `json:"request_cache,omitempty"`
	RetentionLeases *ShardRetentionLeases        `json:"retention_leases,omitempty"`
	Routing         *ShardRouting                `json:"routing,omitempty"`
	Search          *SearchStats                 `json:"search,omitempty"`
	Segments        *SegmentsStats               `json:"segments,omitempty"`
	SeqNo           *ShardSequenceNumber         `json:"seq_no,omitempty"`
	ShardPath       *ShardPath                   `json:"shard_path,omitempty"`
	ShardStats      *ShardsTotalStats            `json:"shard_stats,omitempty"`
	Shards          map[string]json.RawMessage   `json:"shards,omitempty"`
	SparseVector    *SparseVectorStats           `json:"sparse_vector,omitempty"`
	Store           *StoreStats                  `json:"store,omitempty"`
	Translog        *TranslogStats               `json:"translog,omitempty"`
	Warmer          *WarmerStats                 `json:"warmer,omitempty"`
}

// NewIndicesShardStats returns a IndicesShardStats.
func NewIndicesShardStats() *IndicesShardStats {
	r := &IndicesShardStats{
		Indices: make(map[string]IndicesShardStats),
		Shards:  make(map[string]json.RawMessage),
	}

	return r
}
