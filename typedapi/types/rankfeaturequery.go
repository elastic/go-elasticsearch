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

// RankFeatureQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/query_dsl/specialized.ts#L309-L335
type RankFeatureQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// Field `rank_feature` or `rank_features` field used to boost relevance scores.
	Field string `json:"field"`
	// Linear Linear function used to boost relevance scores based on the value of the rank
	// feature `field`.
	Linear *RankFeatureFunctionLinear `json:"linear,omitempty"`
	// Log Logarithmic function used to boost relevance scores based on the value of the
	// rank feature `field`.
	Log        *RankFeatureFunctionLogarithm `json:"log,omitempty"`
	QueryName_ *string                       `json:"_name,omitempty"`
	// Saturation Saturation function used to boost relevance scores based on the value of the
	// rank feature `field`.
	Saturation *RankFeatureFunctionSaturation `json:"saturation,omitempty"`
	// Sigmoid Sigmoid function used to boost relevance scores based on the value of the
	// rank feature `field`.
	Sigmoid *RankFeatureFunctionSigmoid `json:"sigmoid,omitempty"`
}

func (s *RankFeatureQuery) UnmarshalJSON(data []byte) error {

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

		case "boost":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Boost", err)
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "linear":
			if err := dec.Decode(&s.Linear); err != nil {
				return fmt.Errorf("%s | %w", "Linear", err)
			}

		case "log":
			if err := dec.Decode(&s.Log); err != nil {
				return fmt.Errorf("%s | %w", "Log", err)
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "QueryName_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		case "saturation":
			if err := dec.Decode(&s.Saturation); err != nil {
				return fmt.Errorf("%s | %w", "Saturation", err)
			}

		case "sigmoid":
			if err := dec.Decode(&s.Sigmoid); err != nil {
				return fmt.Errorf("%s | %w", "Sigmoid", err)
			}

		}
	}
	return nil
}

// NewRankFeatureQuery returns a RankFeatureQuery.
func NewRankFeatureQuery() *RankFeatureQuery {
	r := &RankFeatureQuery{}

	return r
}
