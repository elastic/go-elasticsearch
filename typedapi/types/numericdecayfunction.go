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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

// NumericDecayFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/_types/query_dsl/compound.ts#L210-L210
type NumericDecayFunction struct {
	DecayFunctionBasedoubledouble map[string]DecayPlacementdoubledouble `json:"-"`
	// MultiValueMode Determines how the distance is calculated when a field used for computing the
	// decay contains multiple values.
	MultiValueMode *multivaluemode.MultiValueMode `json:"multi_value_mode,omitempty"`
}

func (s *NumericDecayFunction) UnmarshalJSON(data []byte) error {

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
				if s.DecayFunctionBasedoubledouble == nil {
					s.DecayFunctionBasedoubledouble = make(map[string]DecayPlacementdoubledouble, 0)
				}
				raw := NewDecayPlacementdoubledouble()
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "DecayFunctionBasedoubledouble", err)
				}
				if raw != nil {
					s.DecayFunctionBasedoubledouble[key] = *raw
				}
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s NumericDecayFunction) MarshalJSON() ([]byte, error) {
	type opt NumericDecayFunction
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]json.RawMessage, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.DecayFunctionBasedoubledouble {
		marshaled, err := json.Marshal(value)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal additional property %q: %w", key, err)
		}
		tmp[fmt.Sprintf("%s", key)] = marshaled
	}
	delete(tmp, "DecayFunctionBasedoubledouble")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewNumericDecayFunction returns a NumericDecayFunction.
func NewNumericDecayFunction() *NumericDecayFunction {
	r := &NumericDecayFunction{
		DecayFunctionBasedoubledouble: make(map[string]DecayPlacementdoubledouble),
	}

	return r
}

type NumericDecayFunctionVariant interface {
	NumericDecayFunctionCaster() *NumericDecayFunction
}

func (s *NumericDecayFunction) NumericDecayFunctionCaster() *NumericDecayFunction {
	return s
}

func (s *NumericDecayFunction) DecayFunctionCaster() *DecayFunction {
	if s == nil {
		return nil
	}
	o := DecayFunction(s)
	return &o
}
