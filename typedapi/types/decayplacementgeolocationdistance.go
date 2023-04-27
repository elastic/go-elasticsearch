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
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// DecayPlacementGeoLocationDistance type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/query_dsl/compound.ts#L77-L82
type DecayPlacementGeoLocationDistance struct {
	Decay  *Float64    `json:"decay,omitempty"`
	Offset *string     `json:"offset,omitempty"`
	Origin GeoLocation `json:"origin,omitempty"`
	Scale  *string     `json:"scale,omitempty"`
}

func (s *DecayPlacementGeoLocationDistance) UnmarshalJSON(data []byte) error {

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

		case "decay":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Decay = &f
			case float64:
				f := Float64(v)
				s.Decay = &f
			}

		case "offset":
			if err := dec.Decode(&s.Offset); err != nil {
				return err
			}

		case "origin":
			if err := dec.Decode(&s.Origin); err != nil {
				return err
			}

		case "scale":
			if err := dec.Decode(&s.Scale); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDecayPlacementGeoLocationDistance returns a DecayPlacementGeoLocationDistance.
func NewDecayPlacementGeoLocationDistance() *DecayPlacementGeoLocationDistance {
	r := &DecayPlacementGeoLocationDistance{}

	return r
}
