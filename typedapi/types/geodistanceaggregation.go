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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geodistancetype"
)

// GeoDistanceAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/bucket.ts#L396-L419
type GeoDistanceAggregation struct {
	// DistanceType The distance calculation type.
	DistanceType *geodistancetype.GeoDistanceType `json:"distance_type,omitempty"`
	// Field A field of type `geo_point` used to evaluate the distance.
	Field *string `json:"field,omitempty"`
	// Origin The origin  used to evaluate the distance.
	Origin GeoLocation `json:"origin,omitempty"`
	// Ranges An array of ranges used to bucket documents.
	Ranges []AggregationRange `json:"ranges,omitempty"`
	// Unit The distance unit.
	Unit *distanceunit.DistanceUnit `json:"unit,omitempty"`
}

func (s *GeoDistanceAggregation) UnmarshalJSON(data []byte) error {

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

		case "distance_type":
			if err := dec.Decode(&s.DistanceType); err != nil {
				return fmt.Errorf("%s | %w", "DistanceType", err)
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
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

		case "ranges":
			if err := dec.Decode(&s.Ranges); err != nil {
				return fmt.Errorf("%s | %w", "Ranges", err)
			}

		case "unit":
			if err := dec.Decode(&s.Unit); err != nil {
				return fmt.Errorf("%s | %w", "Unit", err)
			}

		}
	}
	return nil
}

// NewGeoDistanceAggregation returns a GeoDistanceAggregation.
func NewGeoDistanceAggregation() *GeoDistanceAggregation {
	r := &GeoDistanceAggregation{}

	return r
}
