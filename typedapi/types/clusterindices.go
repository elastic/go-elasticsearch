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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ClusterIndices type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L100-L147
type ClusterIndices struct {
	// Analysis Contains statistics about analyzers and analyzer components used in selected
	// nodes.
	Analysis *CharFilterTypes `json:"analysis,omitempty"`
	// Completion Contains statistics about memory used for completion in selected nodes.
	Completion CompletionStats `json:"completion"`
	// Count Total number of indices with shards assigned to selected nodes.
	Count int64 `json:"count"`
	// DenseVector Contains statistics about indexed dense vector
	DenseVector DenseVectorStats `json:"dense_vector"`
	// Docs Contains counts for documents in selected nodes.
	Docs DocStats `json:"docs"`
	// Fielddata Contains statistics about the field data cache of selected nodes.
	Fielddata FielddataStats `json:"fielddata"`
	// Mappings Contains statistics about field mappings in selected nodes.
	Mappings *FieldTypesMappings `json:"mappings,omitempty"`
	// QueryCache Contains statistics about the query cache of selected nodes.
	QueryCache QueryCacheStats `json:"query_cache"`
	// Search Holds a snapshot of the search usage statistics.
	// Used to hold the stats for a single node that's part of a
	// ClusterStatsNodeResponse, as well as to
	// accumulate stats for the entire cluster and return them as part of the
	// ClusterStatsResponse.
	Search SearchUsageStats `json:"search"`
	// Segments Contains statistics about segments in selected nodes.
	Segments SegmentsStats `json:"segments"`
	// Shards Contains statistics about indices with shards assigned to selected nodes.
	Shards ClusterIndicesShards `json:"shards"`
	// SparseVector Contains statistics about indexed sparse vector
	SparseVector SparseVectorStats `json:"sparse_vector"`
	// Store Contains statistics about the size of shards assigned to selected nodes.
	Store StoreStats `json:"store"`
	// Versions Contains statistics about analyzers and analyzer components used in selected
	// nodes.
	Versions []IndicesVersions `json:"versions,omitempty"`
}

func (s *ClusterIndices) UnmarshalJSON(data []byte) error {

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

		case "analysis":
			if err := dec.Decode(&s.Analysis); err != nil {
				return fmt.Errorf("%s | %w", "Analysis", err)
			}

		case "completion":
			if err := dec.Decode(&s.Completion); err != nil {
				return fmt.Errorf("%s | %w", "Completion", err)
			}

		case "count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "dense_vector":
			if err := dec.Decode(&s.DenseVector); err != nil {
				return fmt.Errorf("%s | %w", "DenseVector", err)
			}

		case "docs":
			if err := dec.Decode(&s.Docs); err != nil {
				return fmt.Errorf("%s | %w", "Docs", err)
			}

		case "fielddata":
			if err := dec.Decode(&s.Fielddata); err != nil {
				return fmt.Errorf("%s | %w", "Fielddata", err)
			}

		case "mappings":
			if err := dec.Decode(&s.Mappings); err != nil {
				return fmt.Errorf("%s | %w", "Mappings", err)
			}

		case "query_cache":
			if err := dec.Decode(&s.QueryCache); err != nil {
				return fmt.Errorf("%s | %w", "QueryCache", err)
			}

		case "search":
			if err := dec.Decode(&s.Search); err != nil {
				return fmt.Errorf("%s | %w", "Search", err)
			}

		case "segments":
			if err := dec.Decode(&s.Segments); err != nil {
				return fmt.Errorf("%s | %w", "Segments", err)
			}

		case "shards":
			if err := dec.Decode(&s.Shards); err != nil {
				return fmt.Errorf("%s | %w", "Shards", err)
			}

		case "sparse_vector":
			if err := dec.Decode(&s.SparseVector); err != nil {
				return fmt.Errorf("%s | %w", "SparseVector", err)
			}

		case "store":
			if err := dec.Decode(&s.Store); err != nil {
				return fmt.Errorf("%s | %w", "Store", err)
			}

		case "versions":
			if err := dec.Decode(&s.Versions); err != nil {
				return fmt.Errorf("%s | %w", "Versions", err)
			}

		}
	}
	return nil
}

// NewClusterIndices returns a ClusterIndices.
func NewClusterIndices() *ClusterIndices {
	r := &ClusterIndices{}

	return r
}
