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

// ElserServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L1269-L1295
type ElserServiceSettings struct {
	// AdaptiveAllocations Adaptive allocations configuration details.
	// If `enabled` is true, the number of allocations of the model is set based on
	// the current load the process gets.
	// When the load is high, a new model allocation is automatically created,
	// respecting the value of `max_number_of_allocations` if it's set.
	// When the load is low, a model allocation is automatically removed, respecting
	// the value of `min_number_of_allocations` if it's set.
	// If `enabled` is true, do not set the number of allocations manually.
	AdaptiveAllocations *AdaptiveAllocations `json:"adaptive_allocations,omitempty"`
	// NumAllocations The total number of allocations this model is assigned across machine
	// learning nodes.
	// Increasing this value generally increases the throughput.
	// If adaptive allocations is enabled, do not set this value because it's
	// automatically set.
	NumAllocations int `json:"num_allocations"`
	// NumThreads The number of threads used by each model allocation during inference.
	// Increasing this value generally increases the speed per inference request.
	// The inference process is a compute-bound process; `threads_per_allocations`
	// must not exceed the number of available allocated processors per node.
	// The value must be a power of 2.
	// The maximum value is 32.
	//
	// > info
	// > If you want to optimize your ELSER endpoint for ingest, set the number of
	// threads to 1. If you want to optimize your ELSER endpoint for search, set the
	// number of threads to greater than 1.
	NumThreads int `json:"num_threads"`
}

func (s *ElserServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "adaptive_allocations":
			if err := dec.Decode(&s.AdaptiveAllocations); err != nil {
				return fmt.Errorf("%s | %w", "AdaptiveAllocations", err)
			}

		case "num_allocations":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumAllocations", err)
				}
				s.NumAllocations = value
			case float64:
				f := int(v)
				s.NumAllocations = f
			}

		case "num_threads":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumThreads", err)
				}
				s.NumThreads = value
			case float64:
				f := int(v)
				s.NumThreads = f
			}

		}
	}
	return nil
}

// NewElserServiceSettings returns a ElserServiceSettings.
func NewElserServiceSettings() *ElserServiceSettings {
	r := &ElserServiceSettings{}

	return r
}
