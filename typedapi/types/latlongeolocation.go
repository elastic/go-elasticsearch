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

// LatLonGeoLocation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/Geo.ts#L120-L129
type LatLonGeoLocation struct {
	// Lat Latitude
	Lat Float64 `json:"lat"`
	// Lon Longitude
	Lon Float64 `json:"lon"`
}

func (s *LatLonGeoLocation) UnmarshalJSON(data []byte) error {

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

		case "lat":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Lat", err)
				}
				f := Float64(value)
				s.Lat = f
			case float64:
				f := Float64(v)
				s.Lat = f
			}

		case "lon":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Lon", err)
				}
				f := Float64(value)
				s.Lon = f
			case float64:
				f := Float64(v)
				s.Lon = f
			}

		}
	}
	return nil
}

// NewLatLonGeoLocation returns a LatLonGeoLocation.
func NewLatLonGeoLocation() *LatLonGeoLocation {
	r := &LatLonGeoLocation{}

	return r
}

type LatLonGeoLocationVariant interface {
	LatLonGeoLocationCaster() *LatLonGeoLocation
}

func (s *LatLonGeoLocation) LatLonGeoLocationCaster() *LatLonGeoLocation {
	return s
}

func (s *LatLonGeoLocation) GeoLocationCaster() *GeoLocation {
	o := GeoLocation(s)
	return &o
}
