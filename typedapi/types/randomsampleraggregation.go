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

// RandomSamplerAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/bucket.ts#L749-L769
type RandomSamplerAggregation struct {
	// Probability The probability that a document will be included in the aggregated data.
	// Must be greater than 0, less than 0.5, or exactly 1.
	// The lower the probability, the fewer documents are matched.
	Probability Float64 `json:"probability"`
	// Seed The seed to generate the random sampling of documents.
	// When a seed is provided, the random subset of documents is the same between
	// calls.
	Seed *int `json:"seed,omitempty"`
	// ShardSeed When combined with seed, setting shard_seed ensures 100% consistent sampling
	// over shards where data is exactly the same.
	ShardSeed *int `json:"shard_seed,omitempty"`
}

func (s *RandomSamplerAggregation) UnmarshalJSON(data []byte) error {

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

		case "probability":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Probability", err)
				}
				f := Float64(value)
				s.Probability = f
			case float64:
				f := Float64(v)
				s.Probability = f
			}

		case "seed":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Seed", err)
				}
				s.Seed = &value
			case float64:
				f := int(v)
				s.Seed = &f
			}

		case "shard_seed":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardSeed", err)
				}
				s.ShardSeed = &value
			case float64:
				f := int(v)
				s.ShardSeed = &f
			}

		}
	}
	return nil
}

// NewRandomSamplerAggregation returns a RandomSamplerAggregation.
func NewRandomSamplerAggregation() *RandomSamplerAggregation {
	r := &RandomSamplerAggregation{}

	return r
}
