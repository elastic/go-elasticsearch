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
// https://github.com/elastic/elasticsearch-specification/tree/b2c13a00c152a97cb41193deda8ed9b37fd06796

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// RankFeatureFunctionSaturation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b2c13a00c152a97cb41193deda8ed9b37fd06796/specification/_types/query_dsl/specialized.ts#L275-L280
type RankFeatureFunctionSaturation struct {
	// Pivot Configurable pivot value so that the result will be less than 0.5.
	Pivot *float32 `json:"pivot,omitempty"`
}

func (s *RankFeatureFunctionSaturation) UnmarshalJSON(data []byte) error {

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

		case "pivot":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Pivot", err)
				}
				f := float32(value)
				s.Pivot = &f
			case float64:
				f := float32(v)
				s.Pivot = &f
			}

		}
	}
	return nil
}

// NewRankFeatureFunctionSaturation returns a RankFeatureFunctionSaturation.
func NewRankFeatureFunctionSaturation() *RankFeatureFunctionSaturation {
	r := &RankFeatureFunctionSaturation{}

	return r
}
