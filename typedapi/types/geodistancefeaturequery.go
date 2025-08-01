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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// GeoDistanceFeatureQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/query_dsl/specialized.ts#L66-L69
type GeoDistanceFeatureQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// Field Name of the field used to calculate distances. This field must meet the
	// following criteria:
	// be a `date`, `date_nanos` or `geo_point` field;
	// have an `index` mapping parameter value of `true`, which is the default;
	// have an `doc_values` mapping parameter value of `true`, which is the default.
	Field string `json:"field"`
	// Origin Date or point of origin used to calculate distances.
	// If the `field` value is a `date` or `date_nanos` field, the `origin` value
	// must be a date.
	// Date Math, such as `now-1h`, is supported.
	// If the field value is a `geo_point` field, the `origin` value must be a
	// geopoint.
	Origin GeoLocation `json:"origin"`
	// Pivot Distance from the `origin` at which relevance scores receive half of the
	// `boost` value.
	// If the `field` value is a `date` or `date_nanos` field, the `pivot` value
	// must be a time unit, such as `1h` or `10d`. If the `field` value is a
	// `geo_point` field, the `pivot` value must be a distance unit, such as `1km`
	// or `12m`.
	Pivot      string  `json:"pivot"`
	QueryName_ *string `json:"_name,omitempty"`
}

func (s *GeoDistanceFeatureQuery) UnmarshalJSON(data []byte) error {

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

		case "boost":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Boost", err)
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
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

		case "pivot":
			if err := dec.Decode(&s.Pivot); err != nil {
				return fmt.Errorf("%s | %w", "Pivot", err)
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "QueryName_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		}
	}
	return nil
}

// NewGeoDistanceFeatureQuery returns a GeoDistanceFeatureQuery.
func NewGeoDistanceFeatureQuery() *GeoDistanceFeatureQuery {
	r := &GeoDistanceFeatureQuery{}

	return r
}

type GeoDistanceFeatureQueryVariant interface {
	GeoDistanceFeatureQueryCaster() *GeoDistanceFeatureQuery
}

func (s *GeoDistanceFeatureQuery) GeoDistanceFeatureQueryCaster() *GeoDistanceFeatureQuery {
	return s
}

func (s *GeoDistanceFeatureQuery) DistanceFeatureQueryCaster() *DistanceFeatureQuery {
	o := DistanceFeatureQuery(s)
	return &o
}
