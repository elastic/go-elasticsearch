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

// GeoBoundsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/Aggregate.ts#L327-L333
type GeoBoundsAggregate struct {
	Bounds GeoBounds `json:"bounds,omitempty"`
	Meta   Metadata  `json:"meta,omitempty"`
}

func (s *GeoBoundsAggregate) UnmarshalJSON(data []byte) error {

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

		case "bounds":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Bounds", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		bounds_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Bounds", err)
				}

				switch t {

				case "bottom", "left", "right", "top":
					o := NewCoordsGeoBounds()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Bounds", err)
					}
					s.Bounds = o
					break bounds_field

				case "bottom_right", "top_left":
					o := NewTopLeftBottomRightGeoBounds()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Bounds", err)
					}
					s.Bounds = o
					break bounds_field

				case "bottom_left", "top_right":
					o := NewTopRightBottomLeftGeoBounds()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Bounds", err)
					}
					s.Bounds = o
					break bounds_field

				case "wkt":
					o := NewWktGeoBounds()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Bounds", err)
					}
					s.Bounds = o
					break bounds_field

				}
			}
			if s.Bounds == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Bounds); err != nil {
					return fmt.Errorf("%s | %w", "Bounds", err)
				}
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		}
	}
	return nil
}

// NewGeoBoundsAggregate returns a GeoBoundsAggregate.
func NewGeoBoundsAggregate() *GeoBoundsAggregate {
	r := &GeoBoundsAggregate{}

	return r
}
