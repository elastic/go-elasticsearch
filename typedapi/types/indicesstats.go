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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexmetadatastate"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// IndicesStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/indices/stats/types.ts#L92-L101
type IndicesStats struct {
	Health    *healthstatus.HealthStatus             `json:"health,omitempty"`
	Primaries *IndexStats                            `json:"primaries,omitempty"`
	Shards    map[string][]IndicesShardStats         `json:"shards,omitempty"`
	Status    *indexmetadatastate.IndexMetadataState `json:"status,omitempty"`
	Total     *IndexStats                            `json:"total,omitempty"`
	Uuid      *string                                `json:"uuid,omitempty"`
}

func (s *IndicesStats) UnmarshalJSON(data []byte) error {

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

		case "health":
			if err := dec.Decode(&s.Health); err != nil {
				return err
			}

		case "primaries":
			if err := dec.Decode(&s.Primaries); err != nil {
				return err
			}

		case "shards":
			if s.Shards == nil {
				s.Shards = make(map[string][]IndicesShardStats, 0)
			}
			if err := dec.Decode(&s.Shards); err != nil {
				return err
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return err
			}

		case "total":
			if err := dec.Decode(&s.Total); err != nil {
				return err
			}

		case "uuid":
			if err := dec.Decode(&s.Uuid); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIndicesStats returns a IndicesStats.
func NewIndicesStats() *IndicesStats {
	r := &IndicesStats{
		Shards: make(map[string][]IndicesShardStats, 0),
	}

	return r
}
