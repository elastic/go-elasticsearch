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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// IndicesShardStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/indices/stats/types.ts#L183-L211
type IndicesShardStats struct {
	Bulk            *BulkStats                 `json:"bulk,omitempty"`
	Commit          *ShardCommit               `json:"commit,omitempty"`
	Completion      *CompletionStats           `json:"completion,omitempty"`
	Docs            *DocStats                  `json:"docs,omitempty"`
	Fielddata       *FielddataStats            `json:"fielddata,omitempty"`
	Flush           *FlushStats                `json:"flush,omitempty"`
	Get             *GetStats                  `json:"get,omitempty"`
	Indexing        *IndexingStats             `json:"indexing,omitempty"`
	Indices         *IndicesStats              `json:"indices,omitempty"`
	Mappings        *MappingStats              `json:"mappings,omitempty"`
	Merges          *MergesStats               `json:"merges,omitempty"`
	QueryCache      *ShardQueryCache           `json:"query_cache,omitempty"`
	Recovery        *RecoveryStats             `json:"recovery,omitempty"`
	Refresh         *RefreshStats              `json:"refresh,omitempty"`
	RequestCache    *RequestCacheStats         `json:"request_cache,omitempty"`
	RetentionLeases *ShardRetentionLeases      `json:"retention_leases,omitempty"`
	Routing         *ShardRouting              `json:"routing,omitempty"`
	Search          *SearchStats               `json:"search,omitempty"`
	Segments        *SegmentsStats             `json:"segments,omitempty"`
	SeqNo           *ShardSequenceNumber       `json:"seq_no,omitempty"`
	ShardPath       *ShardPath                 `json:"shard_path,omitempty"`
	ShardStats      *ShardsTotalStats          `json:"shard_stats,omitempty"`
	Shards          map[string]json.RawMessage `json:"shards,omitempty"`
	Store           *StoreStats                `json:"store,omitempty"`
	Translog        *TranslogStats             `json:"translog,omitempty"`
	Warmer          *WarmerStats               `json:"warmer,omitempty"`
}

func (s *IndicesShardStats) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "bulk":
			if err := dec.Decode(&s.Bulk); err != nil {
				return err
			}

		case "commit":
			if err := dec.Decode(&s.Commit); err != nil {
				return err
			}

		case "completion":
			if err := dec.Decode(&s.Completion); err != nil {
				return err
			}

		case "docs":
			if err := dec.Decode(&s.Docs); err != nil {
				return err
			}

		case "fielddata":
			if err := dec.Decode(&s.Fielddata); err != nil {
				return err
			}

		case "flush":
			if err := dec.Decode(&s.Flush); err != nil {
				return err
			}

		case "get":
			if err := dec.Decode(&s.Get); err != nil {
				return err
			}

		case "indexing":
			if err := dec.Decode(&s.Indexing); err != nil {
				return err
			}

		case "indices":
			if err := dec.Decode(&s.Indices); err != nil {
				return err
			}

		case "mappings":
			if err := dec.Decode(&s.Mappings); err != nil {
				return err
			}

		case "merges":
			if err := dec.Decode(&s.Merges); err != nil {
				return err
			}

		case "query_cache":
			if err := dec.Decode(&s.QueryCache); err != nil {
				return err
			}

		case "recovery":
			if err := dec.Decode(&s.Recovery); err != nil {
				return err
			}

		case "refresh":
			if err := dec.Decode(&s.Refresh); err != nil {
				return err
			}

		case "request_cache":
			if err := dec.Decode(&s.RequestCache); err != nil {
				return err
			}

		case "retention_leases":
			if err := dec.Decode(&s.RetentionLeases); err != nil {
				return err
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return err
			}

		case "search":
			if err := dec.Decode(&s.Search); err != nil {
				return err
			}

		case "segments":
			if err := dec.Decode(&s.Segments); err != nil {
				return err
			}

		case "seq_no":
			if err := dec.Decode(&s.SeqNo); err != nil {
				return err
			}

		case "shard_path":
			if err := dec.Decode(&s.ShardPath); err != nil {
				return err
			}

		case "shard_stats":
			if err := dec.Decode(&s.ShardStats); err != nil {
				return err
			}

		case "shards":
			if s.Shards == nil {
				s.Shards = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Shards); err != nil {
				return err
			}

		case "store":
			if err := dec.Decode(&s.Store); err != nil {
				return err
			}

		case "translog":
			if err := dec.Decode(&s.Translog); err != nil {
				return err
			}

		case "warmer":
			if err := dec.Decode(&s.Warmer); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIndicesShardStats returns a IndicesShardStats.
func NewIndicesShardStats() *IndicesShardStats {
	r := &IndicesShardStats{
		Shards: make(map[string]json.RawMessage, 0),
	}

	return r
}
