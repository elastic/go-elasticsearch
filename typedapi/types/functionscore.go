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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// FunctionScore type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/query_dsl/compound.ts#L107-L127
type FunctionScore struct {
	Exp              DecayFunction                  `json:"exp,omitempty"`
	FieldValueFactor *FieldValueFactorScoreFunction `json:"field_value_factor,omitempty"`
	Filter           *Query                         `json:"filter,omitempty"`
	Gauss            DecayFunction                  `json:"gauss,omitempty"`
	Linear           DecayFunction                  `json:"linear,omitempty"`
	RandomScore      *RandomScoreFunction           `json:"random_score,omitempty"`
	ScriptScore      *ScriptScoreFunction           `json:"script_score,omitempty"`
	Weight           *Float64                       `json:"weight,omitempty"`
}

func (s *FunctionScore) UnmarshalJSON(data []byte) error {

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

		case "exp":
			if err := dec.Decode(&s.Exp); err != nil {
				return err
			}

		case "field_value_factor":
			if err := dec.Decode(&s.FieldValueFactor); err != nil {
				return err
			}

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return err
			}

		case "gauss":
			if err := dec.Decode(&s.Gauss); err != nil {
				return err
			}

		case "linear":
			if err := dec.Decode(&s.Linear); err != nil {
				return err
			}

		case "random_score":
			if err := dec.Decode(&s.RandomScore); err != nil {
				return err
			}

		case "script_score":
			if err := dec.Decode(&s.ScriptScore); err != nil {
				return err
			}

		case "weight":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Weight = &f
			case float64:
				f := Float64(v)
				s.Weight = &f
			}

		}
	}
	return nil
}

// NewFunctionScore returns a FunctionScore.
func NewFunctionScore() *FunctionScore {
	r := &FunctionScore{}

	return r
}
