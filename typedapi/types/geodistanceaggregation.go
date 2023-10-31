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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geodistancetype"
)

// GeoDistanceAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/aggregations/bucket.ts#L380-L403
type GeoDistanceAggregation struct {
	// DistanceType The distance calculation type.
	DistanceType *geodistancetype.GeoDistanceType `json:"distance_type,omitempty"`
	// Field A field of type `geo_point` used to evaluate the distance.
	Field *string  `json:"field,omitempty"`
	Meta  Metadata `json:"meta,omitempty"`
	Name  *string  `json:"name,omitempty"`
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
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		case "origin":
			if err := dec.Decode(&s.Origin); err != nil {
				return err
			}

		case "ranges":
			if err := dec.Decode(&s.Ranges); err != nil {
				return err
			}

		case "unit":
			if err := dec.Decode(&s.Unit); err != nil {
				return err
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
