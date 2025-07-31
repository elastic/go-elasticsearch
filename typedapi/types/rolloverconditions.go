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

// RolloverConditions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/rollover/types.ts#L24-L40
type RolloverConditions struct {
	MaxAge                   Duration `json:"max_age,omitempty"`
	MaxAgeMillis             *int64   `json:"max_age_millis,omitempty"`
	MaxDocs                  *int64   `json:"max_docs,omitempty"`
	MaxPrimaryShardDocs      *int64   `json:"max_primary_shard_docs,omitempty"`
	MaxPrimaryShardSize      ByteSize `json:"max_primary_shard_size,omitempty"`
	MaxPrimaryShardSizeBytes *int64   `json:"max_primary_shard_size_bytes,omitempty"`
	MaxSize                  ByteSize `json:"max_size,omitempty"`
	MaxSizeBytes             *int64   `json:"max_size_bytes,omitempty"`
	MinAge                   Duration `json:"min_age,omitempty"`
	MinDocs                  *int64   `json:"min_docs,omitempty"`
	MinPrimaryShardDocs      *int64   `json:"min_primary_shard_docs,omitempty"`
	MinPrimaryShardSize      ByteSize `json:"min_primary_shard_size,omitempty"`
	MinPrimaryShardSizeBytes *int64   `json:"min_primary_shard_size_bytes,omitempty"`
	MinSize                  ByteSize `json:"min_size,omitempty"`
	MinSizeBytes             *int64   `json:"min_size_bytes,omitempty"`
}

func (s *RolloverConditions) UnmarshalJSON(data []byte) error {

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

		case "max_age":
			if err := dec.Decode(&s.MaxAge); err != nil {
				return fmt.Errorf("%s | %w", "MaxAge", err)
			}

		case "max_age_millis":
			if err := dec.Decode(&s.MaxAgeMillis); err != nil {
				return fmt.Errorf("%s | %w", "MaxAgeMillis", err)
			}

		case "max_docs":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxDocs", err)
				}
				s.MaxDocs = &value
			case float64:
				f := int64(v)
				s.MaxDocs = &f
			}

		case "max_primary_shard_docs":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxPrimaryShardDocs", err)
				}
				s.MaxPrimaryShardDocs = &value
			case float64:
				f := int64(v)
				s.MaxPrimaryShardDocs = &f
			}

		case "max_primary_shard_size":
			if err := dec.Decode(&s.MaxPrimaryShardSize); err != nil {
				return fmt.Errorf("%s | %w", "MaxPrimaryShardSize", err)
			}

		case "max_primary_shard_size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxPrimaryShardSizeBytes", err)
				}
				s.MaxPrimaryShardSizeBytes = &value
			case float64:
				f := int64(v)
				s.MaxPrimaryShardSizeBytes = &f
			}

		case "max_size":
			if err := dec.Decode(&s.MaxSize); err != nil {
				return fmt.Errorf("%s | %w", "MaxSize", err)
			}

		case "max_size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxSizeBytes", err)
				}
				s.MaxSizeBytes = &value
			case float64:
				f := int64(v)
				s.MaxSizeBytes = &f
			}

		case "min_age":
			if err := dec.Decode(&s.MinAge); err != nil {
				return fmt.Errorf("%s | %w", "MinAge", err)
			}

		case "min_docs":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinDocs", err)
				}
				s.MinDocs = &value
			case float64:
				f := int64(v)
				s.MinDocs = &f
			}

		case "min_primary_shard_docs":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinPrimaryShardDocs", err)
				}
				s.MinPrimaryShardDocs = &value
			case float64:
				f := int64(v)
				s.MinPrimaryShardDocs = &f
			}

		case "min_primary_shard_size":
			if err := dec.Decode(&s.MinPrimaryShardSize); err != nil {
				return fmt.Errorf("%s | %w", "MinPrimaryShardSize", err)
			}

		case "min_primary_shard_size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinPrimaryShardSizeBytes", err)
				}
				s.MinPrimaryShardSizeBytes = &value
			case float64:
				f := int64(v)
				s.MinPrimaryShardSizeBytes = &f
			}

		case "min_size":
			if err := dec.Decode(&s.MinSize); err != nil {
				return fmt.Errorf("%s | %w", "MinSize", err)
			}

		case "min_size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinSizeBytes", err)
				}
				s.MinSizeBytes = &value
			case float64:
				f := int64(v)
				s.MinSizeBytes = &f
			}

		}
	}
	return nil
}

// NewRolloverConditions returns a RolloverConditions.
func NewRolloverConditions() *RolloverConditions {
	r := &RolloverConditions{}

	return r
}
