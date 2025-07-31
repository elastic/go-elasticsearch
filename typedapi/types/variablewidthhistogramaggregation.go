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

// VariableWidthHistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/bucket.ts#L1091-L1115
type VariableWidthHistogramAggregation struct {
	// Buckets The target number of buckets.
	Buckets *int `json:"buckets,omitempty"`
	// Field The name of the field.
	Field *string `json:"field,omitempty"`
	// InitialBuffer Specifies the number of individual documents that will be stored in memory on
	// a shard before the initial bucketing algorithm is run.
	// Defaults to `min(10 * shard_size, 50000)`.
	InitialBuffer *int    `json:"initial_buffer,omitempty"`
	Script        *Script `json:"script,omitempty"`
	// ShardSize The number of buckets that the coordinating node will request from each
	// shard.
	// Defaults to `buckets * 50`.
	ShardSize *int `json:"shard_size,omitempty"`
}

func (s *VariableWidthHistogramAggregation) UnmarshalJSON(data []byte) error {

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

		case "buckets":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Buckets", err)
				}
				s.Buckets = &value
			case float64:
				f := int(v)
				s.Buckets = &f
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "initial_buffer":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "InitialBuffer", err)
				}
				s.InitialBuffer = &value
			case float64:
				f := int(v)
				s.InitialBuffer = &f
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		case "shard_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardSize", err)
				}
				s.ShardSize = &value
			case float64:
				f := int(v)
				s.ShardSize = &f
			}

		}
	}
	return nil
}

// NewVariableWidthHistogramAggregation returns a VariableWidthHistogramAggregation.
func NewVariableWidthHistogramAggregation() *VariableWidthHistogramAggregation {
	r := &VariableWidthHistogramAggregation{}

	return r
}
