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
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shardstate"
)

// RepositoryStatsShards type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/cluster/stats/types.ts#L682-L687
type RepositoryStatsShards struct {
	Complete   int                           `json:"complete"`
	Incomplete int                           `json:"incomplete"`
	States     map[shardstate.ShardState]int `json:"states"`
	Total      int                           `json:"total"`
}

func (s *RepositoryStatsShards) UnmarshalJSON(data []byte) error {

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

		case "complete":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Complete", err)
				}
				s.Complete = value
			case float64:
				f := int(v)
				s.Complete = f
			}

		case "incomplete":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Incomplete", err)
				}
				s.Incomplete = value
			case float64:
				f := int(v)
				s.Incomplete = f
			}

		case "states":
			if s.States == nil {
				s.States = make(map[shardstate.ShardState]int, 0)
			}
			if err := dec.Decode(&s.States); err != nil {
				return fmt.Errorf("%s | %w", "States", err)
			}

		case "total":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Total", err)
				}
				s.Total = value
			case float64:
				f := int(v)
				s.Total = f
			}

		}
	}
	return nil
}

// NewRepositoryStatsShards returns a RepositoryStatsShards.
func NewRepositoryStatsShards() *RepositoryStatsShards {
	r := &RepositoryStatsShards{
		States: make(map[shardstate.ShardState]int),
	}

	return r
}
