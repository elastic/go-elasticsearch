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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ChiSquareHeuristic type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_types/aggregations/bucket.ts#L733-L742
type ChiSquareHeuristic struct {
	// BackgroundIsSuperset Set to `false` if you defined a custom background filter that represents a
	// different set of documents that you want to compare to.
	BackgroundIsSuperset bool `json:"background_is_superset"`
	// IncludeNegatives Set to `false` to filter out the terms that appear less often in the subset
	// than in documents outside the subset.
	IncludeNegatives bool `json:"include_negatives"`
}

func (s *ChiSquareHeuristic) UnmarshalJSON(data []byte) error {

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

		case "background_is_superset":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "BackgroundIsSuperset", err)
				}
				s.BackgroundIsSuperset = value
			case bool:
				s.BackgroundIsSuperset = v
			}

		case "include_negatives":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IncludeNegatives", err)
				}
				s.IncludeNegatives = value
			case bool:
				s.IncludeNegatives = v
			}

		}
	}
	return nil
}

// NewChiSquareHeuristic returns a ChiSquareHeuristic.
func NewChiSquareHeuristic() *ChiSquareHeuristic {
	r := &ChiSquareHeuristic{}

	return r
}
