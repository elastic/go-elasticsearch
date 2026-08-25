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
// https://github.com/elastic/elasticsearch-specification/tree/abf9c2c6bb21328339daa197aae15af2ecbc46f0

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Usage statistics for data stream lifecycle (DLM), reported by `_xpack/usage`
// under `data_lifecycle`. Besides `available` and `enabled`, all the following
// statistics are only present when the feature is enabled.
//
// https://github.com/elastic/elasticsearch-specification/blob/abf9c2c6bb21328339daa197aae15af2ecbc46f0/specification/xpack/usage/types.ts#L131-L161
type DataStreamLifecycleUsage struct {
	Available bool `json:"available"`
	// Count The number of data streams that have a lifecycle configured.
	Count *int64 `json:"count,omitempty"`
	// DataRetention Statistics about the explicitly configured data retention across data
	// streams.
	DataRetention *DataStreamLifecycleRetentionStats `json:"data_retention,omitempty"`
	// DefaultRolloverUsed Whether the default rollover configuration is used by at least one data
	// stream.
	DefaultRolloverUsed *bool `json:"default_rollover_used,omitempty"`
	// EffectiveRetention Statistics about the effective retention (configured or derived from global
	// retention) across data streams.
	EffectiveRetention *DataStreamLifecycleEffectiveRetentionStats `json:"effective_retention,omitempty"`
	Enabled            bool                                        `json:"enabled"`
	// FrozenAfter Statistics about the configured `frozen_after` (searchable snapshot) tier
	// threshold across data streams.
	FrozenAfter *DataStreamLifecycleRetentionStats `json:"frozen_after,omitempty"`
	// GlobalRetention Statistics about the cluster's global default and maximum retention settings.
	GlobalRetention *DataStreamLifecycleGlobalRetention `json:"global_retention,omitempty"`
}

func (s *DataStreamLifecycleUsage) UnmarshalJSON(data []byte) error {

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

		case "available":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Available", err)
				}
				s.Available = value
			case bool:
				s.Available = v
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
				s.Count = &value
			case float64:
				f := int64(v)
				s.Count = &f
			}

		case "data_retention":
			if err := dec.Decode(&s.DataRetention); err != nil {
				return fmt.Errorf("%s | %w", "DataRetention", err)
			}

		case "default_rollover_used":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DefaultRolloverUsed", err)
				}
				s.DefaultRolloverUsed = &value
			case bool:
				s.DefaultRolloverUsed = &v
			}

		case "effective_retention":
			if err := dec.Decode(&s.EffectiveRetention); err != nil {
				return fmt.Errorf("%s | %w", "EffectiveRetention", err)
			}

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "frozen_after":
			if err := dec.Decode(&s.FrozenAfter); err != nil {
				return fmt.Errorf("%s | %w", "FrozenAfter", err)
			}

		case "global_retention":
			if err := dec.Decode(&s.GlobalRetention); err != nil {
				return fmt.Errorf("%s | %w", "GlobalRetention", err)
			}

		}
	}
	return nil
}

// NewDataStreamLifecycleUsage returns a DataStreamLifecycleUsage.
func NewDataStreamLifecycleUsage() *DataStreamLifecycleUsage {
	r := &DataStreamLifecycleUsage{}

	return r
}
