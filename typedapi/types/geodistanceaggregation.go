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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geodistancetype"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// GeoDistanceAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/aggregations/bucket.ts#L176-L182
type GeoDistanceAggregation struct {
	DistanceType *geodistancetype.GeoDistanceType `json:"distance_type,omitempty"`
	Field        *string                          `json:"field,omitempty"`
	Meta         Metadata                         `json:"meta,omitempty"`
	Name         *string                          `json:"name,omitempty"`
	Origin       GeoLocation                      `json:"origin,omitempty"`
	Ranges       []AggregationRange               `json:"ranges,omitempty"`
	Unit         *distanceunit.DistanceUnit       `json:"unit,omitempty"`
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
			o := string(tmp)
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
