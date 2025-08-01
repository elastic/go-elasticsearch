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
)

// ShardRetentionLeases type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/stats/types.ts#L156-L160
type ShardRetentionLeases struct {
	Leases      []ShardLease `json:"leases"`
	PrimaryTerm int64        `json:"primary_term"`
	Version     int64        `json:"version"`
}

func (s *ShardRetentionLeases) UnmarshalJSON(data []byte) error {

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

		case "leases":
			if err := dec.Decode(&s.Leases); err != nil {
				return fmt.Errorf("%s | %w", "Leases", err)
			}

		case "primary_term":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryTerm", err)
				}
				s.PrimaryTerm = value
			case float64:
				f := int64(v)
				s.PrimaryTerm = f
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewShardRetentionLeases returns a ShardRetentionLeases.
func NewShardRetentionLeases() *ShardRetentionLeases {
	r := &ShardRetentionLeases{}

	return r
}
