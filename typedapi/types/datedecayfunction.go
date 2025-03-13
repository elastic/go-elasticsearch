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
// https://github.com/elastic/elasticsearch-specification/tree/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/multivaluemode"
)

// DateDecayFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827/specification/_types/query_dsl/compound.ts#L209-L209
type DateDecayFunction struct {
	DecayFunctionBaseDateMathDuration map[string]DecayPlacementDateMathDuration `json:"-"`
	// MultiValueMode Determines how the distance is calculated when a field used for computing the
	// decay contains multiple values.
	MultiValueMode *multivaluemode.MultiValueMode `json:"multi_value_mode,omitempty"`
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

		case "multi_value_mode":
			if err := dec.Decode(&s.MultiValueMode); err != nil {
				return fmt.Errorf("%s | %w", "MultiValueMode", err)
			}
		default:
			if key, ok := t.(string); ok {
				if s.DecayFunctionBaseDateMathDuration == nil {
					s.DecayFunctionBaseDateMathDuration = make(map[string]DecayPlacementDateMathDuration)
				}
				raw := new(DecayPlacementDateMathDuration)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "DecayPlacementDateMathDuration", err)
				}
				s.DecayFunctionBaseDateMathDuration[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s DateDecayFunction) MarshalJSON() ([]byte, error) {
	type opt DateDecayFunction
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.DecayFunctionBaseDateMathDuration {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "DecayFunctionBaseDateMathDuration")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewDateDecayFunction returns a DateDecayFunction.
func NewDateDecayFunction() *DateDecayFunction {
	r := &DateDecayFunction{
		DecayFunctionBaseDateMathDuration: make(map[string]DecayPlacementDateMathDuration, 0),
	}

	return r
}
