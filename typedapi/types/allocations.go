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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Allocations type.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/nodes/_types/Stats.ts#L120-L141
type Allocations struct {
	// CurrentDiskUsageInBytes Current disk usage, in bytes, for the node.
	CurrentDiskUsageInBytes *int64 `json:"current_disk_usage_in_bytes,omitempty"`
	// ForecastedDiskUsageInBytes Forecasted disk usage, in bytes, for the node.
	ForecastedDiskUsageInBytes *int64 `json:"forecasted_disk_usage_in_bytes,omitempty"`
	// ForecastedIngestLoad Forecasted ingest load for the node.
	ForecastedIngestLoad *Float64 `json:"forecasted_ingest_load,omitempty"`
	// Shards Number of shards allocated to the node.
	Shards *int `json:"shards,omitempty"`
	// UndesiredShards Number of shards allocated to the node that are currently undesired.
	UndesiredShards *int `json:"undesired_shards,omitempty"`
}

func (s *Allocations) UnmarshalJSON(data []byte) error {

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

		case "current_disk_usage_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CurrentDiskUsageInBytes", err)
				}
				s.CurrentDiskUsageInBytes = &value
			case float64:
				f := int64(v)
				s.CurrentDiskUsageInBytes = &f
			}

		case "forecasted_disk_usage_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ForecastedDiskUsageInBytes", err)
				}
				s.ForecastedDiskUsageInBytes = &value
			case float64:
				f := int64(v)
				s.ForecastedDiskUsageInBytes = &f
			}

		case "forecasted_ingest_load":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ForecastedIngestLoad", err)
				}
				f := Float64(value)
				s.ForecastedIngestLoad = &f
			case float64:
				f := Float64(v)
				s.ForecastedIngestLoad = &f
			}

		case "shards":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Shards", err)
				}
				s.Shards = &value
			case float64:
				f := int(v)
				s.Shards = &f
			}

		case "undesired_shards":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "UndesiredShards", err)
				}
				s.UndesiredShards = &value
			case float64:
				f := int(v)
				s.UndesiredShards = &f
			}

		}
	}
	return nil
}

// NewAllocations returns a Allocations.
func NewAllocations() *Allocations {
	r := &Allocations{}

	return r
}
