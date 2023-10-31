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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// RankFeatureFunctionSigmoid type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/query_dsl/specialized.ts#L282-L291
type RankFeatureFunctionSigmoid struct {
	// Exponent Configurable Exponent.
	Exponent float32 `json:"exponent"`
	// Pivot Configurable pivot value so that the result will be less than 0.5.
	Pivot float32 `json:"pivot"`
}

func (s *RankFeatureFunctionSigmoid) UnmarshalJSON(data []byte) error {

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

		case "exponent":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Exponent = f
			case float64:
				f := float32(v)
				s.Exponent = f
			}

		case "pivot":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Pivot = f
			case float64:
				f := float32(v)
				s.Pivot = f
			}

		}
	}
	return nil
}

// NewRankFeatureFunctionSigmoid returns a RankFeatureFunctionSigmoid.
func NewRankFeatureFunctionSigmoid() *RankFeatureFunctionSigmoid {
	r := &RankFeatureFunctionSigmoid{}

	return r
}
