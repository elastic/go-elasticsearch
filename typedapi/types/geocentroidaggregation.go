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

// GeoCentroidAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/metric.ts#L125-L128
type GeoCentroidAggregation struct {
	Count *int64 `json:"count,omitempty"`
	// Field The field on which to run the aggregation.
	Field    *string     `json:"field,omitempty"`
	Location GeoLocation `json:"location,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *GeoCentroidAggregation) UnmarshalJSON(data []byte) error {

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

		case "count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = &value
			case float64:
				f := int64(v)
				s.Count = &f
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "location":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Location", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		location_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Location", err)
				}

				switch t {

				case "lat", "lon":
					o := NewLatLonGeoLocation()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Location", err)
					}
					s.Location = o
					break location_field

				case "geohash":
					o := NewGeoHashLocation()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Location", err)
					}
					s.Location = o
					break location_field

				}
			}
			if s.Location == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Location); err != nil {
					return fmt.Errorf("%s | %w", "Location", err)
				}
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return fmt.Errorf("%s | %w", "Missing", err)
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		}
	}
	return nil
}

// NewGeoCentroidAggregation returns a GeoCentroidAggregation.
func NewGeoCentroidAggregation() *GeoCentroidAggregation {
	r := &GeoCentroidAggregation{}

	return r
}
