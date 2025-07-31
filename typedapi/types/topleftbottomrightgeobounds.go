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
)

// TopLeftBottomRightGeoBounds type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/Geo.ts#L161-L164
type TopLeftBottomRightGeoBounds struct {
	BottomRight GeoLocation `json:"bottom_right"`
	TopLeft     GeoLocation `json:"top_left"`
}

func (s *TopLeftBottomRightGeoBounds) UnmarshalJSON(data []byte) error {

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

		case "bottom_right":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "BottomRight", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		bottomright_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "BottomRight", err)
				}

				switch t {

				case "lat", "lon":
					o := NewLatLonGeoLocation()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "BottomRight", err)
					}
					s.BottomRight = o
					break bottomright_field

				case "geohash":
					o := NewGeoHashLocation()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "BottomRight", err)
					}
					s.BottomRight = o
					break bottomright_field

				}
			}
			if s.BottomRight == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.BottomRight); err != nil {
					return fmt.Errorf("%s | %w", "BottomRight", err)
				}
			}

		case "top_left":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "TopLeft", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		topleft_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "TopLeft", err)
				}

				switch t {

				case "lat", "lon":
					o := NewLatLonGeoLocation()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "TopLeft", err)
					}
					s.TopLeft = o
					break topleft_field

				case "geohash":
					o := NewGeoHashLocation()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "TopLeft", err)
					}
					s.TopLeft = o
					break topleft_field

				}
			}
			if s.TopLeft == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.TopLeft); err != nil {
					return fmt.Errorf("%s | %w", "TopLeft", err)
				}
			}

		}
	}
	return nil
}

// NewTopLeftBottomRightGeoBounds returns a TopLeftBottomRightGeoBounds.
func NewTopLeftBottomRightGeoBounds() *TopLeftBottomRightGeoBounds {
	r := &TopLeftBottomRightGeoBounds{}

	return r
}
