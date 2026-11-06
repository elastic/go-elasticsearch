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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// SearchableSnapshotAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/ilm/_types/Phase.ts#L131-L149
type SearchableSnapshotAction struct {
	ForceMergeIndex *bool `json:"force_merge_index,omitempty"`
	// ReplicateFor The period for which the mounted index keeps one replica before the replica
	// is removed. By default, the index is mounted without replicas. Must be
	// positive. Only the first `searchable_snapshot` action in a policy may set
	// this. It must fit before the next phase starts: for example, if this phase
	// has a `min_age` of 5d and the next phase has 12d, it can be at most 7d.
	ReplicateFor       Duration `json:"replicate_for,omitempty"`
	SnapshotRepository string   `json:"snapshot_repository"`
	// TotalShardsPerNode The maximum number of shards, including primaries and replicas, that can be
	// allocated to any single node for the mounted index. Must be at least 1.
	TotalShardsPerNode *int `json:"total_shards_per_node,omitempty"`
}

func (s *SearchableSnapshotAction) UnmarshalJSON(data []byte) error {

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

		case "force_merge_index":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ForceMergeIndex", err)
				}
				s.ForceMergeIndex = &value
			case bool:
				s.ForceMergeIndex = &v
			}

		case "replicate_for":
			if err := dec.Decode(&s.ReplicateFor); err != nil {
				return fmt.Errorf("%s | %w", "ReplicateFor", err)
			}

		case "snapshot_repository":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "SnapshotRepository", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SnapshotRepository = o

		case "total_shards_per_node":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalShardsPerNode", err)
				}
				s.TotalShardsPerNode = &value
			case float64:
				f := int(v)
				s.TotalShardsPerNode = &f
			}

		}
	}
	return nil
}

// NewSearchableSnapshotAction returns a SearchableSnapshotAction.
func NewSearchableSnapshotAction() *SearchableSnapshotAction {
	r := &SearchableSnapshotAction{}

	return r
}

type SearchableSnapshotActionVariant interface {
	SearchableSnapshotActionCaster() *SearchableSnapshotAction
}

func (s *SearchableSnapshotAction) SearchableSnapshotActionCaster() *SearchableSnapshotAction {
	return s
}
