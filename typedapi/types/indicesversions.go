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

// IndicesVersions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L359-L365
type IndicesVersions struct {
	IndexCount        int      `json:"index_count"`
	PrimaryShardCount int      `json:"primary_shard_count"`
	TotalPrimaryBytes int64    `json:"total_primary_bytes"`
	TotalPrimarySize  ByteSize `json:"total_primary_size,omitempty"`
	Version           string   `json:"version"`
}

func (s *IndicesVersions) UnmarshalJSON(data []byte) error {

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

		case "index_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexCount", err)
				}
				s.IndexCount = value
			case float64:
				f := int(v)
				s.IndexCount = f
			}

		case "primary_shard_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryShardCount", err)
				}
				s.PrimaryShardCount = value
			case float64:
				f := int(v)
				s.PrimaryShardCount = f
			}

		case "total_primary_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalPrimaryBytes", err)
				}
				s.TotalPrimaryBytes = value
			case float64:
				f := int64(v)
				s.TotalPrimaryBytes = f
			}

		case "total_primary_size":
			if err := dec.Decode(&s.TotalPrimarySize); err != nil {
				return fmt.Errorf("%s | %w", "TotalPrimarySize", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewIndicesVersions returns a IndicesVersions.
func NewIndicesVersions() *IndicesVersions {
	r := &IndicesVersions{}

	return r
}
