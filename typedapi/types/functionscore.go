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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

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
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/compound.ts#L201-L241
type FunctionScore struct {
	// Exp Function that scores a document with a exponential decay, depending on the
	// distance of a numeric field value of the document from an origin.
	Exp DecayFunction `json:"exp,omitempty"`
	// FieldValueFactor Function allows you to use a field from a document to influence the score.
	// It’s similar to using the script_score function, however, it avoids the
	// overhead of scripting.
	FieldValueFactor *FieldValueFactorScoreFunction `json:"field_value_factor,omitempty"`
	Filter           *Query                         `json:"filter,omitempty"`
	// Gauss Function that scores a document with a normal decay, depending on the
	// distance of a numeric field value of the document from an origin.
	Gauss DecayFunction `json:"gauss,omitempty"`
	// Linear Function that scores a document with a linear decay, depending on the
	// distance of a numeric field value of the document from an origin.
	Linear DecayFunction `json:"linear,omitempty"`
	// RandomScore Generates scores that are uniformly distributed from 0 up to but not
	// including 1.
	// In case you want scores to be reproducible, it is possible to provide a
	// `seed` and `field`.
	RandomScore *RandomScoreFunction `json:"random_score,omitempty"`
	// ScriptScore Enables you to wrap another query and customize the scoring of it optionally
	// with a computation derived from other numeric field values in the doc using a
	// script expression.
	ScriptScore *ScriptScoreFunction `json:"script_score,omitempty"`
	Weight      *Float64             `json:"weight,omitempty"`
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
