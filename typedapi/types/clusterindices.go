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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ClusterIndices type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cluster/stats/types.ts#L63-L94
type ClusterIndices struct {
	// Analysis Contains statistics about analyzers and analyzer components used in selected
	// nodes.
	Analysis CharFilterTypes `json:"analysis"`
	// Completion Contains statistics about memory used for completion in selected nodes.
	Completion CompletionStats `json:"completion"`
	// Count Total number of indices with shards assigned to selected nodes.
	Count int64 `json:"count"`
	// Docs Contains counts for documents in selected nodes.
	Docs DocStats `json:"docs"`
	// Fielddata Contains statistics about the field data cache of selected nodes.
	Fielddata FielddataStats `json:"fielddata"`
	// Mappings Contains statistics about field mappings in selected nodes.
	Mappings FieldTypesMappings `json:"mappings"`
	// QueryCache Contains statistics about the query cache of selected nodes.
	QueryCache QueryCacheStats `json:"query_cache"`
	// Segments Contains statistics about segments in selected nodes.
	Segments SegmentsStats `json:"segments"`
	// Shards Contains statistics about indices with shards assigned to selected nodes.
	Shards ClusterIndicesShards `json:"shards"`
	// Store Contains statistics about the size of shards assigned to selected nodes.
	Store    StoreStats        `json:"store"`
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
				return err
			}

		case "completion":
			if err := dec.Decode(&s.Completion); err != nil {
				return err
			}

		case "count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "docs":
			if err := dec.Decode(&s.Docs); err != nil {
				return err
			}

		case "fielddata":
			if err := dec.Decode(&s.Fielddata); err != nil {
				return err
			}

		case "mappings":
			if err := dec.Decode(&s.Mappings); err != nil {
				return err
			}

		case "query_cache":
			if err := dec.Decode(&s.QueryCache); err != nil {
				return err
			}

		case "segments":
			if err := dec.Decode(&s.Segments); err != nil {
				return err
			}

		case "shards":
			if err := dec.Decode(&s.Shards); err != nil {
				return err
			}

		case "store":
			if err := dec.Decode(&s.Store); err != nil {
				return err
			}

		case "versions":
			if err := dec.Decode(&s.Versions); err != nil {
				return err
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
