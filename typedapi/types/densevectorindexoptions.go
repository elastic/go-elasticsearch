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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/densevectorindexoptionstype"
)

// DenseVectorIndexOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/mapping/DenseVectorProperty.ts#L129-L162
type DenseVectorIndexOptions struct {
	// ConfidenceInterval The confidence interval to use when quantizing the vectors. Can be any value
	// between and including `0.90` and
	// `1.0` or exactly `0`. When the value is `0`, this indicates that dynamic
	// quantiles should be calculated for
	// optimized quantization. When between `0.90` and `1.0`, this value restricts
	// the values used when calculating
	// the quantization thresholds.
	//
	// For example, a value of `0.95` will only use the middle `95%` of the values
	// when calculating the quantization
	// thresholds (e.g. the highest and lowest `2.5%` of values will be ignored).
	//
	// Defaults to `1/(dims + 1)` for `int8` quantized vectors and `0` for `int4`
	// for dynamic quantile calculation.
	//
	// Only applicable to `int8_hnsw`, `int4_hnsw`, `int8_flat`, and `int4_flat`
	// index types.
	ConfidenceInterval *float32 `json:"confidence_interval,omitempty"`
	// EfConstruction The number of candidates to track while assembling the list of nearest
	// neighbors for each new node.
	//
	// Only applicable to `hnsw`, `int8_hnsw`, `bbq_hnsw`, and `int4_hnsw` index
	// types.
	EfConstruction *int `json:"ef_construction,omitempty"`
	// M The number of neighbors each node will be connected to in the HNSW graph.
	//
	// Only applicable to `hnsw`, `int8_hnsw`, `bbq_hnsw`, and `int4_hnsw` index
	// types.
	M *int `json:"m,omitempty"`
	// Type The type of kNN algorithm to use.
	Type densevectorindexoptionstype.DenseVectorIndexOptionsType `json:"type"`
}

func (s *DenseVectorIndexOptions) UnmarshalJSON(data []byte) error {

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

		case "confidence_interval":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "ConfidenceInterval", err)
				}
				f := float32(value)
				s.ConfidenceInterval = &f
			case float64:
				f := float32(v)
				s.ConfidenceInterval = &f
			}

		case "ef_construction":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "EfConstruction", err)
				}
				s.EfConstruction = &value
			case float64:
				f := int(v)
				s.EfConstruction = &f
			}

		case "m":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "M", err)
				}
				s.M = &value
			case float64:
				f := int(v)
				s.M = &f
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// NewDenseVectorIndexOptions returns a DenseVectorIndexOptions.
func NewDenseVectorIndexOptions() *DenseVectorIndexOptions {
	r := &DenseVectorIndexOptions{}

	return r
}
