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

// Hyperparameter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/TrainedModel.ts#L216-L230
type Hyperparameter struct {
	// AbsoluteImportance A positive number showing how much the parameter influences the variation of
	// the loss function. For hyperparameters with values that are not specified by
	// the user but tuned during hyperparameter optimization.
	AbsoluteImportance *Float64 `json:"absolute_importance,omitempty"`
	// Name Name of the hyperparameter.
	Name string `json:"name"`
	// RelativeImportance A number between 0 and 1 showing the proportion of influence on the variation
	// of the loss function among all tuned hyperparameters. For hyperparameters
	// with values that are not specified by the user but tuned during
	// hyperparameter optimization.
	RelativeImportance *Float64 `json:"relative_importance,omitempty"`
	// Supplied Indicates if the hyperparameter is specified by the user (true) or optimized
	// (false).
	Supplied bool `json:"supplied"`
	// Value The value of the hyperparameter, either optimized or specified by the user.
	Value Float64 `json:"value"`
}

func (s *Hyperparameter) UnmarshalJSON(data []byte) error {

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

		case "absolute_importance":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.AbsoluteImportance = &f
			case float64:
				f := Float64(v)
				s.AbsoluteImportance = &f
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "relative_importance":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.RelativeImportance = &f
			case float64:
				f := Float64(v)
				s.RelativeImportance = &f
			}

		case "supplied":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Supplied = value
			case bool:
				s.Supplied = v
			}

		case "value":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Value = f
			case float64:
				f := Float64(v)
				s.Value = f
			}

		}
	}
	return nil
}

// NewHyperparameter returns a Hyperparameter.
func NewHyperparameter() *Hyperparameter {
	r := &Hyperparameter{}

	return r
}
