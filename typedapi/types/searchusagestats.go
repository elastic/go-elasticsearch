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

// SearchUsageStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L149-L155
type SearchUsageStats struct {
	Queries    map[string]int64 `json:"queries"`
	Rescorers  map[string]int64 `json:"rescorers"`
	Retrievers map[string]int64 `json:"retrievers"`
	Sections   map[string]int64 `json:"sections"`
	Total      int64            `json:"total"`
}

func (s *SearchUsageStats) UnmarshalJSON(data []byte) error {

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

		case "queries":
			if s.Queries == nil {
				s.Queries = make(map[string]int64, 0)
			}
			if err := dec.Decode(&s.Queries); err != nil {
				return fmt.Errorf("%s | %w", "Queries", err)
			}

		case "rescorers":
			if s.Rescorers == nil {
				s.Rescorers = make(map[string]int64, 0)
			}
			if err := dec.Decode(&s.Rescorers); err != nil {
				return fmt.Errorf("%s | %w", "Rescorers", err)
			}

		case "retrievers":
			if s.Retrievers == nil {
				s.Retrievers = make(map[string]int64, 0)
			}
			if err := dec.Decode(&s.Retrievers); err != nil {
				return fmt.Errorf("%s | %w", "Retrievers", err)
			}

		case "sections":
			if s.Sections == nil {
				s.Sections = make(map[string]int64, 0)
			}
			if err := dec.Decode(&s.Sections); err != nil {
				return fmt.Errorf("%s | %w", "Sections", err)
			}

		case "total":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Total", err)
				}
				s.Total = value
			case float64:
				f := int64(v)
				s.Total = f
			}

		}
	}
	return nil
}

// NewSearchUsageStats returns a SearchUsageStats.
func NewSearchUsageStats() *SearchUsageStats {
	r := &SearchUsageStats{
		Queries:    make(map[string]int64),
		Rescorers:  make(map[string]int64),
		Retrievers: make(map[string]int64),
		Sections:   make(map[string]int64),
	}

	return r
}
