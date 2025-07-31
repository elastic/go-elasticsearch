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

// MinHashTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L514-L525
type MinHashTokenFilter struct {
	// BucketCount Number of buckets to which hashes are assigned. Defaults to `512`.
	BucketCount *int `json:"bucket_count,omitempty"`
	// HashCount Number of ways to hash each token in the stream. Defaults to `1`.
	HashCount *int `json:"hash_count,omitempty"`
	// HashSetSize Number of hashes to keep from each bucket. Defaults to `1`.
	// Hashes are retained by ascending size, starting with the bucket’s smallest
	// hash first.
	HashSetSize *int    `json:"hash_set_size,omitempty"`
	Type        string  `json:"type,omitempty"`
	Version     *string `json:"version,omitempty"`
	// WithRotation If `true`, the filter fills empty buckets with the value of the first
	// non-empty bucket to its circular right if the `hash_set_size` is `1`. If the
	// `bucket_count` argument is greater than 1, this parameter defaults to `true`.
	// Otherwise, this parameter defaults to `false`.
	WithRotation *bool `json:"with_rotation,omitempty"`
}

func (s *MinHashTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "bucket_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "BucketCount", err)
				}
				s.BucketCount = &value
			case float64:
				f := int(v)
				s.BucketCount = &f
			}

		case "hash_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "HashCount", err)
				}
				s.HashCount = &value
			case float64:
				f := int(v)
				s.HashCount = &f
			}

		case "hash_set_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "HashSetSize", err)
				}
				s.HashSetSize = &value
			case float64:
				f := int(v)
				s.HashSetSize = &f
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		case "with_rotation":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "WithRotation", err)
				}
				s.WithRotation = &value
			case bool:
				s.WithRotation = &v
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s MinHashTokenFilter) MarshalJSON() ([]byte, error) {
	type innerMinHashTokenFilter MinHashTokenFilter
	tmp := innerMinHashTokenFilter{
		BucketCount:  s.BucketCount,
		HashCount:    s.HashCount,
		HashSetSize:  s.HashSetSize,
		Type:         s.Type,
		Version:      s.Version,
		WithRotation: s.WithRotation,
	}

	tmp.Type = "min_hash"

	return json.Marshal(tmp)
}

// NewMinHashTokenFilter returns a MinHashTokenFilter.
func NewMinHashTokenFilter() *MinHashTokenFilter {
	r := &MinHashTokenFilter{}

	return r
}
