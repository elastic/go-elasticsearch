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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// DistanceFeatureQueryBaseGeoLocationDistance type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/query_dsl/specialized.ts#L40-L60
type DistanceFeatureQueryBaseGeoLocationDistance struct {
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

func (s *DistanceFeatureQueryBaseGeoLocationDistance) UnmarshalJSON(data []byte) error {

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
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "origin":
			if err := dec.Decode(&s.Origin); err != nil {
				return err
			}

		case "pivot":
			if err := dec.Decode(&s.Pivot); err != nil {
				return err
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
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

// NewDistanceFeatureQueryBaseGeoLocationDistance returns a DistanceFeatureQueryBaseGeoLocationDistance.
func NewDistanceFeatureQueryBaseGeoLocationDistance() *DistanceFeatureQueryBaseGeoLocationDistance {
	r := &DistanceFeatureQueryBaseGeoLocationDistance{}

	return r
}
