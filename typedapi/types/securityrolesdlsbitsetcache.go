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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// SecurityRolesDlsBitSetCache type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6785a6caa1fa3ca5ab3308963d79dce923a3469f/specification/xpack/usage/types.ts#L322-L359
type SecurityRolesDlsBitSetCache struct {
	// Count Number of entries in the cache.
	Count int `json:"count"`
	// Evictions Total number of cache evictions.
	Evictions int64 `json:"evictions"`
	// Hits Total number of cache hits.
	Hits int64 `json:"hits"`
	// HitsTimeInMillis Total combined time spent in cache for hits in milliseconds.
	HitsTimeInMillis int64 `json:"hits_time_in_millis"`
	// Memory Human-readable amount of memory taken up by the cache.
	Memory ByteSize `json:"memory,omitempty"`
	// MemoryInBytes Memory taken up by the cache in bytes.
	MemoryInBytes uint64 `json:"memory_in_bytes"`
	// Misses Total number of cache misses.
	Misses int64 `json:"misses"`
	// MissesTimeInMillis Total combined time spent in cache for misses in milliseconds.
	MissesTimeInMillis int64 `json:"misses_time_in_millis"`
}

func (s *SecurityRolesDlsBitSetCache) UnmarshalJSON(data []byte) error {

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

		case "count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "evictions":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Evictions", err)
				}
				s.Evictions = value
			case float64:
				f := int64(v)
				s.Evictions = f
			}

		case "hits":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Hits", err)
				}
				s.Hits = value
			case float64:
				f := int64(v)
				s.Hits = f
			}

		case "hits_time_in_millis":
			if err := dec.Decode(&s.HitsTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "HitsTimeInMillis", err)
			}

		case "memory":
			if err := dec.Decode(&s.Memory); err != nil {
				return fmt.Errorf("%s | %w", "Memory", err)
			}

		case "memory_in_bytes":
			if err := dec.Decode(&s.MemoryInBytes); err != nil {
				return fmt.Errorf("%s | %w", "MemoryInBytes", err)
			}

		case "misses":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Misses", err)
				}
				s.Misses = value
			case float64:
				f := int64(v)
				s.Misses = f
			}

		case "misses_time_in_millis":
			if err := dec.Decode(&s.MissesTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "MissesTimeInMillis", err)
			}

		}
	}
	return nil
}

// NewSecurityRolesDlsBitSetCache returns a SecurityRolesDlsBitSetCache.
func NewSecurityRolesDlsBitSetCache() *SecurityRolesDlsBitSetCache {
	r := &SecurityRolesDlsBitSetCache{}

	return r
}
