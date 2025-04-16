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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// SnapshotShardFailure type.
//
// https://github.com/elastic/elasticsearch-specification/blob/52c473efb1fb5320a5bac12572d0b285882862fb/specification/snapshot/_types/SnapshotShardFailure.ts#L22-L29
type SnapshotShardFailure struct {
	Index     string  `json:"index"`
	IndexUuid string  `json:"index_uuid"`
	NodeId    *string `json:"node_id,omitempty"`
	Reason    string  `json:"reason"`
	ShardId   string  `json:"shard_id"`
	Status    string  `json:"status"`
}

func (s *SnapshotShardFailure) UnmarshalJSON(data []byte) error {

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

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "index_uuid":
			if err := dec.Decode(&s.IndexUuid); err != nil {
				return fmt.Errorf("%s | %w", "IndexUuid", err)
			}

		case "node_id":
			if err := dec.Decode(&s.NodeId); err != nil {
				return fmt.Errorf("%s | %w", "NodeId", err)
			}

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Reason", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = o

		case "shard_id":
			if err := dec.Decode(&s.ShardId); err != nil {
				return fmt.Errorf("%s | %w", "ShardId", err)
			}

		case "status":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Status = o

		}
	}
	return nil
}

// NewSnapshotShardFailure returns a SnapshotShardFailure.
func NewSnapshotShardFailure() *SnapshotShardFailure {
	r := &SnapshotShardFailure{}

	return r
}

// false
