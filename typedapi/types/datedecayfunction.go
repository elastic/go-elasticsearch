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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/multivaluemode"

	"fmt"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// DateDecayFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/query_dsl/compound.ts#L92-L94
type DateDecayFunction struct {
	DateDecayFunction map[string]DecayPlacementDateMathDuration `json:"DateDecayFunction,omitempty"`
	MultiValueMode    *multivaluemode.MultiValueMode            `json:"multi_value_mode,omitempty"`
}

func (s *DateDecayFunction) UnmarshalJSON(data []byte) error {

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

		case "DateDecayFunction":
			if s.DateDecayFunction == nil {
				s.DateDecayFunction = make(map[string]DecayPlacementDateMathDuration, 0)
			}
			if err := dec.Decode(&s.DateDecayFunction); err != nil {
				return err
			}

		case "multi_value_mode":
			if err := dec.Decode(&s.MultiValueMode); err != nil {
				return err
			}

		default:

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s DateDecayFunction) MarshalJSON() ([]byte, error) {
	type opt DateDecayFunction
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.DateDecayFunction {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "DateDecayFunction")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewDateDecayFunction returns a DateDecayFunction.
func NewDateDecayFunction() *DateDecayFunction {
	r := &DateDecayFunction{
		DateDecayFunction: make(map[string]DecayPlacementDateMathDuration, 0),
	}

	return r
}
