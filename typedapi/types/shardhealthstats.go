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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
)

// ShardHealthStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cluster/health/types.ts#L36-L43
type ShardHealthStats struct {
	ActiveShards       int                       `json:"active_shards"`
	InitializingShards int                       `json:"initializing_shards"`
	PrimaryActive      bool                      `json:"primary_active"`
	RelocatingShards   int                       `json:"relocating_shards"`
	Status             healthstatus.HealthStatus `json:"status"`
	UnassignedShards   int                       `json:"unassigned_shards"`
}

func (s *ShardHealthStats) UnmarshalJSON(data []byte) error {

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

		case "active_shards":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ActiveShards = value
			case float64:
				f := int(v)
				s.ActiveShards = f
			}

		case "initializing_shards":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.InitializingShards = value
			case float64:
				f := int(v)
				s.InitializingShards = f
			}

		case "primary_active":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.PrimaryActive = value
			case bool:
				s.PrimaryActive = v
			}

		case "relocating_shards":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.RelocatingShards = value
			case float64:
				f := int(v)
				s.RelocatingShards = f
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return err
			}

		case "unassigned_shards":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
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

// NewShardHealthStats returns a ShardHealthStats.
func NewShardHealthStats() *ShardHealthStats {
	r := &ShardHealthStats{}

	return r
}
