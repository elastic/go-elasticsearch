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
)

// CacheStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/enrich/stats/types.ts#L38-L50
type CacheStats struct {
	Count              int    `json:"count"`
	Evictions          int    `json:"evictions"`
	Hits               int    `json:"hits"`
	HitsTimeInMillis   int64  `json:"hits_time_in_millis"`
	Misses             int    `json:"misses"`
	MissesTimeInMillis int64  `json:"misses_time_in_millis"`
	NodeId             string `json:"node_id"`
	SizeInBytes        int64  `json:"size_in_bytes"`
}

func (s *CacheStats) UnmarshalJSON(data []byte) error {

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
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Evictions", err)
				}
				s.Evictions = value
			case float64:
				f := int(v)
				s.Evictions = f
			}

		case "hits":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Hits", err)
				}
				s.Hits = value
			case float64:
				f := int(v)
				s.Hits = f
			}

		case "hits_time_in_millis":
			if err := dec.Decode(&s.HitsTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "HitsTimeInMillis", err)
			}

		case "misses":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Misses", err)
				}
				s.Misses = value
			case float64:
				f := int(v)
				s.Misses = f
			}

		case "misses_time_in_millis":
			if err := dec.Decode(&s.MissesTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "MissesTimeInMillis", err)
			}

		case "node_id":
			if err := dec.Decode(&s.NodeId); err != nil {
				return fmt.Errorf("%s | %w", "NodeId", err)
			}

		case "size_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SizeInBytes", err)
				}
				s.SizeInBytes = value
			case float64:
				f := int64(v)
				s.SizeInBytes = f
			}

		}
	}
	return nil
}

// NewCacheStats returns a CacheStats.
func NewCacheStats() *CacheStats {
	r := &CacheStats{}

	return r
}
