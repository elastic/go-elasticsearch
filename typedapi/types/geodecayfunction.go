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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

// GeoDecayFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6/specification/_types/query_dsl/compound.ts#L212-L215
type GeoDecayFunction struct {
	DecayFunctionBaseGeoLocationDistance map[string]DecayPlacementGeoLocationDistance `json:"-"`
	// MultiValueMode Determines how the distance is calculated when a field used for computing the
	// decay contains multiple values.
	MultiValueMode *multivaluemode.MultiValueMode `json:"multi_value_mode,omitempty"`
}

func (s *GeoDecayFunction) UnmarshalJSON(data []byte) error {

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
				if s.DecayFunctionBaseGeoLocationDistance == nil {
					s.DecayFunctionBaseGeoLocationDistance = make(map[string]DecayPlacementGeoLocationDistance, 0)
				}
				raw := NewDecayPlacementGeoLocationDistance()
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "DecayFunctionBaseGeoLocationDistance", err)
				}
				if raw != nil {
					s.DecayFunctionBaseGeoLocationDistance[key] = *raw
				}
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s GeoDecayFunction) MarshalJSON() ([]byte, error) {
	type opt GeoDecayFunction
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
	for key, value := range s.DecayFunctionBaseGeoLocationDistance {
		marshaled, err := json.Marshal(value)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal additional property %q: %w", key, err)
		}
		tmp[fmt.Sprintf("%s", key)] = marshaled
	}
	delete(tmp, "DecayFunctionBaseGeoLocationDistance")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewGeoDecayFunction returns a GeoDecayFunction.
func NewGeoDecayFunction() *GeoDecayFunction {
	r := &GeoDecayFunction{
		DecayFunctionBaseGeoLocationDistance: make(map[string]DecayPlacementGeoLocationDistance),
	}

	return r
}

type GeoDecayFunctionVariant interface {
	GeoDecayFunctionCaster() *GeoDecayFunction
}

func (s *GeoDecayFunction) GeoDecayFunctionCaster() *GeoDecayFunction {
	return s
}

func (s *GeoDecayFunction) DecayFunctionCaster() *DecayFunction {
	if s == nil {
		return nil
	}
	o := DecayFunction(s)
	return &o
}
