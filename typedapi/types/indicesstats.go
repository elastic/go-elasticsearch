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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/healthstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexmetadatastate"
)

// IndicesStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/stats/types.ts#L95-L110
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
				return fmt.Errorf("%s | %w", "Health", err)
			}

		case "primaries":
			if err := dec.Decode(&s.Primaries); err != nil {
				return fmt.Errorf("%s | %w", "Primaries", err)
			}

		case "shards":
			if s.Shards == nil {
				s.Shards = make(map[string][]IndicesShardStats, 0)
			}
			if err := dec.Decode(&s.Shards); err != nil {
				return fmt.Errorf("%s | %w", "Shards", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "total":
			if err := dec.Decode(&s.Total); err != nil {
				return fmt.Errorf("%s | %w", "Total", err)
			}

		case "uuid":
			if err := dec.Decode(&s.Uuid); err != nil {
				return fmt.Errorf("%s | %w", "Uuid", err)
			}

		}
	}
	return nil
}

// NewIndicesStats returns a IndicesStats.
func NewIndicesStats() *IndicesStats {
	r := &IndicesStats{
		Shards: make(map[string][]IndicesShardStats),
	}

	return r
}
