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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
)

// IndexHealthStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/health/types.ts#L24-L35
type IndexHealthStats struct {
	ActivePrimaryShards     int                         `json:"active_primary_shards"`
	ActiveShards            int                         `json:"active_shards"`
	InitializingShards      int                         `json:"initializing_shards"`
	NumberOfReplicas        int                         `json:"number_of_replicas"`
	NumberOfShards          int                         `json:"number_of_shards"`
	RelocatingShards        int                         `json:"relocating_shards"`
	Shards                  map[string]ShardHealthStats `json:"shards,omitempty"`
	Status                  healthstatus.HealthStatus   `json:"status"`
	UnassignedPrimaryShards int                         `json:"unassigned_primary_shards"`
	UnassignedShards        int                         `json:"unassigned_shards"`
}

func (s *IndexHealthStats) UnmarshalJSON(data []byte) error {

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

		case "active_primary_shards":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ActivePrimaryShards", err)
				}
				s.ActivePrimaryShards = value
			case float64:
				f := int(v)
				s.ActivePrimaryShards = f
			}

		case "active_shards":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ActiveShards", err)
				}
				s.ActiveShards = value
			case float64:
				f := int(v)
				s.ActiveShards = f
			}

		case "initializing_shards":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "InitializingShards", err)
				}
				s.InitializingShards = value
			case float64:
				f := int(v)
				s.InitializingShards = f
			}

		case "number_of_replicas":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfReplicas", err)
				}
				s.NumberOfReplicas = value
			case float64:
				f := int(v)
				s.NumberOfReplicas = f
			}

		case "number_of_shards":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfShards", err)
				}
				s.NumberOfShards = value
			case float64:
				f := int(v)
				s.NumberOfShards = f
			}

		case "relocating_shards":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RelocatingShards", err)
				}
				s.RelocatingShards = value
			case float64:
				f := int(v)
				s.RelocatingShards = f
			}

		case "shards":
			if s.Shards == nil {
				s.Shards = make(map[string]ShardHealthStats, 0)
			}
			if err := dec.Decode(&s.Shards); err != nil {
				return fmt.Errorf("%s | %w", "Shards", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "unassigned_primary_shards":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "UnassignedPrimaryShards", err)
				}
				s.UnassignedPrimaryShards = value
			case float64:
				f := int(v)
				s.UnassignedPrimaryShards = f
			}

		case "unassigned_shards":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "UnassignedShards", err)
				}
				s.UnassignedShards = value
			case float64:
				f := int(v)
				s.UnassignedShards = f
			}

		}
	}
	return nil
}

// NewIndexHealthStats returns a IndexHealthStats.
func NewIndexHealthStats() *IndexHealthStats {
	r := &IndexHealthStats{
		Shards: make(map[string]ShardHealthStats),
	}

	return r
}
