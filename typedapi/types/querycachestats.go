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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// QueryCacheStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/Stats.ts#L192-L226
type QueryCacheStats struct {
	// CacheCount Total number of entries added to the query cache across all shards assigned
	// to selected nodes.
	// This number includes current and evicted entries.
	CacheCount int `json:"cache_count"`
	// CacheSize Total number of entries currently in the query cache across all shards
	// assigned to selected nodes.
	CacheSize int `json:"cache_size"`
	// Evictions Total number of query cache evictions across all shards assigned to selected
	// nodes.
	Evictions int `json:"evictions"`
	// HitCount Total count of query cache hits across all shards assigned to selected nodes.
	HitCount int `json:"hit_count"`
	// MemorySize Total amount of memory used for the query cache across all shards assigned to
	// selected nodes.
	MemorySize ByteSize `json:"memory_size,omitempty"`
	// MemorySizeInBytes Total amount, in bytes, of memory used for the query cache across all shards
	// assigned to selected nodes.
	MemorySizeInBytes int64 `json:"memory_size_in_bytes"`
	// MissCount Total count of query cache misses across all shards assigned to selected
	// nodes.
	MissCount int `json:"miss_count"`
	// TotalCount Total count of hits and misses in the query cache across all shards assigned
	// to selected nodes.
	TotalCount int `json:"total_count"`
}

func (s *QueryCacheStats) UnmarshalJSON(data []byte) error {

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

		case "cache_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.CacheCount = value
			case float64:
				f := int(v)
				s.CacheCount = f
			}

		case "cache_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.CacheSize = value
			case float64:
				f := int(v)
				s.CacheSize = f
			}

		case "evictions":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Evictions = value
			case float64:
				f := int(v)
				s.Evictions = f
			}

		case "hit_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.HitCount = value
			case float64:
				f := int(v)
				s.HitCount = f
			}

		case "memory_size":
			if err := dec.Decode(&s.MemorySize); err != nil {
				return err
			}

		case "memory_size_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MemorySizeInBytes = value
			case float64:
				f := int64(v)
				s.MemorySizeInBytes = f
			}

		case "miss_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MissCount = value
			case float64:
				f := int(v)
				s.MissCount = f
			}

		case "total_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TotalCount = value
			case float64:
				f := int(v)
				s.TotalCount = f
			}

		}
	}
	return nil
}

// NewQueryCacheStats returns a QueryCacheStats.
func NewQueryCacheStats() *QueryCacheStats {
	r := &QueryCacheStats{}

	return r
}
