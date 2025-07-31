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

// DecayPlacementGeoLocationDistance type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/query_dsl/compound.ts#L170-L189
type DecayPlacementGeoLocationDistance struct {
	// Decay Defines how documents are scored at the distance given at scale.
	Decay *Float64 `json:"decay,omitempty"`
	// Offset If defined, the decay function will only compute the decay function for
	// documents with a distance greater than the defined `offset`.
	Offset *string `json:"offset,omitempty"`
	// Origin The point of origin used for calculating distance. Must be given as a number
	// for numeric field, date for date fields and geo point for geo fields.
	Origin GeoLocation `json:"origin,omitempty"`
	// Scale Defines the distance from origin + offset at which the computed score will
	// equal `decay` parameter.
	Scale *string `json:"scale,omitempty"`
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
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Decay", err)
				}
				f := Float64(value)
				s.Decay = &f
			case float64:
				f := Float64(v)
				s.Decay = &f
			}

		case "offset":
			if err := dec.Decode(&s.Offset); err != nil {
				return fmt.Errorf("%s | %w", "Offset", err)
			}

		case "origin":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Origin", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		origin_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Origin", err)
				}

				switch t {

				case "lat", "lon":
					o := NewLatLonGeoLocation()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Origin", err)
					}
					s.Origin = o
					break origin_field

				case "geohash":
					o := NewGeoHashLocation()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Origin", err)
					}
					s.Origin = o
					break origin_field

				}
			}
			if s.Origin == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Origin); err != nil {
					return fmt.Errorf("%s | %w", "Origin", err)
				}
			}

		case "scale":
			if err := dec.Decode(&s.Scale); err != nil {
				return fmt.Errorf("%s | %w", "Scale", err)
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
